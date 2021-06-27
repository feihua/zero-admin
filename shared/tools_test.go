package shared

import (
	"net/url"
	"testing"
)

func TestHttpPostForm(t *testing.T) {
	body, err := HttpPostForm("http://xxxxxxxxxx", url.Values{})
	if err != nil {
		t.Error(err)
	}
	t.Error(string(body))
}
