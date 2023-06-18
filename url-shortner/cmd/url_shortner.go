package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sullyvannunes/url-shortner/cmd/routes"
	"github.com/sullyvannunes/url-shortner/cmd/server"
)

func New() *cobra.Command {
	command := &cobra.Command{
		Use:   "url_shortner",
		Short: "url_shortner is a url shortner service",
		Long:  "url_shortner is a url shortner service",
	}

	command.AddCommand(
		server.NewServer(),
		routes.NewRouteCommand(),
	)

	return command
}
