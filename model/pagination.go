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

type Pagination struct {
	Limit      int `json:"limit"`
	Page       int `json:"page"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}
