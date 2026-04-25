# simpsons

<p align="center">
  <img src="docs/img/bart.png" alt="simpsons" width="180">
</p>

<p align="center">
  <em>"I am so smart! S-M-R-T!" — Bart Simpson, reviewing his Claude Code spend</em>
</p>

<p align="center">
  <a href="https://github.com/rudrakshsisodia/simpsons/blob/main/LICENSE"><img src="https://img.shields.io/github/license/rudrakshsisodia/simpsons" alt="License"></a>
</p>

---

A terminal dashboard that reads your `~/.claude/` folder and shows you everything about your Claude Code sessions — what you spent, what tools ran, which projects ate your budget, and how productive you've been. Offline only. Nothing leaves your machine.

## Install

```sh
brew tap rudrakshsisodia/tap
brew install simpsons
```

Or with Go:

```sh
go install github.com/rudrakshsisodia/simpsons@latest
```

## Start

```sh
simpsons
```

## Tabs

Use `Tab` / `Shift+Tab` or `h` / `l` to move between tabs.

**Analysis** — the big picture. Daily usage bars, an hour-by-hour heatmap, cost trend for the last 30 days, tool call rankings, model breakdown, streaks, and personal bests like longest session and most productive day.

**Projects** — every repo you've used Claude in, sorted by last active. Shows session count, total estimated cost, and when you were last in it. Press `Enter` to drill into a project.

**Sessions** — every conversation, sorted newest first. Shows project, duration, token count, estimated cost, and tool call count. Press `Enter` for the full breakdown: transcript, timeline, file activity, subagent spawns, and per-session tool stats.

**Agents** — subagent usage rolled up across all sessions.

**Tools** — ranked tool call counts across everything, split by built-in and MCP server tools.

## Controls

| Key | What it does |
|-----|-------------|
| `j` / `↓` | Move down |
| `k` / `↑` | Move up |
| `h` / `←` | Go back / previous tab |
| `l` / `→` | Drill in / next tab |
| `Enter` | Open selected item |
| `Esc` | Back up one level |
| `/` | Filter current list |
| `y` | Copy `claude --resume <id>` to clipboard |
| `?` | Show all shortcuts |

## Sessions — export and import

Pack any session into a zip and share it, archive it, or open it on another machine.

| Key | Action |
|-----|--------|
| `e` | Zip the selected session |
| `E` | Zip everything currently visible |
| `i` | Import from a zip |

Zips contain the raw JSONL transcript plus a manifest. They're compatible with `claude --resume` so whoever receives one can pick up right where it left off.

## Cost tracking

Every session shows an estimated cost based on the model and token counts from the transcript. Totals roll up per project and across everything. The Analysis tab shows a daily cost graph for the last 30 days with this-week vs last-week comparison.

Prices are hardcoded against Anthropic's published rates and updated manually when they change.
