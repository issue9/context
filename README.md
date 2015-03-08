context [![Build Status](https://travis-ci.org/issue9/context.svg?branch=master)](https://travis-ci.org/issue9/context)
======

```go
func h(w http.ResponseWriter, req *http.Request) {
    ctx := context.Get(req)
    ctx.Set("key", "val")
    // do somthing...
    var v string
    vi,found := ctx.Get("key")
    if found {
        v = vi.(string)
    }
}
```

### 安装

```shell
go get github.com/issue9/context
```


### 文档

[![Go Walker](http://gowalker.org/api/v1/badge)](http://gowalker.org/github.com/issue9/context)
[![GoDoc](https://godoc.org/github.com/issue9/context?status.svg)](https://godoc.org/github.com/issue9/context)


### 版权

本项目采用[MIT](http://opensource.org/licenses/MIT)开源授权许可证，完整的授权说明可在[LICENSE](LICENSE)文件中找到。
