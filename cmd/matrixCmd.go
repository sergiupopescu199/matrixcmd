package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	mat "goland/matrix"
)

// determinantCmd represents the determinant command
var determinantCmd = &cobra.Command{
	Use:   "det",
	Short: "Calculate the determinant of an NxN matrix",
	Long: `This command calculates the determinant of a square matrix from a file

the commad need an external file to read the matrix
the file must have in first row the dimension of the matrix and in the third row the matrix can be written
eg:

3

1.2 3.6 3
2.5 7.8 3
6 8 3

use the -f or --file flag to chose a file.txt

eg: file.txt or /home/user/file.txt

`,
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		dimension, matrix := mat.ReadFile(file)
		det := mat.Determinant(matrix, dimension)
		fmt.Printf("Determinant: %.4f\n", det)
	},
}

var inverseCmd = &cobra.Command{
	Use:   "inv",
	Short: "Calculate the inverse of an NxN matrix",
	Long: `This command calculates the inverse of a matrix from a file

the commad need an external file to read the matrix
the file must have in first row the dimension of the matrix and in the third row the matrix can be written
eg:

3

1.2 3.6 3
2.5 7.8 3
6 8 3

use the -f or --file flag to chose a file.txt

eg: file.txt or /home/user/file.txt

`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
			in var file i put the falue of the file flag from console
		*/
		file, _ := cmd.Flags().GetString("file")
		dimension, matrix := mat.ReadFile(file)
		det := mat.Determinant(matrix, dimension)
		mat.Inverse(matrix, dimension, det)
	},
}

func init() {
	rootCmd.AddCommand(determinantCmd)
	rootCmd.AddCommand(inverseCmd)
	determinantCmd.Flags().StringP("file", "f", "", "get the file, eg: file.txt or /home/user/file.txt")
	inverseCmd.Flags().StringP("file", "f", "", "get the file, eg: file.txt or /home/user/file.txt")
}
