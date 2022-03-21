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
		// TODO : Should saves to the other backend (do not wait for it to be display on local)
		proxy.PostComment(model.CommentToPost{Message: input.TextEn, Author: input.AuthorId})
		c.JSON(http.StatusCreated, input)
	}
}
