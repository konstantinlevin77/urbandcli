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
	Upvotes string
	Downvotes string
}

func GetRandomDefinition() Definition{


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

	randomDefinition := Definition{}

	randomDefinition.Title = doc.Find("a.word").First().Text()
	randomDefinition.Description = doc.Find("div.meaning").First().Text()
	randomDefinition.Example = doc.Find("div.example").First().Text()
	randomDefinition.AuthorDate = doc.Find("div.contributor").First().Text()
	

	return randomDefinition
}

