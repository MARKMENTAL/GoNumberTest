# GoNumberTest
GoNumberTest is a program made to test and benchmark the Go programming language. It will count numbers to a given limit and also display whether a number is even or odd, the time taken to process and how many multiples of 3 were encountered.

# GCCGO Build & Run Instructions
This program was built with GNU's Go Compiler: GCCGO,
to build from source, from within a folder containing `numtest.go` run these commands in a terminal

`gccgo -c numtest.go`
`gccgo -o numtest numtest.o`

Now you can run the program by typing in
`./numtest`
