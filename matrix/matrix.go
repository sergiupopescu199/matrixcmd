package matrix

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	this function read a file.txt and extracts the fimension and the matrix itself, it returns
	a 2D slice of float64 type
*/
func ReadFile(file string) (int, [][]float64) {
	/*
		variables for the function
	*/
	var dimension int
	var matrix = []float64{}
	/*
		open file
	*/
	openFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	/*
		it will close the file to free memory
	*/
	defer openFile.Close()
	/*
		scan the information stored in the buffer
	*/
	scanner := bufio.NewScanner(openFile)
	/*
		in order to have complete control and have one number per column we need
		to split all the information by word and not by line becouse in index 0
		we can have the entire row of the file
	*/
	scanner.Split(bufio.ScanWords)
	/*
		create a slice and append all the words formatted in string
	*/
	var fileTextLines []string
	for scanner.Scan() {
		fileTextLines = append(fileTextLines, scanner.Text())
	}

	/*
		in dimension are stored the dimension of the matrix
		in convertMatrix are stored all the matrix numbers formatted in float64
		but is one dimensiona
	*/
	for i, line := range fileTextLines {
		if i == 0 {
			dimension, _ = strconv.Atoi(line)
		} else if i > 0 {
			convertMatrix, _ := strconv.ParseFloat(line, 64)
			matrix = append(matrix, convertMatrix)
		}

	}
	/*
		formatMatrix is a 2D slice and it will contain all the data from the matrix 1D slice
		it is the final matrix
	*/
	formatMatrix := make([][]float64, dimension)
	var incr = 0
	for i := 0; i < dimension; i++ {
		formatMatrix[i] = make([]float64, dimension)
		for j := 0; j < dimension; j++ {
			formatMatrix[i][j] = matrix[incr]
			incr++
		}
	}

	return dimension, formatMatrix
}

/*
	this function calculate the determminant of an NxN matrix only
*/
func Determinant(matrix [][]float64, dimension int) float64 {
	var det float64 = 0
	var row, i, j, subRow, subCol int

	if dimension == 1 {
		/*
			determinant of an 1x1 matrix
		*/
		det = matrix[0][0]
	} else if dimension == 2 {
		/*
			determiant of an 2x2 matrix
		*/
		det = matrix[1][1]*matrix[0][0] - matrix[0][1]*matrix[1][0]
	} else {
		/*
			determinant of an matrix upper 2x2 dimension
		*/
		subMatrix := make([][]float64, dimension)
		for row = 0; row < dimension; row++ {
			for i = 0; i < dimension-1; i++ {
				subMatrix[i] = make([]float64, dimension)
				for j = 0; j < dimension-1; j++ {
					if i < row {
						subRow = i
					} else {
						subRow = i + 1
					}
					subCol = j + 1
					subMatrix[i][j] = matrix[subRow][subCol]
				}
			}
			if row%2 == 0 {
				det += matrix[row][0] * Determinant(subMatrix, dimension-1)
			} else {
				det -= matrix[row][0] * Determinant(subMatrix, dimension-1)
			}
		}
	}
	return det
}

/*
	this function calculate the traspose of an NxN matrix
*/
func transpose(minorMatrix [][]float64, dimension int, det float64) {
	var i, j int
	inverseMatrix := make([][]float64, dimension)
	for i = 0; i < dimension; i++ {
		inverseMatrix[i] = make([]float64, dimension)
		for j = 0; j < dimension; j++ {
			inverseMatrix[i][j] = minorMatrix[j][i] / det
		}
	}
	/*
		print the inverse matrix
	*/
	printMatrix(inverseMatrix, dimension)
}

/*
	this function calculate the minor of an NxN matrix
*/
func minor(matrix [][]float64, dimension int, det float64) {
	/*
		alocating memory for two 2D slices
	*/
	extraMatrix := make([][]float64, dimension)
	copyMatrix := make([][]float64, dimension)
	var l, h, r, k, i, j int
	/*
		alocating memory for every column
		if printed the slice it will be like this
		[[] [] [] ... []]
		in each slice i can store a row of a matrix
		[[1 2 3] [3 4 5] [6 7 8]]
	*/
	for j = 0; j < dimension; j++ {
		extraMatrix[j] = make([]float64, dimension)
		copyMatrix[j] = make([]float64, dimension)
	}

	for h = 0; h < dimension; h++ {
		for l = 0; l < dimension; l++ {
			r = 0
			k = 0
			for i = 0; i < dimension; i++ {
				for j = 0; j < dimension; j++ {
					if i != h && j != l {
						copyMatrix[r][k] = matrix[i][j]
						if k < (dimension - 2) {
							k++
						} else {
							k = 0
							r++
						}
					}
				}
			}
			if h%2 == 0 || l%2 == 0 {
				extraMatrix[h][l] += 1 * Determinant(copyMatrix, dimension-1)
			} else {
				extraMatrix[h][l] -= 1 * Determinant(copyMatrix, dimension-1)
			}
		}
	}
	transpose(extraMatrix, dimension, det)
}

/*
	this function calculate the inverse of an NxN matrix
*/
func Inverse(matrix [][]float64, dimension int, det float64) {
	if det == 0 {
		fmt.Println("Matrix is NOT invertible")
	} else if dimension == 1 {
		fmt.Println("Inverse matrix")
		fmt.Printf("%f\n", 1/det)
	} else {
		minor(matrix, dimension, det)
	}
}

/*
	print matrix
*/
func printMatrix(matrix [][]float64, dimension int) {
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			fmt.Printf("%.4f   ", matrix[i][j])
		}
		fmt.Println()
	}
}
