$startupDir = (Get-Location).Path
$destinationDir = "$env:LOCALAPPDATA\Programs"
$startmenuDir = "$env:APPDATA\Microsoft\Windows\Start Menu\Programs"
$powershellPath = "$PSHOME\powershell.exe"

# Create program directory and relocate
Write-Host "Creating program directory in Userprofile..." -ForegroundColor Cyan
$programDir = "ProxySwitcher"
if ($startupDir -ne $destinationDir)
{
        Set-Location $destinationDir
        if (-not (Test-Path $programDir))
        { New-Item -ItemType Directory -Path $programDir | Out-Null
        }
}
Set-Location $programDir
$programPath = "$destinationDir\$programDir"

# Download functional files from github as-is
Write-Host "Downloading program into local directory..." -ForegroundColor Cyan
function DownloadFile
{
        param (
                [string]$url
        )

        $fileName = $url.Split("/")[-1]
        Invoke-RestMethod $url -OutFile $fileName
}
DownloadFile -url https://github.com/Baipyrus/ProxySwitcher/releases/latest/download/ProxySwitcher.exe
DownloadFile -url https://raw.githubusercontent.com/Baipyrus/ProxySwitcher/main/configs.json
DownloadFile -url https://raw.githubusercontent.com/Baipyrus/ProxySwitcher/main/run.ps1

# Create assets directory and relocate
Write-Host "Downloading assets into local directory..." -ForegroundColor Cyan
$assetsDir = "assets"
if (-not (Test-Path $assetsDir))
{ New-Item -ItemType Directory -Path $assetsDir | Out-Null
}
Set-Location $assetsDir
$assetPath = "$programPath\assets\ICON_Enabled.ico"

# Download asset files from github
DownloadFile -url https://raw.githubusercontent.com/Baipyrus/ProxySwitcher/main/assets/ICON_Disabled.ico
DownloadFile -url https://raw.githubusercontent.com/Baipyrus/ProxySwitcher/main/assets/ICON_Enabled.ico

# Create Startmenu Shortcut
Write-Host "Creating shortcuts for easy access..." -ForegroundColor Cyan
$assetsDir = "assets"
$shell = New-Object -comObject WScript.Shell
$shortcutPath = "$startmenuDir\Proxy Switcher.lnk"
$shortcut = $shell.CreateShortcut($shortcutPath)
$shortcut.TargetPath = $powershellPath
$shortcut.WorkingDirectory = $programPath
$shortcut.Arguments = "-ExecutionPolicy Bypass -NonInteractive -NoProfile -WindowStyle Hidden -File ""$programPath\run.ps1"""
$shortcut.IconLocation = $assetPath
$shortcut.WindowStyle = 7
$shortcut.Save()

# Copy shortcut to autostart
Copy-Item -Path $shortcutPath -Destination "$startmenuDir\Startup\" -Force

# Navigate back to starting position
Set-Location $startupDir
Write-Host "Windows setup complete!" -ForegroundColor Green
