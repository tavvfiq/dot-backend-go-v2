package it

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (s *e2eTestSuite) Test_EndToEnd_AddComment() {
	reqStr := `{"user_id":20, "article_id":1, "content": "test e2e comment"}`
	req, err := http.NewRequest(echo.POST, fmt.Sprintf("%s/comment", s.baseUrl), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"code":"","success":true,"message":"comment added"}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_DeleteComment() {
	reqStr := ""
	req, err := http.NewRequest(echo.DELETE, fmt.Sprintf("%s/comment/%d", s.baseUrl, 1), strings.NewReader(reqStr))
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"code":"","success":true,"message":"comment deleted"}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()
}
