//go:build integration

package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/hgcassiopeia/assessment/expenses/drivers"
	"github.com/hgcassiopeia/assessment/expenses/entities"
	"github.com/hgcassiopeia/assessment/expenses/repo"
	"github.com/hgcassiopeia/assessment/expenses/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const serverPort = 2565

func TestAddNewExpenses(t *testing.T) {
	eh := echo.New()
	go func(e *echo.Echo) {
		dbConn, err := drivers.ConnectDB()
		if err != nil {
			log.Fatal(err)
		}

		expensesRepository := repo.InitRepository(dbConn)
		expenseUseCase := service.Init(expensesRepository)
		httpHandler := HttpHandler{UseCase: expenseUseCase}

		e.POST("/expenses", httpHandler.AddNewExpense)
		e.Start(fmt.Sprintf(":%d", serverPort))
	}(eh)
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", serverPort), 30*time.Second)
		if err != nil {
			log.Println(err)
		}
		if conn != nil {
			conn.Close()
			break
		}
	}

	body := bytes.NewBufferString(`{
		"title": "Isakaya Bangna",
		"amount": 899,
		"note": "central bangna", 
		"tags": ["food", "beverage"]
	}`)

	var result entities.Expenses
	res := request(http.MethodPost, uri("expenses"), body)
	err := res.Decode(&result)

	expected := entities.Expenses{
		Id:     1,
		Title:  "Isakaya Bangna",
		Amount: float32(899),
		Note:   "central bangna",
		Tags:   []string{"food", "beverage"},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.NotEqual(t, expected.Id, &result.Id)
		assert.Equal(t, expected.Title, result.Title)
		assert.Equal(t, expected.Amount, result.Amount)
		assert.Equal(t, expected.Note, result.Note)
		assert.ElementsMatch(t, expected.Tags, result.Tags)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)
}

func uri(paths ...string) string {
	host := fmt.Sprintf("http://localhost:%d", serverPort)
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
