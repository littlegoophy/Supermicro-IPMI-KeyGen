package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2{
		fmt.Printf("\033[31mUsage ./keygen 'MAC ADDRESS'\033[0m\n")
		os.Exit(1)
	}
	mac := os.Args[1]
	key, _ := hex.DecodeString("8544E3B47ECA58F9583043F8")
	lic := generateLicense(mac, key)
	b := bytes.Buffer{}

	for i, char := range lic {
		if i < 24 {
			if i%4 == 0 && i!=0 {
				b.WriteString(" " + string(char))
			} else {
				b.WriteString(string(char))
			}
		} else {
			break
		}
	}
	fmt.Printf("\033[32mYour License Key:\033[33m %s\u001B[0m\n", b.String())

}
func generateLicense(mac string, key []byte) string {
	sha := hmac.New(sha1.New, key)
	sha.Write([]byte(mac))
	return hex.EncodeToString(sha.Sum(nil))

}
