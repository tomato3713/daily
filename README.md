# Daily Command

*This project is under constructing*

This command is for daily report.

## Usage

### report
This subcommand is for writing daily report at today.
So, make today daily report file if it is not exists, initialized by template file.
This default file name is "yyyy-mm-dd-daily-report.md".

```sh
daily report
```

### serve
Serve daily report directory.

```sh
daily serve
```

### config
There are sub command for configuration.

```
# make new cofiguration file.
daily config create

# show config file
daily config
```

and this configuration file format is under.

```toml
# File: $HOME/.config/daily/daily.toml
reportDir = "path to daily report file directory"
fileName = "daily-report"
templateFile = "path to template file"
pluginDir = "path to plugin directory"

[Serve]
templateBodyFile = "path to index.html"
assetsDir = "path to assets directory"
```

### Thanks

- https://github.com/spf13/cobra
- https://github.com/spf13/viper