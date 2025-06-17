# Get inetcpl proxy settings
$settings = Get-ItemProperty -Path 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Internet Settings';

# Load required assemblies for Windows Forms
Add-Type -AssemblyName System.Windows.Forms
Add-Type -AssemblyName System.Drawing

# Create a new NotifyIcon
$trayIcon = New-Object System.Windows.Forms.NotifyIcon
# Use a default application icon
$trayIcon.Icon = if ($settings.proxyEnable)
{ "./assets/ICON_Enabled.ico" 
} else
{ "./assets/ICON_Disabled.ico" 
}
$trayIcon.Text = 'ProxySwitcher'
$trayIcon.Visible = $true

# Create the context menu
$contextMenu = New-Object System.Windows.Forms.ContextMenuStrip

# Add "Open Internet Options" menu item
$ctxProperties = New-Object System.Windows.Forms.ToolStripMenuItem
$ctxProperties.Text = 'Configuration'
$ctxProperties.Add_Click({
        Start-Process powershell.exe -WindowStyle Hidden -ArgumentList '-Command', '.\ProxySwitcher.exe config'
    })

# Add "Enable Proxy Setting" menu item
$ctxEnable = New-Object System.Windows.Forms.ToolStripMenuItem
$ctxEnable.Text = 'Enable Proxy'
$ctxEnable.Add_Click({
        $trayIcon.Icon = './assets/ICON_Enabled.ico'
        Start-Process powershell.exe -WindowStyle Hidden -ArgumentList '-Command', '.\ProxySwitcher.exe set'
    })

# Add "Disable Proxy Setting" menu item
$ctxDisable = New-Object System.Windows.Forms.ToolStripMenuItem
$ctxDisable.Text = 'Disable Proxy'
$ctxDisable.Add_Click({
        $trayIcon.Icon = './assets/ICON_Disabled.ico'
        Start-Process powershell.exe -WindowStyle Hidden -ArgumentList '-Command', '.\ProxySwitcher.exe unset'
    })

# Add "Exit" menu item
$ctxExit = New-Object System.Windows.Forms.ToolStripMenuItem
$ctxExit.Text = 'Exit'
$ctxExit.Add_Click({
        # Gracefully exit the application
        [System.Windows.Forms.Application]::Exit()
    })

# Add menu items to the context menu
$contextMenu.Items.Add($ctxProperties)
$contextMenu.Items.Add($ctxEnable)
$contextMenu.Items.Add($ctxDisable)
$contextMenu.Items.Add($ctxSave)
$contextMenu.Items.Add($ctxExit)

# Assign the context menu to the tray icon
$trayIcon.ContextMenuStrip = $contextMenu

# Start a timer to keep the script running
$timer = New-Object System.Windows.Forms.Timer
$timer.Interval = 1000 # 1 second interval to keep the script alive
$timer.Add_Tick({})    # Empty event handler to keep the timer running
$timer.Start()

# Keep the application running until the exit action is triggered
[System.Windows.Forms.Application]::Run()

# Clean up the tray icon when the application is closed
$trayIcon.Dispose()
