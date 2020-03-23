package flames

import (
	"fmt"
	"strings"
)

var flamesMap = map[string]string{
	"F": "Friends",
	"L": "Lovers",
	"A": "Admirer",
	"M": "Married",
	"E": "Enemy",
	"S": "Siblings",
}

func removeCharacters(input string, characters string) string {
	input = strings.ToLower(input)
	characters = strings.ToLower(characters)

	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}

	return strings.Map(filter, input)
}

// Pair returns a concatenated result of name1 and name2.
func Pair(name1, name2 string) string {
	var x, y string

	for _, rune1 := range name1 {
		for _, rune2 := range name2 {
			if rune1 == rune2 {
				x = removeCharacters(name1, string(rune1))
				y = removeCharacters(name2, string(rune2))
			}
		}
	}

	return fmt.Sprintf("%s%s", x, y)
}

// Get returns the `FLAMES` value ot the pair.
func Get(pair string) string {

	var FLAMES = "FLAMES"
	var flamesLength = 6
	var index int

	for len(FLAMES) != 1 {
		index = len(pair) % flamesLength

		if index == 0 {
			FLAMES = removeCharacters(FLAMES, string(FLAMES[len(FLAMES)-1]))
		} else {
			FLAMES = removeCharacters(FLAMES, string(FLAMES[index-1]))
			FLAMES = FLAMES[index-1:len(FLAMES)] + FLAMES[0:index-1]
		}

		flamesLength--
	}
	return flamesMap[strings.ToUpper(FLAMES)]
}
