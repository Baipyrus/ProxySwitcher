# Load required assemblies for Windows Forms
Add-Type -AssemblyName System.Windows.Forms
Add-Type -AssemblyName System.Drawing

# Create a new NotifyIcon
$trayIcon = New-Object System.Windows.Forms.NotifyIcon
# Use a default application icon
$trayIcon.Icon = [System.Drawing.SystemIcons]::Application
$trayIcon.Text = 'ProxySwitcher'
$trayIcon.Visible = $true

# Create the context menu
$contextMenu = New-Object System.Windows.Forms.ContextMenuStrip

# Add "Exit" menu item
$ctxExit = New-Object System.Windows.Forms.ToolStripMenuItem
$ctxExit.Text = 'Exit'
$ctxExit.Add_Click({
        # Gracefully exit the application
        [System.Windows.Forms.Application]::Exit()
    })

# Add menu items to the context menu
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
