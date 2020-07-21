package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func main() {

	http.HandleFunc("/getAllCategory", getAllCategory)
	http.HandleFunc("/getBookByCategory", getBookByCategory)
	http.HandleFunc("/getBookByTitleOrPriceRange", getBookByTitleOrPriceRange)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func getAllCategory(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("#zg_browseRoot", func(e *colly.HTMLElement) {
		metaTags := e.DOM.Find("ul").Find("ul").Find("li")
		metaTags.Each(func(_ int, s *goquery.Selection) {
			fmt.Println(s.Text())
		})
	})

	c.Visit("https://www.amazon.com/best-sellers-books-Amazon/zgbs/books/ref=zg_bs_nav_0")
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func getBookByCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func getBookByTitleOrPriceRange(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.amazon.com/s?k=nintendo+switch&ref=nb_sb_noss_1")
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println("Visiting", r.URL)

}
