# random

A simple Go CLI tool to generate a random integer within a specified range.

## Features

- Specify an upper and lower bound
- Support for shorthand flags
- Error handling for invalid ranges

## Usage

```bash
go run main.go --lower 10 --upper 100
```

Or using shorthand:

```bash
go run main.go -l 10 -u 100
```

## Output

Prints a random number between the given lower (inclusive) and upper (exclusive) bounds.

## License

This project is licensed under the BSD 3-Clause License. See the [LICENSE](./LICENSE) file for details.

## Installation

To install the tool, run the following command:

```bash
go install github.com/TWolfis/random@latest
```
