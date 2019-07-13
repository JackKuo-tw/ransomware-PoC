package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"math/big"
)

func CFBEncrypt(originData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	fmt.Printf("key: %s\n", key)
	fmt.Printf("keylen: %d\n", len(key))
	if err != nil {
		fmt.Println(err)
	}
	// const, BlockSize = 16
	blockSize := block.BlockSize()
	fmt.Printf("blocksize: %d\n", blockSize)
	// Use AES-CFB, init vector = key[:16]

	fmt.Printf("init vector: %x\n", key[:blockSize])
	// padding with number in ASCII which is how long it is
	padding := blockSize - len(originData)%blockSize
	padtext := append(originData, bytes.Repeat([]byte{byte(padding)}, padding)...)
	// make buffer and CFBEncrypt
	encryptedData := make([]byte, len(padtext))
	cipher.NewCFBEncrypter(block, key[:blockSize]).XORKeyStream(encryptedData, padtext)
	return encryptedData, nil
}

func CFBDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	blockSize := block.BlockSize()
	originData := make([]byte, len(crypted))
	cipher.NewCFBDecrypter(block, key[:blockSize]).XORKeyStream(originData, crypted)
	// unpadding
	length := len(originData)
	unpadding := int(originData[length-1])
	originData = originData[:(length - unpadding)]
	return originData, nil
}

func RandSeq() []byte {
	charsetLen := int64(len(charset))
	key := make([]byte, 32)
	// AES-256 CFB needs 32 chars
	for i := 0; i < 32; i++ {
		a, _ := rand.Int(rand.Reader, big.NewInt(charsetLen))
		key[i] = charset[a.Int64()]
	}

	return key
}
