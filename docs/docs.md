# gfc

## Usage
> This cli template shows the date and time in the terminal

gfc

## Description

```
This is a template CLI application, which can be used as a boilerplate for awesome CLI tools written in Go.
This template prints the date or time to the terminal.
```
## Examples

```bash
cli-template date
cli-template date --format 20060102
cli-template time
cli-template time --live
```

## Flags
|Flag|Usage|
|----|-----|
|`--c string`|to clean the local flutter cache|
|`--debug`|enable debug messages|
|`--disable-update-checks`|disables update checks|
|`--raw`|print unstyled raw output (set it if output is written to a file)|

## Commands
|Command|Usage|
|-------|-----|
|`gfc completion`|generate the autocompletion script for the specified shell|
|`gfc help`|Help about any command|
# ... completion
`gfc completion`

## Usage
> generate the autocompletion script for the specified shell

gfc completion

## Description

```

Generate the autocompletion script for gfc for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|`gfc completion bash`|generate the autocompletion script for bash|
|`gfc completion fish`|generate the autocompletion script for fish|
|`gfc completion powershell`|generate the autocompletion script for powershell|
|`gfc completion zsh`|generate the autocompletion script for zsh|
# ... completion bash
`gfc completion bash`

## Usage
> generate the autocompletion script for bash

gfc completion bash

## Description

```

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:
$ source <(gfc completion bash)

To load completions for every new session, execute once:
Linux:
  $ gfc completion bash > /etc/bash_completion.d/gfc
MacOS:
  $ gfc completion bash > /usr/local/etc/bash_completion.d/gfc

You will need to start a new shell for this setup to take effect.
  
```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
`gfc completion fish`

## Usage
> generate the autocompletion script for fish

gfc completion fish

## Description

```

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:
$ gfc completion fish | source

To load completions for every new session, execute once:
$ gfc completion fish > ~/.config/fish/completions/gfc.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
`gfc completion powershell`

## Usage
> generate the autocompletion script for powershell

gfc completion powershell

## Description

```

Generate the autocompletion script for powershell.

To load completions in your current shell session:
PS C:\> gfc completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
`gfc completion zsh`

## Usage
> generate the autocompletion script for zsh

gfc completion zsh

## Description

```

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for every new session, execute once:
# Linux:
$ gfc completion zsh > "${fpath[1]}/_gfc"
# macOS:
$ gfc completion zsh > /usr/local/share/zsh/site-functions/_gfc

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... help
`gfc help`

## Usage
> Help about any command

gfc help [command]

## Description

```
Help provides help for any command in the application.
Simply type gfc help [path to command] for full details.
```


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 20 January 2022**
