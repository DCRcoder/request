package request

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.Config{
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
	UseNumber:              true,
}.Froze()

// Request request 实例
type Request struct {
	Client *http.Client
	Header map[string]string
	Cookie []*http.Cookie
	Body   io.Reader
	Data   map[string]string
	JSON   interface{}
}

// Response response 实例
type Response struct {
	Body       []byte
	StatusCode int
	Header     http.Header
	Cookie     []*http.Cookie
	URL        string
}

// DefaultHeader 默认 header
var DefaultHeader = map[string]string{
	HeaderContentType: MIMEApplicationJSONCharsetUTF8,
}

// NewRequest return new request
func NewRequest() *Request {
	return &Request{
		Client: &http.Client{},
	}
}

// Text return response text
func (res *Response) Text() string {
	return string(res.Body)
}

func (res *Response) Content() []byte {
	return res.Body
}

// JSON return json
func (res *Response) JSON(dst interface{}) error {
	if !reflect.ValueOf(dst).Elem().CanSet() {
		return OBJNotCanSet
	}
	err := json.Unmarshal(res.Body, dst)
	if err != nil {
		return err
	}
	return nil
}

// Timeout set timeout
func (r *Request) Timeout(timeout time.Duration) {
	r.Client.Timeout = timeout
}

// SetHeaderByMap 设置 header
func (r *Request) SetHeaderByMap(header map[string]string) {
	r.Header = header
}

// SetHeader 设置 header
func (r *Request) SetHeader(key, value string) {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header[key] = value
}

// AddCookie add cookie
func (r *Request) AddCookie(cookie *http.Cookie) {
	if r.Cookie == nil {
		r.Cookie = make([]*http.Cookie, 0)
	}
	r.Cookie = append(r.Cookie, cookie)
}

func (r *Request) prepareReq(method, baseURL string, querys map[string]string, body io.Reader) (*http.Request, error) {
	urlPath, err := ParseQueryURL(baseURL, querys)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, urlPath, body)
	if err != nil {
		return nil, err
	}
	for k, v := range r.Header {
		req.Header.Set(k, v)
	}
	for _, c := range r.Cookie {
		req.AddCookie(c)
	}
	return req, nil
}

// Reset reset request data
func (r *Request) Reset() {
	r.Header = nil
	r.Cookie = nil
	r.Data = nil
	r.JSON = nil
}

func (r *Request) finishReq(res *http.Response) (*Response, error) {
	cookie := res.Cookies()
	for _, c := range cookie {
		r.AddCookie(c)
	}
	for k, v := range res.Header {
		r.SetHeader(k, v[0])
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &Response{
		Body:       body,
		StatusCode: res.StatusCode,
		Header:     res.Header,
		Cookie:     cookie,
		URL:        res.Request.URL.String(),
	}, nil
}

// Get get method
func (r *Request) Get(baseURL string, querys map[string]string) (*Response, error) {
	return r.request(http.MethodGet, baseURL, querys)
}

func (r *Request) request(method, baseURL string, querys map[string]string) (*Response, error) {
	if r.JSON != nil {
		jsonByte, err := json.Marshal(r.JSON)
		if err != nil {
			return nil, err
		}
		req, err := r.prepareReq(method, baseURL, querys, bytes.NewBuffer(jsonByte))
		if err != nil {
			return nil, err
		}
		req.Header.Add(HeaderContentType, MIMEApplicationJSONCharsetUTF8)
		req.Header.Add(HeaderContentLength, strconv.Itoa(len(string(jsonByte))))
		resp, err := r.Client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return r.finishReq(resp)
	}
	if r.Data != nil {
		data := url.Values{}
		for k, v := range r.Data {
			data.Set(k, v)
		}
		req, err := r.prepareReq(method, baseURL, querys, strings.NewReader(data.Encode()))
		if err != nil {
			return nil, err
		}
		req.Header.Add(HeaderContentType, MIMEApplicationForm)
		req.Header.Add(HeaderContentLength, strconv.Itoa(len(data.Encode())))
		resp, err := r.Client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return r.finishReq(resp)
	}
	req, err := r.prepareReq(method, baseURL, querys, nil)
	if err != nil {
		return nil, err
	}
	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return r.finishReq(resp)
}

// Post post method
func (r *Request) Post(baseURL string, querys map[string]string) (*Response, error) {
	return r.request(http.MethodPost, baseURL, querys)
}

// Put put method
func (r *Request) Put(baseURL string, querys map[string]string) (*Response, error) {
	return r.request(http.MethodPut, baseURL, querys)
}

// Delete delete method
func (r *Request) Delete(baseURL string, querys map[string]string) (*Response, error) {
	return r.request(http.MethodDelete, baseURL, querys)
}

// Patch patch method
func (r *Request) Patch(baseURL string, querys map[string]string) (*Response, error) {
	return r.request(http.MethodPatch, baseURL, querys)
}

// Options options method
func (r *Request) Options(baseURL string, querys map[string]string) (*Response, error) {
	return r.request(http.MethodOptions, baseURL, querys)
}

// Head head method
func (r *Request) Head(baseURL string, querys map[string]string) (*Response, error) {
	return r.request(http.MethodHead, baseURL, querys)
}

// Connect connect method
func (r *Request) Connect(baseURL string, querys map[string]string) (*Response, error) {
	return r.request(http.MethodConnect, baseURL, querys)
}

// Trace trace method
func (r *Request) Trace(baseURL string, querys map[string]string) (*Response, error) {
	return r.request(http.MethodTrace, baseURL, querys)
}

// Get direct do get method
func Get(baseURL string, querys map[string]string) (*Response, error) {
	r := NewRequest()
	return r.Get(baseURL, querys)
}

// Delete direct do delete method
func Delete(baseURL string, querys map[string]string) (*Response, error) {
	r := NewRequest()
	return r.Delete(baseURL, querys)
}

// PostJSON direct do post method with json data
func PostJSON(baseURL string, querys map[string]string, jsonData interface{}) (*Response, error) {
	r := NewRequest()
	r.JSON = jsonData
	return r.Post(baseURL, querys)
}

// PutJSON direct do put method with json data
func PutJSON(baseURL string, querys map[string]string, jsonData interface{}) (*Response, error) {
	r := NewRequest()
	r.JSON = jsonData
	return r.Put(baseURL, querys)
}

// PatchJSON direct do patch method with json data
func PatchJSON(baseURL string, querys map[string]string, jsonData interface{}) (*Response, error) {
	r := NewRequest()
	r.JSON = jsonData
	return r.Patch(baseURL, querys)
}
