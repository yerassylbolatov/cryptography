package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"time"
)

func hashFunction(m string) []byte {
	h := sha256.New()
	h.Write([]byte(m))
	hash := h.Sum(nil)
	return hash
}

func egKey() (*big.Int, *big.Int, *big.Int, *big.Int) {
	var i, j int64
	fmt.Println("Enter value for prime number p:")
	fmt.Scanln(&i)

	fmt.Println("Enter value for prime number q:")
	fmt.Scanln(&j)

	p := big.NewInt(i)
	q := big.NewInt(j)

	if p.Cmp(q) != 1 {
		fmt.Println("p must be greater than q.")
		os.Exit(1)
	}

	if new(big.Int).GCD(nil, nil, p, q).Cmp(big.NewInt(1)) != 0 {
		fmt.Println("p and q must be mutually prime.")
		os.Exit(1)
	}

	// Choose a random private key, x
	rand.Seed(time.Now().UnixNano())
	x := big.NewInt(int64(rand.Intn(int(p.Int64())-1) + 1))

	// Compute the corresponding public key, y
	y := new(big.Int).Exp(q, x, p)

	return p, q, y, x
}

func egGen(p, q, x *big.Int, hash []byte) (*big.Int, []int64) {
	k := getRandomK(p)
	a := new(big.Int).Exp(q, k, p)
	s := new(big.Int).ModInverse(k, new(big.Int).Sub(p, big.NewInt(1)))

	p1 := p.Int64()
	x1 := x.Int64()
	s1 := s.Int64()
	a1 := a.Int64()

	b := make([]int64, len(hash))
	for i := range hash {
		b[i] = s1 * (int64(hash[i]) - x1*a1) % (p1 - 1)
	}
	return a, b
}

func getRandomK(p *big.Int) *big.Int {
	var k *big.Int
	for {
		rand.Seed(time.Now().UnixNano())
		k = big.NewInt(rand.Int63n(p.Int64() - 1))
		if k.Cmp(big.NewInt(0)) != 0 && new(big.Int).GCD(nil, nil, k, new(big.Int).Sub(p, big.NewInt(1))).Cmp(big.NewInt(1)) == 0 {
			break
		}
	}
	return k
}

func egVer(p, q, y, a *big.Int, b []int64, message string) {
	v1 := make([]int64, len(b))
	v2 := make([]int64, len(b))

	for i := range b {
		var tmp big.Int
		tmp.Exp(y, a, p)
		var tmp1 big.Int
		tmp1.Exp(a, big.NewInt(b[i]), p)
		v1[i] = tmp1.Int64() * tmp.Int64()
	}

	hashed := hashFunction(message)
	fmt.Printf("Hashed message by recipient: %x\n", hashed)
	for i := range hashed {
		var tmp big.Int
		tmp.Exp(q, big.NewInt(int64(hashed[i])), p)
		v2[i] = tmp.Int64()
	}
	fmt.Println(v1)
	fmt.Println(v2)
}

func main() {
	var message string
	fmt.Println("Enter message")
	fmt.Scanln(&message)

	hashed := hashFunction(message)
	fmt.Printf("Hashed message: %x\n", hashed)

	p, q, y, x := egKey()
	fmt.Println("Public keys (p, q, y):", p, q, y)
	fmt.Println("Secret key:", x)

	a, b := egGen(p, q, x, hashed)
	fmt.Println("Signature (a, b):", a, b)

	egVer(p, q, y, a, b, message)
}
