package learn_golang_context

import (
	"context"
	"fmt"
	"testing"
)

type ContextKey string

func TestContextValue(t *testing.T) {
	contextRoot := context.Background()

	contextA := context.WithValue(contextRoot, ContextKey("ctxA"), "A")
	contextB := context.WithValue(contextRoot, ContextKey("ctxB"), "B")

	contextC := context.WithValue(contextA, ContextKey("ctxC"), "C")
	contextD := context.WithValue(contextB, ContextKey("ctxD"), "D")

	fmt.Println(contextC.Value(ContextKey("ctxA")))
	fmt.Println(contextD.Value(ContextKey("ctxB")))
	fmt.Println(contextA.Value(ContextKey("ctxC")))
}
