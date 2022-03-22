package controller

import (
	"fmt"
	"net/http"
	"owlcomments/model"
	"owlcomments/proxy"
	"owlcomments/repository"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/rs/xid"
)

var fakeTargets []string = []string{}

// GetComments return an array of comments matching an targetId
func GetComments(c *gin.Context) {
	targetId := c.Params.ByName("targetId")

	// if targetId found in db
	if found, comment := repository.GetCommentByTargetId(targetId); found {
		c.JSON(http.StatusOK, comment)
	} else {
		c.String(http.StatusNotFound, "")
	}
}

// PostComment save in the db the comment and
func PostComment(c *gin.Context) {

	targetId := c.Params.ByName("targetId")

	fakeTargets = repository.UpdateFakeTargets()

	// Does the asked target exist ?
	for _, target := range fakeTargets {
		if target == targetId {

			// Validate input
			var input model.NewComment
			if err := c.ShouldBindJSON(&input); err != nil {
				c.String(http.StatusBadRequest, "")
			} else {

				if targetId == input.TargetId {
					// Traduction of comments
					if input.TextEn == "" {
						// Traduction fr to en
						input.TextEn = proxy.PostTradution(model.Traduction{TextToTrad: input.TextFr, Source: "fr", Target: "en"})
					} else if input.TextFr == "" {
						// Traduction en to fr
						input.TextFr = proxy.PostTradution(model.Traduction{TextToTrad: input.TextEn, Source: "en", Target: "fr"})
					}

					// Save the PublishedAt attributes if none where provided
					if input.PublishedAt == "" {
						input.PublishedAt = strconv.FormatInt(time.Now().UTC().Unix(), 10)
					}

					// Generate id of the comment
					input.Id = "Comment-" + xid.New().String()

					// Send comment to faultybackend
					go proxy.PostComment(model.CommentToPost{Message: input.TextEn, Author: input.AuthorId})

					// Saving into db
					repository.PostComment(input)

					c.JSON(http.StatusCreated, input)
				} else {
					fmt.Println("targetId is not the same in the path than in the request")
				}
			}
			return
		}
	}

	fmt.Println("No matching targets")

}
