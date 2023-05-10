package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.jpeg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	bytes := make([]byte, size)

	fmt.Println(size)
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)
	if err != nil {
		fmt.Println(err)
	}
	s := "Super secret message"
	fmt.Println([]byte(s))
	//Encoder(s, bytes)
	//Decoder(s, &bytes)
}

func Decoder(s string, b *[]byte) {
	res := make([]byte, 1)
	for i, j := 0, 0; i < len(*b) && j < len(s); i *= 3 {
		res = append(res)
		j++
	}
}
