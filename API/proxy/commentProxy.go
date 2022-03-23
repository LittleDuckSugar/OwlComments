package proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"owlcomments/model"
	"time"
)

var uriBackEnd string = "https://faulty-backend.herokuapp.com/on_comment"

// PostComment post received comment to falty_backend API
func PostComment(comment model.CommentToPost) {
	client := &http.Client{
		Timeout: 12 * time.Second,
	}

	jsonReq, err := json.Marshal(comment)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest(http.MethodPost, uriBackEnd, bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		// If timeout then redo request
		if os.IsTimeout(err) {
			PostComment(comment)
		}
		return
	}

	// If tea poted
	if resp.StatusCode == 418 {
		PostComment(comment)
		return
	}
	fmt.Println(resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var response interface{}

	json.Unmarshal(body, &response)

	fmt.Println(response)
}
