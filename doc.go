// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// context包提供了基于每个http请求期间的数据共享机制。
//
//  func h(w http.ResponseWriter, req *http.Request) {
//      ctx := context.Get(req)
//      ctx.Set("key", "val")
//      // do somthing...
//      var v string
//      vi,found := ctx.Get("key")
//      if found {
//          v = vi.(string)
//      }
//  }
package context

const Version = "0.1.1.150308"
