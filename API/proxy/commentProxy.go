package proxy

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"owlcomments/model"
	"time"
)

var uriBackEnd string = "https://faulty-backend.herokuapp.com/on_comment"

func PostComment(comment model.CommentToPost) {
	client := &http.Client{
		Timeout: 3 * time.Second,
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
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("contextDeadLine: true")
		}
		if os.IsTimeout(err) {
			fmt.Println("IsTimeoutError: true")
		}
		fmt.Println(err)
	}

	for resp.StatusCode != 200 {
		resp, err = client.Do(req)
		fmt.Println("pause")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.StatusCode)
	}

	if resp.StatusCode == 418 {
		fmt.Println("Tea poted")
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
