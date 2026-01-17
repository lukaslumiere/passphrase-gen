package generator

import (
	"crypto/rand"
	"errors"
	"math"
	"math/big"
	"strings"

	"github.com/lukaslumiere/passphrase-gen/pkg/wordlist"
)

const (
	Version      = "1.0.0"
	DefaultWords = 6
	DefaultCount = 1
	DefaultSep   = "-"
	MinWords     = 1
	MaxWords     = 20
	MaxSpecials  = 5
	SpecialChars = "~!#$%^&*()-=+[]{}:;\"'<>?/0123456789"
)

var (
	ErrEmptyWordlist   = errors.New("wordlist is empty")
	ErrInvalidWords    = errors.New("word count must be between 1 and 20")
	ErrInvalidSpecials = errors.New("specials count must be between 0 and 5")
	ErrRandomFailed    = errors.New("failed to generate random number")
)

type Config struct {
	Words     int
	Separator string
	Capitalize bool
	Specials  int
}

func NewConfig() Config {
	return Config{
		Words:      DefaultWords,
		Separator:  DefaultSep,
		Capitalize: true,
	}
}

func randomInt(max int) (int, error) {
	if max <= 0 {
		return 0, nil
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, ErrRandomFailed
	}
	return int(n.Int64()), nil
}

func shuffleIndices(n int) ([]int, error) {
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}
	for i := n - 1; i > 0; i-- {
		j, err := randomInt(i + 1)
		if err != nil {
			return nil, err
		}
		indices[i], indices[j] = indices[j], indices[i]
	}
	return indices, nil
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func selectWords(count int, caps bool) ([]string, error) {
	words := wordlist.Words
	if len(words) == 0 {
		return nil, ErrEmptyWordlist
	}
	if count < MinWords || count > MaxWords {
		return nil, ErrInvalidWords
	}

	indices, err := shuffleIndices(len(words))
	if err != nil {
		return nil, err
	}

	result := make([]string, count)
	for i := 0; i < count; i++ {
		word := words[indices[i]]
		if caps {
			word = capitalize(word)
		}
		result[i] = word
	}
	return result, nil
}

func generateSpecials(count int) (string, error) {
	if count <= 0 {
		return "", nil
	}
	chars := make([]byte, count)
	for i := 0; i < count; i++ {
		idx, err := randomInt(len(SpecialChars))
		if err != nil {
			return "", err
		}
		chars[i] = SpecialChars[idx]
	}
	return string(chars), nil
}

func Generate(cfg Config) (string, error) {
	if cfg.Specials < 0 || cfg.Specials > MaxSpecials {
		return "", ErrInvalidSpecials
	}

	words, err := selectWords(cfg.Words, cfg.Capitalize)
	if err != nil {
		return "", err
	}

	passphrase := strings.Join(words, cfg.Separator)

	if cfg.Specials > 0 {
		specials, err := generateSpecials(cfg.Specials)
		if err != nil {
			return "", err
		}
		passphrase += specials
	}

	return passphrase, nil
}

func GenerateMultiple(cfg Config, count int) ([]string, error) {
	results := make([]string, count)
	for i := 0; i < count; i++ {
		p, err := Generate(cfg)
		if err != nil {
			return nil, err
		}
		results[i] = p
	}
	return results, nil
}

func CalculateEntropy(cfg Config) float64 {
	wordCount := len(wordlist.Words)
	if wordCount == 0 {
		return 0
	}

	bitsPerWord := math.Log2(float64(wordCount))
	entropy := float64(cfg.Words) * bitsPerWord

	if cfg.Specials > 0 {
		entropy += float64(cfg.Specials) * math.Log2(float64(len(SpecialChars)))
	}

	return entropy
}
