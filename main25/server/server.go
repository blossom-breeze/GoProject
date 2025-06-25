package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		if title == "" {
			title = "未知图书"
		}
		response := fmt.Sprintf("您正在查询图书：《%s》", title)
		fmt.Fprint(w, response)
	})
	http.HandleFunc("/comment", func(w http.ResponseWriter, r *http.Request) {
		type Comment struct {
			User    string `json:"user"`
			Comment string `json:"comment"`
		}

		var c Comment
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, "无效的JSON数据", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"message": "评论提交成功",
			"user":    c.User,
			"comment": c.Comment,
		}
		json.NewEncoder(w).Encode(response)
	})
	fmt.Println("服务启动于 http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
