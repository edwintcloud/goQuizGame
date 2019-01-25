# goQuizGame 
A quiz game in Go that generates a quiz from a csv file and presents it via the command line with specified options.

# Compiling The Game
```
go build
```

# Running The Game
There are some command line options for starting the binary. These can also be viewed by running the binary with the `-h` flag.
```
Usage of ./goQuizGame:
  -csv string
    	a csv file in the format of 'question,answer' (default "problems.csv")
  -limit int
    	the time limit for the quiz in seconds (default 30)
```
Run the game by running the compiled binary with specified options. Ex: `./goQuizGame -csv="myproblems.csv"`
