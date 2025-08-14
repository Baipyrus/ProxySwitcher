# ProxySwitcher

ProxySwitcher is a Golang-based application that allows you to easily switch
internet proxy settings for various programs including the Windows system settings.
The application provides a command-line interface (CLI) through the usage of a library
called [cobra](https://github.com/spf13/cobra). This CLI then provides commands to
set and unset both your own as well as the default proxy configurations, or edit
them in a simple GUI using the wonderful Framework at [AllenDang/giu](https://github.com/AllenDang/giu).
Additionally for Windows Users, it provides a system tray icon for quick
access to settings and toggling proxies.

## Features

- Switch internet proxy settings for various applications and the windows
system settings.
- Configuration files stored in [`configs/`](./configs/) for better
organization and flexibility.
- CLI commands (`set`, `unset`, `config`, `debug`) for managing proxy settings.
- Windows system tray icon for fast access to most functionality.

## Installation

It is recommended to manually download all necessary files from the [Release Pages](https://github.com/Baipyrus/ProxySwitcher/releases).
The [`.\install.ps1`](./install.ps1) script will then install the Project and add
it to the users' `%PATH%`. To achieve this, all contents will be installed to
`%LOCALAPPDATA%/Programs` for a user installation.

> [!NOTE]
> The aforementioned [`.\install.ps1`](./install.ps1) script is recommended for
Windows users only. Other platforms will have to use the CLI only.

## Usage

First, you will either need to navigate into the program directory (`C:\Users\[Username]\AppData\Local\Programs\ProxySwitcher\`)
manually or specify a path to any directory containing configuration files using
the flag `-c, --configs string   configurations path (default "configs/")`.
To keep it simple, it is still recommended to use the program in system tray or
directly via code.

In case you want to run the code directly:

- Clone the repository:

    ```powershell
    git clone https://github.com/Baipyrus/ProxySwitcher.git
    ```

- Set environment variables to build for windows:

    ```bash
    CGO_ENABLED=1
    GOOS=windows
    GOARCH=amd64 
    ```

### CLI Commands

- **set**: Enable all saved proxies including windows system proxy.

    ```powershell
    # ProxySwitcher.exe set
    go run . set
    ```

- **unset**: Disable all saved proxies including windows system proxy.

    ```powershell
    # ProxySwitcher.exe unset
    go run . unset
    ```

- **config**: Opens a new GUI Window to configure all available ProxySwitcher settings.

    ```powershell
    # ProxySwitcher.exe config
    go run . config
    ```

- **debug**: Prints all proxy configurations after generating corresponding commands.

    ```powershell
    # ProxySwitcher.exe debug
    go run . save
    ```

### Configuration

Proxy configurations are organized within the [`configs/`](./configs/),
directory, with each JSON file representing a configuration for a specific command
group. You can modify or add your own configurations in any JSON file in this directory
directly, or you could use the `config` command in a CLI or the Windows system tray
to guide you through the settings.
For example configurations, please take a look at the [default config](./configs/).

#### Structure

All of the JSON files provided in the [`configs/`](./configs/) are used to build
commands that the program can run to set and unset the proxy configuration of
other programs. The Windows system proxy will be set automatically upon calling
the respective commands. The actual proxy host, protocol and port will be saved
in a special [`configs/proxy.ini`](./configs/proxy.ini) file. This, you can also
just edit on your own or via the `config` command.

Once in one of these config JSON files, you will need to create an object
(`{}`) with the following properties:

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

### Windows System Tray Icon

The system tray icon provides a context menu for all commands:

- **Configuration**: View and modify proxy configurations.
- **Enable Proxy**: Enable all saved proxies including system proxy.
- **Disable Proxy**: Disable all saved proxies including system proxy.
- **Exit**: Close the application.

## Building

1. Clone the repository as seen above in [Usage](#usage)
2. Use the given [Dockerfile](./Dockerfile) in isolation:

    ```bash
    docker build --target windows -t proxyswitcher:latest .
    docker run --rm -it -v ./build:/app/build proxyswitcher:latest
    ```

3. Or in case you want to do it manually:
    - If running on Windows Subsystem for Linux, set environment
    variables to build for windows:

    ```bash
    CGO_ENABLED=1
    GOOS=windows
    GOARCH=amd64 
    ```

4. Run the following `go build` command:

    ```powershell
    # Remove-Item -Recurse -Force -ErrorAction SilentlyContinue build/
    go build -v -o build/
    ```
