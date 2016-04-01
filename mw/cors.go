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
	"net/http"
	"strconv"
	"strings"

	"github.com/abcum/fibre"
)

// CorsOpts defines options for the Cors middleware.
type CorsOpts struct {
	AllowedMethods      []string
	AllowedHeaders      []string
	AccessControlMaxAge int
}

// Cors defines middleware for setting and checking CORS headers,
func Cors(opts ...*CorsOpts) fibre.MiddlewareFunc {
	return func(h fibre.HandlerFunc) fibre.HandlerFunc {
		return func(c *fibre.Context) error {

			// Set defaults
			if len(opts) == 0 {
				opts = append(opts, &CorsOpts{})
			}

			origin := c.Request().Header().Get("Origin")

			if origin == "" {
				return h(c)
			}

			// Set default values for opts.AllowedMethods
			allowedMethods := opts[0].AllowedMethods
			if len(allowedMethods) == 0 {
				allowedMethods = []string{"GET", "PUT", "POST", "PATCH", "DELETE", "OPTIONS"}
			}

			// Set default values for opts.AllowedHeaders
			allowedHeaders := opts[0].AllowedHeaders
			if len(allowedHeaders) == 0 {
				allowedHeaders = []string{"Accept", "Authorization", "Content-Type", "Origin"}
			}

			// Set default values for opts.AccessControlMaxAge
			accessControlMaxAge := opts[0].AccessControlMaxAge
			if accessControlMaxAge == 0 {
				accessControlMaxAge = 3600
			}

			// Normalize AllowedMethods and make comma-separated-values
			normedMethods := []string{}
			for _, allowedMethod := range allowedMethods {
				normed := http.CanonicalHeaderKey(allowedMethod)
				normedMethods = append(normedMethods, normed)
			}

			// Normalize AllowedHeaders and make comma-separated-values
			normedHeaders := []string{}
			for _, allowedHeader := range allowedHeaders {
				normed := http.CanonicalHeaderKey(allowedHeader)
				normedHeaders = append(normedHeaders, normed)
			}

			c.Response().Header().Set("Access-Control-Allow-Methods", strings.Join(normedMethods, ","))
			c.Response().Header().Set("Access-Control-Allow-Headers", strings.Join(normedHeaders, ","))
			c.Response().Header().Set("Access-Control-Max-Age", strconv.Itoa(accessControlMaxAge))
			c.Response().Header().Set("Access-Control-Allow-Origin", origin)

			return h(c)

		}
	}
}
