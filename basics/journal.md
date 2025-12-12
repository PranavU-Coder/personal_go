so well Go is a statically "strong" typed compiled language developed by Google, i personally seek to learn this language to well truly appreciate the hype behind its "great coroutines" that it offers and being a compiled language i believe the performance would be great.

some basic stuff out of the way: 

(1)"package main" lets the Go compiler know that we want this code to compile and run as a standalone program, as opposed to being a library that's imported by other programs.

(2)import "fmt" imports the fmt (formatting) package from the standard library. It allows us to use fmt.Println to print to the console.

(3)func main() defines the main function, the entry point for a Go program.

so like any other compiled language, it has two standard errors at surface level: compiled-time and runtime error.

Go code generally runs faster than interpreted languages and compiles faster than other compiled languages.

Go has an interesting amount of datatypes for us to play with:

(1)Signed integers : int  int8  int16  int32  int64

(2)Unsigned integers (non-negative) : uint uint8 uint16 uint32 uint64 uintptr

(3)Signed decimal numbers : float32 float64

(4)Complex numbers : complex64 complex128

Go has this special thing called "type inference" with the walrus operator ":=" which allows to declare and assign value to a variable in a single line.

generally when everything revolves around writing super "idiomatic and performant" code , default types are avoided.

statically typed in most general sense means that the types are known before the code ever runs  means any editor and compiler can display type errors before the code is ever run, making development easier and faster.

strong or weak refers to how strict the typechecking is conducted.

the chart mentioned in the image file gives an idea of memory footprint by various languages wrt each other.

Go programs are fairly lightweight. Each program includes a small amount of extra code that's included in the executable binary called the Go Runtime. One of the purposes of the Go runtime is to clean up unused memory at runtime. It includes a garbage collector that automatically frees up memory that's no longer in use.

Java uses a virtual machine to interpret bytecode at runtime and typically allocates more on the heap

Rust and C programs use slightly less memory than Go programs because more control is given to the developer to optimize the memory usage of the program. The Go runtime just handles it for us automatically.

const keywords can't utilize the walrus operator.

In Go, strings are just sequences of bytes: they can hold arbitrary data. However, Go also has a special type, rune, which is an alias for int32. This means that a rune is a 32-bit integer, which is large enough to hold any Unicode code point.

what this means is:

(1)When you need to work with individual characters in a string, you should use the rune type. It breaks strings up into their individual characters, which can be more than one byte long.

(2)We can include a wide variety of Unicode characters in our strings, such as emojis and Chinese characters, and Go will handle them just fine.

Unlike other languages, Go requires you to put the opening brace on the same line as the condition and not on a new line.

A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore, or more precisely, the blank identifier _ hence making it a placeholder.

A return statement without arguments returns the named return values. This is known as a "naked" return. Naked return statements should be used only in short functions. They can harm readability in longer functions.

nil is the zero value of an error. More on this later.

Guard Clauses leverage the ability to return early from a function (or continue through a loop) to make nested conditionals one-dimensional. Instead of using if/else chains, we just return early from the function at the end of each conditional block.

Go supports first-class and higher-order functions, which are just fancy ways of saying "functions as values". Functions are just another type -- like ints and strings and bools.

Anonymous functions are true to form in that they have no name. They're useful when defining a function that will only be used once or to create a quick closure.

The defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically just before its enclosing function returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

Deferred functions are typically used to clean up resources that are no longer being used. Often to close database connections, file handlers and the like.

A closure is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.

Function currying is a concept from functional programming and involves partial application of functions. It allows a function with multiple arguments to be transformed into a sequence of functions, each taking a single argument.
