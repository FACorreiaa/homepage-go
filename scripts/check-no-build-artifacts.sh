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
