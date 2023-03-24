package iteration

func Repeat(s string) (r string) {
	for i := 0; i < 5; i++ {
		r += s
	}
	return
}
