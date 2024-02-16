package backend

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/mmcdole/gofeed"
	"strings"
)

type Data struct {
	Title  string
	Author string
	Rating string
	Review string
	Date   string
}

// The function `RSSToJSON` parses an RSS feed from Letterboxd and extracts reviews, converting them
// into a JSON format.
func RSSToJSON() []Data{
	fp := gofeed.NewParser()

	feed, err := fp.ParseURL("https://letterboxd.com/kyouko_/rss/")
	if err != nil {
		panic(err)
	}

	//Getting only the reviews from the RSS feed
	var reviews []Data
	for _, item := range feed.Items {
		if strings.Contains(item.GUID, "review") {
			newAppend := Data{
				Title:  item.Extensions["letterboxd"]["filmTitle"][0].Value,
				Author: item.Author.Name,
				Rating: item.Extensions["letterboxd"]["memberRating"][0].Value,
				Review: item.Description,
				Date:   item.Published,
			}
			reviews = append(reviews, newAppend)
			// fmt.Println("------------------------------------")
			// spew.Dump(item.Extensions["letterboxd"]["filmTitle"])
			// spew.Dump(item)
			// fmt.Println("------------------------------------")
		}
	}
	for _, v := range reviews {
		fmt.Println("------------------------------------")
		spew.Dump(v)
		fmt.Println("------------------------------------")
	}
	return reviews
}
