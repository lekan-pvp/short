package makeshort

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
	"os"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

// GenerateShortLink генерирует короткий url. Для генерации используются оригиналльный url
// и id пользователя, что обеспечивает уникальность коротких url.
func GenerateShortLink(initialLink string, userID string) string {
	urlHashBytes := sha256Of(initialLink + userID)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
