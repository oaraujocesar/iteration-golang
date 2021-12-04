package iteration

import "strings"

func Repeat(letter string, times int) (repeatedChar string) {

	repeatedChar = strings.Repeat(letter, times)

	return
}
