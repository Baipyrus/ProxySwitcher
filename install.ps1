$startupDir = (Get-Location).Path
$startmenuDir = "$env:APPDATA\Microsoft\Windows\Start Menu\Programs"

# Create Local Program Directory if not exists
$destinationDir = "$env:LOCALAPPDATA\Programs"
if (-not (Test-Path $destinationDir))
{ New-Item -ItemType Directory -Path $destinationDir | Out-Null
}

# Detect Powershell Version and Path
Write-Host "Detecting Powershell Executable..." -ForegroundColor Cyan
$powershellPath = "$PSHOME"
if (Test-Path "$powershellPath\pwsh.exe")
{ $powershellPath = "$powershellPath\pwsh.exe"
} else
{ $powershellPath = "$powershellPath\powershell.exe"
}

# Detect if current dir is release asset
$isRelease = Test-Path ".\ProxySwitcher.exe"
$releaseDir = (Get-Location).Path

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

if ($isRelease)
{
        # Copy release assets to program dir
        Write-Host "Copying program into local directory..." -ForegroundColor Cyan
        Copy-Item -Path "$releaseDir\*" -Destination $programPath -Recurse -Force
} else
{
        # Specify temporary output Path
        $tmpPRSWzip = "$env:TMP\ProxySwitcher.zip"

        # Download release files from github as-is
        Write-Host "Downloading program into local directory..." -ForegroundColor Cyan
        $releases = Invoke-RestMethod "https://api.github.com/repos/Baipyrus/ProxySwitcher/releases/latest"
        $assets = $releases[0].assets
        $archive = $assets | Where-Object name -match "zip"
        Invoke-RestMethod $archive.browser_download_url -OutFile $tmpPRSWzip

        # Expand Archive to program directory
        Expand-Archive $tmpPRSWzip -DestinationPath $programPath -Force
}

# Unblock any files after download, if necessary
Get-ChildItem $programPath | Unblock-File

# Add program to PATH for cli application
$userpath = [System.Environment]::GetEnvironmentVariable("PATH", "User")
if ($userpath.Split(";") -notcontains $programPath)
{
        $userpath = $userpath + ";$programPath"
        [System.Environment]::SetEnvironmentVariable("PATH", $userpath, "User")
}

# Create Startmenu Shortcut
Write-Host "Creating shortcuts for easy access..." -ForegroundColor Cyan
$shell = New-Object -comObject WScript.Shell
$shortcutPath = "$startmenuDir\Proxy Switcher.lnk"
$shortcut = $shell.CreateShortcut($shortcutPath)
$shortcut.TargetPath = $powershellPath
$shortcut.WorkingDirectory = $programPath
$shortcut.Arguments = "-ExecutionPolicy Bypass -NonInteractive -NoProfile -WindowStyle Hidden -File ""$programPath\run.ps1"""
$shortcut.IconLocation = "$programPath\assets\ICON_Enabled.ico"
$shortcut.WindowStyle = 7
$shortcut.Save()

# Copy shortcut to autostart
Copy-Item -Path $shortcutPath -Destination "$startmenuDir\Startup\" -Force

# Navigate back to starting position
Set-Location $startupDir
Write-Host "Windows setup complete!" -ForegroundColor Green
