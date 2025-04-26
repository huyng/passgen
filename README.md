# PassGen - Password Generator

PassGen is a simple command-line password generation tool. It provides various options to customize the generated password, such as length, inclusion of special characters, lowercasing.

It allows you to directly copy the password to the system clipboard.

This tool is designed to be **simple**, **auditable**, and with **zero external dependencies**.

## Features

- Generate random passwords
- Include special characters
- Generate passwords with only lowercase letters.
- Copy the generated password to the clipboard (supports macOS, Linux, and Windows).

## Installation

#### Method 1: Install to custom location on your $PATH

```bash
git clone https://github.com/huyng/passgen.git
cd passgen
go build
cp ./passgen $HOME/.local/bin/
```

#### Method 2: Install into $GOPATH/bin
```bash
go install github.com/huyng/passgen@latest
```

## Usage

```
Password Generator - Creates secure random passwords
Usage: passgen [options]

Options:
  -n N    Set password length (default: 10, minimum: 4)
  -c      Copy the generated password to the clipboard
  -s      Include special characters in the password
  -l      Use only lowercase letters in the password
  -h      Show this help message

By default, the password is printed to stdout.
```

### Examples

Generate a password of length 12:
```bash
./passgen -n 12
```

Generate a password with special characters:
```bash
./passgen -s
```

Generate a password and copy it to the clipboard:
```bash
./passgen -c
```

Generate a password with only lowercase letters:
```bash
./passgen -l
```

Show the help message:
```bash
./passgen -h
```

## Platform Support

- **macOS**: Uses `pbcopy` to copy passwords to the clipboard.
- **Linux**: Supports `xclip` or `xsel` for clipboard functionality.
- **Windows**: Uses the `clip` command for clipboard functionality.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the project.

## Author

Huy Nguyen