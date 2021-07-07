# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.0.5] 2021-07-06
### Changed
- Ignore `q` param in Accept header sorting.

## [0.0.4] 2021-07-03
### Changed
- Return from Negotiate function mime type struct, instead of string with mime type. It's needed to get params, like charset.

## [0.0.3] 2021-07-02
### Added
- Added examples

## [0.0.2] 2021-06-20
### Added
- Added specific error types
- Added benchmarks

### Changed
- Changed README.md a little bit

## [0.0.1] 2021-06-18
### Added
- Mime typ and mime header types are implemented
- Accept header parser, sorter and matchers are implemented
