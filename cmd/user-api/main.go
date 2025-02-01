package main

import (
	"context"

	"user-api/internal/app"
)

func main() {

	a := app.New()

	ctx := context.Background()

	if err := a.Run(ctx); err != nil {
		panic(err)
	}
}
