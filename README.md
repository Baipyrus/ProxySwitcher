# ProxySwitcher

ProxySwitcher is a Golang-based Windows application that allows you to easily switch
internet proxy settings for various programs including the system settings. The
application provides a command-line interface (CLI) through the usage of a library
called [cobra](https://github.com/spf13/cobra). This CLI then provides commands to
set, unset, and save proxy configurations, as well as a system tray icon for quick
access to settings and toggling proxies.

## Features

- Switch internet proxy settings for various applications and the system.
- Configuration files stored in `configs/` for better organization and flexibility.
- CLI commands (`set`, `unset`, `save`, `debug`) for managing proxy settings.
- System tray icon for fast access to all functionality.

## Installation

This project offers two installation methods: via PowerShell or manually. In the
end, you will always be using the [`.\install.ps1`](./install.ps1) script which will
automatically download the Project if necessary and adds it to `%PATH%`. All contents
will then be installed to `%LOCALAPPDATA%/Programs` for a user installation.

### PowerShell Installation

Run the following command in Windows PowerShell:

```powershell
irm 'https://raw.githubusercontent.com/Baipyrus/ProxySwitcher/main/install.ps1' | iex
```

### Manual Installation

1. Download the [latest release](https://github.com/Baipyrus/ProxySwitcher/releases).
2. Extract the archive to somewhere accessible to your user.
3. Run the [`.\install.ps1`](./install.ps1) script within a Windows Powershell instance.

## Usage

First, you will either need to navigate into the program directory (`C:\Users\[Username]\AppData\Local\Programs\ProxySwitcher\`)
manually or specify a path to any directory containing configuration files using
the flag `-c, --configs string   configurations path (default "configs/")`.
To keep it simple, it is still recommended to use the program in system tray or
directly via code.

In case you want to run the code directly:

- Clone the repository:

    ```powersell
    git clone https://github.com/Baipyrus/ProxySwitcher.git
    ```

### CLI Commands

- **set**: Enable all saved proxies including system proxy.

    ```powersell
    # ProxySwitcher.exe set
    go run . set
    ```

- **unset**: Disable all saved proxies including system proxy.

    ```powersell
    # ProxySwitcher.exe unset
    go run . unset
    ```

- **save**: Saves a new configuration to set a proxy for.

    ```powersell
    # ProxySwitcher.exe save
    go run . save
    ```

- **debug**: Prints all proxy configurations after generating corresponding commands.

    ```powersell
    # ProxySwitcher.exe debug
    go run . save
    ```

### Configuration

Proxy configurations are organized within the [`configs/`](https://github.com/Baipyrus/ProxySwitcher/tree/main/configs),
directory, with each JSON file representing a configuration for a specific command
group. You can modify or add your own configurations in any JSON file in this directory
directly, or you could use the `save` command in a CLI to save settings for you.
For examples configurations, please take a look at the [default config](https://github.com/Baipyrus/ProxySwitcher/tree/main/configs).

#### Structure

Any of these JSON files are used to build commands that the program can run to `set`
and `unset` the proxy configuration of other programs. The system proxy will be set
automatically upon calling the respective commands. The actual proxy address and
port will be automatically detected from your system settings, if available.

Once in one of these JSON files, you will need to create an object (`{}`) with the
following properties:

```javascript
{
    // Either "name" or "cmd" or both is required:
    "name": "npm",
    // Use a custom '$PRSW_ARG' variable to inject 
    // arguments into "cmd" at given positions:
    // "cmd": "npm",

    "set": [
        {
            "args": [
                "config",
                "set",
                "proxy"
            ],
            // Optionally specify a separator between the
            // last argument and the injected proxy string:
            "equator": "=", // Default: " "
            // Optionally specify a surrounding character
            // for injected proxy string:
            "surround": "\"",
            // Optionally specify the type of command. If
            // using '$PRSW_ARG' in "cmd", set to "variable".
            "type": "variable", // Default: "text"
            // Optionally choose to skip injecting the
            // proxy string for pre-config commands.
            "discard": "true" // Default: "false"
        }
    ],

    // Use the same structure as "set":
    "unset" []
}
```

### System Tray Icon

The system tray icon provides a context menu for all commands:

- **Properties**: View and modify system proxy settings.
- **Enable Proxy**: Enable all saved proxies including system proxy.
- **Disable Proxy**: Disable all saved proxies including system proxy.
- **Save New Config**: Open a prompt to save a new proxy configuration.
- **Exit**: Close the application.

## Building

1. Clone the repository as seen above in [Usage](#usage)
2. If running on Windows Subsystem for Linux:
    - Set environment variables to build for windows:

    ```bash
    GOOS=windows
    GOARCH=amd64 
    ```

3. Run the following command:

    ```powershell
    # Remove-Item -Recurse -Force -ErrorAction SilentlyContinue build/
    go build -o build/ -v ./...
    ```
