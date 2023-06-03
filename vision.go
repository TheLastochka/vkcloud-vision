package vision

import "net/http"

// visionClient is a client for vk cloud vision api
type visionClient struct {
	token   string
	client  *http.Client
	persons persons
	objects objects
}

type persons struct {
	domain string
}

type objects struct {
	domain string
}

func NewVisionClient(client *http.Client, token string) *visionClient {
	providerUrl := "https://smarty.mail.ru/api/v1"
	return &visionClient{
		token:   token,
		client:  client,
		persons: persons{domain: providerUrl + "/persons"},
		objects: objects{domain: providerUrl + "/objects"},
	}
}
