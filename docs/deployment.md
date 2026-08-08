# Deployment

Production runs on k3s, reconciled by ArgoCD. There is no SSH step, no
`docker compose` step, and nothing to run by hand on a server. **Every change
to production is a git commit.** If you find yourself typing `kubectl apply`
or `kubectl set image`, stop — `selfHeal: true` will revert it within seconds.

## The moving parts

| Thing | Where | Does what |
|---|---|---|
| App source | `FACorreiaa/homepage-go` (this repo) | the site |
| Build + push | `.github/workflows/build-push.yml` | builds the image, pushes to GHCR, bumps the tag in the infra repo |
| Cluster config | `LuminaVault/LuminaVaultInfra` | Helm values, ArgoCD Applications, SealedSecrets |
| Chart | `LuminaVaultInfra/charts/norviq-app` | renders Deployment, Service, Ingress, PVCs, middlewares |
| This app's values | `LuminaVaultInfra/apps/facorreia/values.yaml` | everything specific to facorreia.com |
| ArgoCD Application | `LuminaVaultInfra/argocd/apps/facorreia.yaml` | watches `main`, syncs automatically |
| Ingress | Traefik (bundled with k3s) | TLS via cert-manager, security headers middleware |

Namespace is `facorreia`, Helm release name is `site`. So the Deployment is
`site` in namespace `facorreia` — that is why every `kubectl` command below
reads `-n facorreia ... deploy/site`.

---

## 1. Deploying a code change

1. **Merge to `main`.** That is the only manual step.
2. CI builds the Dockerfile and pushes two tags to GHCR:
   `ghcr.io/facorreiaa/homepage-go:<git-sha>` and `:latest`.
3. CI clones `LuminaVaultInfra` using the `INFRA_TOKEN` PAT, `sed`s the new SHA
   into `apps/facorreia/values.yaml`, commits as `facorreia-ci`, and pushes to
   `main`.
