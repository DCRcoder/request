package request

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestGet(t *testing.T) {
	req := NewRequest()
	res, _ := req.Get("https://www.baidu.com", nil)
	assert.Equal(t, res.StatusCode, 200)
}

func TestDelete(t *testing.T) {
	req := NewRequest()
	res, _ := req.Delete("https://www.baidu.com", nil)
	assert.Equal(t, res.StatusCode, 200)
}
