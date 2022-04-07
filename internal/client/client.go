/*
 * This file is part of the dupman/sdk project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>
 */

package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/dupman/sdk/model"
)

type RequestOptions struct {
	Method string
	Path   string
	Query  url.Values
	Body   io.Reader
	Header http.Header
}

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
}

func New(baseURL string) (*Client, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		httpClient: &http.Client{},
		baseURL:    URL,
	}, nil
}

func (c *Client) NewRequest(options RequestOptions) (*http.Request, error) {
	path := &url.URL{Path: options.Path}
	if options.Query != nil {
		path.RawQuery = options.Query.Encode()
	}

	URL := c.baseURL.ResolveReference(path)

	request, err := http.NewRequest(options.Method, URL.String(), options.Body)
	if err != nil {
		return nil, err
	}

	if options.Header != nil {
		request.Header = options.Header
	}

	return request, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (res *http.Response, err error) {
	res, err = c.httpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer res.Body.Close()

	if res.StatusCode >= http.StatusBadRequest {
		var httpError model.HTTPError

		err = json.NewDecoder(res.Body).Decode(&httpError)
		if err != nil {
			httpError.ErrorRaw = err
		}

		if httpError.Code == 0 {
			httpError.Code = res.StatusCode
		}

		return nil, &httpError
	}

	err = json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		return nil, err
	}

	return res, nil
}
