# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.1] - 2022-05-9

### Added

- Don't panic when invalid regex is supplied

## [1.1.0] - 2022-04-27

### Added

- Search values with Regex patterns with -r
- Now accepts input via stdin (pipes)

## [1.0.2] - 2021-05-22

### Added

- Now outputs a JQ-compatible path

## [1.0.1] - 2020-06-20

### Added

- Added this change log!
- Added Makefile to build for all systems
- Added docstrings to all exported types and functions
- Added License

### Changed

- jf now prints out help when invalid arguments are supplied
- jf is now able to consume JSON files of various structures
- Various project structure changes to be more consistent with standard Go packages

## [1.0.0] - 2020-06-18

### Added

- Initial release of jf
