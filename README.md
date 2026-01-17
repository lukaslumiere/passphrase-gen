# passphrase-gen

Secure passphrase generator following the [Diceware](http://diceware.com/) method, using the EFF wordlist and cryptographically secure randomness.

## Installation

```bash
go install github.com/lukaslumiere/passphrase-gen@latest
```

> **Note:** If the command is not found after installation, run `hash -r` or open a new terminal to refresh your shell's command cache.

Or build from source:

```bash
git clone https://github.com/lukaslumiere/passphrase-gen.git
cd passphrase-gen
make build
```

## Usage

```bash
passphrase-gen [options]
```

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `-n`, `--num` | Number of words (1-20) | 6 |
| `-d`, `--delimiter` | Word separator: - _ . , : ; \| + or empty | - |
| `-c`, `--caps` | Capitalize first letter of each word | true |
| `--no-caps` | Disable capitalization | - |
| `-s`, `--specials` | Append special characters (0-5) | 0 |
| `--count` | Generate multiple passphrases | 1 |
| `-v`, `--version` | Show version | - |
| `-h`, `--help` | Show help | - |

### Examples

```bash
# Default: 6 capitalized words
passphrase-gen
# Output: Horse-Battery-Correct-Staple-Apple-Window

# 8 words with 2 special characters at the end
passphrase-gen -n 8 -s 2
# Output: Mountain-Ocean-Keyboard-Shoe-River-House-Light-Sun7=

# No capitalization, no delimiter (like classic Diceware)
passphrase-gen --no-caps -d ""
# Output: horsebatterycorrectstapleapplewindow

# Generate 5 passphrases
passphrase-gen --count 5

# Use underscore as delimiter
passphrase-gen -d _
# Output: Horse_Battery_Correct_Staple_Apple_Window
```

## Security

- Uses `crypto/rand` exclusively (never `math/rand`)
- 2500+ words from EFF Large Wordlist (~11.3 bits/word)
- Fisher-Yates shuffle with secure random
- Special characters from: `~!#$%^&*()-=+[]{}:;"'<>?/0123456789`

### Entropy

| Words | Entropy |
|-------|---------|
| 4 | ~45 bits |
| 5 | ~57 bits |
| 6 | ~68 bits |
| 7 | ~79 bits |
| 8 | ~90 bits |

## License

MIT
