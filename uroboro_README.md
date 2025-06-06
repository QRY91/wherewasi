# wherewasi 🧠💫

**The AI Context Rope Start - Less Explaining. More Building.**

Stop wasting time on context overhead. One pull, perfect AI context, back to building.

> *"I love my work. I want to DO my work. Communication should be collaboration, not context setup."*

## 🔥 THE PROBLEM

**Context overhead steals focus from building:**
- New AI chat = 10 minutes re-explaining your project stack
- "I'm working on X using Y, just implemented Z..."
- Lost momentum, broken flow state  
- Communication energy wasted on overhead, not collaboration

## ⚡ THE SOLUTION

```bash
# Track your work automatically
wherewasi start

# Pull the rope - instant context engine startup
wherewasi pull --hours 2
# → Dense, AI-ready summary. Zero mental overhead.

# Paste into AI chat and GO
# → More building, less explaining
```

## 🎯 CORE PHILOSOPHY

**"Less explaining. More building."**

- ✅ **Communication AS work** - Pair programming, reviews, brainstorming
- ❌ **Communication TO work** - Context setup, project re-explanation  
- 🎯 **Focus your energy** - Collaborate meaningfully, minimize overhead

## 🎯 CORE FEATURES

- **🧠 Smart Context Capture** - Git commits, file changes, terminal commands
- **⚡ Instant Ripchord** - 2-second context generation for AI chats
- **🔒 Privacy-First** - Everything stays local, zero cloud dependencies
- **🎨 Dense Summaries** - AI-optimized format, not human verbose
- **⏰ Time-Based Context** - Last hour, day, sprint, whatever you need

## 🚀 QUICK START

```bash
# Install
go install github.com/QRY91/wherewasi@latest

# Start tracking (background daemon)
wherewasi start

# Work normally... git commits, file changes tracked

# Need AI context? Pull the rope!
wherewasi pull --hours 4 | pbcopy
# → Perfect context copied, paste and build

# Check what's tracked
wherewasi status
```

## 🎪 USE CASES (WHERE WHEREWASI SHINES!)

- **Multi-Project Developers**: Context switch between uroboro, doggowoof, slopsquid without explanation fatigue
- **AI-Powered Workflows**: Start fresh chats with perfect context instantly
- **Remote Teams**: Share dense project summaries without meetings
- **Documentation**: Auto-generate "what I did today" summaries

## 🔧 INTEGRATION (PLAYS WELL WITH OTHERS!)

- **Git**: Automatic commit tracking and analysis
- **File System**: Smart file change detection
- **Terminal**: Command history with context awareness
- **AI Chats**: Optimized output format for Claude, GPT, etc.
- **uroboro**: Cross-pollinate insights for content generation

## 📊 LOCAL DATA (YOUR DATA STAYS HOME!)

Everything stays on YOUR machine:
- **SQLite database**: Work history, context summaries
- **Local analysis**: Pattern recognition, smart filtering
- **Privacy-first**: NO TELEMETRY, NO CLOUD DEPENDENCIES!

## 🛠️ ARCHITECTURE (SIMPLE BUT SMART!)

```
┌─ Go CLI ────────────┐    ┌─ Background Daemon ─┐    ┌─ SQLite DB ────────┐
│ • wherewasi start   │ ←→ │ • Git monitoring    │ ←→ │ • Work history     │
│ • wherewasi pull    │    │ • File watching     │    │ • Context cache    │
│ • wherewasi status  │    │ • Command tracking  │    │ • Smart summaries  │
└─────────────────────┘    └─────────────────────┘    └────────────────────┘
```

## 🎯 RIPCHORD OUTPUT EXAMPLE

```
Working on wherewasi (Go CLI for AI context generation). 
Recent: Created project structure, implemented git monitoring, 
added SQLite storage. Current focus: pull command implementation.
Tech stack: Go, SQLite, file watching. Recent files: main.go, 
git_monitor.go, database.go. Last commits: "Add git monitoring", 
"Implement basic CLI structure". Working directory: /projects/wherewasi.
```

## 🗺️ ROADMAP (THE PLAN!)

- [x] Project genesis and wine-fueled README
- [ ] Go CLI foundation with basic commands
- [ ] Git monitoring and commit analysis
- [ ] File system watching (smart filtering)
- [ ] Terminal command tracking
- [ ] SQLite storage and retrieval
- [ ] Pull context generation
- [ ] Time-based filtering (hours, days, sprints)
- [ ] Integration with uroboro ecosystem
- [ ] Landing page (wherewasi.dev)

## 🚨 WHY WHEREWASI? 🚨

Because your AI conversations deserve:
- ✅ **INSTANT CONTEXT** - No more 10-minute explanations
- ✅ **SMART SUMMARIES** - Dense, AI-optimized format
- ✅ **PRIVACY FIRST** - Your work stays on your machine
- ✅ **SEAMLESS WORKFLOW** - One command, perfect context
- ✅ **MULTI-PROJECT SUPPORT** - Handle complex project ecosystems

---

## 🍷 WINE-TIME DEVELOPMENT NOTES

*Built during a late-night wine session because context switching pain is REAL.*

**Philosophy**: If you're explaining the same project context to AI chats multiple times per day, you're doing it wrong.

**Target User**: Developers juggling multiple projects who use AI for coding assistance and need to stop wasting time on context setup.

---

***wherewasi: BECAUSE YOUR AI CHATS DESERVE PERFECT CONTEXT FROM SECOND ONE*** 🧠💫🚀