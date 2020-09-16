# Daily Command

*This project is under constructing*

This command is for daily report.

# Installation

1. Run command: `go get github.com/tomato3713/daily`.
2. Run Command: `daily config create`.
3. Set your favorite text editor to `Editor` in the configuration file.
4. start writing your daily report. Run Command: `daily report`. If occured any error, please open issues.

## Usage

### report
This subcommand is for writing daily report at today.
So, make today daily report file if it is not exists, initialized by template file.
This file name is "yyyy-mm-dd-daily-report.md".

```sh
daily report
```

### serve
*Not implemented*

Serve daily report directory.

```sh
daily serve
```

### config
*no test code*

There are sub command for configuration.

```sh
# make new cofiguration file in $HOME/.config/daily/daily.toml if not exists.
daily config create

# show config file
daily config
```

and this configuration file format is under.

```toml
# File: $HOME/.config/daily/daily.toml
reportDir = "path to daily report file directory"
templateFile = "path to template file"
pluginDir = "path to plugin directory"
Editor = "code"

[Serve]
templateBodyFile = "path to index.html"
assetsDir = "path to assets directory"
```

### list
*Partly implemented*

Show daily report file path list and part of contents.
```sh
daily list
```

### cat
*partly implemented*

Prints daily report you want.
This command find in `reportDir` in configuration file.

```sh
daily cat

# print daily report content of yyyy/mm/dd.
daily cat yyyy mm dd
```

### Global Option

- `---config`
Set the path to configuration file.
Example:
```
daily config --config .
```

### Thanks

- https://github.com/spf13/cobra
- https://github.com/spf13/viper
