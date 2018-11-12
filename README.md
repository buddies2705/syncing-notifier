# syncing-notifier

[![Build Status](https://travis-ci.org/pavel-kiselyov/syncing-notifier.svg?branch=master)](https://travis-ci.org/pavel-kiselyov/syncing-notifier) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Sends Slack incoming webhook about Geth node syncing status.

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
   --interval value     notifications interval (ms) (default: 60000)
   --webhook-url value  Slack incoming webhook URL
   --nodes value        Ethereum node RPC entrypoints
   --help, -h           show help
   --version, -v        print the version
```
