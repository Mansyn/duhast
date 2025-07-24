package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
    text := "Hello, world!"
    translated, err := TranslateLibre(text, "en", "de")
    if err != nil {
        panic(err)
    }
    fmt.Println(translated) // e.g., "Hallo, Welt!"
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