#Requires -Version 5.1
<#
.SYNOPSIS
  Creates Desktop\OpenNox-Playtest on this PC (e.g. killputer) and a desktop shortcut.

.DESCRIPTION
  Run once from an elevated or normal PowerShell window after you have built or copied
  OpenNox binaries. Detects a common GOG Nox install path when possible.

  Example:
    powershell -ExecutionPolicy Bypass -File .\Setup-PlaytestDesktop.ps1
    powershell -ExecutionPolicy Bypass -File .\Setup-PlaytestDesktop.ps1 -GogNoxPath "D:\Games\Nox"
#>
[CmdletBinding()]
param(
    [string]$GogNoxPath = "",
    [string]$OpenNoxBinPath = ""
)

$ErrorActionPreference = "Stop"

$Desktop = [Environment]::GetFolderPath("Desktop")
$PlaytestRoot = Join-Path $Desktop "OpenNox-Playtest"
$BinDir = Join-Path $PlaytestRoot "bin"
$ConfigPath = Join-Path $PlaytestRoot "opennox.yml"

function Find-GogNox {
    $candidates = @(
        "C:\GOG Games\Nox",
        "C:\Program Files (x86)\GOG Galaxy\Games\Nox",
        "D:\GOG Games\Nox",
        "D:\Games\Nox"
    )
    foreach ($p in $candidates) {
        if (Test-Path (Join-Path $p "gamedata.bin")) { return $p }
    }
    return $null
}

New-Item -ItemType Directory -Force -Path $PlaytestRoot, $BinDir | Out-Null

if (-not $GogNoxPath) {
    $GogNoxPath = Find-GogNox
}
if (-not $GogNoxPath -or -not (Test-Path (Join-Path $GogNoxPath "gamedata.bin"))) {
    Write-Warning "GOG Nox folder not found. Set -GogNoxPath to your install (must contain gamedata.bin)."
    $GogNoxPath = Read-Host "Enter full path to GOG Nox folder"
}
if (-not (Test-Path (Join-Path $GogNoxPath "gamedata.bin"))) {
    throw "Invalid Nox path: $GogNoxPath"
}

if (-not $OpenNoxBinPath) {
    $scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
    $repoBuild = Resolve-Path (Join-Path $scriptDir "..\..\..\build") -ErrorAction SilentlyContinue
    if ($repoBuild -and (Test-Path (Join-Path $repoBuild "opennox-hd.exe"))) {
        $OpenNoxBinPath = $repoBuild
    }
}

if ($OpenNoxBinPath -and (Test-Path $OpenNoxBinPath)) {
    Copy-Item -Force -Path (Join-Path $OpenNoxBinPath "opennox-hd.exe") -Destination $BinDir -ErrorAction SilentlyContinue
    Copy-Item -Force -Path (Join-Path $OpenNoxBinPath "opennox.exe") -Destination $BinDir -ErrorAction SilentlyContinue
    Copy-Item -Force -Path (Join-Path $OpenNoxBinPath "*.dll") -Destination $BinDir -ErrorAction SilentlyContinue
}

$gameDataRel = $GogNoxPath
try {
    $gameDataRel = [System.IO.Path]::GetRelativePath($PlaytestRoot, $GogNoxPath)
} catch {
    # PowerShell 5.1 fallback: keep absolute path
}

@"
# OpenNox playtest config (TravelerHacker fork)
game:
  data: "$($gameDataRel -replace '\\','/')"
video:
  size:
    width: 2560
    height: 1440
  gamma: 1.0
  filtering: true
  stretch: false
"@ | Set-Content -Encoding UTF8 -Path $ConfigPath

$launcher = Join-Path $PlaytestRoot "Launch-OpenNox-HD.ps1"
@'
$Root = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location $Root
$exe = Join-Path $Root "bin\opennox-hd.exe"
if (-not (Test-Path $exe)) {
    Write-Error "Missing bin\opennox-hd.exe. Build OpenNox (see docs/build-windows.md) and copy binaries into bin\, or re-run Setup-PlaytestDesktop.ps1 with -OpenNoxBinPath."
}
Start-Process -FilePath $exe -WorkingDirectory $Root -ArgumentList @("-config", (Join-Path $Root "opennox.yml"))
'@ | Set-Content -Encoding UTF8 -Path $launcher

$shortcutPath = Join-Path $Desktop "OpenNox Playtest.lnk"
$WshShell = New-Object -ComObject WScript.Shell
$Shortcut = $WshShell.CreateShortcut($shortcutPath)
$Shortcut.TargetPath = "powershell.exe"
$Shortcut.Arguments = "-NoProfile -ExecutionPolicy Bypass -File `"$launcher`""
$Shortcut.WorkingDirectory = $PlaytestRoot
$Shortcut.IconLocation = (Join-Path $BinDir "opennox-hd.exe")
$Shortcut.Description = "OpenNox HD playtest (TravelerHacker fork)"
$Shortcut.Save()

Write-Host "Playtest folder: $PlaytestRoot"
Write-Host "Desktop shortcut: $shortcutPath"
Write-Host "Game data: $GogNoxPath"
Write-Host "When binaries are ready, place opennox-hd.exe and DLLs in: $BinDir"
