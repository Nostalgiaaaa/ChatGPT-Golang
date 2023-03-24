package model

import "github.com/sashabaranov/go-openai"

type ChatRequest struct {
	Prompt  string             `json:"prompt"`
	Options ChatRequestOptions `json:"options,omitempty"`
}
type ChatRequestOptions struct {
	ParentMessageId string `json:"parentMessageId"`
}

type VerifyRequest struct {
	Token string `json:"token"`
}

type ChatResponse struct {
	Role            string                              `json:"role"`
	Id              string                              `json:"id"`
	ParentMessageId string                              `json:"parentMessageId"`
	Delta           string                              `json:"delta"`
	Text            string                              `json:"text"`
	Detail          openai.ChatCompletionStreamResponse `json:"detail"`
}

type ChatConfig struct {
	Message string         `json:"message"`
	Data    ChatConfigData `json:"data"`
	Status  string         `json:"status"`
}
type ChatConfigData struct {
	APIModel     string `json:"apiModel"`
	ReverseProxy string `json:"reverseProxy"`
	TimeoutMs    int    `json:"timeoutMs"`
	SocksProxy   string `json:"socksProxy"`
	HttpsProxy   string `json:"httpsProxy"`
	Balance      string `json:"balance"`
}
