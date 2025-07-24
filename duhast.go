package duhast

var dictionary = map[string]string{
    "hello": "hallo",
    "goodbye": "auf wiedersehen",
    "please": "bitte",
    "thank you": "danke",
    "yes": "ja",
    "no": "nein",
}

func Translate(word string) string {
    if german, ok := dictionary[word]; ok {
        return german
    }
    return ""
}