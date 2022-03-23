package model

// CommentToPost stores attributes asked by faulty_backend
type CommentToPost struct {
	Message string `json:"message"`
	Author  string `json:"author"`
}
