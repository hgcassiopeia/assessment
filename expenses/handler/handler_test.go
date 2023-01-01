//go:build integration

package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hgcassiopeia/assessment/expenses/entities"
	"github.com/stretchr/testify/assert"
)

func TestAddNewExpenses(t *testing.T) {
	body := bytes.NewBufferString(`{
		"title": "Isakaya Bangna",
		"amount": 899,
		"note": "central bangna", 
		"tags": ["food", "beverage"]
	}`)

	var expense entities.Expenses
	res := request(http.MethodPost, uri("expenses"), body)
	err := res.Decode(&expense)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.NotEqual(t, 0, expense.Id)
		assert.Equal(t, "Isakaya Bangna", expense.Title)
		assert.Equal(t, float32(899), expense.Amount)
		assert.Equal(t, "central bangna", expense.Note)
		assert.ElementsMatch(t, []string{"food", "beverage"}, expense.Tags)
	}
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
