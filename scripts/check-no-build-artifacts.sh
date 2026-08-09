#!/usr/bin/env bash
# Fails CI if generated/build/runtime artifacts are tracked in git.
set -euo pipefail

patterns=(
  '*_templ.go'
  'assets/css/output.css'
  'assets/css/sources.generated.css'
  'app.db'
  'blog_stats.json'
  'vault/raw/*.md'
  'vault/raw/**/*.md'
  # air's scratch dir. tmp/main is a compiled binary; it was tracked for a long
  # time and is most of why .git is 81MB.
  'tmp/*'
  'tmp/**'
)

bad=0
for p in "${patterns[@]}"; do
  matches=$(git ls-files -- "$p")
  if [[ -n "$matches" ]]; then
    echo "ERROR: tracked build/runtime artifacts match '$p':" >&2
    echo "$matches" | head -5 >&2
    bad=1
  fi
done

if [[ "$bad" -ne 0 ]]; then
  echo "Remove them with: git rm --cached <file> (and ensure .gitignore covers them)" >&2
  exit 1
fi

echo "OK: no tracked build artifacts."
