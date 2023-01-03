//go:build integration

package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/hgcassiopeia/assessment/expenses/drivers"
	"github.com/hgcassiopeia/assessment/expenses/entities"
	"github.com/hgcassiopeia/assessment/expenses/repo"
	"github.com/hgcassiopeia/assessment/expenses/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExpensesTestSuite struct {
	suite.Suite
	app echo.Echo
}

func TestExpensesTestSuite(t *testing.T) {
	suite.Run(t, &ExpensesTestSuite{})
}

func (s *ExpensesTestSuite) SetupSuite() {
	dbConn, err := drivers.ConnectDB()
	s.Nil(err)
	err = drivers.InitTable(dbConn)
	s.Nil(err)

	expensesRepository := repo.InitRepository(dbConn)
	expenseUseCase := service.Init(expensesRepository)
	httpHandler := HttpHandler{UseCase: expenseUseCase}

	s.app = *echo.New()
	s.app.POST("/expenses", httpHandler.AddNewExpense)
	s.app.GET("/expenses/:id", httpHandler.GetExpenseDetail)
	s.app.PUT("/expenses/:id", httpHandler.UpdateExpense)

	go func() {
		serverPort := fmt.Sprintf(":%v", os.Getenv("PORT"))
		s.app.Start(serverPort)
	}()
}

func (s *ExpensesTestSuite) TearDownSuite() {
	s.app.Logger.Info("tear down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := s.app.Shutdown(ctx)
	assert.NoError(s.T(), err)
}

func (s *ExpensesTestSuite) TestAddNewExpenses() {
	body := bytes.NewBufferString(`{
		"title": "Isakaya Bangna",
		"amount": 899,
		"note": "central bangna", 
		"tags": ["food", "beverage"]
	}`)

	var result entities.Expenses
	response := request(http.MethodPost, uri("expenses"), body)
	err := response.Decode(&result)

	expected := entities.Expenses{
		Id:     0,
		Title:  "Isakaya Bangna",
		Amount: float32(899),
		Note:   "central bangna",
		Tags:   []string{"food", "beverage"},
	}

	if err == nil {
		assert.Equal(s.T(), http.StatusCreated, response.StatusCode)
		assert.NotNil(s.T(), &result.Id)
		assert.NotEqual(s.T(), expected.Id, result.Id)
		assert.Equal(s.T(), expected.Title, result.Title)
		assert.Equal(s.T(), expected.Amount, result.Amount)
		assert.Equal(s.T(), expected.Note, result.Note)
		assert.ElementsMatch(s.T(), expected.Tags, result.Tags)
	}
}

func (s *ExpensesTestSuite) TestAddNewExpenses_BadRequest() {
	body := bytes.NewBufferString(`{
		"title": "Isakaya Bangna",
		"amount": "899",
		"note": "central bangna", 
		"tags": ["food", "beverage"]
	}`)

	var exp entities.Expenses
	response := request(http.MethodPost, uri("expenses"), body)
	err := response.Decode(&exp)

	if err != nil {
		assert.Equal(s.T(), http.StatusBadRequest, response.StatusCode)
	}
}

func (s *ExpensesTestSuite) TestGetExpenseDetail() {
	var exp entities.Expenses

	given := "1"
	response := request(http.MethodGet, uri("expenses", given), nil)
	err := response.Decode(&exp)

	if err == nil {
		assert.Equal(s.T(), http.StatusOK, response.StatusCode)
	}
}

func (s *ExpensesTestSuite) TestUpdateExpense() {
	body := bytes.NewBufferString(`{
		"id": 1,
		"title": "Isakaya Bangna",
		"amount": 1000,
		"note": "Central bangna near by BigC", 
		"tags": ["food", "beverage"]
	}`)

	id := "1"
	var result entities.Expenses
	response := request(http.MethodPut, uri("expenses", id), body)
	err := response.Decode(&result)

	expected := entities.Expenses{
		Id:     1,
		Title:  "Isakaya Bangna",
		Amount: float32(1000),
		Note:   "Central bangna near by BigC",
		Tags:   []string{"food", "beverage"},
	}

	if err == nil {
		assert.Equal(s.T(), http.StatusOK, response.StatusCode)
		assert.Equal(s.T(), expected.Id, result.Id)
		assert.Equal(s.T(), expected.Title, result.Title)
		assert.Equal(s.T(), expected.Amount, result.Amount)
		assert.Equal(s.T(), expected.Note, result.Note)
		assert.ElementsMatch(s.T(), expected.Tags, result.Tags)
	}
}

func (s *ExpensesTestSuite) TestUpdateExpense_BadRequest() {
	id := "1"
	body := bytes.NewBufferString(`{
		"id": "1",
		"title": "Isakaya Bangna",
		"amount": "899",
		"note": "central bangna",
		"tags": ["food", "beverage"]
	}`)

	var exp entities.Expenses
	response := request(http.MethodPut, uri("expenses", id), body)
	err := response.Decode(&exp)

	if err != nil {
		assert.Equal(s.T(), http.StatusBadRequest, response.StatusCode)
	}
}

func uri(paths ...string) string {
	host := fmt.Sprintf("http://localhost:%v", os.Getenv("PORT"))
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
