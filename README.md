# ✦ Gacha Pulls & Banner Planner ✦

> *A sleek, Cyberpunk / Catppuccin Mocha TUI application for tracking gacha savings, soft/hard pity counters, and banner guarantee readiness across **Honkai: Star Rail** & **Reverse: 1999**.*

![Go Version](https://img.shields.io/badge/Go-1.27%2B-00ADD8?style=for-the-badge&logo=go)
![TUI Framework](https://img.shields.io/badge/Charm-Bubbletea%20%26%20Bubbles-cba6f7?style=for-the-badge)
![Theme](https://img.shields.io/badge/Theme-Catppuccin%20Mocha-fab387?style=for-the-badge)

---

## 🌟 Features

- 📜 **Timeless Generic Timeframe Presets**: Zero-maintenance preset selector for banner timeframes (`Current Banner Phase (+14d)`, `Next Phase (+21d)`, `Next Patch (+42d)`, `In 2 Patches (+84d)`, `Anniversary Event (+126d)`, or `Custom Manual Input`). Never becomes outdated with game version updates!
- 🎲 **Probability & Win Rate Calculator**: Real-time cumulative probability percentage calculation for target characters considering pity, soft pity ramp, hard pity, and 50/50 state.
- 📆 **Daily Income & Banner Countdown Projection**: Project total pull budget when banner ends (`Current Pulls + (Days * Daily Income / Rate)`) with clear readiness badges (`✅ Target Achievable` / `⚡ Short by X Pulls`).
- 📝 **Interactive Form Input**: Real-time form fields built with `bubbles/textinput` for editing Stellar Jades, Special Passes, Clear Drops, Unilogs, Current Pity, and Guaranteed 50/50 status.
- 📊 **Side-by-Side Dashboard**: High-level overview displaying total savings, total combined pulls, current pity states, and visual readiness indicators.
- 💾 **Real-Time JSON Persistence**: Auto-saves user savings & pity profile at `~/.config/gacha-planner/data.json` immediately as you type.
- 🎨 **Catppuccin Mocha Aesthetics**: Custom color palette designed with Lipgloss for dark & vibrant terminal aesthetics.

---

## 🛠️ Tech Stack

| Component | Technology | Description |
|---|---|---|
| **Language** | [Go](https://go.dev/) (1.27+) | High performance compiled language |
| **TUI Engine** | [Charm Bubbletea](https://github.com/charmbracelet/bubbletea) | Elm architecture for terminal UIs |
| **Form & Viewport** | [Charm Bubbles](https://github.com/charmbracelet/bubbles) | Interactive `textinput` & scrollable `viewport` |
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
| `←` / `→` or `Space` | Cycle **Timeframe Presets** (when focused on Timeframe Preset row) |
| `Type (0-9)` | Edit numeric values in real-time |
| `Space` or `g` | Toggle Guaranteed status (when focused on Guaranteed row) |
| `PgUp` / `PgDn` | Scroll viewport content |
| `Esc`, `Ctrl+C` or `q` | Auto-save profile & quit application |

---

## 📚 References & Resources

- **Reverse: 1999 CN Banner History & Schedule Spreadsheet**:
  [Prydwen Reverse: 1999 CN Banner History](https://docs.google.com/spreadsheets/u/1/d/e/2PACX-1vSZ_xS_5dfuvV7-yfitOqCw7e8VVkqtfWCAkcFpXLYTvx7XXR7lg2e1wjTLBZJJAleeEKm-f1Y7gNI_/pubhtml) — Use this community-maintained schedule reference to check exact estimated days and target banner releases.

---

## 💾 Storage & Configuration Path

User data is stored locally in JSON format:
```text
~/.config/gacha-planner/data.json
```

---

## 📄 License

Distributed under the MIT License.
