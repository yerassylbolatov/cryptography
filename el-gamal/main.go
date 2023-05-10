package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"time"
)

func main() {
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
	fmt.Printf("\n========================\n")
	fmt.Printf("Generated private key %d", x)
	fmt.Printf("\n========================\n")

	// Compute the corresponding public key, y
	y := new(big.Int).Exp(q, x, p)
	fmt.Println("y =", y)

	// Encrypt a message, m
	var enc int64
	fmt.Println("\nEnter message to encrypt")
	fmt.Scanln(&enc)
	m := big.NewInt(enc)
	k := getRandomK(p)

	a := new(big.Int).Exp(q, k, p)
	s := new(big.Int).Exp(y, k, p)
	b := s.Mul(s, m)
	b.Mod(b, p)

	// Print the encrypted message
	fmt.Println("Encrypted message: (a =", a, ", b =", b, ")\n")

	//Decrypt the message
	for {
		fmt.Print("\nENTER PRIVATE KEY 'X' TO DECRYPT MESSAGE: ")
		var user_x int64
		fmt.Scanln(&user_x)
		res := new(big.Int).Exp(a, big.NewInt(user_x), p)
		res.ModInverse(res, p)
		res.Mul(res, b)
		res.Mod(res, p)
		fmt.Printf("Decrypted message: %d", res)
		if x.Int64() == user_x {
			fmt.Printf("\nSuccess!\n")
			break
		}
		fmt.Println()
	}

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
