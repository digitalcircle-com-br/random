package random

import (
	"encoding/base64"
	"math/big"
	"math/rand"
	"time"
	"unsafe"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const upperBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const letterNumsBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const letterNumsUpperBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func strFromSrc(n int, srcChars string) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(srcChars) {
			b[i] = srcChars[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

//Str - creates randon str w len of n. Letter only (upper and lower).
func Str(n int) string {
	return strFromSrc(n, letterBytes)
}

//StrUpper - same as Str, but only Uppercase Letters.
func StrUpper(n int) string {
	return strFromSrc(n, upperBytes)
}

//StrLetterNum - same as Str, but Letters and Numbers.
func StrLetterNum(n int) string {
	return strFromSrc(n, letterNumsBytes)
}

//StrLetterNumUpper - same as Str, but Letters Upper only and Numbers.
func StrLetterNumUpper(n int) string {
	return strFromSrc(n, letterNumsUpperBytes)
}

//TSNano - returns now in nanosecs encoded b64.
func TSNano() string {
	e := time.Now().UnixNano()
	eb := big.NewInt(int64(e))
	return base64.RawURLEncoding.EncodeToString(eb.Bytes())
}

//TSNano - returns now in secs encoded b64.
func TS() string {
	e := time.Now().Unix()
	eb := big.NewInt(int64(e))
	return base64.RawURLEncoding.EncodeToString(eb.Bytes())
}

//StrTSNano - Same as Str,but adds UnixNano in Base64 as prefix.
func StrTSNano(n int) string {

	return TSNano() + "-" + Str(n)
}

//StrTS - Same as Str,but adds Unix in Base64 as prefix.
func StrTS(n int) string {
	return TS() + "-" + Str(n)
}
