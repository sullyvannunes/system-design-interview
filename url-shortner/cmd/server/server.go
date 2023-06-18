package server

import (
	"github.com/spf13/cobra"
	"github.com/sullyvannunes/url-shortner/pkg/api"
)

func NewServer() *cobra.Command {
	srvCommand := &cobra.Command{
		Use:   "server",
		Short: "server options starts url shortner's server",
		Long:  "server options starts url shortner's server",
		RunE: func(cmd *cobra.Command, args []string) error {
			srv := api.NewServer()
			return api.StartServer(srv)
		},
	}

	return srvCommand
}
