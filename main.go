package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Article struct {
	ID      uint64 `json:id`
	Title   string `json:string`
	Content string `json:string`
}

var articles = []Article{
	{ID: 1, Title: "title 1", Content: "content 1"},
	{ID: 2, Title: "title 2", Content: "content 2"},
	{ID: 3, Title: "title 3", Content: "content 3"},
	{ID: 4, Title: "title 4", Content: "content 4"},
}

func main() {
	r := gin.Default()
	// go glob not support **, reference: https://github.com/gin-gonic/gin/issues/2705
	r.LoadHTMLGlob("templates/*/*.tmpl")

	r.Static("/assets", "./assets")

	r.GET("/", index)
	r.GET("/articles/:id", showArticle)

	r.Run(":8080")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
		"articles": &articles,
		"content":  "test",
	})
}

func showArticle(c *gin.Context) {
	id := c.Param("id")
	var result Article
	for _, article := range articles {
		if strconv.FormatUint(article.ID, 10) == id {
			result = article
		}
	}

	c.HTML(http.StatusOK, "articles/show.tmpl", gin.H{
		"article": result,
	})
}
