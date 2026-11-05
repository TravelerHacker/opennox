# Playtesting on Windows (killputer / GOG)

Cloud agents cannot write to your home PC desktop directly. Use the script below **once on killputer** (or any Windows 11 gaming PC) after you clone or pull this fork.

## One-time setup

1. Install build deps if you will compile locally: [build-windows.md](build-windows.md) (Go, MSYS2 MinGW, SDL2, OpenAL).
2. Open **PowerShell** (not necessarily admin).
3. From the repo root:

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\windows\killputer\Setup-PlaytestDesktop.ps1
```

Optional arguments:

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\windows\killputer\Setup-PlaytestDesktop.ps1 `
  -GogNoxPath "C:\GOG Games\Nox" `
  -OpenNoxBinPath "C:\Dev\opennox\build"
```

This creates:

| Path | Purpose |
|------|---------|
| `%USERPROFILE%\Desktop\OpenNox-Playtest\` | Playtest root, config, launcher |
| `%USERPROFILE%\Desktop\OpenNox-Playtest\bin\` | Put `opennox-hd.exe` and MinGW DLLs here |
| `%USERPROFILE%\Desktop\OpenNox Playtest.lnk` | Desktop shortcut to launch when ready |

Default config targets **2560×1440** with filtering on and stretch off (good starting point for an RTX 4070 Ti Super). Edit `opennox.yml` in the playtest folder for ultrawide or 4K.

## Build `opennox-hd.exe`

From `src\` in MSYS2 MinGW32 shell:

```bash
go run ./internal/noxbuild client-hd
```

Copy from `src\` (or `build\` after Docker/Makefile):

- `opennox-hd.exe`, required `.dll` files (see [build-windows.md](build-windows.md))

## Launch

Double-click **OpenNox Playtest** on the desktop, or run `Launch-OpenNox-HD.ps1` inside `OpenNox-Playtest`.

## What to verify

- Options menu no longer resets custom resolution (PR #1 on this fork).
- If keyboard input dies after a cutscene, run console `unstick` or enable `extensions.input.auto_recover_stuck` (PR #7).
- Campaign load, multiplayer join, and chat with non-ASCII characters if you use localized strings.
- Compare FPS with in-game perf overlay; expect CPU-bound rendering at high resolutions.
