[![Run on Repl.it](https://repl.it/badge/github/matt4biz/go-class-contains)](https://repl.it/github/matt4biz/go-class-contains)

# Go class: JSON contains example
This class demonstrates using reflection to test a JSON object to see if it contains certain desired fields (without unmarshalling the entire thing into a struct, for example).

The main program gets some JSON from Typicode's test server:

```shell
$ go run ./cmd
{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}

$ go run ./cmd 2
2020/07/26 11:03:08 id unmatched (1)
exit status 1

$ go run ./cmd 404
2020/07/26 11:03:11 404 Not Found
exit status 1
```

The test program runs over some samples, checking for things that exist and don't.

### Exercise

Add a couple of test snippets that increase code coverage of error cases in `contains()` where it checks the map case.

You can get code coverate with

```shell
$ go test ./... -cover
```
or better 

```shell
$ go test ./... -coverprofile=c.out -covermode=count
```

and see it with 

```shell
$ go tool cover -html=c.out
```

which will open a browser window for you.

### Another exercise

Add support for bool and array types, and handle nulls.
