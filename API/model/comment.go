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
	Id          string `json:"id" bson:"id"`
	TextFr      string `json:"textFr" bson:"textFr"`
	TextEn      string `json:"textEn" bson:"textEn"`
	PublishedAt string `json:"publishedAt" bson:"publishedAt"`
	AuthorId    string `json:"authorId" binding:"required" bson:"authorId"`
	TargetId    string `json:"targetId" binding:"required" bson:"targetId"`
}
