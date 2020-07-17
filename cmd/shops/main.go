package main

import (
	"context"
	"fmt"
	"os"

	"github.com/thockin/go-build-template/pkg/shops/cmd"
)

func main() {
	ctx := context.Background()
	if err := cmd.Root().ExecuteContext(ctx); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
