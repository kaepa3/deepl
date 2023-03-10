package deepler

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type DeepLResponse struct {
	Translations []Translated
}

type Translated struct {
	DetectedSourceLaguage string `json:"detected_source_language"`
	Text                  string `json:"text"`
}

type Deepler struct {
	url        string
	sourceLang string
	targetLang string
	token      string
}

func getSourceLang(sLang string) string {
	if sLang == "" {
		return "EN"
	}
	return sLang
}

func getTargetLang(tLang string) string {
	if tLang == "" {
		return "Ja"
	}
	return tLang
}

func NewDeepler(token string, sLang string, tLang string) *Deepler {
	url := "https://api-free.deepl.com/v2/translate"
	d := Deepler{
		url:        url,
		sourceLang: getSourceLang(sLang),
		targetLang: getTargetLang(tLang),
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
