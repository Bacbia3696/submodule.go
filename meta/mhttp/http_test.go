package mhttp_test

import (
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/submodule-org/submodule.go"
	"github.com/submodule-org/submodule.go/meta/mhttp"
)

type Hello struct{}

func (h *Hello) AdaptToHTTPHandler(m *http.ServeMux) {
	m.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
}

var HelloRoute = submodule.Resolve(&Hello{})

func TestMHTTP(t *testing.T) {

	t.Run("simple http server", func(t *testing.T) {
		var e error
		s := submodule.CreateScope()

		_, e = HelloRoute.SafeResolveWith(s)
		assert.Nil(t, e)

		go func() {
			mhttp.StartIn(s)
		}()

		defer func() {
			mhttp.StopIn(s)
		}()

		time.Sleep(200 * time.Millisecond)
		r, e := http.Get("http://localhost:8080/hello")

		assert.Nil(t, e)
		assert.Equal(t, 200, r.StatusCode)

		body, e := io.ReadAll(r.Body)
		assert.Nil(t, e)
		assert.Equal(t, "hello", string(body))
	})

	t.Run("change port", func(t *testing.T) {
		ov := os.Getenv("HTTP_ADDR")
		defer os.Setenv("HTTP_ADDR", ov)

		var e error
		e = os.Setenv("HTTP_ADDR", ":30001")
		assert.Nil(t, e)

		s := submodule.CreateScope()

		_, e = HelloRoute.SafeResolveWith(s)
		assert.Nil(t, e)

		go func() {
			mhttp.StartIn(s)
		}()

		defer func() {
			mhttp.StopIn(s)
		}()

		time.Sleep(200 * time.Millisecond)
		r, e := http.Get("http://localhost:30001/hello")

		assert.Nil(t, e)
		assert.Equal(t, 200, r.StatusCode)

		body, e := io.ReadAll(r.Body)
		assert.Nil(t, e)
		assert.Equal(t, "hello", string(body))
	})

}
