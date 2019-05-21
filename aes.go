package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path"
)

var dataDir string

func getSha40String(pwd string) string {
	bp := []byte(pwd)
	buf := bytes.NewBuffer(make([]byte, 0, len(bp)*40))
	for i := 0; i < 40; i++ {
		buf.Write(bp)
	}
	res := sha256.Sum256(buf.Bytes())
	return fmt.Sprintf("%x", res)
}

func getSha4(pwd string) []byte {
	bp := []byte(pwd)
	buf := bytes.NewBuffer(make([]byte, 0, len(bp)*4))
	for i := 0; i < 4; i++ {
		buf.Write(bp)
	}
	res := sha256.Sum256(buf.Bytes())
	return res[:]
}

func getRealKeyString(pwd string) (string, error) {

	rkey := make([]byte, 32)

	io.ReadFull(rand.Reader, rkey)
	p1 := getSha4(pwd)
	aesBlock, err := aes.NewCipher(p1)
	if err != nil {
		return "", err
	}
	iv := make([]byte, aes.BlockSize)
	io.ReadFull(rand.Reader, iv)

	aesEncoder := cipher.NewCTR(aesBlock, iv)
	var res = make([]byte, 32)
	aesEncoder.XORKeyStream(res, rkey)
	buf := bytes.NewBuffer(iv)
	buf.Write(res)
	return fmt.Sprintf("%x", buf.Bytes()), nil
}

func newRealKeyString(rkey []byte, pwd string) (string, error) {
	p1 := getSha4(pwd)
	aesBlock, err := aes.NewCipher(p1)
	if err != nil {
		return "", err
	}
	iv := make([]byte, aes.BlockSize)
	io.ReadFull(rand.Reader, iv)

	aesEncoder := cipher.NewCTR(aesBlock, iv)
	var res = make([]byte, 32)
	aesEncoder.XORKeyStream(res, rkey)
	buf := bytes.NewBuffer(iv)
	buf.Write(res)
	return fmt.Sprintf("%x", buf.Bytes()), nil
}

func decodeRealKey(key, data, iv []byte) ([]byte, error) {
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesDecoder := cipher.NewCTR(aesBlock, iv)
	var res = make([]byte, 32)
	aesDecoder.XORKeyStream(res, data)
	return res, nil
}

func encodeToFile(data []byte, filename string, key []byte) error {
	dataPath := path.Join(dataDir, filename)
	fp, err := os.Create(dataPath)
	if err != nil {
		return err
	}
	defer fp.Close()

	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	iv := make([]byte, aesBlock.BlockSize())
	io.ReadFull(rand.Reader, iv)
	_, err = fp.Write(iv)
	if err != nil {
		return err
	}
	aesEncoder := cipher.NewCTR(aesBlock, iv)
	wr := cipher.StreamWriter{S: aesEncoder, W: fp}
	_, err = wr.Write(data)
	return err
}

func decodeFromFile(filename string, key []byte) ([]byte, error) {
	dataPath := path.Join(dataDir, filename)
	fp, err := os.Open(dataPath)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, aesBlock.BlockSize())
	_, err = io.ReadFull(fp, iv)
	if err != nil {
		return nil, err
	}
	aesStream := cipher.NewCTR(aesBlock, iv)
	rd := cipher.StreamReader{S: aesStream, R: fp}
	res := bytes.NewBufferString("")
	_, err = io.Copy(res, rd)
	return res.Bytes(), err
}

func encodeToPathName(data []byte, filename string, key []byte) error {
	dataPath := filename
	fp, err := os.Create(dataPath)
	if err != nil {
		return err
	}
	defer fp.Close()

	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	iv := make([]byte, aesBlock.BlockSize())
	io.ReadFull(rand.Reader, iv)
	_, err = fp.Write(iv)
	if err != nil {
		return err
	}
	aesEncoder := cipher.NewCTR(aesBlock, iv)
	wr := cipher.StreamWriter{S: aesEncoder, W: fp}
	_, err = wr.Write(data)
	return err
}

func decodeFromPathName(filename string, key []byte) ([]byte, error) {
	dataPath := filename
	fp, err := os.Open(dataPath)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, aesBlock.BlockSize())
	_, err = io.ReadFull(fp, iv)
	if err != nil {
		return nil, err
	}
	aesStream := cipher.NewCTR(aesBlock, iv)
	rd := cipher.StreamReader{S: aesStream, R: fp}
	res := bytes.NewBufferString("")
	_, err = io.Copy(res, rd)
	return res.Bytes(), err
}
