## hobot completion powershell

Generate the autocompletion script for powershell

### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	hobot completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
hobot completion powershell [flags]
```

### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --config string    config file (default is "/Users/fabian/.config/hobot/config.ini")
      --context string   default is read from state file or "default"
      --state string     state file (default is "/Users/fabian/.local/state/hobot/state.ini")
```

### SEE ALSO

* [hobot completion](hobot_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 14-May-2025
