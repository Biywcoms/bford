package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const Token = "b5dee807f02c4eccc01525e7bcd82b02214c0576573549d970a77b"
const ApiUrl = "http://api.tushare.pro"

type PostBody struct {
	TOKEN    string            `json:"token"`
	API_NAME string            `json:"api_name"`
	PARAMS   map[string]string `json:"params"`
	FIELDS   string            `json:"fields"`
}

func main() {

	post := PostBody{Token, "", nil, ""}
	post.API_NAME = "fina_mainbz"
	post.PARAMS = map[string]string{
		"ts_code":    "000001.SZ",
		"start_date": "20000101",
		"end_date":   "20001231",
	}
	post.FIELDS = ""

	jsonBytes, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", ApiUrl, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(body), &data); err == nil {
		fmt.Println(data["data"])
	}

	//	fmt.Println("response Body:", string(body))

}
