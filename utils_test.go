package request

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestPraseUrl(t *testing.T) {
	baseURL := "https://www.123.com/query"
	param := map[string]string{
		"123":  "123",
		"home": "youye",
	}

	u, _ := ParseQueryURL(baseURL, param)
	assert.Equal(t, u, "https://www.123.com/query?123=123&home=youye")
	u, _ = ParseQueryURL(baseURL, nil)
	assert.Equal(t, u, baseURL)
}
