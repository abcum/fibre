// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mw

import (
	"github.com/abcum/fibre"
	"github.com/abcum/surreal/util/uuid"
)

// UniqOpts defines options for the Uniq middleware.
type UniqOpts struct {
	HeaderKey string
}

// Uniq defines middleware for assigning a unique request id.
func Uniq(opts ...*UniqOpts) fibre.MiddlewareFunc {
	return func(h fibre.HandlerFunc) fibre.HandlerFunc {
		return func(c *fibre.Context) error {

			// Set defaults
			if len(opts) == 0 {
				opts = append(opts, &UniqOpts{})
			}

			// Set default values for opts.HeaderKey
			headerKey := opts[0].HeaderKey
			if headerKey == "" {
				headerKey = "Request-Id"
			}

			id := uuid.NewV4().String()

			c.Response().Header().Set(headerKey, id)

			return h(c)

		}
	}
}
