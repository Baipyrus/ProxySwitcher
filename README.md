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

1. Clone the repository:

    ```powersell
    git clone https://github.com/Baipyrus/ProxySwitcher.git
    ```

## Usage

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
