package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("key"))

	anotherContext := context.WithValue(ctx, "key", "anotherValue")
	doAnother(anotherContext)

	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("key"))

}

func doAnother(ctx context.Context) {
	fmt.Printf("doAnother: myKey's value is %s\n", ctx.Value("key"))
}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "key", "value-test")
	doSomething(ctx)
}
