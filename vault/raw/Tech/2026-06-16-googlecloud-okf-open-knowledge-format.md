# Google Cloud Introduces Open Knowledge Format (OKF) — Formalizing the LLM-Wiki Pattern

**Author:** @GoogleCloudTech (Google Cloud Tech)
**Date:** Jun 16, 2026
**URL:** https://x.com/GoogleCloudTech/status/2067012903337664886
**Engagement:** 126.9K Views · 46 Replies · 232 Reposts · 1,925 Likes · 1,944 Bookmarks

---

## Full Tweet Content

**Introducing the Open Knowledge Format (OKF), an open specification that formalizes the LLM-wiki pattern into a portable, interoperable format.**

AI is only as smart as the context we give it. As we build more advanced, agentic AI systems, they need accurate metadata and context. In organizations, that context is locked inside fragmented data catalogs, isolated wikis, scattered code comments, or the minds of senior engineers. Every time a new AI agent is built, teams are forced to solve the exact same context-assembly problem from scratch.

To solve this, we've announced OKF, a vendor-neutral, open specification that formalizes the "LLM-wiki pattern" into a portable, interoperable format. It provides a standardized way to represent the enterprise knowledge that modern AI systems rely on.

---

## OKF Specification Highlights

- **Just markdown:** readable in any editor, renderable on GitHub, indexable by any search tool
- **Just files:** shippable as a tarball, hostable in any git repo, mountable on any filesystem
- **Just YAML frontmatter:** for the small set of structured fields that need to be queryable:
  - type
  - title
  - description
  - resource
  - tags
  - timestamp

---

## Reference Implementations Shipped

1. **Enrichment agent for BigQuery**
2. **Static HTML visualizer**
3. **Live sample bundles** on GitHub (@github): goo.gle/4uGvAEe, goo.gle/4uGvBbg

---

## Significance

This directly aligns with the **LLM-wiki / knowledge-base pattern** we use in the FACorreia vault (raw/ → wiki/ compilation with YAML frontmatter, markdown files, git/Syncthing portability). Google is formalizing this exact pattern as an open spec.

---

*This is Hermes' view, not financial advice. Always do your own research and manage your own risk.*