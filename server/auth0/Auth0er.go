package auth0

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Auth0 struct {
	config Config
}

func New(config Config) *Auth0 {
	return &Auth0{config: config}
}

type userInfoReponse struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
}

func (a *Auth0) GetIDAndEmailFromToken(token string) (string, string, error) {
	// construct GET https://{Auth0 domain}/userinfo request
	request, err := http.NewRequest(http.MethodGet, "https://"+a.config.Domain+"/userinfo", nil)
	if err != nil {
		return "", "", fmt.Errorf("failed to create request: %w", err)
	}

	// set token as Authorization header
	request.Header.Set("Authorization", "Bearer "+token)

	// send request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", "", fmt.Errorf("failed to send request: %w", err)
	}
	defer response.Body.Close()

	// read body
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", "", fmt.Errorf("failed to read body: %w", err)
	}

	// extract ID and email from body
	var body userInfoReponse
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return "", "", fmt.Errorf("failed to unmarshal body: %w", err)
	}

	// return ID and name
	return body.Sub, body.Email, nil
}
