# ProxySwitcher

ProxySwitcher is a Golang-based Windows application that allows you to easily switch
internet proxy settings for various programs including the system settings. The
application provides a command-line interface (CLI) through the usage of a library
called [cobra](https://github.com/spf13/cobra). This CLI then provides commands to
set, unset, and save proxy configurations, as well as a system tray icon for quick
access to settings and toggling proxies.

## Features

- Switch internet proxy settings for various applications and the system.
- Manage configurations through a `configs.json` file.
- Command-line interface with commands like `set`, `unset`, and `save`.
- System tray icon with options to enable and disable proxies, open settings,
save configurations, or exit the application.

## Installation

- Simply run the following command in a Windows PowerShell:

    ```powershell
    # Using 'Invoke-RestMethod' and 'Invoke-Expression'
    irm 'https://raw.githubusercontent.com/Baipyrus/ProxySwitcher/main/install.ps1'
     | iex
    ```

## Usage

Generally, the [Installation](#installation) step will install both the program,
its assets, and a shortcut for the Windows Startmenu for your current userprofile.
You *could* go into the program directory (`C:\Users\[Username]\AppData\Local\Programs\ProxySwitcher\`)
manually and run the program from your CLI in there, but it is recommended to simply
use the program in system tray or directly via code. This latter option will be explained
next:

- Clone the repository:

    ```powersell
    git clone https://github.com/Baipyrus/ProxySwitcher.git
    ```

### CLI Commands

- **set**: Enable all saved proxies including system proxy.

    ```powersell
    go run . set
    ```

- **unset**: Disable all saved proxies including system proxy.

    ```powersell
    go run . unset
    ```

- **save**: Saves a new configuration to set a proxy for.

    ```powersell
    go run . save
    ```

### Configuration

The programs for which the proxy settings should be managed are stored in a `configs.json`
file, which can be modified directly or through the `save` command. For examples,
please take a look at the [default config](./configs.json) or at the following block:

```js
[
    {
        "name": "test1", // Default CMD Name
        "cmd": "echo 'Hello, World!'", // Optional
        "set": [
            {
                "args": [
                    "http"
                ], // Optionally empty
                "type": "text" // Default; Optional
                "equator": " " // Default; Optional
            }, // Writes "echo 'Hello, World!' http ", followed
               // by your system proxy, to a powershell process.
            {
                "args": [
                    "https"
                ], // Optionally empty
                "equator": "=" // Optional
            }  // Writes "echo 'Hello, World!' https=", followed
               // by your system proxy, to a powershell process.
        ], // Optional
        "unset": []
        // "unset" has the contents as "set" above
        // and it is also optional.
    }, {
        "name": "test2", // Default CMD Name
        "cmd": "echo '$PRSW_ARG $PRSW_ARG'", // Optional
        "set": [
            {
                "args": [
                    "https"
                ], // Optionally empty
                "type": "variable" // Optional
            }  // Writes "echo 'https <Your System Proxy>'"
               // to a powershell process.
        ] // Optional
    }
]
```

### System Tray Icon

Right-click the system tray icon to:

- **Properties**: View and modify system proxy settings.
- **Enable Proxy**: Enable all saved proxies including system proxy.
- **Disable Proxy**: Disable all saved proxies including system proxy.
- **Save New Config**: Open a prompt to save a new proxy configuration.
- **Exit**: Close the application.

## Building

1. Clone the repository as seen above in [Usage](#usage)
2. Run the following command:

    ```powershell
    GOOS=windows GOARCH=amd64 go build -o build/ -v ./...
    ```
