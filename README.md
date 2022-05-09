# JSONFind

![Build and Test](https://github.com/WTFox/jsonfind/workflows/Build%20and%20Test/badge.svg?branch=master&event=push)
[![Go Report Card](https://goreportcard.com/badge/github.com/wtfox/jsonfind)](https://goreportcard.com/report/github.com/wtfox/jsonfind)
[![GitHub stars](https://img.shields.io/github/stars/wtfox/jsonfind)](https://github.com/wtfox/jsonfind/stargazers)
[![GitHub license](https://img.shields.io/github/license/wtfox/jsonfind)](https://github.com/wtfox/jsonfind/blob/master/LICENSE)

A fast and lightweight utility to easily find paths to values in JSON files.

---

## Installation

```sh
brew install wtfox/tap/jsonfind
```

---

## Usage

![usage](./assets/usage.png)

```text
NAME:
   jf - JSONFind

USAGE:
   jf <valueToFind> <jsonFile>

VERSION:
   1.1.0

DESCRIPTION:
   Search a JSON file for a specified value and output full paths of each occurrence found

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --parent-paths, -p  Renders the parent paths only (default: false)
   --first-only, -f    Returns the first occurrence only (default: false)
   --use-regex, -r     Use pattern matching via regex expression (default: false)
   --help, -h          show help (default: false)
   --version, -v       print the version (default: false)

```

---

## Contributing

Bugs likely exist so issues and pull requests are welcomed!
