// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package context

import (
	"errors"
)

var (
	ErrNotFound   = errors.New("未找到对应的值")
	ErrNotConvert = errors.New("不能转换成目标类型")
)

func (ctx *Context) String(key interface{}) (string, error) {
	v, found := ctx.Get(key)
	if !found {
		return "", ErrNotFound
	}

	switch val := v.(type) {
	case string:
		return val, nil
	case []byte:
		return string(val), nil
	case []rune:
		return string(val), nil
	default:
		return "", ErrNotConvert
	}
}

func (ctx *Context) MustString(key interface{}, def string) string {
	switch val := ctx.MustGet(key, def).(type) {
	case string:
		return val
	case []byte:
		return string(val)
	case []rune:
		return string(val)
	default:
		return def
	}
}

func (ctx *Context) Int64(key interface{}) (int64, error) {
	v, found := ctx.Get(key)
	if !found {
		return 0, ErrNotFound
	}

	val, ok := v.(int64)
	if !ok {
		return 0, ErrNotConvert
	}
	return val, nil
}

func (ctx *Context) MustInt64(key interface{}, def int64) int64 {
	v := ctx.MustGet(key, def)

	val, ok := v.(int64)
	if !ok {
		return def
	}
	return val
}

func (ctx *Context) Int32(key interface{}) (int32, error) {
	v, found := ctx.Get(key)
	if !found {
		return 0, ErrNotFound
	}

	val, ok := v.(int32)
	if !ok {
		return 0, ErrNotConvert
	}
	return val, nil
}

func (ctx *Context) MustInt32(key interface{}, def int32) int32 {
	v := ctx.MustGet(key, def)

	val, ok := v.(int32)
	if !ok {
		return def
	}
	return val
}

func (ctx *Context) Int16(key interface{}) (int16, error) {
	v, found := ctx.Get(key)
	if !found {
		return 0, ErrNotFound
	}

	val, ok := v.(int16)
	if !ok {
		return 0, ErrNotConvert
	}
	return val, nil
}

func (ctx *Context) MustInt16(key interface{}, def int16) int16 {
	v := ctx.MustGet(key, def)

	val, ok := v.(int16)
	if !ok {
		return def
	}
	return val
}

func (ctx *Context) Int8(key interface{}) (int8, error) {
	v, found := ctx.Get(key)
	if !found {
		return 0, ErrNotFound
	}

	val, ok := v.(int8)
	if !ok {
		return 0, ErrNotConvert
	}
	return val, nil
}

func (ctx *Context) MustInt8(key interface{}, def int8) int8 {
	v := ctx.MustGet(key, def)

	val, ok := v.(int8)
	if !ok {
		return def
	}
	return val
}

func (ctx *Context) Int(key interface{}) (int, error) {
	v, found := ctx.Get(key)
	if !found {
		return 0, ErrNotFound
	}

	val, ok := v.(int)
	if !ok {
		return 0, ErrNotConvert
	}
	return val, nil
}

func (ctx *Context) MustInt(key interface{}, def int) int {
	v := ctx.MustGet(key, def)

	val, ok := v.(int)
	if !ok {
		return def
	}
	return val
}

func (ctx *Context) Uint64(key interface{}) (uint64, error) {
	v, found := ctx.Get(key)
	if !found {
		return 0, ErrNotFound
	}

	val, ok := v.(uint64)
	if !ok {
		return 0, ErrNotConvert
	}
	return val, nil
}

func (ctx *Context) MustUint64(key interface{}, def uint64) uint64 {
	v := ctx.MustGet(key, def)

	val, ok := v.(uint64)
	if !ok {
		return def
	}
	return val
}

func (ctx *Context) Uint32(key interface{}) (uint32, error) {
	v, found := ctx.Get(key)
	if !found {
		return 0, ErrNotFound
	}

	val, ok := v.(uint32)
	if !ok {
		return 0, ErrNotConvert
	}
	return val, nil
}

func (ctx *Context) MustUint32(key interface{}, def uint32) uint32 {
	v := ctx.MustGet(key, def)

	val, ok := v.(uint32)
	if !ok {
		return def
	}
	return val
}

func (ctx *Context) Uint16(key interface{}) (uint16, error) {
	v, found := ctx.Get(key)
	if !found {
		return 0, ErrNotFound
	}

	val, ok := v.(uint16)
	if !ok {
		return 0, ErrNotConvert
	}
	return val, nil
}

func (ctx *Context) MustUint16(key interface{}, def uint16) uint16 {
	v := ctx.MustGet(key, def)

	val, ok := v.(uint16)
	if !ok {
		return def
	}
	return val
}

func (ctx *Context) Uint8(key interface{}) (uint8, error) {
	v, found := ctx.Get(key)
	if !found {
		return 0, ErrNotFound
	}

	val, ok := v.(uint8)
	if !ok {
		return 0, ErrNotConvert
	}
	return val, nil
}

func (ctx *Context) MustUint8(key interface{}, def uint8) uint8 {
	v := ctx.MustGet(key, def)

	val, ok := v.(uint8)
	if !ok {
		return def
	}
	return val
}

func (ctx *Context) Uint(key interface{}) (uint, error) {
	v, found := ctx.Get(key)
	if !found {
		return 0, ErrNotFound
	}

	val, ok := v.(uint)
	if !ok {
		return 0, ErrNotConvert
	}
	return val, nil
}

func (ctx *Context) MustUint(key interface{}, def uint) uint {
	v := ctx.MustGet(key, def)

	val, ok := v.(uint)
	if !ok {
		return def
	}
	return val
}
