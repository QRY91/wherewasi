# wherewasi 🪂

**AI Context Generation CLI - The Ripcord for Developer Flow**

Pull the cord, get context, keep building. Invisible until you need it.

> *"Stop explaining your project again. Pull the ripcord instead."*

## 🤖 AI Collaboration Transparency

This project documentation and development has been enhanced through systematic AI collaboration following QRY Labs methodology:

- **Human-Centered Development**: All core functionality, architecture decisions, and strategic direction remain human-controlled
- **AI-Enhanced Documentation**: AI assistants help improve documentation quality and systematic presentation
- **Transparent Attribution**: AI collaboration is acknowledged openly as part of QRY's commitment to ethical technology use
- **Context Generation Focus**: wherewasi specifically enables better AI collaboration by providing systematic context
- **Systematic Methodology**: AI collaboration follows structured procedures documented in `/ai/` directory

**Core Principle**: AI enhances human capability rather than replacing human judgment. wherewasi exemplifies this by generating context that enables more effective human-AI collaboration while keeping humans in control of strategic decisions.

## 🎯 What It Actually Does

**wherewasi** tracks your development ecosystem passively and generates dense AI context summaries on demand. Like a ripcord - invisible until you need perfect context for AI collaboration.

**Current reality (working MVP):**
- Scans multiple git projects in your ecosystem
- Finds recent activity and key files automatically  
- Generates structured context for AI handoffs
- Copies to clipboard or saves to database
- Cross-project search with file:line precision

## 🪂 Quick Start (30 seconds)

```bash
# Build and install
go build -o wherewasi .

# Initialize ripcord tracking
./wherewasi start

# Check ecosystem status  
./wherewasi status

# Pull the ripcord - get AI context
./wherewasi pull

# Context now in clipboard → paste into AI chat
```

## ⚡ Core Commands

```bash
# Start passive tracking
wherewasi start

# Get instant AI context (clipboard ready)
wherewasi pull

# Check ecosystem status
wherewasi status

# Search across projects  
wherewasi pull --keyword "whisper" --project "miqro"

# View context history
wherewasi pull --history
```

## 🔍 What Gets Tracked

**File Intelligence:**
- Recent git commits and file changes
- Key project files (README, main files, configs)
- Chat history files with conversation ranges
- Active session detection (recent modifications)

**Cross-Project Awareness:**
- 13+ projects in QRY ecosystem currently tracked
- File:line precision in search results  
- Chat conversation context with line ranges
- Project relationships and dependencies

## 📊 Sample Output

```
🪂 WHEREWASI CONTEXT - miqro project
═══════════════════════════════════════

📍 PROJECT: miqro (AI voice transcription tool)
📅 TIMEFRAME: Last 7 days  
⚡ ACTIVE SESSION: Recent file modifications detected

🎯 PROJECT SUMMARY:
Local-first audio transcription using Whisper AI. Working MVP with successful 
test results. Part of QRY ecosystem for hands-free development workflow.

📝 RECENT ACTIVITY:
• Enhanced audio processing pipeline  
• Integrated clipboard functionality
• Added ecosystem integration patterns

🔍 KEY FILES:
main.py → Core transcription logic with Whisper integration
README.md → Project documentation and test results
requirements.txt → Dependencies (whisper, pyaudio, etc.)

💬 RECENT DISCUSSIONS:
cursor_miqro_setup.md:245-267 → Whisper installation and config
cursor_debug_session.md:89-102 → Audio input troubleshooting
```

## 🏗️ Current Architecture

**Local-First Design:**
- SQLite database in `~/.local/share/wherewasi/`
- XDG-compliant directory handling  
- No cloud dependencies or telemetry
- Cross-platform compatibility (Go)

**Testing & CI:**
- Database tests covering all operations
- CLI integration tests for core workflows
- GitHub Actions pipeline with 3 stages
- Coverage reporting and quality gates

## 🎯 Honest Status Report

**What's Working Today:**
- ✅ Context generation from git + file analysis
- ✅ Cross-project ecosystem intelligence  
- ✅ Clipboard integration for instant AI handoff
- ✅ Persistent context storage and search
- ✅ Chat history scanning with line precision
- ✅ Basic CI/CD pipeline with test coverage

**What's Still Rough:**
- 🔄 No background daemon (manual `start` required)
- 🔄 Basic search (no semantic/AI-powered matching)
- 🔄 Limited file type intelligence  
- 🔄 No integration with other QRY tools yet

**What's Planned:**
- Background daemon for truly passive tracking
- Smarter pattern recognition across projects
- Integration with uroboro and doggowoof
- Enhanced context density optimization

## 🛠️ Development

```bash
# Run tests
go test -v ./...

# Run specific test suites
go test -v ./internal/database  # Database tests
go test -v .                    # CLI integration tests

# Build binary
go build -o wherewasi .
```

## 📈 The Vision (Long-term)

**Ripcord for AI Collaboration:**
- Pull cord → get perfect context → continue building
- Works across complex multi-project ecosystems
- Eliminates "explain my project again" overhead  
- Maintains developer flow state during AI handoffs

**Philosophy**: Be invisible until needed. Deploy instantly when pulled. Save developers from context loss crashes.

## 🔗 QRY Ecosystem Integration

**Complementary Tools:**
- **uroboro**: Content generation from captured work
- **doggowoof**: Alert monitoring and triage  
- **osmotic**: World state awareness and intelligence

**Shared Principles:**
- Local-first privacy protection
- Respect for developer flow states
- Tools that work quietly in background
- No cloud dependencies or data sharing

---

**Status**: Working MVP with basic ripcord functionality  
**Next**: Background daemon and smarter context intelligence  
**Philosophy**: Underpromise, overdeliver. Show the working tool, not AI buzzwords.

*Pull the cord. Get context. Keep building.* 🪂 