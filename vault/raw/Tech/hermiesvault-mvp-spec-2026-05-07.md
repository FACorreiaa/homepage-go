# HermiesVault MVP Specification

## Project Context
- **App Name:** HermiesVault (working title)
- **Type:** iOS second brain / note-taking with AI assistance
- **Tech Stack:** SwiftUI, Swift on Server (Hummingbird), Docker VPS, Hermes, Postgres, PGVector
- **Design Principles:** Simplicity, Privacy, Offline-first, Sci-fi/Futurism UI

## Core Value Proposition
"Obsidian for iPhone with an AI that learns from your notes"

## MVP Features (Minimal Viable Product)

### 1. Note-Taking Core
- **Create notes** with Obsidian-flavored markdown
- **Edit notes** with basic formatting (headers, lists, bold, italic)
- **Organize notes** with folders/tags
- **Search notes** (local only)
- **Offline-first** - all notes stored locally with CoreData/SQLite

### 2. AI Assistance (The "Smart" Part)
- **AI summarization** - Summarize long notes
- **AI expansion** - Help elaborate on ideas
- **Question answering** - Ask questions about your notes (RAG-based)
- **All AI processing happens locally** (using Hermes/local LLM) for privacy

### 3. Sync & Backup (Optional MVP, maybe v1)
- **Manual export** to JSON/Markdown
- **iCloud sync** (if feasible within MVP)
- **Local backup** to computer via file sharing

## UI/UX Design (Sci-Fi/Futurism)
- **Dark mode default** with neon accents (cyan/magenta)
- **Clean typography** - Inter or SF Compact fonts
- **Smooth animations** - spring physics, transitions
- **HUD-like elements** - subtle borders, glowing effects on active elements
- **Custom transitions** between notes and AI interface

## Architecture

### Client (iOS)
```
├── SwiftUI App (App + Scene)
├── CoreData Stack (Local storage)
│   ├── Note entity
│   ├── Tag entity
│   └── Relationship management
├── AI Service (Local)
│   ├── Embedding generation (Hermes/local model)
│   ├── RAG pipeline (PGVector locally)
│   └── Text generation (Hermes/local model)
├── Network Service (Optional)
│   ├── Sync API (if implementing sync)
│   └── Backup API
└── UI Components
    ├── Note editor (Markdown)
    ├── Note list (Searchable)
    ├── AI Assistant panel
    └── Dashboard (Recent/Starred notes)
```

### Server (Swift on Server - Hummingbird)
*Note: Server may be optional for MVP if keeping everything local*

```
├── API Server (Hummingbird)
│   ├── Authentication (optional)
│   ├── Sync endpoints (if needed)
│   └── Backup endpoints
├── Database (Postgres + PGVector)
│   ├── Notes table
│   ├── Embeddings table
│   └── Metadata tables
└── Background Workers
    ├── Embedding generation
    └── Cleanup tasks
```

## Technology Decisions

### LLM Choice
- **Primary:** Hermes (local model) for privacy
- **Fallback:** Smaller quantized model for faster inference
- **Why local?** Privacy is a core principle - user data never leaves device

### Vector Database
- **PGVector** embedded in Postgres
- Embeddings generated on-device, stored locally
- Optional: sync vectors to server if sync is implemented

### UI Framework
- **SwiftUI** for modern, declarative UI
- **Hummingbird** for Swift on Server if needed
- **Swift Concurrency** for async operations

## Privacy & Security
- **No data collection** - all notes stay on device
- **Optional sync** - user chooses to backup/export
- **On-device AI** - no cloud API calls for processing
- **Local encryption** - optional passcode/biometric protection

## Success Metrics (MVP)
- User can create and save a note
- User can ask AI a question about their notes and get relevant answer
- App works offline
- Smooth, responsive UI (60fps)

## Stretch Goals (Post-MVP)
- Voice input for notes
- Image uploads with OCR
- Cross-device sync (with end-to-end encryption)
- Plugin system (like Obsidian)
- Web version

## Risks & Mitigations
- **LLM performance on device:** Use quantized models, optimize inference
- **Storage limits:** CoreData with SQLite, implement cleanup policies
- **User acquisition:** Position as privacy-focused alternative to Notion/Obsidian
- **Complexity:** Start with MVP scope, resist feature creep

---
*Document created: 2026-05-07*
*Next: Project timeline and resource allocation*