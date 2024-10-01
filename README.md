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

You should download the [latest release](https://github.com/Baipyrus/ProxySwitcher/releases)
archive, **extract all contents** to a dedicated directory and then simply execute
the [`.\install.ps1`](./install.ps1) script within a Windows Powershell.

Alternatively, you could simply run the following command in a Windows PowerShell:

```powershell
# Using 'Invoke-RestMethod' and 'Invoke-Expression'
irm 'https://raw.githubusercontent.com/Baipyrus/ProxySwitcher/main/install.ps1' | iex
```

## Usage

Generally, the [Installation](#installation) step will install both the program,
its assets, and a shortcut for the Windows Startmenu for your current userprofile.
Additionally, the program will also be added to the user's `%PATH%` Variable and
will this be executable from within the command-line. However, you will either need
to navigate into the program directory (`C:\Users\[Username]\AppData\Local\Programs\ProxySwitcher\`)
manually or specify a path to any directory containing configuration files using
the flag `-c, --configs string   configurations path (default "configs/")`.
To keep it simple, it is still recommended to use the program in system tray or
directly via code. This latter option will be explained next:

- Clone the repository:

    ```powersell
    git clone https://github.com/Baipyrus/ProxySwitcher.git
    ```

### CLI Commands

- **set**: Enable all saved proxies including system proxy.

    ```powersell
    # .\ProxySwitcher.exe set
    go run . set
    ```

- **unset**: Disable all saved proxies including system proxy.

    ```powersell
    # .\ProxySwitcher.exe unset
    go run . unset
    ```

- **save**: Saves a new configuration to set a proxy for.

    ```powersell
    # .\ProxySwitcher.exe save
    go run . save
    ```

- **debug**: Prints all proxy configurations after generating corresponding commands.

    ```powersell
    # .\ProxySwitcher.exe debug
    go run . save
    ```

### Configuration

The programs for which the proxy settings should be managed are stored in [`configs/`](https://github.com/Baipyrus/ProxySwitcher/tree/main/configs),
wherein you can then create a JSON file per command group. These files can easily
be modified directly or generated through the `save` command. For examples,
please take a look at the [default config](https://github.com/Baipyrus/ProxySwitcher/tree/main/configs)
or at the following block:

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
