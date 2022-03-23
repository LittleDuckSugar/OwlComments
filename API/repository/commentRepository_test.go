package repository

import (
	"owlcomments/model"
	"reflect"
	"testing"
)

func TestGetComments(t *testing.T) {
	foundWant := true
	commentWant := []model.Comment{
		{
			Id:          "Comment-kjh784fgevdhhdwhh7563",
			TextFr:      "Bonjour ! je suis un commentaire.",
			TextEn:      "Hi ! Im a comment.",
			PublishedAt: "1639477064",
			AuthorId:    "User-kjh784fgevdhhdwhh7563",
			TargetId:    "Photo-bdgetr657434hfggrt8374",
			Replies: []model.Comment{
				{
					Id:          "Comment-1234abcd",
					TextFr:      "Je suis une réponse au commentaire",
					TextEn:      "Im a reply!",
					PublishedAt: "1639477064",
					AuthorId:    "User-5647565dhfbdshs",
					TargetId:    "Comment-kjh784fgevdhhdwhh7563",
					Replies:     nil,
				},
				{
					Id:          "Comment-c8t3oian3s3a0u64l8ag",
					TextFr:      "Hé ! Voici un autre commentaire",
					TextEn:      "Hey ! Here is yet an other comment",
					PublishedAt: "1647983689",
					AuthorId:    "Antoine",
					TargetId:    "Comment-kjh784fgevdhhdwhh7563",
					Replies: []model.Comment{
						{
							Id:          "Comment-c8t3q4qn3s3a9jlfj4mg",
							TextFr:      "Un commentaire sur un commentaire !",
							TextEn:      "A comment on a comment !",
							PublishedAt: "1647983891",
							AuthorId:    "Antoine",
							TargetId:    "Comment-c8t3oian3s3a0u64l8ag",
							Replies: []model.Comment{
								{
									Id:          "Comment-c8t3s6qn3s3ai9r3vgn0",
									TextFr:      "Un commentaire sur un commentaire sur un commentaire ?",
									TextEn:      "A comment on a comment on a comment ?",
									PublishedAt: "1647984155",
									AuthorId:    "Antoine",
									TargetId:    "Comment-c8t3q4qn3s3a9jlfj4mg",
									Replies:     nil,
								},
							},
						},
					},
				},
			},
		},
	}

	foundGot, commentGot := GetCommentByTargetId("Comment-kjh784fgevdhhdwhh7563")

	if foundWant != foundGot {
		t.Error("got", foundGot, "wanted", foundWant)
	}
	if !reflect.DeepEqual(commentGot, commentWant) {
		t.Errorf("got %q, wanted %q", commentGot, commentWant)
	}
}
