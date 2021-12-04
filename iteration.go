package iteration

const repetitions = 5

func Repeat(letter string) (repeatedChar string) {

	for i := 0; i < repetitions; i++ {
		repeatedChar += letter
	}

	return
}
