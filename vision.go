package vision

// visionClient is a client for vk cloud vision api
type visionClient struct {
	token   string
	persons persons
	objects objects
}

type persons struct {
	domain string
}

type objects struct {
	domain string
}

func NewVisionClient(token string) *visionClient {
	providerUrl := "https://smarty.mail.ru/api/v1"
	return &visionClient{
		token:   token,
		persons: persons{domain: providerUrl + "/persons"},
		objects: objects{domain: providerUrl + "/objects"},
	}
}
