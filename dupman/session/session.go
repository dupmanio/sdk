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

package session

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/dupman/sdk/internal/client"
	"github.com/dupman/sdk/model"
)

type Session struct {
	client *client.Client

	Token *model.AuthToken
}

func New(username, password, endpoint string) (*Session, error) {
	httpClient, err := client.New(endpoint)
	if err != nil {
		return nil, err
	}

	session := &Session{
		client: httpClient,
	}

	token, err := session.authenticate(username, password)
	if err != nil {
		return nil, err
	}

	session.Token = token

	return session, nil
}

func (s *Session) MakeRequest(options client.RequestOptions, v interface{}) (res *http.Response, err error) {
	req, err := s.client.NewRequest(options)
	if err != nil {
		return res, err
	}

	return s.client.Do(req, v)
}

func (s *Session) authenticate(username, password string) (token *model.AuthToken, err error) {
	body := url.Values{}
	body.Add("username", username)
	body.Add("password", password)
	body.Add("grant_type", "password")
	body.Add("scope", "")
	body.Add("client_id", "")
	body.Add("client_secret", "")

	header := http.Header{}
	header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, err = s.MakeRequest(client.RequestOptions{
		Method: http.MethodPost,
		Path:   "/auth/token",
		Body:   strings.NewReader(body.Encode()),
		Header: header,
	}, &token)

	return token, err
}
