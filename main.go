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
	d := NewDeepler(os.Getenv("DEEPL_TOKEN"))
	if text, err := d.Translate("Hello"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(text)
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

type Deepler struct {
	url        string
	sourceLang string
	targetLang string
	token      string
}

func NewDeepler(token string) *Deepler {
	url := "https://api-free.deepl.com/v2/translate"
	d := Deepler{
		url:        url,
		sourceLang: "EN",
		targetLang: "Ja",
		token:      token,
	}
	return &d
}

func (d *Deepler) Translate(text string) (string, error) {
	params := url.Values{}
	params.Add("auth_key", d.token)
	params.Add("source_lang", d.sourceLang)
	params.Add("target_lang", d.targetLang)
	params.Add("text", text)
	resp, err := http.PostForm(d.url, params)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var respData DeepLResponse
	decoder.Decode(&respData)
	return respData.Translations[0].Text, nil
}
