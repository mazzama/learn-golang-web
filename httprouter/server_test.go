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

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Product with id: " + params.ByName("id"))
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/product/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product with id: 2", string(body))
}

func TestRouterNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id/item/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Product with id: " + params.ByName("id") + " ItemId: " + params.ByName("itemId"))
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/product/1/item/sku-1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product with id: 1 ItemId: sku-1", string(body))
}

func TestRouterCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Image location: " + params.ByName("image"));
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image location: /small/profile.png", string(body))
}