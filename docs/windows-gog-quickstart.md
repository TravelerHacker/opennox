# Windows 11 + GOG Nox quickstart (TravelerHacker fork)

## Build `opennox-hd.exe`

1. Install [Go](https://go.dev/dl/), [MSYS2 MinGW32](https://www.msys2.org/), and deps from [build-windows.md](build-windows.md) (`mingw-w64-i686-SDL2`, `mingw-w64-i686-openal`).
2. Clone this fork, branch `dev` (or your PR branch).
3. In **MSYS2 MinGW 32-bit** shell:

```bash
cd /c/path/to/opennox/src
go run ./internal/noxbuild client-hd
```

Output: `opennox-hd.exe`, `opennox.exe`, plus required `.dll` from MSYS `mingw32/bin`.

## Point at GOG data

Either install OpenNox next to GOG files, or set in `opennox.yml`:

```yaml
game:
  data: "C:/GOG Games/Nox"
```

## Playtest desktop setup (killputer)

Run once in PowerShell from repo root:

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\windows\killputer\Setup-PlaytestDesktop.ps1
```

See [playtest-killputer-windows.md](playtest-killputer-windows.md).

## Recommended PR stack to merge first

1. Resolution/config (#1)  
2. CI (#3)  
3. Vendored libs / vsync (#4)  
4. Playtest scripts (#2)  
5. Input UI (#5), experimental video (#6)
