---
title: "Homepage on the cluster, Forgejo on the way"
summary: "I wired getHomepage into k3s and exposed it over Tailscale. Next up: self-hosting git with Forgejo so every prod app lives on my own cluster."
category: engineering
date: "2026-08-20"
---

Today was the day I finally added [getHomepage](https://gethomepage.dev/installation/k8s/) to my cluster.

It went in as a normal k8s workload, and I made it reachable through my tailnet instead of poking another hole in the firewall. That means the dashboard — services, status, bookmarks, all of it — is just there when I'm on the network, without being internet-facing. Clean.

The bigger shift this points at: I'm going to move all the apps I have in prod to self-hosted git on the cluster too. The plan is Forgejo. Right now my repos live elsewhere, which means the cluster depends on an external service I don't control for the thing that defines what runs on it. Closing that loop — code, CI, and runtime all on hardware I own — is the point. Self-hosted git on the same tailnet as the workload is the obvious next step.

One app at a time. Homepage is the beachhead. Forgejo is next.
