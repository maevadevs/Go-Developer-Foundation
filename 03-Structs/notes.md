# Structs

## Struct basics

- A Data Structure used widely throughout Go
- A collection of properties that are related together
- More flexible properties than just a single string type
- Similar to Python Dictionary, but not an *Object*
  - The type of each field inside the struct can be different
- Can be used to represent complex data structures
- We can embbed one struct inside of another struct

```go
// Declaring new structs
type contactInfo struct {
    email    string
    zip_code int
}

type person struct {
    first_name string
    last_name  string
    contact    contactInfo
}
```

- Now, we can start making use of the structs in the `main()` function
- The order of the fields can be used when defining a struct instance
  - However, this is not a recommended approach

```go
func main() {

    // Creating a value of type person
    // Using field order: This is not the recommended approach
    alex := person{
        "Alex",                     // first_name
        "Anderson",                 // last_name
        contactInfo{                // contact
            "alex@anderson.com",
            12345,
        },
    }
    fmt.Println(alex)

}
```

- Instead of relying on the order of fields, use named fields
  - This is the recommended approach

```go
func main() {

    // Creating a value of type person
    // Using named fields: This is the recommended approach
    maria := person{
        first_name: "Maria",
        last_name:  "Anderson",
        contact: contactInfo{
            email:    "maria@anderson.com",
            zip_code: 98765,
        },
    }
    fmt.Println(maria)

}
```

- We can also declare then assign later
- The variable would be assigned its zero-value by default before a new value is assigned to it
  - string -> `""`
  - int -> `0`
  - float64 -> `0`
  - bool -> `false`

```go
func main() {

    // Declare variable: Default to null-value
    var julie person

    // To print a struct with key:value format, use %+v
    fmt.Printf("%+v\n", julie)

    // Assigning values
    julie.first_name = "Julie"
    julie.last_name = "Arkorius"

    // Check again
    fmt.Printf("%+v\n", julie)

}
```

## Struct with receiver function

- Similr to *Types*, we can define receiver functions to be attached to structs as well
- **Note:** The code below has issue because of *Pointers*. We will review that below

```go
// Receiver function for "person"
func (p person) print() {
    fmt.Printf("%+v\n", p)
}

func (p person) updateName(new_first_name string) {
    p.first_name = new_first_name
}

func main() {

    // Defining a new person variable
    jim := person{
        first_name: "Jim",
        last_name:  "Patterson",
        contact: contactInfo{
            email:    "jim@patterson.com",
            zip_code: 98765,
        },
    }

    // Calling a receiver function
    jim.print()
    jim.updateName("Jimmy")
    jim.print()

    // Why did it not update? --> Pointers

}
```

## Pointers and values

- Previously, we saw that calling `jim.updateName()` did not update the variable
- Why did it not update the name to `"Jimmy"`?
  - Because of Pointers in Go
  - A struct variable is a pointer type
    - Points to a value stored in memory
- By default, Go is a *pass-by-value* language
  - With the way `jim.updateName()` is defined, the *value* is copied and stored in a new location in memory
  - However, the variable `jim` is still *pointing* to the old *value*
  - We are not updating the original *struct*
  - Instead, we are copying the newly assigned *value* into a different location is memory
- To solve this issue, we need to use *Pointers* to force Go to *pass-by-reference*
  - We use the `&` sign to create a reference
  - `*var` - Operator to access the value that exist at the memory address (pointer)
  - `*<type>` - A type of Pointer that point to a type

```go
// Using pointer as receiver allows us to pass-by-reference
func (ptr *person) updateNamePointer(new_first_name string) {
    (*ptr).first_name = new_first_name
}
```

- Now, in `main()`, we can make use of this receiver function
- We use the `&` sign to create a reference
  - `&` - Operator to access the memory address that the variable is pointing to

```go
func main() {

    jim := person{
        first_name: "Jim",
        last_name:  "Patterson",
        contact: contactInfo{
            email:    "jim@patterson.com",
            zip_code: 98765,
        },
    }

    // Checking jim before calling a receiver function
    jim.print()

    // Updating jim
    (&jim).updateNamePointer("Jimmy")

    // Checking jim after calling a receiver function
    jim.print()

}
```

- However, with Go, it is possible to substitute a pointer with its root variable
- So the following still work, even if the receiver requires a pointer type

```go
// Updating jim
jim.updateNamePointer("Big ol'Jim")
jim.print()
```

### About pointers

- `&` - Operator to access the memory address that the variable is pointing to
- Variable Pointer (Reference)
  - Points to a memory address
  - Turn into equivalent Value using *
- Variable Value
  - Contains the actual value stored in a memory address
  - Turn into equivalent Pointer using &
- `*var` - Operator to access the value that exist at the memory address (pointer)
- `*<type>` - A Pointer type that point to a memory address whose value is the type

### Pointer gotchas

- Go is typically a *Pass-By-Value* language
- However, Slices are by default *passed-by-reference (Pointer)* for better performance
  - **NOTE:** *Arrays* are *passed-by-value* though!!
- When we create a slice, Go internally creates 2 separate data structure:
  - The Slice Data Structure
    - Pointer to the underlying array -> Points to a different address in memory
    - Capacity
    - Length
  - An underlying Array
    - Contains the actual list of items
- When we modifying the array, the Slice Data Structure gets moved around, but it is still pointing to the same underlying Array

```go
funct main() {

    mySlice := []string{"Hi", "there", "how", "are", "you"}
    myArray := [4]string{"This", "is", "an", "Array"}

    updateSlice(mySlice) 
    // Works because passed by reference
    // Updating on the argument = updating the same memory address

    updateArray(myArray) 
    // Does not work because passed by value
    // Updating on the argument = updating a different copied value
    // Need to use pointer to update the same variable

    fmt.Println(mySlice)
    fmt.Println(myArray)

}

// Helper Functions
func updateSlice(s []string) {
    s[0] = "Bye"
}

func updateArray(arr [4]string) {
    arr[0] = "What"
}
```

#### List of Data Structures that are passed BY REFERENCE (REFERENCE TYPES)

No need to use pointers with these

- Slice
- Maps
- Channels
- Pointers
- Functions

#### List of Data Structures that are passed BY VALUE (VALUES TYPES)

Use pointers to change the underlying values for these in functions

- int
- float64
- string
- bool
- struct
