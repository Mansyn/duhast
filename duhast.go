package duhast

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Translate(word string) string {
	translated, err := TranslateLibre(word, "en", "de")
	if err != nil {
		return "(error)"
	}
	return translated
}

func TranslateLibre(text, from, to string) (string, error) {
	url := "https://libretranslate.com/translate"
	payload := map[string]string{
		"q":      text,
		"source": from,
		"target": to,
		"format": "text",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		TranslatedText string `json:"translatedText"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.TranslatedText, nil
}