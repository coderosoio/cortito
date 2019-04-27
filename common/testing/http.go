package testing

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

type HTTPSuite struct {
	Suite
	ServiceClient client.Client
	Router        *gin.Engine

	defaultHeaders map[string]string
}

func (suite *HTTPSuite) Init() {
	suite.Suite.Init()

	registryName := os.Getenv("MICRO_REGISTRY")
	registryAddress := os.Getenv("MICRO_REGISTRY_ADDRESS")

	switch registryName {
	case "consul":
		consulRegistry := consul.NewRegistry(
			registry.Addrs(registryAddress),
		)
		suite.ServiceClient = client.NewClient(
			client.Registry(consulRegistry),
		)
	case "msdn":
		fallthrough
	default:
		suite.ServiceClient = client.DefaultClient
	}

	suite.defaultHeaders = map[string]string{
		"Content-Type": "application/json",
	}
}

func (suite *HTTPSuite) Request(path string, method string, payload interface{}, headers map[string]string) (*httptest.ResponseRecorder, error) {
	var (
		r    *http.Request
		body io.Reader

		w = httptest.NewRecorder()
	)

	if payload != nil {
		data, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(data))
	}
	r, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	suite.addHeaders(r, suite.defaultHeaders)
	suite.addHeaders(r, headers)

	suite.Router.ServeHTTP(w, r)

	return w, nil
}

func (suite *HTTPSuite) addHeaders(request *http.Request, headers map[string]string) {
	for key, value := range headers {
		request.Header.Add(key, value)
	}
}
