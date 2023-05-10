package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	arr := Matrix{{1, 2, 3}, {2, 5, 2}, {2, 2, 1}}
	texted := "Оружие"
	//arr := Matrix{{1, 3, 2}, {2, 1, 5}, {3, 2, 1}}
	//texted := "Приказ"
	texted = strings.ToUpper(texted)
	alpha := "АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
	alphaMap := map[rune]float64{}
	k := 1
	for _, i := range alpha {
		if i == 0 {
			continue
		}
		alphaMap[i] = float64(k)
		k++
	}
	text := []float64{}
	for _, i := range texted {
		if i == 0 {
			continue
		}
		text = append(text, alphaMap[i])
	}
	fmt.Println("Given text:", texted)
	fmt.Println("Periodic number:", text)
	ciphered := [6]float64{}

	for i := 0; i < 3; i++ {
		ciphered[i] = arr[i][0]*text[0] + arr[i][1]*text[1] + arr[i][2]*text[2]
	}

	for i := 0; i < 3; i++ {
		ciphered[i+3] = arr[i][0]*text[3] + arr[i][1]*text[4] + arr[i][2]*text[5]
	}

	fmt.Println("Ciphertext:", ciphered)

	inv, det := arr.Inverse()

	fmt.Println("Determinant", det)
	fmt.Println("Inverted", inv)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			inv[i][j] /= det
		}
	}

	decrypted := [6]float64{}

	for i := 0; i < 3; i++ {
		decrypted[i] = inv[i][0]*ciphered[0] + inv[i][1]*ciphered[1] + inv[i][2]*ciphered[2]
	}

	for i := 0; i < 3; i++ {
		decrypted[i+3] = inv[i][0]*ciphered[3] + inv[i][1]*ciphered[4] + inv[i][2]*ciphered[5]
	}
	for i := 0; i < 6; i++ {
		decrypted[i] = math.Round(decrypted[i])
	}
	fmt.Println("Decrypted:", decrypted)

	decText := ""
	for i := range decrypted {
		for v, key := range alphaMap {
			if key == decrypted[i] {
				decText += string(v)
			}
		}
	}

	fmt.Println("Decrypted text:", decText)

}

type Matrix [3][3]float64

func (m Matrix) Inverse() (Matrix, float64) {
	var inv Matrix
	inv[0][0] = m[1][1]*m[2][2] - m[1][2]*m[2][1]
	inv[1][0] = m[1][0]*m[2][2] - m[1][2]*m[2][0]
	inv[1][0] *= -1
	inv[2][0] = m[1][0]*m[2][1] - m[2][0]*m[1][1]
	inv[0][1] = m[0][1]*m[2][2] - m[0][2]*m[2][1]
	inv[0][1] *= -1
	inv[1][1] = m[0][0]*m[2][2] - m[0][2]*m[2][0]
	inv[2][1] = m[0][0]*m[2][1] - m[0][1]*m[2][0]
	inv[2][1] *= -1
	inv[0][2] = m[0][1]*m[1][2] - m[0][2]*m[1][1]
	inv[1][2] = m[0][0]*m[1][2] - m[0][2]*m[1][0]
	inv[1][2] *= -1
	inv[2][2] = m[0][0]*m[1][1] - m[1][0]*m[0][1]

	determinant := m[0][0]*inv[0][0] + m[0][1]*inv[1][0] + m[0][2]*inv[2][0]

	return inv, determinant
}
