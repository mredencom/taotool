package thash

type THash struct {
}

type Hash interface {
	// AdditiveHash 加法hash
	// key 字符串
	// prime 一个质数
	AdditiveHash(key string, prime int) int
}

func (THash) AdditiveHash(key string, prime int) int {
	//var hash, i int64
	j := len(key)
	for i := 0; i < j; i++ {
		key += string(rune(i))
	}
	return len(key) % prime
}
