# matrix_transformer
A basic web server written in GoLang with the ability to perform various operations on a matrix given in the input file.

## Usages

Put the path of matrix_transformer under the go path.

Run web server
```
go run .
```

Send request
```
curl -F 'file=@/resources/matrix.csv' "localhost:8080/echo"
curl -F 'file=@/resources/matrix.csv' "localhost:8080/invert"
```

## Unit Test

Run the unit tests: 
```
go test
```