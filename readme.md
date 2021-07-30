# net Description

```
Manages network connections
aims to be stateful and highly user focused
wraps iwd & iwctl under the hood

Usage:
  net [flags]
  net [command]

Available Commands:
  config      command set for viewing and editing program configuration
  ctrl        simply invokes iwctl
  doctor      Checks various driver, interfaces, routes (items necessary for network connection)
  help        Help about any command
  list        scan for and list available networks
  test        pings google 2x

Flags:
      --completion string   generate shell completion
      --config string       config file (default is $HOME/.config/net.yaml)
  -h, --help                help for net

Use "net [command] --help" for more information about a command.
```

# Installation
* net in under construction and is not yet stable

### prerequisites
* you must have go installed and configured as a prerequesite
* you must install and configure [iwd](https://wiki.archlinux.org/title/Iwd)

### install
* currently it is only installable from source

```
cd "${GOPATH}/src"
git clone git@github.com:mrgarelli/net.git
cd net
make build
net --completion <shell> # zsh or bash or ...
# restart terminal
```

# Configuration
* all configuration can be handeled with:
	> net config
* config file should be located at ```~/.config/net.yaml```
* [initial config file](./.config/net.yaml)
