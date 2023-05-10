package main

import (
	"fmt"
)

func xorEncode(message, key []byte) []byte {
	encodedMessage := make([]byte, len(message))

	for i := range message {
		encodedMessage[i] = message[i] ^ key[i]
	}
	return encodedMessage
}

func xorDecode(encodedMessage, key []byte) []byte {
	decodedMessage := make([]byte, len(encodedMessage))

	for i := range encodedMessage {
		decodedMessage[i] = encodedMessage[i] ^ key[i]
	}
	return decodedMessage
}

func main() {
	fmt.Println("Enter message to encrypt")
	var message string
	fmt.Scan(&message)
	fmt.Println("Enter key")
	var key string
	fmt.Scan(&key)

	l1, l2 := len(key), len(message)
	for l1 != l2 {
		key += " " + key
		l1++
	}

	encodedMessage := xorEncode([]byte(message), []byte(key))
	fmt.Printf("Encoded message: %d\n", encodedMessage)
	decodedMessage := xorDecode(encodedMessage, []byte(key))
	fmt.Printf("Decoded message: %s\n", string(decodedMessage))
}
