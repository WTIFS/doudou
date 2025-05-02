package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// WebhookRequest 定义请求结构
type WebhookRequest struct {
	Event     string `json:"event"`
	ClientKey string `json:"client_key"`
	Content   struct {
		Challenge int `json:"challenge"`
	} `json:"content"`
}

// ResponseBody 定义响应结构
type ResponseBody struct {
	Challenge int `json:"challenge"`
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req WebhookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 创建响应
	response := ResponseBody{
		Challenge: req.Content.Challenge,
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 返回响应
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/webhook", handleWebhook)

	log.Println("Server starting on :10080...")
	if err := http.ListenAndServe(":10080", nil); err != nil {
		log.Fatal(err)
	}
}
