package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("1 - Polybius Square")
	fmt.Println("2 - Caesar encryption")
	fmt.Println("3 - Trisemus algorithm")
	fmt.Println("CHOOSE ONE ALGORITHM FOR ENCRYPTION")

	var algo int
	fmt.Scan(&algo)
	if algo == 1 {
		polybius()
	} else if algo == 2 {
		ceasereus()
	} else if algo == 3 {
		trisemus()
	}
}

func ceasereus() {
	s := ""
	fmt.Print("Enter word to encrypt\n")
	fmt.Scan(&s)
	s = strings.ToLower(s)
	var key int
	fmt.Print("Specify K to shift\n")
	fmt.Scan(&key)

	fmt.Print("Encrypted: ")
	enc := rotate(s, key)
	fmt.Println(enc)

	fmt.Print("Decrypted: ")
	fmt.Println(rotate(enc, -key))
}

func rotate(text string, shift int) string {
	shift = (shift%26 + 26) % 26
	b := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		t := text[i]
		var a int
		switch {
		case 'a' <= t && t <= 'z':
			a = 'a'
		case 'A' <= t && t <= 'Z':
			a = 'A'
		default:
			b[i] = t
			continue
		}
		b[i] = byte(a + ((int(t)-a)+shift)%26)
	}
	return string(b)
}

func polybius() {
	s := ""
	fmt.Print("Enter word to encrypt\n")
	fmt.Scan(&s)
	s = strings.ToLower(s)

	alphabet := "abcdefghiklmnopqrstuvwxyz"
	arr := [5][5]rune{}
	i, j := 0, 0
	for _, l := range alphabet {
		arr[i][j] = l
		j++
		if j != 0 && j%5 == 0 {
			i++
			j = 0
		}
	}

	enc := ""
	for _, l := range s {
		for a := 0; a < 5; a++ {
			for b := 0; b < 5; b++ {
				if arr[a][b] == l {
					if a == 4 {
						enc += string(arr[0][b])
					} else {
						enc += string(arr[a+1][b])
					}
				}
			}
		}
	}
	fmt.Print("encrypted: ")
	fmt.Println(enc)

	dec := ""
	for _, l := range enc {
		for a := 0; a < 5; a++ {
			for b := 0; b < 5; b++ {
				if arr[a][b] == l {
					if a == 0 {
						dec += string(arr[4][b])
					} else {
						dec += string(arr[a-1][b])
					}
				}
			}
		}
	}
	fmt.Print("decrypted: ")
	fmt.Println(dec)
}

func trisemus() {
	s := ""
	fmt.Print("Enter word to encrypt\n")
	fmt.Scan(&s)
	s = strings.ToLower(s)

	key := ""
	fmt.Print("Enter key\n")
	fmt.Scan(&key)
	key = strings.ToLower(key)
	arr := [5][5]rune{}

	//fmt.Println(arr)
	hm := make(map[rune]bool)
	i := 0
	j := 0
	for _, letter := range key {
		if _, ok := hm[letter]; !ok {
			arr[i][j] = letter
			hm[letter] = true
			j++
		}
		if j != 0 && j%5 == 0 {
			i++
			j = 0
		}
	}

	alphabet := "abcdefghiklmnopqrstuvwxyz"

	for _, letter := range alphabet {
		if _, ok := hm[letter]; !ok {
			arr[i][j] = letter
			j++
		}
		if i != 4 && j != 0 && j%5 == 0 {
			i++
			j = 0
		}
	}

	res := ""
	for _, l := range s {
		for a := 0; a < 5; a++ {
			for b := 0; b < 5; b++ {
				if arr[a][b] == l {
					if a == 4 {
						res += string(arr[0][b])
					} else {
						res += string(arr[a+1][b])
					}
				}
			}
		}
	}

	fmt.Print("encrypted: ")
	fmt.Println(res)

	decrypted := ""
	for _, l := range res {
		for a := 0; a < 5; a++ {
			for b := 0; b < 5; b++ {
				if arr[a][b] == l {
					if a == 0 {
						decrypted += string(arr[4][b])
					} else {
						decrypted += string(arr[a-1][b])
					}
				}
			}
		}
	}
	fmt.Print("decrypted: ")
	fmt.Println(decrypted)
}
