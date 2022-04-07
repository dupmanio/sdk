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

package system

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dupman/sdk/dupman/session"
	"github.com/dupman/sdk/internal/client"
	"github.com/dupman/sdk/service/system/model"
)

type System struct {
	session *session.Session
}

func New(s *session.Session) *System {
	return &System{
		session: s,
	}
}

func (s *System) GetWebsites(publicKey string, page int) (websites *model.WebsitesResponse, err error) {
	header := http.Header{}
	header.Set("X-Public-Key", publicKey)
	header.Set("Authorization", s.session.Token.String())

	query := url.Values{}
	query.Add("limit", "50")
	query.Add("page", fmt.Sprintf("%d", page))

	_, err = s.session.MakeRequest(client.RequestOptions{
		Method: http.MethodGet,
		Path:   "/system/websites",
		Query:  query,
		Header: header,
	}, &websites)

	return websites, err
}
