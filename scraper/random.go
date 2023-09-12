package scraper

import (
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)


type Definition struct{
	Title string
	Description string
	Example string
	AuthorDate string
}

func GetRandomEntries()  []Definition{

	res,err := http.Get("https://www.urbandictionary.com/random.php")
	if err!=nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("HTTP Error: %d %s",res.StatusCode,res.Status)
	}

	doc,err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}


	sevenDefs := make([]Definition,7)

	doc.Find("a.word").Each(func(i int, s *goquery.Selection) {
		sevenDefs[i].Title = s.Text()
	})

	doc.Find("div.meaning").Each(func(i int, s *goquery.Selection) {
		sevenDefs[i].Description = s.Text()
	})

	doc.Find("div.example").Each(func(i int, s *goquery.Selection) {
		sevenDefs[i].Example = s.Text()
	})

	doc.Find("div.contributor").Each(func(i int, s *goquery.Selection) {
		sevenDefs[i].AuthorDate = s.Text()
	})

	return sevenDefs
}