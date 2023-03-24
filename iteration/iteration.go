package iteration

const repeatCount = 5

func Repeat(charactor string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += charactor
	}
	return
}
