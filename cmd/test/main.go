package main

import (
	"context"

	"github.com/jorgesanchez-e/test/internal/rest"
)

func main() {
	ctx := context.Background()

	rs := rest.NewServer(ctx)

	rs.Start()
}
