// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package context

import (
	"net/http"
)

// Free函数的http.Handler接口包装。
func FreeHandler(h http.Handler) http.Handler {
	freeHandler := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		Free(r)
	}

	return http.HandlerFunc(freeHandler)
}
