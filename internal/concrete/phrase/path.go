package phrase

import "strings"

type path []string

func (p path) reverse() path {

	reversed := make([]string, len(p))

	for total, i, j := len(p), 0, len(p)-1; i < total; i, j = i+1, j-1 {
		reversed[j] = p[i]
	}

	return reversed
}

func (p path) toString() string {
	return strings.Join(p, " ")
}
