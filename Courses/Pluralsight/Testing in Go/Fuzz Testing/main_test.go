package main

import (
	"strconv"
	"strings"
	"testing"
)

func toUpperOld(s string) string {
	var b strings.Builder
	for i := range s {
		if _, err := strconv.Atoi(string(s[i])); err == nil {
			b.WriteByte(s[i])
			continue
		}
		if s[i] >= 'A' && s[i] <= 'Z' {
			b.WriteByte(s[i])
			continue
		}
		b.WriteByte(byte(s[i]) - 32)
	}
	return b.String()
}

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func TestUpper(t *testing.T) {
	have := "hello"
	expect := "HELLO"
	got := toUpper(have)
	t.Log(got, expect)
	if expect != got {
		t.Fail()
	}
}

// go test -v -fuzz .
// go test -v -fuzz . -fuzztime 30s # Setea el tiempo del test, por default es inifnito
// go test -v -fuzz . -fuzztime 3s -fuzzminimizetime 30s # Setea el maximo tiempo de falla de optimizacion
func FuzzToUpper(f *testing.F) {
	// Agrega los argumentos que deberian ser pasados al fuzz test
	f.Add("hello")

	// Solo 1 Fuzz por Fuzz test
	f.Fuzz(func(t *testing.T, s string) {
		out := toUpper(s)
		if out != strings.ToUpper(s) {
			t.Fail()
		}
	}) // Registramos el fuzz
}
