package transateITE

import (
	gtranslate "github.com/gilang-as/google-translate"
)

func TranslateITE(toTranslate string) string {
	var ruText string
	value := gtranslate.Translate{
		Text: toTranslate,
		From: "en",
		To:   "ru",
	}
	translated, err := gtranslate.Translator(value)
	if err != nil {
		panic(err)
	} else {
		ruText = translated.Text
	}
	return ruText
}
