# ✦ Gacha Pulls & Banner Planner ✦

> *A sleek, Cyberpunk / Catppuccin Mocha TUI application for tracking gacha savings, soft/hard pity counters, and banner guarantee readiness across **Honkai: Star Rail** & **Reverse: 1999**.*

![Go Version](https://img.shields.io/badge/Go-1.27%2B-00ADD8?style=for-the-badge&logo=go)
![TUI Framework](https://img.shields.io/badge/Charm-Bubbletea%20%26%20Bubbles-cba6f7?style=for-the-badge)
![Theme](https://img.shields.io/badge/Theme-Catppuccin%20Mocha-fab387?style=for-the-badge)

---

## 🌟 Features

- 📝 **Interactive Form Input**: Real-time form fields built with `bubbles/textinput` for editing Stellar Jades, Special Passes, Clear Drops, Unilogs, Current Pity, and Guaranteed 50/50 status.
- 📊 **Side-by-Side Dashboard**: High-level overview displaying total savings, total combined pulls, current pity states, and visual readiness indicators (`✅ Ready for Guarantee` / `⚡ Need X More Pulls`).
- 🌌 **Honkai: Star Rail Planner**:
  - Currency conversion: `160 Stellar Jades = 1 Special Pass`.
  - Pity Engine: Soft pity at `74`, Hard pity at `90`.
  - Worst-case guarantee calculator (`90` pulls if Guaranteed, `180` pulls if 50/50).
- 📜 **Reverse: 1999 Planner**:
  - Currency conversion: `180 Clear Drops = 1 Unilog`.
  - Pity Engine: Soft pity at `60`, Hard pity at `70`.
  - Worst-case guarantee calculator (`70` pulls if Guaranteed, `140` pulls if 50/50).
- 💾 **Real-Time JSON Persistence**: Auto-saves user savings & pity profile at `~/.config/gacha-planner/data.json` immediately as you type.
- 🎨 **Catppuccin Mocha Aesthetics**: Custom color palette designed with Lipgloss for dark & vibrant terminal aesthetics.

---

## 🛠️ Tech Stack

| Component | Technology | Description |
|---|---|---|
| **Language** | [Go](https://go.dev/) (1.27+) | High performance compiled language |
| **TUI Engine** | [Charm Bubbletea](https://github.com/charmbracelet/bubbletea) | Elm architecture for terminal UIs |
| **Form Component** | [Charm Bubbles](https://github.com/charmbracelet/bubbles) | Interactive `textinput` form components |
| **Styling Engine** | [Charm Lipgloss](https://github.com/charmbracelet/lipgloss) | Declarative terminal styling & layouts |

---

## 🚀 Quick Start

### Prerequisites
Make sure you have [Go](https://go.dev/dl/) installed (version 1.21+ or 1.27+).

### 1. Clone & Run Directly
```bash
# Clone the repository
git clone https://github.com/Elio-AGR/gacha-planner-tui.git
cd gacha-planner-tui

# Run directly
go run main.go
```

### 2. Build Executable Binary
```bash
# Compile binary
go build -o gacha-planner-tui main.go

# Run the binary
./gacha-planner-tui
```

### 3. Run Unit Tests
```bash
go test -v ./...
```

---

## ⌨️ Controls & Keyboard Shortcuts

| Key | Description |
|---|---|
| `1` / `2` / `3` | Switch directly to **Dashboard**, **R1999**, or **HSR** tab |
| `Tab` / `Shift+Tab` | Cycle forward / backward through tabs |
| `↑` / `↓` or `Enter` | Navigate between form input fields |
| `Type (0-9)` | Edit numeric values in real-time |
| `Space` or `g` | Toggle Guaranteed status (when on Guaranteed row) |
| `q` or `Ctrl+C` | Auto-save profile & quit application |

---

## 💾 Storage & Configuration Path

User data is stored locally in JSON format:
```text
~/.config/gacha-planner/data.json
```

---

## 📄 License

Distributed under the MIT License.
