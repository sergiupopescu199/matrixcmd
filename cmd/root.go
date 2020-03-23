package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

/*
	rootCmd represents the base command when called without any subcommands
*/
var rootCmd = &cobra.Command{
	Use:   "matrix",
	Short: "calulate determinant and inverse of a matrix",
	Long: `Matrix is a small utility to calculate the determinant and the inverse of a matrix

exaple: go run matrix.go det -f ./1.txt
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if err := cmd.Help(); err != nil {
				fmt.Println(err)
			}
		}
	},
}

/*
	Execute adds all child commands to the root command and sets flags appropriately.
	This is called by main.main(). It only needs to happen once to the rootCmd.
*/
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
