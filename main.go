package main

import (
	"fmt"
	"os"
	"strings"
	transateITE "teststexts/translateITE"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

var cliOutput bool = false     // false - без вывода в консоль; true с выводом
var displayCounter bool = true // false - без счётчика; true c счётчиком

var translator transateITE.Translator

//mb stoit TO PARSE
// https://itexamanswers.net/ccna-2-v7-0-final-exam-answers-full-switching-routing-and-wireless-essentials.html
// https://itexamanswers.net/ccna-1-v5-1-v6-0-practice-final-exam-answers-100-full.html
// https://itexamanswers.net/7-4-4-module-quiz-dhcpv4-answers.html
// https://itexamanswers.net/ewan-v4-chapter-7-check-your-understanding-ip-addressing-services.html
// https://itexamanswers.net/ccna-4-exploration-v4-0-chapter-7-quiz-answers.html
// https://itexamanswers.net/ccna-2-v7-modules-5-6-redundant-networks-exam-answers.html
// https://itexamanswers.net/ccna-3-final-exam-answers-v5-0-3-v6-0-scaling-networks.html

// PARSED
// https://itexamanswers.net/5-4-2-module-quiz-stp-answers.html
// https://itexamanswers.net/ccna-3-practice-final-exam-answers-v5-0-3-v6-0-full-100.html
// https://itexamanswers.net/ccna-3-pretest-exam-answers-v5-0-3-v6-0-full-100.html
// https://itexamanswers.net/ccna-2-v7-modules-5-6-redundant-networks-exam-answers.html
// https://itexamanswers.net/chapter-3-quiz-advanced-spanning-tree-tuning-answers-ccnpv8-encor.html
// https://itexamanswers.net/ccna-200-301-certification-practice-exam-answers-ensa-v7-0.html
// https://itexamanswers.net/ccna-2-v7-modules-1-4-switching-concepts-vlans-and-intervlan-routing-exam-answers.html

func main() {
	url := "https://itexamanswers.net/5-4-2-module-quiz-stp-answers.html" // ССЫЛКУ СЮДА
	translator.TranslateMech = "yandex"                                   // yandex или google
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{url},
		ParseFunc: parseTestsITE,
	}).Start()
}

func parseTestsAncient(g *geziyor.Geziyor, r *client.Response) {
	fileEn, err := os.Create("curEN.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer fileEn.Close()
	fileRu, err := os.Create("curRU.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer fileRu.Close()

	content := r.HTMLDoc.Find("ol")
	work := content.Find("ol > li")
	work.Each(func(i int, s *goquery.Selection) {
		if displayCounter {
			fmt.Printf("%d/%d\n", i+1, work.Length()+1)
		}
		question := s.Find("strong").Text()
		fileEn.WriteString(question + "\n")
		fileRu.WriteString(translator.TranslateITE(question) + "\n")
		if cliOutput {
			fmt.Println(question)
		}
		s.Find("li").Each(func(i int, s *goquery.Selection) {
			liText := s.Text()
			if len(s.Children().Text()) != 0 {
				liText = "**" + liText
			}
			fileEn.WriteString("\t" + liText + "\n")
			fileRu.WriteString("\t" + translator.TranslateITE(liText) + "\n")
			if cliOutput {
				fmt.Println("\t" + liText)
			}
		})
		fileEn.WriteString("\n\n")
		if cliOutput {
			fmt.Println()
			fmt.Println()
		}
	})

}

func parseTestsITE(g *geziyor.Geziyor, r *client.Response) {
	fileEn, err := os.Create("curEN.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer fileEn.Close()
	fileRu, err := os.Create("curRU.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer fileRu.Close()
	// parsedQuestions := make([]string, 0, 2)
	// parsedAnswers := make(map[int][]string)
	content := r.HTMLDoc.Find("div.post-single-content.box.mark-links.entry-content")
	work := content.Find("p + ul")
	work.Each(func(i int, s *goquery.Selection) {
		if displayCounter {
			fmt.Printf("%d/%d\n", i+1, work.Length())
		}
		p := s.Prev()
		//вопрос всегда над ответом
		for len(p.Text()) == 0 {
			p = p.Prev()
		}
		ul := s
		pText := strings.ReplaceAll(p.Text(), "\n", " ")
		pTextRu := translator.TranslateITE(pText)
		answers := make([]string, 0, 4)
		ul.Find("li").Each(func(i int, s *goquery.Selection) {
			liText := s.Text()
			if len(s.Children().Text()) != 0 || s.HasClass("correct_answer") {
				liText = "**" + liText
			}
			answers = append(answers, liText)
		})

		if cliOutput {
			fmt.Println(pText)
		}
		fileEn.WriteString(pText + "\n")
		if cliOutput {
			fmt.Println(pTextRu)
		}
		fileRu.WriteString(pTextRu + "\n")
		for _, answer := range answers {
			if cliOutput {
				fmt.Println("\t" + answer)
			}
			fileEn.WriteString("\t" + answer + "\n")
			fileRu.WriteString("\t" + translator.TranslateITE(answer) + "\n")
		}
		fileEn.WriteString("\n\n")
		fileRu.WriteString("\n\n")
		if cliOutput {
			fmt.Println()
			fmt.Println()
		}
	})

	// for _, v := range peshki.EachIter() {
	// 	fmt.Println(v.Text())
	// }
	// for _, v := range ulki.EachIter() {
	// 	fmt.Println(v.Text())
	// }
	// content.Find("p").Each(func(i int, s *goquery.Selection) {
	// 	var peshka = s.Text()
	// 	ulki := s.SiblingsFiltered("ul")
	// 	if ulki.Is("ul") {
	// 		fmt.Println("p", peshka)
	// 		fmt.Println("ul", ulki.Text())
	// 	}
	// 	if len(peshka) != 0 && unicode.IsDigit(rune(peshka[0])) && {

	// 		parsedQuestions = append(parsedQuestions, peshka)
	// 	}
	// 	// fmt.Println(peshka)
	// })

	// questionNumber := 0
	// content.Find("ul").Each(func(i int, s *goquery.Selection) {
	// 	parsedPossibleAnswers := make([]string, 0, 2)
	// 	s.Find("li").Each(func(i int, s *goquery.Selection) {
	// 		liText := s.Text()
	// 		// fmt.Println()
	// 		if len(s.Children().Text()) != 0 {
	// 			liText = "**" + liText
	// 		}

	// 		parsedPossibleAnswers = append(parsedPossibleAnswers, liText)
	// 	})
	// 	parsedAnswers[questionNumber] = parsedPossibleAnswers
	// 	questionNumber++
	// })

	// for i := 0; i < len(parsedQuestions); i++ {
	// 	fmt.Println(parsedQuestions[i])
	// }
	// for i := 0; i < len(parsedQuestions); i++ {
	// 	fmt.Println(parsedQuestions[i])
	// 	file.WriteString(parsedQuestions[i] + "\n")
	// 	for j := 0; j < len(parsedAnswers[i]); j++ {
	// 		fmt.Println(parsedAnswers[i][j])
	// 		file.WriteString("\t" + parsedAnswers[i][j] + "\n")
	// 	}
	// }
}
