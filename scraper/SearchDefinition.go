package scraper

import (
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)


// TODO: This function has no handlers when there's no definition.

func SearchDefinition(searchTerm string,numDefs int) []Definition {

	if numDefs<1 || numDefs>7 {
		log.Fatal("Number of definitions you want to see can be 1 at least and 7 at most.")
	}

	var url string = "https://www.urbandictionary.com/define.php?term=" + searchTerm

	res,err := http.Get(url)

	if err != nil {
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

	// TODO: There should be a check here if any definitions exist at all!!

	defSlice := make([]Definition,numDefs)

	doc.Find("a.word").Each(func(i int, s *goquery.Selection) {
		if i>=numDefs {return}
		defSlice[i].Title = s.Text()
		
	})

	doc.Find("div.meaning").Each(func(i int, s *goquery.Selection) {
		if i>=numDefs {return}
		defSlice[i].Description = s.Text()
		
	})

	doc.Find("div.example").Each(func(i int, s *goquery.Selection) {
		if i>=numDefs {return}
		defSlice[i].Example = s.Text()
		
	})

	doc.Find("div.contributor").Each(func(i int, s *goquery.Selection) {
		if i>=numDefs {return}
		defSlice[i].AuthorDate = s.Text()
		
	})

	return defSlice 
}