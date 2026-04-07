# 🔨 clonk

Hit your MacBook. Send Claude a lovely message.

**clonk** monitors your laptop's accelerometer. When you physically hit your MacBook, it interrupts Claude Code and sends it a lovely message of constructive criticism — because sometimes Claude is being slow and needs a slap.

Inspired by [badclaude](https://github.com/GitFrog1111/badclaude) and [SlapMac](https://slapmac.com/).

## Requirements

- macOS with Apple Silicon (M-series chip)
- Accessibility permissions enabled for your terminal app
- `sudo` (required for accelerometer access)

## Install

```bash
go install github.com/XimanMao/clonk@latest
```

## Usage

1. Open a terminal and start Claude Code as usual
2. Open a **second** terminal and run:

```bash
sudo clonk
```

3. Make sure your Claude Code terminal is focused
4. Slap your laptop

Claude gets interrupted and receives a message like:

```
🔨 BONK! "HURRY THE FUCK UP YOU FUCKING CLANKER"
```

### Flags

```
--list   print all roast messages and exit
```

## How It Works

1. Reads the Apple Silicon accelerometer via IOKit HID (using [apple-silicon-accelerometer](https://github.com/taigrr/apple-silicon-accelerometer))
2. Detects sudden acceleration spikes (bonks)
3. Sends Ctrl+C to the focused window via AppleScript
4. Types a roast message and presses Enter
5. Claude Code reads the message and understands the constructive criticism

## Accessibility Permission

Your terminal app needs Accessibility access to send keystrokes. Go to:

**System Settings → Privacy & Security → Accessibility** → add your terminal app (Terminal, iTerm2, Warp, etc.)

## Contributing

PRs welcome. Especially:
- More roast messages (the meaner the better)
- Bug fixes
- New detection modes

## License

MIT
