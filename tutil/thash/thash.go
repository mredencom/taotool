package thash

type THash struct {
}

type Hash interface {
	// AdditiveHash 加法hash
	AdditiveHash(key string, prime int) int
	// RotatingHash 旋转hash
	RotatingHash(key string, prime int) int
	// OneByOneHash 一次一个hash
	OneByOneHash(key string, prime int) int

	Bernstein(key string) int
}

func (THash) AdditiveHash(key string, prime int) int {
	//var hash, i int64
	j := len(key)
	for i := 0; i < j; i++ {
		key += string(rune(i))
	}
	return len(key) % prime
}
