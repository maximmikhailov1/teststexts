package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

func main() {
	url := "https://itexamanswers.net/ccna-2-v7-modules-1-4-switching-concepts-vlans-and-intervlan-routing-exam-answers.html" // ССЫЛКУ СЮДА
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{url},
		ParseFunc: parseTestsITE,
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()
}

func parseTestsAncient(g *geziyor.Geziyor, r *client.Response) {
	file, err := os.Create("cur.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	content := r.HTMLDoc.Find("ol")
	content.Find("ol > li").Each(func(i int, s *goquery.Selection) {
		question := s.Find("strong").Text()
		file.WriteString(question + "\n")
		fmt.Println(question)
		s.Find("li").Each(func(i int, s *goquery.Selection) {
			liText := s.Text()
			if len(s.Children().Text()) != 0 {
				liText = "**" + liText
			}
			file.WriteString("\t" + liText + "\n")
			fmt.Println("\t" + liText)
		})
		file.WriteString("\n\n")
		fmt.Println()
		fmt.Println()
	})

}

func parseTestsITE(g *geziyor.Geziyor, r *client.Response) {
	file, err := os.Create("cur.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	// parsedQuestions := make([]string, 0, 2)
	// parsedAnswers := make(map[int][]string)
	content := r.HTMLDoc.Find("div.post-single-content.box.mark-links.entry-content")
	content.Find("p + ul").Each(func(i int, s *goquery.Selection) {
		p := s.Prev()
		//вопрос всегда над ответом
		for len(p.Text()) == 0 {
			p = p.Prev()
		}
		ul := s
		pText := strings.ReplaceAll(p.Text(), "\n", " ")
		answers := make([]string, 0, 4)
		ul.Find("li").Each(func(i int, s *goquery.Selection) {
			liText := s.Text()
			if len(s.Children().Text()) != 0 || s.HasClass("correct_answer") {
				liText = "**" + liText
			}
			answers = append(answers, liText)
		})

		fmt.Println(pText)
		file.WriteString(pText + "\n")
		for _, v := range answers {
			fmt.Println("\t" + v)
			file.WriteString("\t" + v + "\n")
		}
		file.WriteString("\n\n")
		fmt.Println()
		fmt.Println()
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
