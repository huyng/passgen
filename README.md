# PassGen - Password Generator

PassGen is a simple command-line tool written in Go that generates secure random passwords. It provides various options to customize the generated password, such as length, inclusion of special characters, lowercasing.

It allows you to directly copy the password to the system clipboard.

This tool is designed to be **simple**, **auditable**, and with **zero external dependencies**.

## Features

- Generate random passwords
- Include special characters
- Generate passwords with only lowercase letters.
- Copy the generated password to the clipboard (supports macOS, Linux, and Windows).

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/huyng/passgen.git
   cd passgen
   ```

2. Build the binary:
   ```bash
   go build
   ```

3. Run the binary:
   ```bash
   ./passgen
   ```

3. Install it somewhere into your system path
   ```bash
   cp ./passgen /usr/local/bin/
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

### Options

- `-n N`
  Set password length (default: 10, minimum: 4).

- `-c`
  Copy the generated password to the clipboard.

- `-s`
  Include special characters in the password.

- `-l`
  Use only lowercase letters in the password.

- `-h`
  Show the help message.

### Examples

1. Generate a password of length 12:
   ```bash
   ./passgen -n 12
   ```

2. Generate a password with special characters:
   ```bash
   ./passgen -s
   ```

3. Generate a password and copy it to the clipboard:
   ```bash
   ./passgen -c
   ```

4. Generate a password with only lowercase letters:
   ```bash
   ./passgen -l
   ```

5. Show the help message:
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