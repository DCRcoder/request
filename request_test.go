package request

import (
	"net/http"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGet(t *testing.T) {
	data := make(map[string]interface{})
	req := NewRequest()
	res, _ := req.Get("http://httpbin.org/get", nil)
	assert.Equal(t, res.StatusCode, 200)
	assert.NotEqual(t, res.Text(), "")
	_ = res.JSON(&data)
	assert.Equal(t, data["url"], "http://httpbin.org/get")

	res, _ = req.Get("http://httpbin.org/get123", nil)
	assert.Equal(t, res.StatusCode, 404)

	res, _ = req.Get("http://httpbin.org/delete", nil)
	assert.Equal(t, res.StatusCode, 405)
}

func TestHeader(t *testing.T) {
	req := NewRequest()
	req.SetHeaderByMap(map[string]string{"123": "123"})
	assert.Equal(t, req.Header["123"], "123")
	req.SetHeader("123", "yui")
	assert.Equal(t, req.Header["123"], "yui")
}

func TestDelete(t *testing.T) {
	data := make(map[string]interface{})
	req := NewRequest()
	res, _ := req.Delete("http://httpbin.org/delete", nil)
	assert.Equal(t, res.StatusCode, 200)
	assert.NotEqual(t, res.Text(), "")
	_ = res.JSON(&data)
	assert.Equal(t, data["url"], "http://httpbin.org/delete")

	res, _ = req.Delete("http://httpbin.org/delete123", nil)
	assert.Equal(t, res.StatusCode, 404)

	res, _ = req.Delete("http://httpbin.org/get", nil)
	assert.Equal(t, res.StatusCode, 405)
}

func TestPost(t *testing.T) {
	req := NewRequest()
	data := map[string]interface{}{
		"123": "123",
	}
	req.JSON = data
	res, _ := req.Post("http://httpbin.org/post", nil)
	assert.Equal(t, res.StatusCode, 200)
	da := make(map[string]interface{})
	_ = res.JSON(&da)
	assert.Equal(t, da["json"], data)
	assert.Equal(t, res.Header.Get(HeaderContentType), MIMEApplicationJSON)

	data1 := map[string]string{
		"123": "gogog",
	}
	req.Data = data1
	res, _ = req.Post("http://httpbin.org/post", nil)
	assert.Equal(t, res.StatusCode, 200)
}

func TestPut(t *testing.T) {
	req := NewRequest()
	data := map[string]interface{}{
		"123": "123",
	}
	req.JSON = data
	res, _ := req.Put("http://httpbin.org/put", nil)
	assert.Equal(t, res.StatusCode, 200)
	da := make(map[string]interface{})
	_ = res.JSON(&da)
	assert.Equal(t, da["json"], data)
	assert.Equal(t, res.Header.Get(HeaderContentType), MIMEApplicationJSON)

	data1 := map[string]string{
		"123": "gogog",
	}
	req.Data = data1
	res, _ = req.Put("http://httpbin.org/put", nil)
	assert.Equal(t, res.StatusCode, 200)
}

func TestPatch(t *testing.T) {
	req := NewRequest()
	data := map[string]interface{}{
		"123": "123",
	}
	req.JSON = data
	res, _ := req.Patch("http://httpbin.org/patch", nil)
	assert.Equal(t, res.StatusCode, 200)
	da := make(map[string]interface{})
	_ = res.JSON(&da)
	assert.Equal(t, da["json"], data)
	assert.Equal(t, res.Header.Get(HeaderContentType), MIMEApplicationJSON)

	data1 := map[string]string{
		"123": "gogog",
	}
	req.Data = data1
	res, _ = req.Patch("http://httpbin.org/patch", nil)
	assert.Equal(t, res.StatusCode, 200)
}

type HttpHeaderTest struct {
	Cookie string `json:"Cookie"`
}

type CookieTest struct {
	Header *HttpHeaderTest `json:"headers"`
}

func TestCookie(t *testing.T) {
	r := NewRequest()
	r.AddCookie(&http.Cookie{Name: "test", Value: "test"})
	resp, _ := r.Get("http://httpbin.org/get", nil)
	assert.Equal(t, resp.StatusCode, 200)
	tt := &CookieTest{}
	_ = resp.JSON(tt)
	assert.Equal(t, tt.Header.Cookie, "test=test")
}
