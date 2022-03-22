package controller

import (
	"net/http"
	"owlcomments/model"
	"owlcomments/proxy"
	"owlcomments/repository"

	"github.com/gin-gonic/gin"
)

// GetComments return an array of comments matching an targetId
func GetComments(c *gin.Context) {
	targetId := c.Params.ByName("targetId")

	// if targetId not found in db
	if found, comment := repository.GetByTarget(targetId); found {
		if found, commentsReplies := repository.GetReplies(comment.Id); found {
			comment.Replies = commentsReplies
		} else {
			comment.Replies = make([]model.Comment, 0)
		}
		c.JSON(http.StatusOK, comment)
	} else {
		c.String(http.StatusNotFound, "")
	}
}

// PostComment save in the db the comment and
func PostComment(c *gin.Context) {

	// targetId := c.Params.ByName("targetId")

	// Validate input
	var input model.NewComment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.String(http.StatusBadRequest, "")
	} else {

		// Convertion of comments
		if input.TextEn == "" {
			// Convert fr to en
			input.TextEn = proxy.PostTradution(model.Traduction{TextToTrad: input.TextFr, Source: "fr", Target: "en"})
		} else if input.TextFr == "" {
			// Convert en to fr
			input.TextFr = proxy.PostTradution(model.Traduction{TextToTrad: input.TextEn, Source: "en", Target: "fr"})
		}

		// Send comment to faultybackend
		go proxy.PostComment(model.CommentToPost{Message: input.TextEn, Author: input.AuthorId})
		c.JSON(http.StatusCreated, input)
	}
}
