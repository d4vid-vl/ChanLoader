package cmd

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

var domains = []string{"boards.4chan.org", "archived.moe"}

func scrapeurl(url string, id string) []Post {
	c := colly.NewCollector(
		colly.AllowedDomains(domains...),
	)
	var posts []Post

	c.OnHTML(".postContainer", func(h *colly.HTMLElement) {
		post := Post{}

		if f := h.ChildAttr("blockquote", "id"); strings.Contains(f, id) {
			post.OP = true
		} else {
			post.OP = false
		}

		subject_list := h.ChildTexts(".subject")
		if subject_list != nil {
			post.Subject = subject_list[0]
		} else {
			post.Subject = ""
		}

		name_list := h.ChildTexts(".name")
		date_list := h.ChildTexts(".dateTime")

		post.Media = h.ChildAttr("a.fileThumb", "href")
		post.Name = name_list[0]
		post.Date = date_list[0]
		post.PostID = h.ChildAttr("blockquote", "id")
		post.Message = h.ChildText(".postMessage")

		posts = append(posts, post)
	})

	c.Visit(url)

	return posts
}
