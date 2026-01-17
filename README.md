# passphrase-gen

Secure passphrase generator following the [Diceware](http://diceware.com/) method, using a Portuguese wordlist and cryptographically secure randomness.

## Installation

```bash
go install github.com/lukaslumiere/passphrase-gen@latest
```

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
# Output: Cavalo-Bateria-Correto-Grampo-Abacate-Janela

# 8 words with 2 special characters at the end
passphrase-gen -n 8 -s 2
# Output: Montanha-Oceano-Teclado-Sapato-Rio-Casa-Luz-Sol7=

# No capitalization, no delimiter (like classic Diceware)
passphrase-gen --no-caps -d ""
# Output: cavalobateriacorretogrampoabacatejanela

# Generate 5 passphrases
passphrase-gen --count 5

# Use underscore as delimiter
passphrase-gen -d _
# Output: Cavalo_Bateria_Correto_Grampo_Abacate_Janela
```

## Security

- Uses `crypto/rand` exclusively (never `math/rand`)
- 2300+ words Portuguese wordlist (~11.17 bits/word)
- Fisher-Yates shuffle with secure random
- Special characters from: `~!#$%^&*()-=+[]{}:;"'<>?/0123456789`

### Entropy

| Words | Entropy |
|-------|---------|
| 4 | ~45 bits |
| 5 | ~56 bits |
| 6 | ~67 bits |
| 7 | ~78 bits |
| 8 | ~89 bits |

## License

MIT
