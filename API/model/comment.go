package model

type Comment struct {
	Id          string    `json:"id"`
	TextFr      string    `json:"textFr"`
	TextEn      string    `json:"textEn"`
	PublishedAt string    `json:"publishedAt"`
	AuthorId    string    `json:"authorId"`
	TargetId    string    `json:"targetId"`
	Replies     []Comment `json:"replies"`
}

type NewComment struct {
	TextFr      string `json:"textFr" binding:"required"`
	TextEn      string `json:"textEn" binding:"required"`
	PublishedAt string `json:"publishedAt" binding:"required"`
	AuthorId    string `json:"authorId" binding:"required"`
	TargetId    string `json:"targetId" binding:"required"`
}
