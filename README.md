## Closure example for go serve mux

To run:
```
 git clone https://github.com/guygrigsby/closure-example.git
 cd closure-example
 go run main.go
```

Endoint:
| Method | endpoint |
|--------|----------|
| GET | /standard |
| GET | /closure |

Example:
```
curl localhost:8080/standard
```
or
```
curl localhost:8080/closure
```

Excersises for the reader:

Introduce another fixed value to the closure that could be important to the business logic of an http handler.
Examples:
 - Pass is some config values.
   - A domain name
   - A list of users
 - Pass in a data structure to emulate a database, a map for example
 To create and use a map in GO:

 ```go
 func main() {
	m := make(map[string]string)
	m["food"] = "burrito"

	println(m["food"]) // prints "burrito"
}
 ```
