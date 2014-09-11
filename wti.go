package webtranslateit_go_client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
	"errors"
)

const DEFAULT_API_URL = "https://webtranslateit.com/api/"

var WtiNil = errors.New("wti: seems that wti pointer in project is nil, it is possible if you have not received the project from api call, please set wti pointer and try again")

type WebTranslateIt struct {
	apiUrl	string
	token	string
}

func NewWebTranslateIt(token string) *WebTranslateIt {
	return &WebTranslateIt{
		apiUrl:	DEFAULT_API_URL,
		token:	token,
	}
}

func (self *WebTranslateIt) GetToken() string {
	return self.token
}

func (self *WebTranslateIt) SetToken(token string) *WebTranslateIt {
	self.token = token
	return self
}

func (self *WebTranslateIt) GetApiUrl() string {
	return self.apiUrl
}

func (self *WebTranslateIt) SetApiUrl(apiUrl string) *WebTranslateIt {
	self.apiUrl = apiUrl
	return self
}

func (self *WebTranslateIt) GetProject() (*Project, error) {
	body, err := self.requestGet(fmt.Sprintf("%sprojects/%s.json", self.apiUrl, self.token))
	if err != nil {
		return nil, err
	}

	// Extract project
	var projectResponse ProjectResponse
	if err := json.Unmarshal(body, &projectResponse); err != nil {
		return nil, fmt.Errorf("wti: error decoding project data, %v", err)
	}

	// Save self pointer to the project
	projectResponse.Result.wti = self

	// Save project pointer to each file
	if len(projectResponse.Result.ProjectFiles) > 0 {
		for i, _ := range projectResponse.Result.ProjectFiles {
			projectResponse.Result.ProjectFiles[i].project = &projectResponse.Result
		}
	}

	return &projectResponse.Result, nil
}

func (self *WebTranslateIt) isError(body []byte) (*ErrorResponse, bool) {
	var errorResponse ErrorResponse
	if err := json.Unmarshal(body, &errorResponse); err == nil {
		if errorResponse.Message != "" {
			return &errorResponse, true
		}
	}
	return nil, false
}

func (self *WebTranslateIt) requestGet(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("wti: error sending request %q, %v", url, err)
	}
	defer response.Body.Close()

	// Get body of response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("wti: error reading the response data of request %q, %v", url, err)
	}

	// Check 200 status
	if response.StatusCode != 200 {
		// Check error response
		if errRes, isError := self.isError(body); isError {
			return nil, errRes
		}
		return nil, fmt.Errorf("wti: error unexpected status code %d", response.StatusCode)
	}

	return body, nil
}
