# Random number generator in Python and Go

## Getting started

### Python
Run Python example
```
$ make example-python
...
-1
2
0
1
... 
```

Run Python tests
```
$ make test-python
cd ./python && python3 -m unittest
...
----------------------------------------------------------------------
Ran 3 tests in 0.000s

OK
```

### Go
Run Go example
```
$ make example-go
...
-1
2
0
1
... 
```

Run Go tests
```
$ make test-go
cd ./go && go test -v -cover ./...
        github.com/ATMackay/go/randomgen                coverage: 0.0% of statements
=== RUN   Test_RandomGenInit
--- PASS: Test_RandomGenInit (0.00s)
=== RUN   Test_CumulativeProbability
--- PASS: Test_CumulativeProbability (0.00s)
=== RUN   Test_RandomGenNextNum
--- PASS: Test_RandomGenNextNum (0.00s)
PASS
coverage: 94.1% of statements
ok      github.com/ATMackay/go/randomgen/randomgen      0.003s  coverage: 94.1% of statements
```