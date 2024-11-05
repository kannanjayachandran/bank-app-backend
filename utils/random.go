package utils

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init()  {
	// rand.Seed(time.Now().UnixNano())
}

// returns a random integer between min and max
func GenerateRandomInt(min, max int64) int64  {
	return min + rand.Int63n(max - min + 1)	
}

// return a random string of n characters
func GenerateRandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string  {
	return GenerateRandomString(8)
}

func RandomMoney() int64  {
	return GenerateRandomInt(0, 2000)
}

func RandomCurrency() string  {
	currencies := [] string{"INR", "USD", "YEN"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

