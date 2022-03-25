package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/go-cmp/cmp"
	"io"
	"net/http"
	"net/http/httptest"
	"potato/simple-rest/entities/dto"
	"potato/simple-rest/exports/rest"
	"testing"
)

func TestBasicAPI(t *testing.T) {
	ts := httptest.NewServer(rest.NewServer())
	defer ts.Close()

	// test simpleHandler
	t.Run("QueryString", testSimpleHandler(ts))
	t.Run("PathParam", testPathParamHandler(ts))
	t.Run("PathParams", testPathParamsHandler(ts))
	t.Run("RequestBody", testRequestBodyHandler(ts))
}

func testSimpleHandler(ts *httptest.Server) func(t *testing.T) {
	return func(t *testing.T) {
		if res, err := http.Get(ts.URL + "?name=potato"); err != nil {
			t.Errorf(err.Error())
		} else {
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				t.Errorf("Http Status : %v", res.StatusCode)
			}

			var response dto.BasicResponse

			if data, err := io.ReadAll(res.Body); err != nil {
				t.Errorf("Could not read response : %v", err.Error())
			} else if err := json.Unmarshal(data, &response); err != nil {
				t.Errorf("Could not translate json to response : %v", err.Error())
			}

			if response.Code != 0 || response.Message != fmt.Sprintf("potato 입니다.") {
				t.Errorf("Could not translate json to response : %v , %v", err.Error(), response)
			}

		}
	}
}

func testPathParamHandler(ts *httptest.Server) func(t *testing.T) {
	return func(t *testing.T) {
		if res, err := http.Get(ts.URL + "/potato"); err != nil {
			t.Errorf(err.Error())
		} else {
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				t.Errorf("Http Status : %v", res.StatusCode)
			}

			var response dto.BasicResponse

			if data, err := io.ReadAll(res.Body); err != nil {
				t.Errorf("Could not read response : %v", err.Error())
			} else if err := json.Unmarshal(data, &response); err != nil {
				t.Errorf("Could not translate json to response : %v", err.Error())
			}

			if response.Code != 0 || response.Message != fmt.Sprintf("potato 좋아요.") {
				t.Errorf("Could not translate json to response : %v , %v", err.Error(), response)
			}

		}
	}
}

func testPathParamsHandler(ts *httptest.Server) func(t *testing.T) {
	return func(t *testing.T) {
		if res, err := http.Get(ts.URL + "/potato/2"); err != nil {
			t.Fatalf(err.Error())
		} else {
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				t.Errorf("Http Status : %v", res.StatusCode)
			}

			var response dto.BasicResponse

			if data, err := io.ReadAll(res.Body); err != nil {
				t.Errorf("Could not read response : %v", err.Error())
			} else if err := json.Unmarshal(data, &response); err != nil {
				t.Errorf("Could not translate json to response : %v", err.Error())
			}

			if response.Code != 0 || response.Message != fmt.Sprintf("potato, 2개 주세요.") {
				t.Errorf("Could not translate json to response : %v , %v", err.Error(), response)
			}

		}
	}
}

func testRequestBodyHandler(ts *httptest.Server) func(t *testing.T) {
	return func(t *testing.T) {

		simpleRequest := dto.Simple{
			Id:   0,
			Name: "potato",
		}

		_req, _ := json.Marshal(simpleRequest)
		request := bytes.NewBuffer(_req)
		if res, err := http.Post(ts.URL, binding.MIMEJSON, request); err != nil {
			t.Errorf(err.Error())
		} else {
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				t.Errorf("Http Status : %v", res.StatusCode)
			}

			var response dto.Simple

			if data, err := io.ReadAll(res.Body); err != nil {
				t.Errorf("Could not read response : %v", err.Error())
			} else if err := json.Unmarshal(data, &response); err != nil {
				t.Errorf("Could not translate json to response : %v", err.Error())
			}

			if !cmp.Equal(simpleRequest, response) {
				t.Errorf("Could match response value : %v ", response)
			}
		}
	}
}
