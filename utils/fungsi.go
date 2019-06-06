package utils

import (
	"math/rand"
)

func If_else(a, b string, c string, d string) string {
	if a == b {
		return c
	}
	return d
}

func If_else_int_con(a int, b int, c string, d string) string {
	if a == b {
		return c
	}
	return d
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
