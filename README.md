# syncing-notifier

[![Build Status](https://travis-ci.org/pavel-kiselyov/syncing-notifier.svg?branch=master)](https://travis-ci.org/pavel-kiselyov/syncing-notifier) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Sends Slack incoming webhook about Geth node syncing status.

<p align="center">
    <img src="https://i.imgur.com/N6dKA0C.png" width="50%" height="50%" alt="Screenshot" title="Example usage">
</p>

# Building

    $ git clone github.com/pavel-kiselyov/syncing-notifier
    $ cd syncing-notifier && make build

# Usage

```
NAME:
   syncing-notifier - Sends Slack incoming webhook about Geth node syncing status

USAGE:
   syncing-notifier [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --oneshot            send a single notification and quit
   --interval value     notifications interval (ms) (default: 60000)
   --webhook-url value  Slack incoming webhook URL
   --nodes value        Ethereum node RPC entrypoints
   --help, -h           show help
   --version, -v        print the version
```
