package toolkit

import "crypto/rand"

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is the type used to instantiate this module.
// Any variable of type Tools will have access to all the methods with the receiver *Tools
type Tools struct{}

// RandomString returns a list of random characters of length n, using randomStrinSource as the source of the string
func (t *Tools) RandomString(n int) string {
	s := make([]rune, n)
	r := []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}
