# Interpreter
Interpreter for Monkey programming language.

## Get Started
Just execute following command:
```
go run main.go
```

And you will see a REPL environment:
```
Hello {username}! This is the Monkey programming language!
Feel free to type in commands
>>
```

## Developing and Testing
It's pretty interesting to add new features. This project is **Test-Deiven**. 
You should write test cases first, and then run those tests with:
```shell
go test -coverpkg=./... -cover -covermode=atomic -gcflags=all=-l ./... -coverprofile=coverage.txt
```

Now you should be able to see some error logs. Fix them with new features, and run the test again.