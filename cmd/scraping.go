package cmd

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

func scrapeurl(url string, id string) []Post {
	c := colly.NewCollector(
		colly.AllowedDomains("boards.4chan.org"),
	)
	var posts []Post

	c.OnHTML(".postContainer", func(h *colly.HTMLElement) {
		post := Post{}

		if f := h.ChildAttr("blockquote", "id"); strings.Contains(f, id) {
			post.OP = true
		} else {
			post.OP = false
		}

		post.Subject = h.ChildText(".subject")
		post.Image = h.ChildAttr("a.fileThumb", "href")
		post.Name = h.ChildText(".name")
		post.Date = h.ChildText(".dateTime")
		post.PostID = h.ChildAttr("blockquote", "id")
		post.Message = h.ChildText(".postMessage")

		posts = append(posts, post)
	})

	c.Visit(url)

	return posts
}
