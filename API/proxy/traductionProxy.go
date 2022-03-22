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

var uriTraductionEndpoint string = "https://translate.argosopentech.com/translate"

func PostTradution(toTranslate model.Traduction) string {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	jsonReq, err := json.Marshal(toTranslate)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest(http.MethodPost, uriTraductionEndpoint, bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		// If timeout then redo request
		if os.IsTimeout(err) {
			textTranslated := PostTradution(toTranslate)
			return textTranslated
		}
	}

	// If tea poted
	if resp.StatusCode == 418 {
		textTranslated := PostTradution(toTranslate)
		return textTranslated
	}
	fmt.Println(resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var response model.TraductionResults

	json.Unmarshal(body, &response)

	return response.TextTranslated
}
