package user

import (
	"bytes"
	"encoding/json"
	"net/http"
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
