# Going

Dumping ground for tutorials and other nonsense coming out of learning to code with the Go programming language.

## "Learn GO Fast: Full Tutorial" by Alex Mux

[https://www.youtube.com/watch?v=8uiZC0l4Ajw](https://www.youtube.com/watch?v=8uiZC0l4Ajw)

After installing Go and creating the project folder, we run `go mod init (name of our module)` to produce the go.mod file.

This file will be used to track packages (files in our project) and modules (collections of packages), and probably dependencies?

We create the `main.go` with the `package main` and `func main` to use as the entry point into the app. The tutorial has us put that in the folders: `cmd/tutorial_1`

We import the fmt package built into Go with `import "fmt"`, then use it to print a line to the console.

Thus, we've got a Hello World.

To build it we run `go build cmd/tutorial_1/main.go`, that produces a `main.exe` file.
We can run that using just the exe's name `./main`.

We can also, instead of building and running the exe, simply use the `go run` command to compile and run our project with that one command `go run cmd/tutorial_1/main.go`.

## "Learn Go" course by Boot.dev (from Back-end Developer Path)

[https://www.boot.dev/tracks/backend-python-golang](https://www.boot.dev/tracks/backend-python-golang)

After installing and setting things up I checked out this course and just started appending it's exercises to the main function in 'cmd\tutorial_1'. The first section in the course is for variables, so I'm going to get through that and probably restructure the project. Especially when I learn how to create external modules or like... something to organize methods into (does Go have classes?). With that I'd be putting different lessons into their own files and probably having a little i/o looping routine to access them at runtime.

### Notes: 

Using the ['fmt' package](https://pkg.go.dev/fmt) to print text to the terminal. I'm impressed by the verbosity of the golang documentation. I wish I had infinite time to browse through the [standard library docs](https://pkg.go.dev/std).

Go is a compiled language and it generally compiles much faster than other languages like Rust, Zig and C ... but it's execution performance is slower than them. By virtue of being a compile language it's execution performance is still faster than interpreted languages like python, javascript and ruby.

``String access``: Getting the first character of a string isn't as simple as other langs. It looks like using the array accessor like someStr[0] actually accesses the byte at that location. So that means it has to be cast back to a string if you're going to do string stuff with that afterward, since the result is a byte.

``Optional parameters``: Looks like Go doesn't allow you to use optional parameters in functions, so from what I'm seeing online a common solution is to either overload the function or to create a struct that defines a set of properties that may be passed in as options.

``Random Links``
- [func params](https://www.w3schools.com/go/go_function_parameters.php)
- [accessing first char in str](https://www.bacancytechnology.com/qanda/golang/idiomatic-go-equivalent-of-cs-ternary-operator)
- [accessing first char in str (reddit)](https://www.reddit.com/r/golang/comments/krxahr/how_do_i_make_optional_function_parameter_in_go/)
- [accessing first char in str (article)](https://tipseason.com/how-to-get-first-character-in-string-golang/)
- [learning about interfaces](https://go.dev/doc/effective_go#interfaces)

['Which Type Should I Use?'](https://www.boot.dev/lessons/98e60d90-0111-4626-a690-70124be1e0ba) : Basically just use the default Go types unless you're really concerned about performance. 