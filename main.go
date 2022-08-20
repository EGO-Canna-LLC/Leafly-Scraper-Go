package main

import (
	"github.com/EGO-Canna-LLC/Leafly-Scrapper-Go/cache"
)

func main() {
	url := "https://www.leafly.com/brands/ego-canna/products/ego-canna-thco-d9-designer-gummies-10ct-gummies"
	err := cache.SaveHTML(url, "./reviews.html")
	if err != nil {
		return
	}
	// extract <script type="application/ld+json"> from example.html
	// and save to example.json
	err = cache.ExtractJSON("reviews.html", "reviews.json")
	if err != nil {
		return
	}
}
