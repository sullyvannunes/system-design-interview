package routes

import (
	"github.com/spf13/cobra"
)

func NewRouteCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "routes",
		Short: "Prints out all routes the server handle",
		Long:  "Prints out all routes the server handle",
		Run: func(cmd *cobra.Command, args []string) {
			// for _, route := range api.Routes() {
			// 	fmt.Printf("%6s %s\n", route.Method, route.Url)
			// }
		},
	}
}
