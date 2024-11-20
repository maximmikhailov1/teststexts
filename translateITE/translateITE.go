package transateITE

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"teststexts/glossary"

	gtranslate "github.com/gilang-as/google-translate"
	"github.com/joho/godotenv"
)

type Translator struct {
	TranslateMech string
}

func (t Translator) TranslateITE(toTranslate string) string {
	var ruText string
	switch t.TranslateMech {
	case "yandex":
		ruText = yandexTranslate(toTranslate)
	case "google":
		ruText = googleTranslate(toTranslate)
	}

	return ruText
}

func yandexTranslate(toTranslate string) string {
	//https://yandex.cloud/ru/docs/translate/api-ref/Translation/translate
	var ruText string = ""
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error downloading .env file")
	}

	iam := os.Getenv("YANDEX_IAM_TOKEN")
	folderId := os.Getenv("YANDEX_FOLDER_ID")
	targetLanguage := "ru"
	sourceLanguageCode := "en"
	texts := []string{toTranslate}
	//Объявление глоссария
	glossaryPairs := glossary.Glossary
	glossaryData := map[string]any{"glossaryPairs": glossaryPairs}
	glossaryConfig := map[string]any{"glossaryData": glossaryData}

	body := map[string]any{
		"sourceLanguageCode": sourceLanguageCode,
		"targetLanguageCode": targetLanguage,
		"texts":              texts,
		"folderId":           folderId,
		"glossaryConfig":     glossaryConfig,
	}

	bodyPOST, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Failed to Marshal body with err: %s", err.Error())
	}
	responseBody := bytes.NewBuffer(bodyPOST)
	requestToAPI, err := http.NewRequest(http.MethodPost,
		"https://translate.api.cloud.yandex.net/translate/v2/translate",
		responseBody)

	if err != nil {
		log.Fatalf("Failed to create a request to api with error: %s", err.Error())
	}
	requestToAPI.Header.Add("Content-Type", "application/json")
	requestToAPI.Header.Add("Authorization", fmt.Sprintf("Bearer %s", iam))

	response, err := http.DefaultClient.Do(requestToAPI)
	if err != nil {
		log.Fatalf("Failed to send request to url: %s", err.Error())
	}
	defer response.Body.Close()

	ResponseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response bytes: %s", err.Error())
	}
	// bodyString := string(ResponseBodyBytes)
	// fmt.Println("Response:", bodyString)
	var responseText map[string][]map[string]string
	json.Unmarshal(ResponseBodyBytes, &responseText)
	fmt.Println(responseText["translations"][0]["text"])
	return ruText
}

func googleTranslate(toTranslate string) string {
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
