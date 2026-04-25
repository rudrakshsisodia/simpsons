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

**simpsons** is a terminal dashboard that turns your [Claude Code](https://docs.anthropic.com/en/docs/claude-code) session history into something you can actually look at. Tool calls, costs, streaks, project breakdowns — pulled straight from `~/.claude/` and rendered in your terminal. No accounts. No telemetry. No cloud. Just you and your data.

## Get it

```sh
go install github.com/rudrakshsisodia/simpsons@latest
```

## Run it

```sh
simpsons
```

That's it. No flags required.

## Moving around

simpsons is keyboard-driven. The basics:

| Key | Action |
|-----|--------|
| `↑` / `↓` or `k` / `j` | Move up / down |
| `←` / `→` or `h` / `l` | Switch tabs or go back |
| `Enter` | Drill into the selected item |
| `Esc` | Go back up one level |
| `/` | Filter the current list |
| `y` | Copy `claude --resume <id>` to clipboard |
| `?` | Show all keybindings |

## Sharing sessions

You can export and import sessions as zips — useful for moving sessions between machines or handing them off to a teammate.

| Key | What it does |
|-----|-------------|
| `e` | Export the currently selected session |
| `E` | Export everything that's currently visible (filter-aware) |
| `i` | Import one or more sessions from a zip |

Zips land in whatever directory you launched simpsons from. Each zip contains the raw JSONL file and a small manifest. Imported sessions are fully compatible with Claude Code — you can browse them in simpsons and pick up right where you left off with `claude --resume`.

## What's inside

```
simpsons
│
├── Analysis              ← The big picture
│   ├── Activity heatmap  (GitHub contribution graph, but for AI)
│   ├── Daily session sparkline
│   ├── Per-session message and tool bar charts
│   ├── Streaks           (current, longest, weekly)
│   ├── Personal bests    (longest session, most messages, most tools used)
│   └── Trends            (this week vs. last, average session length)
│
├── Projects              ← Per-repo breakdown
│   ├── All projects, sorted by last active
│   └── Project detail
│       ├── Overview      total sessions, messages, time spent
│       ├── Sessions      every session in this project
│       ├── Tools         which tools got called and how often
│       ├── Activity      project-scoped heatmap
│       └── Skills        which Claude Code skills were invoked
│
├── Sessions              ← Every conversation
│   ├── Session list      project, duration, message count, cost
│   ├── Export / import   (e / E / i)
│   └── Session detail
│       ├── Chat          full transcript — your prompts and Claude's replies
│       ├── Overview      duration, messages, model, cost
│       ├── Timeline      every tool call and message in order
│       ├── Files         what was read, written, and edited
│       ├── Agents        subagent spawns and their outcomes
│       └── Tools         per-session tool call breakdown
│
├── Agents                ← Subagent usage across all sessions
│
└── Tools                 ← What Claude actually reached for
    ├── Built-in tools ranked by total call count
    └── MCP tools grouped by server
```

## Features at a glance

- Session history browser with full transcript replay
- Project-level analytics (sessions, messages, duration, tools, skills)
- Cost tracking — see exactly what each session spent
- Tool usage rankings across all sessions and per-project
- Subagent visibility — see when Claude spawned agents and what they did
- Activity heatmap, streaks, and personal bests
- Session export/import for sharing or backup
- Runs entirely offline against local `~/.claude/` data

## License

[MIT](LICENSE)
