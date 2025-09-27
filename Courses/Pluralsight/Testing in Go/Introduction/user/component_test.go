package user

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockResponseWriter struct {
	bytes.Buffer
}

func (mrw mockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (mrw mockResponseWriter) WriteHeader(status int) {}

func (mrw mockResponseWriter) getData() []byte {
	return mrw.Bytes()
}

func TestHandler(t *testing.T) {
	users = []User{
		{
			ID:       1,
			Username: "adent",
		},
		{
			ID:       2,
			Username: "tmacmillan",
		},
	}

	req, err := http.NewRequest(http.MethodGet, "/users", nil)

	if err != nil {
		t.Fatal(err)
	}
	rw := mockResponseWriter{}
	expect, err := json.Marshal(users)
	if err != nil {
		t.Fatal(err)
	}

	Handler(&rw, req)

	got := rw.getData()

	if !bytes.Equal(expect, got) {
		t.Fail()
	}
}

func TestHandler_httptest(t *testing.T) {
	users = []User{
		{
			ID:       1,
			Username: "adent",
		},
		{
			ID:       2,
			Username: "tmacmillan",
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	expect, err := json.Marshal(users)
	if err != nil {
		t.Fatal(err)
	}

	Handler(rec, req)

	res := rec.Result()
	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// assert
	if !bytes.Equal(expect, got) {
		t.Fail()
	}
}

func TestHandler_server(t *testing.T) {
	users = []User{
		{
			ID:       1,
			Username: "adent",
		},
		{
			ID:       2,
			Username: "tmacmillan",
		},
	}
	expect, err := json.Marshal(users)
	if err != nil {
		t.Fatal(err)
	}

	// creamos el servidor
	s := httptest.NewServer(http.HandlerFunc(Handler))
	// configuramos un cliente
	c := s.Client()
	// hacemos la petición
	res, err := c.Get(s.URL + "/users")
	if err != nil {
		t.Fatal(err)
	}
	// leemos la respuesta
	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	// assert
	if !bytes.Equal(expect, got) {
		t.Fail()
	}
	// cerramos el servidor
	s.Close()
}

func BenchmarkHandler(b *testing.B) {
	users = []User{
		{
			ID:       1,
			Username: "adent",
		},
		{
			ID:       2,
			Username: "tmacmillan",
		},
	}

	req, err := http.NewRequest(http.MethodGet, "/users", nil)

	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.Log(b.N)

	for i := 0; i < b.N; i++ {
		rw := mockResponseWriter{}
		Handler(&rw, req)
	}

}

func BenchmarkHandlerParallel(b *testing.B) {
	users = []User{
		{
			ID:       1,
			Username: "adent",
		},
		{
			ID:       2,
			Username: "tmacmillan",
		},
	}

	req, err := http.NewRequest(http.MethodGet, "/users", nil)

	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.Log(b.N)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rw := mockResponseWriter{}
			if err != nil {
				b.Fatal(err)
			}
			Handler(&rw, req)
		}
	})
}
