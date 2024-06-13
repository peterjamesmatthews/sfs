package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Auth0 interface {
	// GetIDAndNameFromToken gets the ID and name of a user from their opaque token.
	GetIDAndNameFromToken(token string) (id string, name string, err error)
}

func (a *App) GetIDAndNameFromToken(token string) (id string, name string, err error) {
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
	body := struct {
		ID   string `json:"sub"`
		Name string `json:"name"`
	}{}
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return "", "", fmt.Errorf("failed to unmarshal body: %w", err)
	}

	// return ID and name
	return body.ID, body.Name, nil
}