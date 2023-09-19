package scraper

import (
	"log"
	"net/http"

	"strings"
	"github.com/PuerkitoBio/goquery"
)



func SearchDefinition(searchTerm string,numDefs int) []Definition {

	if numDefs<1 || numDefs>7 {
		log.Fatal("Number of definitions you want to see can be 1 at least and 7 at most.")
	}

	var processedSearchTerm string

	if strings.Contains(searchTerm," ") {

		for _,word := range strings.Fields(searchTerm) {
			processedSearchTerm =  processedSearchTerm + "+" + word
			
		}
	} else {
		processedSearchTerm = searchTerm
	}

	
	var url string = "https://www.urbandictionary.com/define.php?term=" + processedSearchTerm

	res,err := http.Get(url)

	if err != nil {
		
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {

		// If page not found, function returns an empty slice to show
		// that there are no definitions.
		if res.StatusCode == 404 {
			return make([]Definition,0)
		}


		log.Println("Bok yemiş bülbül??")
		log.Fatalf("HTTP Error: %d %s",res.StatusCode,res.Status)
	}

	doc,err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		
		log.Fatal(err)
	}


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