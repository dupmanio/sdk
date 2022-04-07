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

type HTTPError struct {
	Code     int         `json:"code"`
	ErrorRaw interface{} `json:"error"`
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%d: %v", e.Code, e.ErrorRaw)
}
