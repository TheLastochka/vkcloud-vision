package vision

import (
	"log"
)

type ImageMeta struct {
	Name     string `json:"name"`
	PersonId int    `json:"person_id"`
}

//==========================================================
//==========================SET=============================
//==========================================================

type MetaSet struct {
	Space  string      `json:"space"`
	Images []ImageMeta `json:"images"`
}

type ResponseSetOk struct {
	Status int `json:"status"`
	Body   struct {
		Objects []struct {
			Name   string `json:"name"`
			Status int    `json:"status"`
			Error  string `json:"error"`
		} `json:"objects"`
	} `json:"body"`
}

func (vision *visionClient) PersonsSet(meta MetaSet) (*ResponseSetOk, *ResponseError) {
	requestUrl := vision.persons.domain + "/set?" + "oauth_provider=mcs&oauth_token=" + vision.token
	body := sendPostRequest(vision.client, requestUrl, meta)

	var respOk *ResponseSetOk
	var respErr *ResponseError
	err := unmarshalResponse(body, &respOk, &respErr)
	if err != nil {
		log.Panicln("Error unmarshaling body: " + string(body))
	}
	return respOk, respErr
}

//==========================================================
//========================RECOGNIZE=========================
//==========================================================

type MetaRecognize struct {
	Space           string      `json:"space"`
	Images          []ImageMeta `json:"images"`
	CreateNew       bool        `json:"create_new"`
	UpdateEmbedding bool        `json:"update_embedding"`
}

type ResponseRecognizeOk struct {
	Status         int  `json:"status"`
	AliasesChanged bool `json:"aliases_changed"`
	Body           struct {
		Objects []struct {
			Status         int    `json:"status"`
			Error          string `json:"error"`
			Name           string `json:"name"`
			CountByDensity int    `json:"count_by_density"`
			Persons        []struct {
				Tag         string   `json:"tag"`
				Coord       []int    `json:"coord"`
				Aliases     []string `json:"aliases"`
				Confidence  float64  `json:"confidence"`
				Similarity  float64  `json:"similarity"`
				Awesomeness float64  `json:"awesomeness"`

				Sex     string  `json:"sex"`
				Age     float64 `json:"age"`
				Emotion string  `json:"emotion"`
				Valence float64 `json:"valence"`
				Arousal float64 `json:"arousal"`
			} `json:"persons"`
		} `json:"objects"`
	} `json:"body"`
}

func (vision *visionClient) PersonsRecognize(meta MetaRecognize) (*ResponseRecognizeOk, *ResponseError) {
	requestUrl := vision.persons.domain + "/recognize?" + "oauth_provider=mcs&oauth_token=" + vision.token
	body := sendPostRequest(vision.client, requestUrl, meta)

	var respOk *ResponseRecognizeOk
	var respErr *ResponseError
	err := unmarshalResponse(body, &respOk, &respErr)
	if err != nil {
		log.Panicln("Error unmarshaling body: " + string(body))
	}
	return respOk, respErr
}

//==========================================================
//==========================DELETE==========================
//==========================================================

type MetaDelete struct {
	Space  string      `json:"space"`
	Images []ImageMeta `json:"images"`
}

type ResponseDeleteOk struct {
	Status int `json:"status"`
	Body   struct {
		Objects []struct {
			Name   string `json:"name"`
			Status int    `json:"status"`
			Error  string `json:"error"`
		} `json:"objects"`
	} `json:"body"`
}

func (vision *visionClient) PersonsDelete(meta MetaDelete) (*ResponseDeleteOk, *ResponseError) {
	requestUrl := vision.persons.domain + "/delete?" + "oauth_provider=mcs&oauth_token=" + vision.token
	body := sendPostRequest(vision.client, requestUrl, meta)

	var respOk *ResponseDeleteOk
	var respErr *ResponseError
	err := unmarshalResponse(body, &respOk, &respErr)
	if err != nil {
		log.Panicln("Error unmarshaling body: " + string(body))
	}
	return respOk, respErr
}

//==========================================================
//=========================TRUNCATE=========================
//==========================================================

type MetaTruncate struct {
	Space string `json:"space"`
}

type ResponseTruncateOk struct {
	Status int      `json:"status"`
	Body   struct{} `json:"body"`
}

func (vision *visionClient) PersonsTruncate(meta MetaTruncate) (*ResponseTruncateOk, *ResponseError) {
	requestUrl := vision.persons.domain + "/truncate?" + "oauth_provider=mcs&oauth_token=" + vision.token
	body := sendPostRequest(vision.client, requestUrl, meta)

	var respOk *ResponseTruncateOk
	var respErr *ResponseError
	err := unmarshalResponse(body, &respOk, &respErr)
	if err != nil {
		log.Panicln("Error unmarshaling body: " + string(body))
	}
	return respOk, respErr
}
