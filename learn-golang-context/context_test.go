package learn_golang_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {

	ctxBg := context.Background()
	fmt.Println(ctxBg)

	ctxTodo := context.TODO()
	fmt.Println(ctxTodo)
}
