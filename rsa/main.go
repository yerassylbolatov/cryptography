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

	fmt.Println(indexed)

	p, q := 0, 0
	fmt.Println("Enter p")
	fmt.Scanln(&p)
	fmt.Println("Enter q")
	fmt.Scanln(&q)

	if !isPrime(p) || !isPrime(q) {
		fmt.Println("Numbers are not prime")
		os.Exit(1)
	}

	n := p * q
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
		if ((e*d)%z == 1) && (e < n) && e != d {
			break
		}
	}
	fmt.Println("e =", e)
	fmt.Println("d =", d)
	open_key := []int64{int64(e), int64(n)}
	secret_key := []int64{int64(d), int64(n)}

	fmt.Println("open key:", open_key)
	fmt.Println("secret key:", secret_key)

	encrypted := make([]*big.Int, len(indexed))
	for i := 0; i < len(indexed); i++ {
		var f big.Int
		encrypted[i] = f.Exp(big.NewInt(indexed[i]), big.NewInt(open_key[0]), big.NewInt(open_key[1]))
		//encrypted[i] = int64(math.Pow(float64(indexed[i]), float64(open_key[0]))) % int64(open_key[1])
	}

	fmt.Println("Encrypted", encrypted)

	decrypted := make([]*big.Int, len(indexed))
	for i := 0; i < len(decrypted); i++ {
		var k big.Int
		decrypted[i] = k.Exp(encrypted[i], big.NewInt(secret_key[0]), big.NewInt(secret_key[1]))
	}

	fmt.Println("Decrypted", decrypted)

	res := ""

	for i := range decrypted {
		res += string(alpha[decrypted[i].Int64()-1])
	}

	fmt.Println("Decrypted text", res)
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
