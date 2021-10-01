package it

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (s *e2eTestSuite) Test_EndToEnd_CreateUser() {
	reqStr := fmt.Sprintf(`{"name":"e2euser%d"}`, s.r.Intn(100))
	req, err := http.NewRequest(echo.POST, fmt.Sprintf("%s/user", s.baseUrl), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"code":"","success":true,"message":"user created"}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()
}
