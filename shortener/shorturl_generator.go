package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
	"os"
)

// 算法：
// url+userId 进行sha256
// 然后进行base58计算
// 最后取base58的前8位作为短链接

func GenerateShortLink(initiaLink string, userId string) string {
	urlHashBytes := sha256Of(initiaLink + userId)
	generatorNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encode([]byte(fmt.Sprintf("%d", generatorNumber)))
	return finalString[:8]
}

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encode(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}
