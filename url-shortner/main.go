package main

import (
	"fmt"

	"github.com/sullyvannunes/url-shortner/cmd"
)

func main() {
	rootCmd := cmd.New()

	if err := rootCmd.Execute(); err != nil {
		panic(fmt.Errorf("failed to execute root command %w", err))
	}
}
