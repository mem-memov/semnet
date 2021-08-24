package word

type path []rune

func (p path) reverse() path {

	reversed := make([]rune, len(p))

	for total, i, j := len(p), 0, len(p)-1; i < total; i, j = i+1, j-1 {
		reversed[j] = p[i]
	}

	return reversed
}

func (p path) toString() string {
	return string(p)
}
