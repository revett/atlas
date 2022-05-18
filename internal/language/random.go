package language

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

const stringChunkLength = 8

// RandomPhrase generates a random identifier made up of english words of a
// given length.
//
// e.g. "unpaste-scripturalist-toxicology-presidentship"
//
// It makes use of "/usr/share/dict/words" which is an UNIX specific word list,
// thus if this does not exist then it will generate random letter combinations
// instead.
//
// e.g. "ohylwsib-ttgfjwlj-hwqbxlyj-owyrlces".
// TODO: this should not have to be concerned about the dictionary, refactor.
func RandomPhrase(size int) (string, error) {
	gf := generatePhrase
	if _, err := os.Stat("/usr/share/dict/words"); errors.Is(err, os.ErrNotExist) {
		gf = generateString
	}

	return gf(size)
}

func generatePhrase(size int) (string, error) {
	f, err := os.Open("/usr/share/dict/words")
	if err != nil {
		return "", fmt.Errorf("failed to open '/usr/share/dict/words': %w", err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf(
			"failed to read contents of '/usr/share/dict/words': %w", err,
		)
	}

	t := strings.Split(string(bytes), "\n")
	var phrase []string

	for i := 0; i < size; i++ {
		phrase = append(phrase, t[rand.Int()%len(t)])
	}

	return strings.ToLower(
		strings.Join(phrase, "-"),
	), nil
}

func generateString(size int) (string, error) {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyz")

	var str []string

	for i := 0; i < size; i++ {
		b := make([]rune, stringChunkLength)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		str = append(str, string(b))
	}

	return strings.Join(str, "-"), nil
}