4. ArgoCD sees the changed tag and rolls the deployment (`Recreate`, see
   [Expect downtime](#expect-downtime-on-every-rollout)).

Verify:

```sh
kubectl -n facorreia get deploy site -o jsonpath='{.spec.template.spec.containers[0].image}'; echo
kubectl -n facorreia rollout status deploy/site
```

Or from outside the cluster, which needs no kubeconfig:

```sh
curl -sS -o /dev/null -w '%{http_code}\n' https://facorreia.com/healthz
```

**Rollback** — edit `tag:` in `apps/facorreia/values.yaml` to the previous SHA
and push. Don't `kubectl set image`; `selfHeal: true` reverts it within
seconds. Git is the only lever.

The deployment pins a SHA, never `:latest` — so a rollout only happens when a
commit says so.

### The CI ↔ values.yaml contract

Step 3 is a blind `sed`:

```sh
sed -i "s|^  tag: .*|  tag: ${GITHUB_SHA} # CI overwrites with the git SHA on each main push|" \
  apps/facorreia/values.yaml
```

It rewrites **every** line that begins with exactly two spaces followed by
`tag: `. Today that is one line, under `image:`. If you ever add another
two-space-indented `tag:` key anywhere in that file, CI will silently
overwrite it too. Check before you commit:

```sh
grep -c '^  tag: ' apps/facorreia/values.yaml   # must be 1
```

---

## 2. Adding a plain (non-secret) env var

Anything in `env:` is **plain text in a public git repo**. Config only. For
anything sensitive see [section 3](#3-adding-a-secret-env-var).

1. Read it in the app the way the rest of the code does, so a missing value
   degrades instead of crashing:

   ```go
   region := getEnvOrDefault("SERVING_REGION", "EU")
   ```

2. Add it to `LuminaVaultInfra/apps/facorreia/values.yaml`:

   ```yaml
   env:
     - name: PORT
       value: "8090"
     - name: SERVING_REGION
       value: "Falkenstein, DE"   # quote anything with a comma or colon
   ```

   Numbers must be quoted — `value:` is a string field and Kubernetes rejects
   a bare integer.

3. **Render it before committing.** This catches YAML and type mistakes that
   would otherwise fail in-cluster:

   ```sh
   cd LuminaVaultInfra
   helm template site charts/norviq-app -f apps/facorreia/values.yaml \
     --namespace facorreia | grep -A1 SERVING_REGION
   ```

4. Commit and push to `main`. ArgoCD picks it up; changing `env` changes the
   pod template, so it triggers a rollout.

5. Verify:

   ```sh
   kubectl -n facorreia exec deploy/site -- printenv SERVING_REGION
   ```

Also update `.env.example` in this repo so local development matches
production.

---

## 3. Adding a secret env var

Secrets go through SealedSecrets — encrypted with the cluster's public key, so
the sealed file is safe to commit and only the in-cluster controller can
decrypt it.

**The order matters. Do these as two separate commits.**

### Commit 1 — create and seal the secret

```sh
kubectl create secret generic site-env -n facorreia \
  --from-literal=SOME_TOKEN=the-real-value \
  --dry-run=client -o yaml \
  | kubeseal --controller-name sealed-secrets-controller \
             --controller-namespace kube-system -o yaml \
  > secrets/facorreia/site-env.yaml
```

Commit **only** that file. `argocd/apps/secrets.yaml` watches `path: secrets`
with `recurse: true`, so a new subdirectory is picked up with no extra
Application. Wait for the controller to materialize the real Secret:

```sh
kubectl -n facorreia get secret site-env
```

Add the file to the inventory table in `LuminaVaultInfra/secrets/README.md`
while you are there.

### Commit 2 — point the deployment at it

```yaml
# apps/facorreia/values.yaml
envFromSecret: site-env
```

Every key in the Secret becomes an env var.

### Why the order matters

`envFrom` against a Secret that does not exist is a **hard failure** — the pod
sits in `CreateContainerConfigError` and never becomes ready. This chart has no
`optional: true` escape hatch. Committing both at once means the values change
can reach the cluster before the Secret does, and the site goes down until the
Secret lands.

To add a key to an **existing** secret, re-seal the whole Secret with all its
keys — `kubeseal` output is not mergeable, and a partial re-seal drops the keys
you left out.

### When not to use a Secret

If the value only has to be *stable and private*, not *shared*, consider
generating it in the app and persisting it to the PVC instead. That is what
`VISIT_SALT` does: the app writes `visit-salt` next to the SQLite database on
first boot and reuses it forever. No secret store, nothing in git, and rotating
it is deleting one file. See `resolveVisitSalt` in
`internal/service/visits.go`.

---

## 4. Changing resources, probes, or storage

All in `apps/facorreia/values.yaml`; the chart's defaults live in
`charts/norviq-app/values.yaml`.

```yaml
resources:
  requests: {cpu: 50m, memory: 64Mi}
  limits:   {memory: 128Mi}

probes:
  readiness: {path: /healthz, initialDelaySeconds: 5, periodSeconds: 10, failureThreshold: 6}
  liveness:  {path: /healthz, initialDelaySeconds: 15, periodSeconds: 20, failureThreshold: 3}
```

**Growing a volume**: raising `persistence.size` does **not** resize an
existing PVC — Kubernetes will not shrink or, with most storage classes,
expand one in place. `local-path` does not support expansion at all. You have
to create a new PVC and copy the data across. Plan the size up front.

**Do not raise `replicaCount`.** It is pinned to `1` on purpose:

- SQLite lives on a `ReadWriteOnce` volume, so only one pod can mount it.
- `/stats` streams live visitor events over SSE from an **in-process**
  broadcaster. A second replica would show each client only the traffic that
  happened to hit its own pod.

Scaling this app out means moving off SQLite and moving the broadcaster to
something shared (Redis pub/sub, NATS). It is not a values change.

---

## 5. Changing HTTP headers or TLS

Security headers are a Traefik middleware, rendered by
`charts/norviq-app/templates/security-headers-middleware.yaml` and enabled in
`apps/facorreia/values.yaml`:

```yaml
securityHeaders:
  enabled: true
  contentSecurityPolicy: >-
    default-src 'self';
    script-src 'self' 'unsafe-inline' 'unsafe-eval';
    ...
```

The `Caddyfile` in this repo is **dead config** — it has not been read since
the k3s cutover in `173c83a`. Editing it does nothing. Change headers in the
infra repo.

If you add a third-party script, font, or API call, the CSP must be widened or
the browser will silently block it. Check the console after any such change:

```sh
curl -sS -o /dev/null -D - https://facorreia.com/ | grep -i content-security-policy
```

---

## Expect downtime on every rollout

The chart switches to `strategy: Recreate` whenever persistence is enabled:

```yaml
{{- if or .Values.persistence.enabled .Values.additionalPersistence }}
strategy:
  type: Recreate
```

It has to. The PVC is `ReadWriteOnce`, so a new pod cannot mount the volume
while the old one still holds it — a rolling update would deadlock, with the
new pod stuck in `ContainerCreating` forever.

So the old pod is **terminated first**, then the new one starts. Measured on
the 2026-08-08 rollout: roughly 20–40 seconds of 503, ~90 seconds end to end
from `git push` to serving. Fine for a personal site; just don't expect
zero-downtime deploys, and don't deploy while someone is looking at it.

Because the process restarts, anything held only in memory resets:

- `VISITORS NOW` drops to 0 and refills as people arrive.
- The SSE broadcaster drops its subscribers; browsers reconnect on their own
  (`EventSource` retries automatically).
- Uptime restarts, which is what `UPTIME` on the landing page is showing.

Anything on the PVC survives: SQLite, blog stats, and the visitor salt.

---

## Troubleshooting

```sh
# what ArgoCD thinks
kubectl -n argocd get app facorreia
kubectl -n argocd describe app facorreia | tail -30

# pod state and recent events
kubectl -n facorreia get pods
kubectl -n facorreia describe pod -l app.kubernetes.io/name=site | tail -40
kubectl -n facorreia logs deploy/site --tail=100

# force a sync instead of waiting for the ~3 min reconcile
argocd app sync facorreia
```

| Symptom | Usual cause |
|---|---|
| `CreateContainerConfigError` | `envFromSecret` points at a Secret that does not exist yet — see [section 3](#3-adding-a-secret-env-var) |
| `ContainerCreating` forever | two pods fighting over one `ReadWriteOnce` PVC; check `replicaCount` and that `strategy` is `Recreate` |
| `ImagePullBackOff` | the tag in values.yaml is not in GHCR — check the build-push run actually finished |
| Deploy "succeeded" but the site is old | ArgoCD synced the chart but CI never bumped the tag; check `INFRA_TOKEN` has not expired |
| A change reverts by itself | you edited live state with `kubectl`; `selfHeal: true` undid it. Commit it instead |
| Site fine, one feature dead, console shows blocked resource | the CSP — see [section 5](#5-changing-http-headers-or-tls) |
