package it

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (s *e2eTestSuite) Test_EndToEnd_CreateArticle() {
	reqStr := `{"title":"e2eTitle", "subtitle":"e2e subtitle", "content": "e2eContent", "author_id":20}`
	req, err := http.NewRequest(echo.POST, fmt.Sprintf("%s/article", s.baseUrl), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"code":"","success":true,"message":"article created"}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_GetAllArticle() {
	reqStr := ""
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("%s/article?title=&page=1&limit=1", s.baseUrl), strings.NewReader(reqStr))
	s.NoError(err)
	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)
}

func (s *e2eTestSuite) Test_EndToEnd_GetDetailArticle() {
	reqStr := ""
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("%s/article/detail/%d", s.baseUrl, 1), strings.NewReader(reqStr))
	s.NoError(err)
	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)
}

func (s *e2eTestSuite) Test_EndToEnd_UpdateArticle() {
	reqStr := `{"title":"e2eTitle updated"}`
	req, err := http.NewRequest(echo.PATCH, fmt.Sprintf("%s/article/%d", s.baseUrl, 1), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"code":"","success":true,"message":"article updated"}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_DeleteArticle() {
	reqStr := ""
	req, err := http.NewRequest(echo.DELETE, fmt.Sprintf("%s/article/%d", s.baseUrl, 2), strings.NewReader(reqStr))
	s.NoError(err)
	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)
	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"code":"","success":true,"message":"article deleted"}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()
}
