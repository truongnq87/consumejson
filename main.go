package main

import (
	"fmt"
	"time"

	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

type Author struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Attachment struct {
	Url         string `json:"url,omitempty"`
	MimeType    string `json:"mime_type,omitempty"`
	SizeInBytes int    `json:"size_in_bytes,omitempty"`
}

type Item struct {
	Title         string       `json:"title,omitempty"`
	ID            string       `json:"id,omitempty"`
	Url           string       `json:"url,omitempty"`
	ContentHTML   string       `json:"content_html,omitempty"`
	DatePublisher time.Time    `json:"date_publisher,omitempty"`
	Attachments   []Attachment `json:"attachments,omitempty"`
}

type Feed struct {
	Version     string  `json:"version"`
	UserComment string  `json:"user_comment"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	HomePageURL string  `json:"home_page_url,omitempty"`
	FeedURL     string  `json:"feed_url,omitempty"`
	Author      Author  `json:"author,omitempty"`
	Items       []*Item `json:"items,omitempty"`
}

func main() {
	r := gin.Default()
	r.POST("/bind-json", processJSON)
	r.Run()
}

func processJSON(c *gin.Context) {
	replacement := "*** ***"
	feed := Feed{}
	err := c.ShouldBindJSON(&feed)
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("got error when unmarshal err %v", err))
	}

	// Replace text exclude string in anchor tag (<a></a>)
	reg := regexp2.MustCompile(`(?!<a[^>]*?>)(Brent[ ]*Simmons|Brent|Chris[ ]*Parrish|Chris)(?![^<]*?</a>)`, regexp2.IgnoreCase)

	// If we want to change any kind of text even it has been presented in an A tag
	//reg := regexp2.MustCompile(`(?!<a[^>]*>)(Brent[ ]*Simmons|Brent|Chris[ ]*Parrish|Chris)(?![^<]*a>)`, regexp2.IgnoreCase)
	for _, item := range feed.Items {
		replacedContentHTML, err := reg.Replace(item.ContentHTML, replacement, -1, -1)
		if err != nil {
			c.AbortWithError(400, fmt.Errorf("got error when replace title"))
		}
		replacedTitle, err := reg.Replace(item.Title, replacement, -1, -1)
		if err != nil {
			c.AbortWithError(400, fmt.Errorf("got error when replace title"))
		}
		item.ContentHTML = replacedContentHTML
		item.Title = replacedTitle
	}

	replacedName, err := reg.Replace(feed.Author.Name, replacement, -1, -1)
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("got error when replace description"))
	}
	feed.Author.Name = replacedName

	replacedDescription, err := reg.Replace(feed.Description, replacement, -1, -1)
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("got error when replace description"))
	}
	feed.Description = replacedDescription

	// If you want to replace any new fields, just follow the above logic.

	c.JSON(200, gin.H{
		"result": feed,
	})
}
