package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	// ctx := context.WithValue(context.Background(), "number", 9)
	ctx, _ := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "number", 9)
	ctx = context.WithValue(ctx, "keword", "value")
	go square(ctx)
	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square: %d", n*n)
	}
	wg.Done()
}
