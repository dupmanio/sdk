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

import "fmt"

type AuthToken struct {
	Token     string `json:"access_token"`
	Type      string `json:"token_type"`
	ExpiresIn int    `json:"expires_in"`
}

func (t *AuthToken) String() string {
	return fmt.Sprintf("%s %s", t.Type, t.Token)
}
