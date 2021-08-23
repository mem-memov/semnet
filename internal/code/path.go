package code

type path []bool

func (p path) reverse() path {

	reversed := make([]bool, len(p))

	for total, i, j := len(p), 0, len(p)-1; i < total; i, j = i+1, j-1 {
		reversed[j] = p[i]
	}

	return reversed
}
