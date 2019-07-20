package main

// import (
// 	"bytes"
// 	"crypto/sha256"
// 	"fmt"
// )

// func getSha40String(pwd string) string {
// 	bp := []byte(pwd)
// 	buf := bytes.NewBuffer(make([]byte, 0, len(bp)*40))
// 	for i := 0; i < 40; i++ {
// 		buf.Write(bp)
// 	}
// 	res := sha256.Sum256(buf.Bytes())
// 	return fmt.Sprintf("%x", res)
// }

// func main() {
// 	fmt.Printf("%s\n", getSha40String("abcd"))
// }
