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

package model

import (
	"github.com/dupman/sdk/model"
	"github.com/google/uuid"
)

type Website struct {
	ID    uuid.UUID `json:"id"`
	URL   string    `json:"url"`
	Token string    `json:"token"`
}

type WebsitesResponse struct {
	Code       int              `json:"code"`
	Websites   []Website        `json:"data"`
	Pagination model.Pagination `json:"pagination"`
}
