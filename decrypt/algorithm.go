// Decrypt package consists of all decryption algorithms
package decrypt

// Decrypts by reducing the ascii code by 3 for each character.
func Nimbus(str string) string {

	encryptedStr := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode - 3)
		encryptedStr += character
	}
	return encryptedStr
}
