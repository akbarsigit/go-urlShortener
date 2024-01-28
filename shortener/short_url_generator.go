package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// hashing of the url
func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))

	return algorithm.Sum(nil)
}

// turn the binary to text encoding. Base 58 used its not confusing (have high entropy)
func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(encoded)
}


// to generate the final short url, we need to hash the initialUrl + userId.
// with this design, we can create different result even with same URL. 
// From the hash, then we create an Integer representation and then encode it with base58
// and choose the 8 first character.
func GenerateShortLink(initialUrl string, userId string) string {
	urlHashBytes := sha256Of(initialUrl + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%v", generatedNumber)))
	return finalString[:8]
}