// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package context

import (
	"net/http"
	"sync"
)

var (
	ctxs    = make(map[*http.Request]*Context)
	ctxsMux sync.Mutex
)

// 获取或是新建Context。
// 当传递nil值时，会触发panic。
func Get(r *http.Request) *Context {
	if r == nil {
		panic("参数r不能为空")
	}

	ctxsMux.Lock()

	ctx, found := ctxs[r]
	if !found {
		ctx = &Context{items: make(map[interface{}]interface{})}
		ctxs[r] = ctx
	}

	ctxsMux.Unlock() // Unlock
	return ctx
}

// 释放Context。
func Free(r *http.Request) {
	if r == nil {
		return
	}

	ctxsMux.Lock()
	delete(ctxs, r)
	ctxsMux.Unlock()
}

// 与http.Request相关联的上下文环境
type Context struct {
	sync.Mutex
	items map[interface{}]interface{}
}

// 查找key对应的值。在没有查到的情况下，found返回false
func (ctx *Context) Get(key interface{}) (val interface{}, found bool) {
	ctx.Lock()
	val, found = ctx.items[key]
	ctx.Unlock()
	return
}

// 功能与Get()相同，在没有找到相关值的情况下，会返回def，
// 但该值不会写入到Context中，下次用Get()依然会返回false
func (ctx *Context) MustGet(key, def interface{}) interface{} {
	ctx.Lock()
	val, found := ctx.items[key]
	ctx.Unlock()

	if !found {
		return def
	} else {
		return val
	}
}

// 设置或是添加一个键值对。
func (ctx *Context) Set(key, val interface{}) {
	ctx.Lock()
	ctx.items[key] = val
	ctx.Unlock()
}

// 添加一个键值对，若该键名已经存在，则不作任何操作，
// 并且ok返回false，表示操作没有成功
func (ctx *Context) Add(key, val interface{}) (ok bool) {
	ctx.Lock()

	_, found := ctx.items[key]
	if !found {
		ctx.items[key] = val
	}

	ctx.Unlock()
	return !found
}
