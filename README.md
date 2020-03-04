# matrix_transformer
A basic web server written in GoLang with the ability to perform various operations on a matrix given in the input file.


## Usages

Run web server
```
go run .
```

Send request
```
curl -F 'file=@/resources/matrix.csv' "localhost:8080/echo"
curl -F 'file=@/resources/matrix.csv' "localhost:8080/invert"
curl -F 'file=@/resources/matrix.csv' "localhost:8080/sum"
curl -F 'file=@/resources/matrix.csv' "localhost:8080/flatten"
curl -F 'file=@/resources/matrix.csv' "localhost:8080/multiply"
```

## Unit Test

Run the unit tests: 
```
go test -v
```