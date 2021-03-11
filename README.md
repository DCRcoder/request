request
=======
A friendly HTTP request library for Golang like [Python-Requests](https://github.com/kennethreitz/requests).

Installation
------------

```
go get -u github.com/DCRcoder/request
```

Usage
-------

```go
import (
    "github.com/DCRcoder/request"
)
```

**GET**:

```go
req := request.NewRequest(c)
resp, err := req.Get("http://httpbin.org/get", nil)

// 输出文本
resp.Text()

// 可以序列化为 map 或者 struct
resp.JSON(&data)

// 设置 querys
req.Get("http://httpbin.org/get", map[string]string{"test": "test"})
```

**POST、PUT、PATCH**:

```go
// 以 post 为例
req := request.NewRequest()
// form reuqest
req.Data = map[string]string{
    "key": "value",
    "a":   "123",
}
resp, err := req.Post("http://httpbin.org/post", nil)

// json request
req.JSON = map[string]string{
    "key": "value",
    "a":   "123",
}
resp, err := req.Post("http://httpbin.org/post", nil)
```

**Headers**:

```go
req := request.NewRequest(c)
// 设置整个 header
req.SetHeaderByMap(map[string]string{
    "Accept-Encoding": "gzip,deflate,sdch",
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
})
resp, err := req.Get("http://httpbin.org/get", nil)

// 设置单个 header
req.SetHeader("content", "json")
```

TODO
-------
[ ] FIEL REQUEST

[ ] AUTH

[ ] HOOK

[ ] PROXY

License
---------

Under the MIT License.