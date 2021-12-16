package iteration

func Iterate(character string, times int) string {
	var repetitions string

	for i := 0; i < times; i++ {
		repetitions += character
	}

	return repetitions
}
