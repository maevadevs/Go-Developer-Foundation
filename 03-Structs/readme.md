# Structs

---

- [Struct Basics](#struct-basics)
- [Struct With Receiver Function](#struct-with-receiver-function)
- [Pointers and Values](#pointers-and-values)
  - [Using Pointer](#using-pointer)
  - [About Pointers](#about-pointers)
  - [Pointer Gotchas](#pointer-gotchas)
    - [List Of Data Structures That Are Passed BY REFERENCE (REFERENCE TYPES)](#list-of-data-structures-that-are-passed-by-reference-reference-types)
    - [List Of Data Structures That Are Passed BY VALUE (VALUES TYPES)](#list-of-data-structures-that-are-passed-by-value-values-types)

---

## Struct Basics

- A Data Structure used widely throughout Go
- A collection of different properties that are related
- More flexible properties than just a single value type (e.g. string or number)
- Similar to Python Dictionary, but it is not an *Object*
- **The type of each field inside the struct can be different**
- Can be used to represent complex data structures
- **We can embbed one struct inside another struct to build complex data structures**
  - We can reuse structs in any way we want
  - This allows *Composition* of different data structures
  - **When embedding another struct, we can also skip the `fieldName` and only use the struct type**
    - This will have the effect of creating the `fieldName` the same name as the struct name
    - *If the `fieldName` is different than the struct name, then it needs to be explicit*

```go
// Declaring new structs
type contactInfo struct {
    email   string
    zipCode int
}

type person struct {
    firstName string
    lastName  string
    // Embedded struct
    contact   contactInfo
}
```

- Now, we can start making use of the struct in the `main()` function
- The order of the fields can be used when defining a struct instance
  - However, this is not the recommended approach

```go
func main() {
    // Creating a value of type person
    // Using field order: This is not the recommended approach
    alex := person{
        "Alex",        // firstName
        "Anderson",    // lastName
        // Embedded struct
        // When embedding another struct, we can also skip the fieldName
        // We can just use the struct type
        // This has the effect of creating the fieldName the same name as the struct name
        contactInfo{   // contactInfo
            "alex@anderson.com", // email
            12345                // zipCode
        }
    }

    // Show the struct
    fmt.Println(alex)
}
```

- **Instead of relying on the order of fields, use *named fields***
  - This is the recommended approach

```go
func main() {
    // Creating a value of type person
    // Using named fields: This is the recommended approach
    maria := person{
        firstName: "Maria",
        lastName:  "Anderson",
        // Embedded struct
        contact: contactInfo{
            email:   "maria@anderson.com",
            zipCode: 98765
        }
    }

    // Show the struct
    fmt.Println(maria)
}
```

- **We can also declare then assign later**
  - The struct's fields would be assigned their zero-value by default before a new value is assigned to it

Data Type|Zero-Value
:-|:-
`struct`|`null`
`string`|`""`
`int`|`0`
`float64`|`0`
`bool`|`false`

```go
func main() {
    // Declare variable:
    // Fields default to null-value
    var julie person

    // To print a struct with key:value format, use %+v
    fmt.Printf("julie before assignment: %+v\n", julie)

    // Assigning / Re-assigning values
    julie.firstName = "Julie"
    julie.lastName = "Arkorius"
    julie.contact.email = "j.arkorius@somemail.com"
    julie.contact.zipCode = 12345

    // Check again
    fmt.Printf("julie after assignment: %+v\n", julie)
}
```

## Struct With Receiver Function

- Similar to *Types*, we can define receiver functions to be attached to *Structs* as well
- **Note: The code below has some issues because of *Pointers*. We will review that below**

```go
// Receiver function for "person"
func (p person) print() {
    fmt.Printf("%+v\n", p)
}

// This function has some issues because of pointer
func (p person) updateFirstName(newFirstName string) {
    p.firstName = newFirstName
}

func main() {
    // Defining a new person variable
    jim := person{
        firstName: "Jim",
        lastName:  "Patterson",
        contact: contactInfo{
            email:    "jim@patterson.com",
            zipCode: 98765,
        },
    }

    // Calling a receiver function
    jim.print()
    jim.updateFirstName("Jimmy")
    jim.print()

    // Why did it not update?
    // Because of how Pointers work
}
```

## Pointers and Values

- Previously, we saw that calling `jim.updateFirstName()` did not update `jim.firstName` field
- Why did it not update the first name to `"Jimmy"`?
  - Because of Pointers in Go
  - **A struct variable is a pointer type**
  - Points to a value stored in memory
- **By default, Go is a *pass-by-value* language**
  - With the way `jim.updateFirstName()` is defined, the *value stored in `jim` is copied* and stored in a new location in memory, which is then used by `updateFirstName()`
  - However, the original variable `jim` is still *pointing* to the old *value*
  - So, we are not updating the original *struct* but a copy of it in a different memory location
- *Posible Solution: We could re-assign back to the original variable*
  - However, this is not preferred because it duplicates values in memory, thus inneficient
  - Also, if the struct is large, we have to move a large data 2 times in memory
    - 1. Original -> Copy
    - 2. Modify Copy
    - 3. Copy -> Original
- **Actual Solution: To solve this issue, we need to use *Pointers* to force Go to *pass-by-reference* and update the original**

### Using Pointer

- We use the `&varName` to create a reference (memory address) to `varName`

```go
// JimPtr contains the memory address pointing to jim
jimPtr := &jim
```

- We use the *pointer type* to pass to the function that wants to update the original variable

```go
// Before
// jim.updateFirstName("Jimmy")

// After
jimPtr.updateFirstName("Jimmy")
```

- This requires the receiver function definition to change as well: *It needs to take a Pointer-type as argument instead*

Syntax|Decription
:-|:-
`varName *type`|A Pointer-type variable that points to a value of the type (i.e. contains the reference/memory address to the value of the type)

```go
// updateFirstNamePointer takes a Pointer-type that points to a value of type person
func (ptrPers *person) updateFirstNamePointer(newFirstName string) {
    ...
}
```

- Within the receiver function definition, we *de-reference the pointer* (i.e. get back the value it is pointing at) using the `*varName` or `(*varName)` syntax

Syntax|Decription
:-|:-
`*varName` or `(*varName)`|Operator to access the value that exists at the memory address (pointer)

```go
// Using pointer as receiver allows us to pass-by-reference
func (ptrPers *person) updateFirstNamePointer(newFirstName string) {
    // De-reference the pointer type
    // This means updating the value at the reference/pointer
    (*ptrPers).firstName = newFirstName
}
```

- **However, Go can automatically de-reference a pointer as well**
- So the following also works

```go
// Using pointer as receiver allows us to pass-by-reference
func (ptrPers *person) updateFirstNamePointer(newFirstName string) {
    // Go can automatically de-reference the pointer
    ptrPers.firstName = newFirstName
}
```

- Now, in `main()`, we can make use of this receiver function
- We use the `&` sign to create a reference

Syntax|Decription
:-|:-
`&varName` or `(&varName)`|Operator to get the reference (memory address) that points to variable

```go
func main() {
    jim := person{
        firstName: "Jim",
        lastName:  "Patterson",
        contact: contactInfo{
            email:    "jim@patterson.com",
            zipCode: 98765,
        },
    }

    // Checking jim before calling a receiver function
    jim.print()

    // Updating jim: Passing By Reference
    (&jim).updateFirstNamePointer("Jimmy")

    // Checking jim after calling a receiver function
    jim.print()
}
```

- **However, with Go, it is possible to substitute a pointer with its root variable**
- So the following still work, even if the receiver requires a pointer type

```go
// Updating jim
jim.updateFirstNamePointer("Big ol'Jim")
jim.print()
```

- TLDR
  - **Make sure to define the parameter(s) of the receiver function to be of Pointer-type `varName *type`**
  - Optional: Within the receiver function, we de-reference the pointer with `*varName` or `*(varName)` to access/update the value at the Pointer-type's address
    - But Go can automatically de-reference, so using just `varName` also works
  - Optional: When calling the receiver function (in `main`), pass a memory reference to it using `&varName` or `&(varName)`
    - But Go can automatically reference, so using just `varName` also works

### About Pointers

Syntax|Decription
:-|:-
`&varName` or `(&varName)`|Operator to get the reference (memory address) that points to variable's value: This is *Referencing a Pointer*
`*varName` or `(*varName)`|Operator to access the value that exists at the memory address/pointer `&varName`: This is *De-referencing a Pointer*
`varName`|The original variable that we asigned the value to: Turns into its equivalent memory address reference using `&`
`*type`|A Pointer type that point to a memory address whose value is the type: `&varName` is of type `*type`

```go
var jon person = person{firstName: "Jon", lastName: "Leukippos"}
var ptrJon *person = &jon
var jonVal person = *ptrJon
var jonCopy person = jon
```

`varName`|varType|Memory Address (Reference) `&varName`|Value (Dereference) `*ptrVarName`|Pointer `*type`
:-:|:-:|:-:|:-:|:-:
`jon`|`person`| `&jon` same as `ptrJon`<br><br>`0x1F4CB3`|`*ptrJon` same as `jonVal`<br><br>`{firstName: "Jon", lastName: "Leukippos"}`<br><br>Same value as `jonCopy` but different *objects*: Original vs Copy|`*person`<br><br>Memory Address type referencing to a value of type `person`

### Pointer Gotchas

- Go is typically a *Pass-By-Value* language
- **However, Slices *seems* to by default *passed-by-reference (Pointer)* for better performance**
  - **NOTE:** *Arrays* are *passed-by-value* though!!
- In reality though, Go is *always* passed by value, but slices use a trick
- When we create a slice, Go internally creates 2 separate data structure:
  - The *Slice* Data Structure
    - Pointer to the underlying array -> Points to a different address in memory
    - Capacity
    - Length
  - An *underlying Array*
    - Contains the actual list of items
- **When we modify the array, the Slice Data Structure gets moved around *by value*, but it is still pointing to the same underlying Array**
  - The *Slice Data Structure* is passed around by value
  - But it always point to the *underlying array*
  - The *Slice Data Structure* gets modified (by value) when passed around
  - But the reference to the *undelying array* data structure remains

```go
func main() {
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

#### List Of Data Structures That Are Passed BY REFERENCE (REFERENCE TYPES)

No need to use pointers with these

- Slice
- Map
- Channel
- Pointer
- Function

#### List Of Data Structures That Are Passed BY VALUE (VALUES TYPES)

Use pointers to change the underlying values for these in functions

- int
- float
- string
- bool
- struct
- Array
