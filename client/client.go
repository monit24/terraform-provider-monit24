package client

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ResourceNotFound struct {
	message string
}

func (r ResourceNotFound) Error() string {
	return "Resource not found: " + r.message
}

const baseURL = "https://api.monit24.pl/v3"

type Client struct {
	client    *http.Client
	basicAuth string
	ownerID   int
}

func NewBasicAuthClient(ctx context.Context, user string, password string) (Client, error) {
	client := &http.Client{}

	c := Client{
		basicAuth: basicAuth(user, password),
		client:    client,
	}

	account, err := c.getMyAccount(ctx)
	if err != nil {
		return Client{}, err
	}

	if account.ParentAccountID != nil {
		c.ownerID = *account.ParentAccountID
	} else {
		c.ownerID = account.ID
	}

	return c, nil
}

type AccountResponse struct {
	ID              int  `json:"id"`
	ParentAccountID *int `json:"parent_account_id"`
}

func (c Client) getMyAccount(ctx context.Context) (AccountResponse, error) {
	resp, err := c.get(ctx, "/accounts/my_account")
	if err != nil {
		return AccountResponse{}, err
	}

	var account AccountResponse
	err = json.Unmarshal(resp, &account)
	if err != nil {
		return AccountResponse{}, err
	}

	return account, nil
}

func (c Client) OwnerID() int {
	return c.ownerID
}

func (c Client) get(ctx context.Context, path string) ([]byte, error) {
	return c.request(ctx, "GET", path, http.StatusOK)
}

func (c Client) post(ctx context.Context, path string, payload interface{}) ([]byte, error) {
	return c.requestWithPayload(ctx, "POST", path, payload, http.StatusCreated)
}

func (c Client) put(ctx context.Context, path string, payload interface{}) error {
	_, err := c.requestWithPayload(ctx, "PUT", path, payload, http.StatusNoContent)
	return err
}

func (c Client) delete(ctx context.Context, path string) error {
	_, err := c.request(ctx, "DELETE", path, http.StatusNoContent)
	return err
}

func (c Client) request(ctx context.Context, method string, path string, expectedStatusCode int) ([]byte, error) {
	return c.rawRequest(ctx, method, path, nil, expectedStatusCode)
}

func (c Client) requestWithPayload(ctx context.Context, method string, path string, payload interface{}, expectedStatusCode int) ([]byte, error) {
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(payload)
	if err != nil {
		return nil, err
	}

	return c.rawRequest(ctx, method, path, body, expectedStatusCode)
}

func (c Client) rawRequest(ctx context.Context, method string, path string, requestBody io.Reader, expectedStatusCode int) ([]byte, error) {
	req, err := http.NewRequest(method, baseURL+path, requestBody)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+c.basicAuth)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != expectedStatusCode {
		if resp.StatusCode == http.StatusNotFound {
			return nil, ResourceNotFound{string(body)}
		}
		return nil, fmt.Errorf("unexpected HTTP status code: %v message: %v", resp.StatusCode, string(body))
	}

	return body, nil
}

func basicAuth(user, password string) string {
	auth := user + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
