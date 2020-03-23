#  Inverse matrix
## Small cmd programm, it calculate the determinant and the inverse of an NxN matrix from file

I just started to study the Go programming language, to understand how slices work a wrote a small program, and i learned a lot

- clone the repository
- enter in the repository folder and execute `go mod vendor` to download the dependecy (I used cobra package)
- execute `go run matrix.go` to se the help
- this program needs an file.txt in `/samples` there is an example
- the first row is the dimension of the matrix then write your matrix
- to calculate the determinant of the matrix in the `/samples` folder execute `go run matrix.go det -f ./samples/1.txt`
- to calculate the inverse of an matrix instead of `det` command use the `inv`
- if you are mor confortable use `go build -o matrix` to have the executable
