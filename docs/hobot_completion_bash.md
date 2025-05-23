## hobot completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(hobot completion bash)

To load completions for every new session, execute once:

#### Linux:

	hobot completion bash > /etc/bash_completion.d/hobot

#### macOS:

	hobot completion bash > $(brew --prefix)/etc/bash_completion.d/hobot

You will need to start a new shell for this setup to take effect.


```
hobot completion bash
```

### Options

```
  -h, --help              help for bash
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
