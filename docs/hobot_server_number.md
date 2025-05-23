## hobot server number

Get server number

### Synopsis

Get server number (the identifier)

```
hobot server number [flags]
```

### Examples

```
hobot server number --ip 12.34.56.78
hobot server number --name marvin
hobot server number --name macos --ignore-case
hobot server number --name "Deep Thought"
hobot server number --name marvin --no-headers
hobot server number --name marvin -o json
hobot server number --name marvin -o yaml
```

### Options

```
  -h, --help            help for number
  -i, --ignore-case     Ignore case when matching (default is false)
      --ip string       Server IP (can be IPv4 or IPv6)
  -n, --name string     Server Name
      --no-headers      Do not print headers in the output
  -o, --output string   Output format. One of (table, json, yaml). (default "table")
```

### Options inherited from parent commands

```
      --config string    config file (default is "/Users/fabian/.config/hobot/config.ini")
      --context string   default is read from state file or "default"
      --state string     state file (default is "/Users/fabian/.local/state/hobot/state.ini")
```

### SEE ALSO

* [hobot server](hobot_server.md)	 - Manage servers

###### Auto generated by spf13/cobra on 14-May-2025
