package cmd

import (
	"dermsnap/app"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var Port string
var SeedData bool

func init() {
	rootCmd.AddCommand(runApiCmd)
	runApiCmd.Flags().StringVarP(&Port, "port", "p", "8080", "Port to listen on")
	runApiCmd.Flags().BoolVarP(&SeedData, "seed", "s", false, "Seed the database with admin user")
}

var runApiCmd = &cobra.Command{
	Use:   "run-app",
	Short: "Runs the dermsnap app",
	Long:  `Runs the dermsnap app`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = godotenv.Load()
		app := app.NewApp(SeedData)
		app.Listen(fmt.Sprintf(":%s", Port))
	},
}
