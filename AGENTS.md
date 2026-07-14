<claude-mem-context>
# Memory Context

# [facorreia-site-go] recent context, 2026-07-14 9:38am GMT+1

Legend: 🎯session 🔴bugfix 🟣feature 🔄refactor ✅change 🔵discovery ⚖️decision 🚨security_alert 🔐security_note
Format: ID TIME TYPE TITLE
Fetch details: get_observations([IDs]) | Search: mem-search skill

Stats: 32 obs (14,491t read) | 357,867t work | 96% savings

### Jul 9, 2026
15040 2:26p ⚖️ Portfolio Projects Showcase and Footer Redesign Planned
15041 2:28p 🔵 Portfolio Site Architecture and Showcase/Footer Structure Mapped
15105 7:14p 🔵 Projects Page Template Structure in facorreia-site-go
15106 7:15p 🔵 Project Data Model and Handler Architecture in facorreia-site-go
15107 " 🔵 CSS Design System: OKLCH Tokens, Geist Fonts, and Scroll Reveal Pattern
15108 7:16p ⚖️ Implementation Plan: Elevate Projects Showcase Page and Footer
15111 7:17p 🟣 Implementation Plan Written to Disk: Elevate Projects Showcase + Footer
15112 7:18p ⚖️ Implementation Plan Approved: Showcase + Footer Elevation Begins Execution
15113 7:19p 🟣 Showcase + Footer Elevation Work Broken Into 6 Tracked Tasks
15114 " 🟣 Proof-Slot Types Added to ProjectItem Model
15115 " 🟣 Featured Project Copy Rewritten to Outcome-Led Framing (Norviq First)
15120 " 🟣 All 4 Featured Project Descriptions Rewritten with Outcome-Led Copy
15121 7:20p 🔴 CSS Insertion Point for New Component Classes Confirmed
15122 " 🟣 New CSS Component Classes Added to input.css
15130 7:21p 🟣 Task 4 Begun: projects.templ Hero Section Updated
15131 7:23p 🔴 Build Error: System `templ` (v0.3.1020) Newer Than go.mod Pin (v0.3.1001) — Fixed by Using `go tool templ`
15132 " 🔴 Tasks 4 + 5 Completed — projects.templ and footer.templ Fully Reworked
15133 " 🔴 Build Clean and Generated Files Confirmed Correct — _templ.go Files Are Gitignored
15134 7:24p 🔴 Tailwind Build Pipeline Confirmed — output.css Must Be Regenerated for New Classes to Take Effect
15135 " 🔴 Tailwind CSS Rebuilt — All 5 New Component Classes Confirmed in output.css
15136 " 🔴 Proof-Slot Smoke Test Injected into Norviq Project — Temporary Data for Visual Verification
15145 " 🔴 Curl Smoke Test Returned All Zeros — Server Port Detection Failed, HTML Was Empty String
15146 7:25p 🔴 Server Did Start on Port 8090 — Curl Failure Was Race Condition (Server Killed Before Request)
15148 " 🔴 Smoke Test Data Reverted — Norviq Project Restored to Zero-Value Proof Slots
15150 7:26p 🔴 Clean-State Verification Hit Stale Server — Smoke Test Data Still Appearing from Previous Binary
15152 7:27p 🔴 Clean-State Verification Confirmed — All 6 Tasks Complete, Ready to Commit
S1163 Elevate /projects showcase page and global footer of facorreia-site-go Go portfolio site for freelance/consulting clients and recruiters — editorial redesign with proof slots, alternating layout, and footer conversion (Jul 9 at 7:27 PM)
### Jul 10, 2026
15368 7:31p ⚖️ Portfolio Projects Showcase & Footer Redesign Planned
15369 7:32p 🟣 CTA Button Micro-interaction Polish: Press Scale + Full Transition
15370 " 🔵 active:scale-[0.97] Not Emitted by Tailwind CSS Compiler
15372 " 🔵 Tailwind CSS Scan Gap: active:scale-[0.97] and duration-[250ms] Missing; cubic-bezier Present as Literal
15373 7:33p 🔵 active:scale-[0.97] IS Correctly Emitted — Prior Grep Was False Negative
15374 " 🟣 Live Server Verification: All Micro-interaction Classes Render Correctly on /projects Page
S1174 Portfolio UI Polish: Professional projects showcase and footer redesign with micro-interaction improvements and motion accessibility fixes (Jul 10 at 7:33 PM)
**Investigated**: - footer.templ: CTA button classes, transition utilities, hover/active states
    - ui/pages/about.templ: All CTA buttons (hero, contact section), aboutContactLink arrow animation
    - assets/css/output.css: Verified Tailwind compilation output for scale, duration, motion-safe, and cubic-bezier classes
    - Live server render at localhost:8090/projects: Confirmed all new HTML attributes present in rendered output
    - Go build pipeline: templ code generation, Go compilation, Tailwind CSS compilation

**Learned**: - Tailwind grep false negatives: bracket characters [ ] in class names like active:scale-[0.97] are regex metacharacters; must grep for raw values (0.97, scale-) to confirm emission
    - active:scale-[0.97] IS correctly emitted — 5 occurrences of "0.97" and 35 of "scale-" confirmed in output.css
    - motion-safe: Tailwind prefix works correctly — 8 confirmed hits in CSS output, and motion-safe:group-hover:translate-x-1 renders in live HTML
    - .templ files ARE being scanned by Tailwind content config (classes from them appear in compiled output)
    - The site uses Go + templ + HTMX + Alpine.js + Tailwind CSS stack

**Completed**: - Footer "Request proposal" button: transition-opacity → transition + active:scale-[0.97]
    - Footer "Email Fernando" button: transition-colors → transition + active:scale-[0.97]
    - About hero "Request proposal" button: transition-opacity active:scale-[0.99] → transition active:scale-[0.97]
    - About hero "Email me" button: transition-colors active:scale-[0.99] → transition active:scale-[0.97]
    - About contact section "Request proposal" link: transition-opacity → transition + active:scale-[0.97]
    - aboutContactLink arrow icon: transition-transform group-hover:translate-x-1 → motion-safe:transition-transform motion-safe:group-hover:translate-x-1 (accessibility fix)
    - Full review of 10 animation/motion issues completed: mobile drawer slide-in, scroll-top fade, smooth scroll gating, card hover, ungated arrow translates, imperceptible active scale, missing :active on CTAs, scroll reveal timing, dead marquee CSS
    - All fixes verified: go build clean, Tailwind emits correct classes, live server renders correct HTML

**Next Steps**: - Manual visual QA on real device recommended: verify 250ms/200ms drawer slide timing feels right for panel width, and check if card hover lift without shadow bloom reads crisper or flat
    - If card hover lift feels too flat without shadow, a pre-rendered shadow layer fix requires a small DOM change (card has overflow-hidden which prevents box-shadow animation)
    - Session appears to be wrapping up after successful build + live verification


Access 358k tokens of past work via get_observations([IDs]) or mem-search skill.
</claude-mem-context>