package meowfacts

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout is zero(0)")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}
func (c *Client) GetAsset(count int) error {
	url := fmt.Sprint("https://meowfacts.herokuapp.com/?count=", count)
	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var as AssetData
	if err = json.Unmarshal(body, &as); err != nil {

		return err
	}
	for i, asset := range as.Data {

		fmt.Println(i+1, ":", asset)
	}
	return nil
}
func (c *Client) GetAssetOnLang(lang string) error {

	url := fmt.Sprint("https://meowfacts.herokuapp.com/?lang=", lang)
	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var as AssetData
	if err = json.Unmarshal(body, &as); err != nil {

		return err
	}
	for i, asset := range as.Data {

		fmt.Println(i+1, ":", asset)
	}
	return nil

}
func (c *Client) GetAssetsOnLang(count int, lang string) error {
	if count <= 0 {
		return errors.New("count is zero or less then zero")
	}
	var as AssetData
	url := fmt.Sprint("https://meowfacts.herokuapp.com/?lang=", lang)

	for i := 0; i < count; i++ {
		resp, err := c.client.Get(url)
		if err != nil {
			return err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(body, &as); err != nil {

			return err
		}
		defer resp.Body.Close()
		fmt.Println(i+1, ":", as.Data[0])
	}

	return nil
}
