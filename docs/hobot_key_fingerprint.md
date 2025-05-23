## hobot key fingerprint

Get fingerprint of SSH key

### Synopsis

Get fingerprint of SSH key

```
hobot key fingerprint [NAME] [flags]
```

### Examples

```
hobot key fingerprint topfstedt
hobot key fingerprint topfstedt --no-headers
hobot key fingerprint topfstedt -o json
hobot key fingerprint topfstedt -o yaml
```

### Options

```
      --comment string      Comment to add to the key (optional)
  -d, --directory string    Output directory where the keys are stored (default is current working directory)
  -h, --help                help for fingerprint
      --no-headers          Do not print headers in the output
  -o, --output string       Output format (table, json, yaml). (default "table")
      --passphrase string   Passphrase to encrypt the key with (optional)
```

### Options inherited from parent commands

```
      --config string    config file (default is "/Users/fabian/.config/hobot/config.ini")
      --context string   default is read from state file or "default"
      --state string     state file (default is "/Users/fabian/.local/state/hobot/state.ini")
```

### SEE ALSO

* [hobot key](hobot_key.md)	 - Manage SSH keys

###### Auto generated by spf13/cobra on 14-May-2025
