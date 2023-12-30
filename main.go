package main

import (
	"fmt"

	"github.com/Milad75Rasouli/cryptit/decrypt"
	"github.com/Milad75Rasouli/cryptit/encrypt"
)

func main() {
	text := "dude!"
	en := encrypt.Nimbus(text)
	fmt.Println("Encrypted Text:", en)
	de := decrypt.Nimbus(en)
	fmt.Println("Decrypted Text:", de)
}
