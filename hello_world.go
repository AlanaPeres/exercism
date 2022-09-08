package greeting

import "testing"

// HelloWorld greets the world.
func HelloWorld() string {
	return "Hello, World!"
}
func TestHelloWord(t *testing.T) {
	expectativa := "Hello, World!"
	if expectativa != HelloWorld() { 
		t.Error("O tipo recebido é diferente do esperado!")
	}
}
