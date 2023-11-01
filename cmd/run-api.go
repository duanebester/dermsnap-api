package cmd

import (
	"dermsnap/api"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var Port string

func init() {
	rootCmd.AddCommand(runApiCmd)
	runApiCmd.Flags().StringVarP(&Port, "port", "p", "3000", "Port to listen on")
}

var runApiCmd = &cobra.Command{
	Use:   "run-api",
	Short: "Runs the dermsnap API",
	Long:  `Runs the dermsnap API`,
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}

		app := api.NewApp()
		app.Listen(fmt.Sprintf(":%s", Port))
	},
}
