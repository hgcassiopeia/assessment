package expenses

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNewExpenses(t *testing.T) {
	body := bytes.NewBufferString(`{
		"title": "Isakaya Bangna",
		"amount": 899,
		"note": "central bangna", 
		"tags": ["food", "beverage"]
	}`)
	var exp Expenses
	res := request(http.MethodPost, uri("expenses"), body)
	err := res.Decode(&exp)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, exp.Id)
	assert.Equal(t, "Isakaya Bangna", exp.Title)
	assert.Equal(t, float32(899), exp.Amount)
	assert.Equal(t, "central bangna", exp.Note)
	assert.ElementsMatch(t, []string{"food", "beverage"}, exp.Tags)
}

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", "November 10, 2009")
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)

	return &Response{res, err}
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}
