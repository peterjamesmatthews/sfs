package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type auth0UserInfoResponse struct {
	Sub  string `json:"sub"`
	Name string `json:"name"`
}

func (a *App) getIDAndNameFromToken(token string) (id string, name string, err error) {
	// construct GET https://{Auth0 domain}/userinfo request
	request, err := http.NewRequest(http.MethodGet, "https://"+a.config.AUTH0_DOMAIN+"/userinfo", nil)
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
	var body auth0UserInfoResponse
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return "", "", fmt.Errorf("failed to unmarshal body: %w", err)
	}

	// return ID and name
	return body.Sub, body.Name, nil
}
