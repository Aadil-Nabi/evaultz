package main

import (
	"crypto/mlkem"
	"encoding/base64"
	"fmt"
)

func main() {
	// Alice generates a new key pair and sends the encapsulation key to Bob.
	decapsulationKey, err := mlkem.GenerateKey768()
	if err != nil {
		fmt.Println(err)
	}
	encapsulationKey := decapsulationKey.EncapsulationKey().Bytes()

	// Bob uses the encapsulation key to encapsulate a shared secret, and sends
	// back the ciphertext to Alice.
	ciphertext := Client(encapsulationKey)

	// Alice decapsulates the shared secret from the ciphertext.
	sharedSecret, err := decapsulationKey.Decapsulate(ciphertext)
	if err != nil {
		fmt.Println(err)

	}
	// Alice and Bob now share a secret.
	secret := base64.StdEncoding.EncodeToString(sharedSecret)
	fmt.Println("Alice's Secret:", secret)

}

func Client(encapsulationKey []byte) (ciphertext []byte) {
	// Client(Bob) encapsulates a shared secret using the encapsulation key.
	ek, err := mlkem.NewEncapsulationKey768(encapsulationKey)
	if err != nil {
		fmt.Println(err)
	}
	sharedSecret, ciphertext := ek.Encapsulate()

	// Alice and Bob now share a secret.
	secret := base64.StdEncoding.EncodeToString(sharedSecret)
	fmt.Println("Bob's Secret:", secret)

	// Bob sends the ciphertext to Alice.
	return ciphertext
}
