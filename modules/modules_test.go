package modules

import (
	"regexp"
	"testing"
)

// Para correr archivos de test uso `go test` o `go test -v` para ver los detalles de los test
func TestHello(t *testing.T) {
	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello("Juan")
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}
