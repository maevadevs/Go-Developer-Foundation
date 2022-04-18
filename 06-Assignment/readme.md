# Assignments

## Shapes

- Write a program that creates 2 custom struct types:
  - `triangle`
  - `square`
- The `square` type should be a struct with a field called `sideLength` of type `float64`
- The `triangle` type should be a struct with a field called `height` of type `float64` and a field called `base` of type `float64`
- Both types should have function called `getArea` that returns the calculated area of the square or triangle
  - Area of Triangle = 0.5 * base * height
  - Area of Square = sideLength * sideLength
- Add a `shape` interface that defines a function called `printArea`
  - This function should calculate the area of the given shape and print out to the terminal
  - Design the interface so that the `printArea` function can be called with either a triangle or a square

## Files

- Create a program that reads the contents of a text file then print its contents to the terminal
  - The file to open should be provider as an argument to the program when it is executed at the terminal
  - For example, `go run main.go myfile.txt`
- To read-in the arguments provided to a program, you can reference the variable `os.Args`, which is a slice of strings
- To open a file, check out the documentation for the `Open` function in the `os` package
- What interfaces does the `File` type implement?
- If the `File` type implements the `Reader` interface, you might be able to reuse that `io.Copy` function