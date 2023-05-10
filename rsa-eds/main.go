package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789"

	runed_alpha := []rune(alpha)

	text := "Yerassyl Bolatov"
	text = strings.ToUpper(text)
	runed_text := []rune(text)
	fmt.Println("Text to encrypt", text)

	indexed := make([]int64, len(runed_text))
	for idx, i := range runed_text {
		for id, j := range runed_alpha {
			if i == j {
				indexed[idx] = int64(id + 1)
			}
		}
	}

	fmt.Println("INDEXED:", indexed)

	p, q := 0, 0
	fmt.Println("Enter p")
	fmt.Scanln(&p)
	fmt.Println("Enter q")
	fmt.Scanln(&q)

	if !isPrime(p) || !isPrime(q) {
		fmt.Println("Numbers are not prime")
		os.Exit(1)
	}

	n := int64(p * q)
	z := (p - 1) * (q - 1)
	fmt.Println("n =", n)
	fmt.Println("z =", z)

	rand.Seed(time.Now().UnixNano())

	e := 0
	for {
		e = rand.Intn(z-1) + 1
		if GCD(e, z) == 1 && (1 < e && e < z) && z%e != 0 {
			break
		}
	}

	d := 0
	for {
		d++
		if ((e*d)%z == 1) && (e < int(n)) && e != d {
			break
		}
	}
	fmt.Println("e =", e)
	fmt.Println("d =", d)
	open_key := []int64{int64(e), int64(n)}
	secret_key := []int64{int64(d), int64(n)}

	fmt.Println("open key:", open_key)
	fmt.Println("secret key:", secret_key)

	h := make([]*big.Int, len(indexed))
	h0 := big.NewInt(int64(rand.Intn(20) + 1))
	for i := 0; i < len(indexed); i++ {
		var f big.Int
		var c big.Int
		if i == 0 {
			h[i] = f.Exp(c.Add(big.NewInt(indexed[i]), h0), big.NewInt(2), big.NewInt(n))
		} else {
			h[i] = f.Exp(c.Add(big.NewInt(indexed[i]), h[i-1]), big.NewInt(2), big.NewInt(n))
		}
	}
	fmt.Println("hash:", h)

	m := h[len(h)-1]
	fmt.Println("m =", m)

	var fi big.Int
	s := fi.Exp(h[len(h)-1], big.NewInt(secret_key[0]), big.NewInt(secret_key[1]))
	fmt.Println("s =", s)

	var user int64
	fmt.Println("Enter signature:")
	fmt.Scanln(&user)

	var fp big.Int
	ms := fp.Exp(big.NewInt(user), big.NewInt(open_key[0]), big.NewInt(open_key[1]))
	fmt.Println("ms =", ms)

	if ms.Int64() == m.Int64() {
		fmt.Println("VALID")
	} else {
		fmt.Println("INVALID")
	}
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func GCD(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}
