package language

import (
	"math/rand"
	"strings"
)

const stringChunkLength = 8

type generatorFunc func(int) (string, error)

// RandomPhrase generates a random identifier made up of english words of a
// given length.
//
// e.g. "unpaste-scripturalist-toxicology-presidentship".
//
// It makes use of "/usr/share/dict/words" which is an UNIX specific word list,
// thus if this does not exist then it will generate random letter combinations
// instead.
//
// e.g. "ohylwsib-ttgfjwlj-hwqbxlyj-owyrlces".
func RandomPhrase(dictionary []string, size int) (string, error) {
	gf := generatePhrase(dictionary)
	if dictionary == nil {
		gf = generateString
	}

	return gf(size)
}

func generatePhrase(dictionary []string) generatorFunc {
	return func(size int) (string, error) {
		var phrase []string

		for i := 0; i < size; i++ {
			phrase = append(phrase, dictionary[rand.Int()%len(dictionary)])
		}

		return strings.ToLower(
			strings.Join(phrase, "-"),
		), nil
	}
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
