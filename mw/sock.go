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
)

// Sock defines middleware for logging websocket connections.
func Sock() fibre.MiddlewareFunc {
	return func(h fibre.HandlerFunc) fibre.HandlerFunc {
		return func(c *fibre.Context) (err error) {

			if !c.IsSocket() {
				return h(c)
			}

			ip := c.IP()
			req := c.Request()
			url := req.URL().Path

			log := c.Fibre().Logger().WithFields(map[string]interface{}{
				"prefix": c.Fibre().Name(),
				"ip":     ip,
				"url":    url,
				"method": "WS",
			})

			if id := c.Get("id"); id != nil {
				log = log.WithField("id", id)
			}

			log.Info("Opening websocket")

			return h(c)

		}
	}
}
