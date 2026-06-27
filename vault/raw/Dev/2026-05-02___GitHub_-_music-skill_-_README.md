---
source: GitHub
url: https://github.com/KaleLjl/music-skill
date: 2026-05-02
tags: [Dev, Hermes, Skill]
project: music-skill
description: MIDI analysis/modification + FluidSynth rendering skill for Hermes Agent — generates music programmatically
---

# music-skill

# music-skill

music-skill is an agent-ready toolkit for arranging workflows with MIDI files. It is designed to help an agent modify MIDI songs with Python (mainly through mido-based scripts), then render edited MIDI files into playable audio with FluidSynth.

This project is being prepared for the Hermes May 2026 hackathon. It is designed as a reusable music skill package that can be used by Hermes agents, other AI-agent runtimes, or directly by developers from the command line.

## What It Can Do

| Capability                         | Description                                                                          |
| ---------------------------------- | ------------------------------------------------------------------------------------ |
| Modify MIDI instruments by channel | Use a Python script (based on mido) to update channel program changes in MIDI files. |
| Analyze MIDI instruments           | Read `program_change` events and produce a channel/program/instrument summary.       |
| MIDI to WAV                        | Use FluidSynth and a SoundFont to render MIDI into playable WAV audio.               |

## Requirements

- `uv` for isolated Python execution and dependency management.
- Python `>=3.10,<3.12`.
- `mido` for MIDI parsing and modification in Python scripts.
- FluidSynth for MIDI-to-WAV rendering.
- At least one `.sf2` SoundFont under `assets/soundfront/`.

Install FluidSynth with one of:

```bash
# macOS
brew install fluidsynth

# Ubuntu/Debian
sudo apt install fluidsynth

# Fedora/RHEL
sudo dnf install fluidsynth

# Windows, with Chocolatey
choco install fluidsynth
```
## Install steps

Clone this project into your agent skill folder, then ask your agent to install the dependencies for you.

## Notes for Hackathon Use

This package is meant to make music-file manipulation predictable during the Hermes May 2026 hackathon. Keep new workflows repeatable, preserve original MIDI files, write generated outputs next to their source files, and prefer structured JSON summaries when passing music metadata between agents or tools.

## See the demo 
https://x.com/CalebLele_go/status/2050102655196754319?s=20  # Limit README size
