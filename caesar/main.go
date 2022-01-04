package main

import "fmt"

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	var ret []rune
	for _, ch := range input {
		ret = append(ret, cipher(ch, delta))
	}
	fmt.Println(string(ret))
}

const (
	start, end       = 'a', 'z'
	startCap, endCap = 'A', 'Z'
)

func cipher(r rune, delta int) rune {
	if r >= startCap && r <= endCap {
		return rotate(r, startCap, delta)
	}
	if r >= start && r <= end {
		return rotate(r, start, delta)
	}
	return r
}

func rotate(r rune, base, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}
