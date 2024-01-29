package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	ServerURL string
	Token     string
}

func (c *Client) Get(path string) (map[string]string, error) {
	req, err := http.NewRequest("GET", c.ServerURL+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Token", c.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("%s", resp.Status)
		}
		return nil, fmt.Errorf("%s: %s", resp.Status, b)
	}
	var res map[string]string
	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}

func (c *Client) Send(path string, args map[string]string) (map[string]string, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.ServerURL+path, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Token", c.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("%s", resp.Status)
		}
		return nil, fmt.Errorf("%s: %s", resp.Status, b)
	}
	var res map[string]string
	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}
