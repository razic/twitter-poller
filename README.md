# twitter-poller

`twitter-poller` polls a list of urls, reporting an aggregate of success count
by application/version.

## Installation

`twitter-poller` supports the following platforms:

* Darwin (amd64)
* Linux (amd64)

There are pre-built executables ready for usage located in the `bin/` folder at
the root of this repository.

## Usage

```bash
$ ./bin/twitter-poller-linux --help
NAME:
   twitter-poller -
  _          _ _   _                         _ _
 | |___ __ _(_) |_| |_ ___ _ _ ___ _ __  ___| | |___ _ _
 |  _\ V  V / |  _|  _/ -_) '_|___| '_ \/ _ \ | / -_) '_|
  \__|\_/\_/|_|\__|\__\___|_|     | .__/\___/_|_\___|_|
                                  |_|

polls a list of urls, reporting an aggregate of success count by
application/version


USAGE:
   twitter-poller [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file value     file containing newline delimited list of urls to poll (default: "servers.txt")
   --pollers value  number of pollers to launch (default: 2)
   --help, -h       show help
   --version, -v    print the version
```

## Development

### Prerequisites

* Go 1.8

### Building from Source

```bash
make
```

### Running Tests

```bash
make test
```
