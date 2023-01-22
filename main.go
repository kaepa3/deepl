package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kaepa3/deepl/deepler"
)

const (
	confFile = ".deepl"
)

func readConfig() {
	paths := []string{
		filepath.Join(".env"),
		filepath.Join(os.Getenv("HOME"), confFile),
		filepath.Join(os.Getenv("HOMEPATH"), confFile),
	}
	for _, v := range paths {
		if err := godotenv.Load(v); err == nil {
			return
		}
		log.Println(v)
	}
	log.Println("Error loading .env file", paths)
	return
}

func main() {
	flag.Parse()
	text := flag.Args()
	if len(text) == 0 {
		fmt.Println("error", text)
		return
	}
	readConfig()
	token := os.Getenv("DEEPL_TOKEN")
	sLang := os.Getenv("SOURCE_LANG")
	tLang := os.Getenv("TARGET_LANG")
	d := deepler.NewDeepler(token, sLang, tLang)
	if text, err := d.Translate(strings.Join(text, ",")); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(text)
	}
}
