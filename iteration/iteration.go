package iteration

func Repeat(charactor string, repeatCount int) (repeated string) {
	repeated = charactor
	for i := 1; i < repeatCount; i++ {
		repeated += charactor
	}
	return
}
