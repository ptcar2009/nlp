package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set"
	"github.com/gocolly/colly"
)

func main() {
	subsSet := mapset.NewSet()
	subsFile, err := os.Open("subs.txt")

	file, err := os.Create("ptbrclosure.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("id1, id2, weight")
	if os.IsNotExist(err) {
		subsFile, _ = os.Create("subs.txt")
		for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
			mainCollector := colly.NewCollector()
			mainCollector.OnHTML(".line > a", func(e *colly.HTMLElement) {

				subsSet.Add(strings.ToLower(e.ChildText("b")))
			})
			mainCollector.OnRequest(func(r *colly.Request) {
				log.Println("Subs: " + r.URL.String())
			})
			mainCollector.OnHTML(".pagination > li > a", func(e *colly.HTMLElement) {
				if e.Text == "»" {
					e.Request.Visit(e.Attr("href"))
				}
			})
			mainCollector.Visit("https://dicionario.aizeta.com/verbetes/substantivo/" + string(letter))
		}
		for _, word := range subsSet.ToSlice() {
			subsFile.WriteString((word.(string)) + "\n")
		}
	} else {
		scanner := bufio.NewScanner(subsFile)
		for scanner.Scan() {
			subsSet.Add(strings.ToLower(scanner.Text()))
		}
	}
	log.Println(subsSet.Cardinality())
	for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
		mainCollector := colly.NewCollector()
		hyperCollector := mainCollector.Clone()

		mainCollector.OnHTML(".line > a", func(e *colly.HTMLElement) {
			hyperCollector.Visit(e.Attr("href"))
		})
		mainCollector.OnHTML(".pagination > li > a", func(e *colly.HTMLElement) {
			if e.Text == "»" {
				e.Request.Visit(e.Attr("href"))
			}
		})
		hyperCollector.OnHTML(".main", func(e *colly.HTMLElement) {
			hyper := []string{}
			definitions := e.ChildText("blockquote")
			split := strings.Split(definitions, ";")
			// log.Println(split)
			for _, def := range split {
				for _, word := range strings.Split(def, " ") {
					if !subsSet.Contains(strings.ToLower(word)) {
						continue
					}
					hyper = append(hyper, word)
					break
				}
			}
			title := e.ChildText("h1")
			splitTitle := strings.Split(title, " ")
			for _, word := range hyper {
				file.WriteString(strings.ToLower(splitTitle[len(splitTitle)-1]) + ", " + strings.ToLower(word) + ", 1\n")
			}

		})
		hyperCollector.OnRequest(func(r *colly.Request) {
			log.Println("Hyper: " + r.URL.String())
		})
		mainCollector.OnRequest(func(r *colly.Request) {
			log.Println("Main: " + r.URL.String())
		})
		mainCollector.Visit("https://dicionario.aizeta.com/verbetes/substantivo/" + string(letter))
	}
}
