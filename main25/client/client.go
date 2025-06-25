package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	testBookQuery()
	testCommentSubmit()
}

func testBookQuery() {
	resp, err := http.Get("http://localhost:8080/book?title=三体")
	if err != nil {
		fmt.Println("GET请求出错:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("图书查询响应:", string(body))
}

func testCommentSubmit() {
	data := map[string]string{
		"user":    "小明",
		"comment": "这本书真棒！",
	}
	jsonData, _ := json.Marshal(data)

	resp, err := http.Post(
		"http://localhost:8080/comment",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println("POST请求出错:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("评论提交响应:", string(body))
}
