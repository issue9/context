// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package context

import (
	"net/http"
	"testing"

	"github.com/issue9/assert"
)

func TestContext(t *testing.T) {
	a := assert.New(t)

	// 测试request为nil的情况
	a.Panic(func() {
		Get(nil)
	})

	req1, err := http.NewRequest("GET", "/abc/", nil)
	a.NotError(err).NotNil(req1)
	req2, err := http.NewRequest("GET", "/abc/", nil)
	a.NotError(err).NotNil(req2)

	ctx1 := Get(req1)
	ctx2 := Get(req2)

	// ctx1.Add()
	ok := ctx1.Add("key", "val")
	a.Equal(1, len(ctx1.items)).Equal(0, len(ctx2.items)) // 不影响ctx2
	a.True(ok).Equal(ctx1.MustGet("key", "default").(string), "val")
	a.Equal(ctx2.MustGet("key", "default").(string), "default")

	// 添加一个相同的值，失败
	a.False(ctx1.Add("key", "val1"))
	a.Equal(len(ctx1.items), 1)

	// 测试set
	ctx1.Set("key2", "val")
	ctx1.Set("key2", "val2")
	v, found := ctx1.Get("key2")
	a.True(found).Equal(v, "val2")

	// 同一个Request，应该是相同的值
	newCtx1 := Get(req1)
	a.Equal(newCtx1.MustGet("key", "default").(string), "val")

	// Free(nil)，应该不会释放任何内容
	Free(nil)
	a.Equal(2, len(ctx1.items)).Equal(0, len(ctx2.items))

	Free(req1)
	Free(req2)
	ctx1 = Get(req1)
	a.Equal(0, len(ctx1.items)).Equal(0, len(ctx2.items))
}

// BenchmarkGetFree	 5000000	       334 ns/op
func BenchmarkGetFree(b *testing.B) {
	a := assert.New(b)
	req, err := http.NewRequest("GET", "/api", nil)
	a.NotError(err).NotNil(req)

	for i := 0; i < b.N; i++ {
		c := Get(req)
		if c != nil {
			Free(req)
		}
	}
}

// BenchmarkContext	 1000000	      1192 ns/op
func BenchmarkContext(b *testing.B) {
	ctx := &Context{
		items: make(map[interface{}]interface{}, 0),
	}

	for i := 0; i < b.N; i++ {
		ctx.Set(i, i)
		// 应该是读多于写
		_, found := ctx.Get(i)
		_, found = ctx.Get(i)
		_, found = ctx.Get(i)
		if !found {
			b.Error("!found")
		}
	}
}
