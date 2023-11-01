package cmd

import (
	"dermsnap/app"
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
	Use:   "run-app",
	Short: "Runs the dermsnap app",
	Long:  `Runs the dermsnap app`,
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}

		app := app.NewApp()
		app.Listen(fmt.Sprintf(":%s", Port))
	},
}
