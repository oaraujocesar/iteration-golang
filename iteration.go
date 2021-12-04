package iteration

func Repeat(letter string, times int) (repeatedChar string) {

	for i := 0; i < times; i++ {
		repeatedChar += letter
	}

	return
}
