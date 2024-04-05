# Maps

---

- [Maps Basics](#maps-basics)
  - [Declaring A Map](#declaring-a-map)
  - [Add Value To An Existing Map](#add-value-to-an-existing-map)
  - [Delete From An Existing Map](#delete-from-an-existing-map)
- [Maps Iteration](#maps-iteration)
- [Map vs Struct](#map-vs-struct)

---

## Maps Basics

- Very similar to Structs but a collection of Key-Value Pairs
- Similar to a *Dictionary* in Python
- But how is it different from a Struct? (Since Structs are also Key-Value Pairs)
  - **In Maps, both the Keys and Values are statically-typed**
  - All the keys must be of the same exact type
  - All the values must be of the same exact type
  - For example: `key int => value string`

```go
// A map of key<string> => value<string>
baseColors := map[string]string{
    "red":   "ff0000",
    "green": "00ff00",
    "blue":  "0000ff",
}
```

### Declaring A Map

- Can we also declare a map using `var`?
- Default zero-value for map is an empty-map: `map[]`
- However, there is a gotcha

```go
// DO NOT USE THIS METHOD!
var colors map[string]string
fmt.Println("colors:", colors)
```

- Why that did not work?
  - Map types are reference types, like pointers or slices
  - Default value of reference types is `nil`
  - The value of above is `nil`: it does not point to an initialized map
  - **A `nil` map behaves like an empty map when reading, but attempting to write to a `nil` map will cause a runtime `panic`**

```go
var colors map[string]string
colors["red"] = "ff0000"
// panic: assignment to entry in nil map
```

- **Do not do that! Instead, to initialize a map, use the built in `make()` function**

```go
// USE THIS METHOD INSTEAD
colors := make(map[string]string)
colors["red"] = "ff0000"
```

- `make()` allocates and initializes a hashmap data structure
  - Returns a `map` value that points to it
  - The specifics of that data structure are an implementation detail of the runtime

### Add Value To An Existing Map

- To add values to an existing map, we can use `[]` syntax
- **We cannot use the `dot` syntax with maps because all the keys are typed**

```go
colors["white"] = "ffffff"
colors["black"] = "000000"
fmt.Println("colors:", colors)
```

### Delete From An Existing Map

- To delete existing mappings inside a map, we can use the built-in function `delete()`
- Pass the key for the item to delete

```go
delete(colors, "white")
fmt.Println("colors:", colors)
```

## Maps Iteration

- Very similar to iterating over slices
- To iterate over a map, we will first create a `custommap` type
- This would allow us for an easier manipulation of the map

```go
type custommap map[string]string
```

- Then, we can add a *receiver* function to this type that would allow us to handle iteration over the map
- We will also use this as a way to print the map to the screen

```go
func (cm custommap) print() {
    for k, v := range cm {
        fmt.Printf("Colors{%s: %s}\n", k, v)
    }
}
```

## Map vs Struct

- In the vast majority, we usually end up using Structs than Maps
- But it really depends on each use cases

MAPS (~Dictionaries/HashMaps)|STRUCTS (~Objects)
:-|:-
All keys must be of the same type|Keys are not strongly type: The are pre-defined fields
All values must be of the same type|Values can be of different types
Keys are indexed and are iterable|Keys are not indexed, cannot iterate over
To represent a collection of *related* properties|To represent an *object* with a lot of different properties
**Dynamic Keys**: Don't need to know all the keys at compile time: Can add and delete keys|**Static Keys**: Need to know all the different fields at compile time: Fixed number of keys
**Reference Type: No need to use pointers with it**|**Value Type: Might need to use pointers to access original**
