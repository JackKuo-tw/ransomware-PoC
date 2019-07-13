package main

import (
	"flag"
	"io/ioutil"
	"os"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+")
var key []byte
var iv []byte
var path = "."
var ch = make(chan string, 3)
var extension = []string{".txt", ".jpg", ".jpeg", ".png"}

const (
	ENC = iota
	DEC
)

func main() {
	// default value, can be rand by -R flag
	key = []byte("11111111111111111111111111111111")
	flag.Parse()
	flagHandle()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func unsecret(fileName string, key []byte) {
	encryptedData, err := ioutil.ReadFile(fileName)
	check(err)
	originFile, err := CFBDecrypt(encryptedData, key)
	check(err)
	length := len(fileName)
	encryptedFile, err := os.Create(fileName[:length-8])
	check(err)
	defer encryptedFile.Close()
	encryptedFile.Write(originFile)
	encryptedFile.Sync()
}

func makesecret(fileName string, key []byte) {
	originFile, err := ioutil.ReadFile(fileName)
	check(err)
	encryptedData, err := CFBEncrypt(originFile, key)
	check(err)
	encryptedFile, err := os.Create(fileName + ".CFBEncrypt")
	check(err)
	defer encryptedFile.Close()
	encryptedFile.Write(encryptedData)
	encryptedFile.Sync()
}
