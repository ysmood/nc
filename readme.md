# nc

A simple tool bind tcp and udp at the same time with sane defaults.

I use it to detect if a machine's tcp/udp is accessible.

Install executable: `curl -L https://git.io/fjaxx | repo=ysmood/nc bin=nc sh`

```txt
usage: nc [<flags>] <command> [<args> ...]

a simple netcat tool

Flags:
  --help     Show context-sensitive help (also try --help-long and --help-man).
  --version  Show application version.

Args:
  [<address>]  the host and port address

Commands:
  help [<command>...]
    Show help.


  serve [<address>]
    run as server for both tcp and udp on the same port


  send [<flags>] [<address>]
    send tcp or udp package

    -p, --protocol="tcp"  protocol to use
```
