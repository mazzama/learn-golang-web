package httprouter

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, r *http.Request)  {
	fmt.Println("Receive Request")
	middleware.Handler.ServeHTTP(writer, r)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Product with id: " + params.ByName("id"))
	})
	middleware := LogMiddleware{Handler: router}

	request := httptest.NewRequest("GET", "http://localhost:3000/product/2", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product with id: 2", string(body))
}
