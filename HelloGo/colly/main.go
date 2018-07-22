package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Article struct {
	ArticleTitle string
	URL          string
	CodeSnippets []string
}

func main() {

	c := colly.NewCollector(
		colly.AllowedDomains("edmundmartin.com"),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  ".*edmundmartin.*",
		Parallelism: 1,
		Delay:       1 * time.Second,
	})

	detailCollector := c.Clone()

	allArticles := []Article{}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML(`a[href]`, func(e *colly.HTMLElement) {
		foundURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Contains(foundURL, "python") {
			detailCollector.Visit(foundURL)
		} else {
			c.Visit(foundURL)
		}
	})

	detailCollector.OnHTML(`div.post-inner-content`, func(e *colly.HTMLElement) {
		fmt.Println("Scraping Content ", e.Request.URL.String())
		article := Article{}
		article.URL = e.Request.URL.String()
		article.ArticleTitle = e.ChildText("h1")

		e.ForEach("div.crayon-main", func(_ int, el *colly.HTMLElement) {
			codeSnip := el.ChildText("table.crayon-table")
			article.CodeSnippets = append(article.CodeSnippets, codeSnip)
		})
		fmt.Println("Found: ", article)
		allArticles = append(allArticles, article)
	})

	c.Visit("http://edmundmartin.com")
}
