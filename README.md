#  Inverse matrix
## Small CLI programm, it calculate the determinant and the inverse of an NxN matrix from file

I just started to study the Go programming language, to understand how slices work I wrote a small program to calvulate the determinant and the inverse of an NxN matrix.

- clone the repository
- enter in the repository folder and execute `go mod vendor` to download dependecy (I used cobra package)
- execute `go run matrix.go` to se the help
- this program needs an file.txt,  in `/samples` there is an example
- the first row is the dimension of the matrix, next keep an empty row and then write your NxN matrix
- to calculate the determinant of the matrix in the `/samples` folder execute `go run matrix.go det -f ./samples/1.txt`
- to calculate the inverse of an matrix instead of `det` command use the `inv`
- if you are more confortable use `go build -o matrix` to have the executable
