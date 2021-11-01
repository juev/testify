package testify

func F(num int) (float32, error) {
	if (num+1)%3 == 0 {
		panic("panic")
	}
	return float32(1) / float32(num-5), nil
}
