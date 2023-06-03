package vision

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

type ResponseError struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}

// This structure is needed to get an array of images
// from meta that has no type (interface{})
type structWithImages struct {
	Images []struct {
		Name string `json:"name"`
	} `json:"images"`
}

// buildBody builds multipart/form-data request body
// meta - struct with request params (meta)
// boundary - boundary for multipart/form-data
// if meta contains images, they would be added to request body
func buildBody(meta interface{}, boundary string) io.Reader {
	metaBytes, _ := json.Marshal(meta)
	var strct structWithImages
	err := json.Unmarshal(metaBytes, &strct)
	if err != nil {
		log.Fatal("Error unmarshaling meta: ", err)
	}
	var imageNames []string
	for _, image := range strct.Images {
		imageNames = append(imageNames, image.Name)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.SetBoundary(boundary)

	metaField := textproto.MIMEHeader{}
	metaField.Set("Content-Disposition", `form-data; name="meta"`)
	part, _ := writer.CreatePart(metaField)
	part.Write(metaBytes)

	for _, imageName := range imageNames {
		imageField := textproto.MIMEHeader{}
		imageField.Set("Content-Disposition", `form-data; name="`+imageName+`"; filename="`+imageName+`"`)
		imageField.Set("Content-Type", "image/jpeg")
		imageField.Set("Content-Transfer-Encoding", "binary")

		part, _ := writer.CreatePart(imageField)
		file, _ := os.Open(imageName)
		defer file.Close()
		io.Copy(part, file)
	}
	writer.Close()

	return body
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// sendPostRequest sends POST request to requestUrl with meta as body
// meta - struct with request params (meta)
// if meta contains images, they would be added to request body
func sendPostRequest(client *http.Client, requestUrl string, meta interface{}) []byte {
	boundary := randStringBytes(30)
	requestBody := buildBody(meta, boundary)
	req, err := http.NewRequest("POST", requestUrl, requestBody)
	if err != nil {
		log.Fatal("Error creating request: ", err)
	}
	req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}
	return body
}

// unmarshal body of response
// respOk - pointer to struct with ok response
// respErr - pointer to struct with error response
func unmarshalResponse(body []byte, respOk interface{}, respErr interface{}) error {
	err := json.Unmarshal(body, &respOk)
	if err != nil {
		err = json.Unmarshal(body, &respErr)
		if err != nil {
			return err
		}
	}
	return nil
}
