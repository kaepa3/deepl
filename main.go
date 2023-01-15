package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	confFile = "deepl"
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
	readConfig()
	if err := next(os.Getenv("DEEPL_TOKEN")); err != nil {
		fmt.Println(err)
	}
}

type DeepLResponse struct {
	Translations []Translated
}

type Translated struct {
	DetectedSourceLaguage string `json:"detected_source_language"`
	Text                  string `json:"text"`
}

func next(token string) error {
	u := "https://api-free.deepl.com/v2/translate"
	params := url.Values{}
	params.Add("auth_key", token)
	params.Add("source_lang", "EN")
	params.Add("target_lang", "Ja")
	params.Add("text", "Hello")
	resp, err := http.PostForm(u, params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var respData DeepLResponse
	decoder.Decode(&respData)
	fmt.Printf("resp:%v", respData)
	return nil
}
