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
=== RUN   TestEchoHandler
--- PASS: TestEchoHandler (0.00s)
=== RUN   TestInvertHandler
--- PASS: TestInvertHandler (0.00s)
=== RUN   TestFlattenHandler
--- PASS: TestFlattenHandler (0.00s)
=== RUN   TestSumHandler
--- PASS: TestSumHandler (0.00s)
=== RUN   TestMultiplyHandler
--- PASS: TestMultiplyHandler (0.00s)
=== RUN   TestEcho
--- PASS: TestEcho (0.00s)
=== RUN   TestInvert
--- PASS: TestInvert (0.00s)
=== RUN   TestFlatten
--- PASS: TestFlatten (0.00s)
=== RUN   TestSum
--- PASS: TestSum (0.00s)
=== RUN   TestMultiply
--- PASS: TestMultiply (0.00s)
PASS
ok
```