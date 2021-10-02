package main

import (
	"log"
	"os"

	"sitemap-generator/sitemap"
)

func main() {
	xmlSitemapFile := os.Getenv("XML_SITEMAP_FILE")

	stm := sitemap.New(xmlSitemapFile)

	err := stm.Add(sitemap.URL{Location: "https://exampe.com"})
	if err != nil {
		log.Fatal(err)
	}

	err = stm.Add(sitemap.URL{
		Location: "https://exampe.com/about.html",
		LastModified: "2021-02-08T18:37:15+01:00",
	})
	if err != nil {
		log.Fatal(err)
	}

	err = stm.Add(sitemap.URL{
		Location: "https://exampe.com/products.html",
		ChangeFrequently: "monthly",
		LastModified: "2021-02-08T18:37:15+01:00",
	})
	if err != nil {
		log.Fatal(err)
	}

	err = stm.Add(sitemap.URL{
		Location: "https://exampe.com/articles.html",
		ChangeFrequently: "weekly",
		LastModified: "2021-02-08T18:37:15+01:00",
		Priority: 0.5,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err = stm.Write(); err != nil {
		log.Fatal(err)
	}
}
