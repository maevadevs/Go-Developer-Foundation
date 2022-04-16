# Maps

## Maps Basics

- Very similar to Structs but a collection of Key-Value Pairs
- Similar to a Dictionaries in Python
- But how is it different from a Struct? (Since Structs are also Key-Value Pairs)
  - **In Maps, both the Keys and Values are statically-typed**
  - All the keys must be of the same exact type
  - All the values must be of the same exact type
  - For example: `key<int> => value<string>`

```go
// A map of <string> => <string>
base_colors := map[string]string{
    "red":   "ff0000",
    "green": "00ff00",
    "blue":  "0000ff",
}
```

### Declaring a map

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
  - The value of above is `nil`: it does not point to an initialized map
  - A `nil` map behaves like an empty map when reading, but attempting to write to a `nil` map will cause a runtime `panic`
  - Do not do that! Instead, to initialize a map, use the built in `make()` function

```go
// USE THIS METHOD INSTEAD
colors := make(map[string]string)
fmt.Println("colors:", colors)
```

### Add value to an existing map

- To add values to an existing map, we can use `[]` syntax
- We cannot use the `dot`-syntax with maps because all the keys are typed

```go
colors["white"] = "ffffff"
colors["black"] = "000000"
fmt.Println("colors:", colors)
```

### Delete from an existing map

- To delete existing mappings inside a map, we can use the built-in function `delete()`

```go
delete(colors, "white")
fmt.Println("colors:", colors)
```

## Maps Iteration

- Very similar to iterating over slices
- To iterate over a map, we will first create a `hashmap` type
- This would allow us for an easier manipulation of the map

```go
type hashmap map[string]string
```

- Then, we can add a *receiver* function to this type that would allow us to handle iteration over the map
- We will also use this as a way to print the hashmap to the screen

```go
func (c hashmap) print() {

    for color, hex := range c {
        fmt.Printf("Colors{%s: %s}\n", color, hex)
    }
    
}
```

## Map vs Struct

- In the vast majority, we usually end up using Structs than Maps
- But it really depends on each use cases

MAPS (~Dictionaries/HashMaps)|STRUCTS (~Objects)
:-|:-
All keys must be of the same type|Values can be of different types
All values must be of the same type|Values can be of different types
Keys are indexed and are iterable|Keys are not indexed, cannot iterate over
To represent a collection of *related* properties|To represent an *object* with a lot of different properties
Don't need to know all the keys at compile time|Need to know all the different fields at compile time
**Reference Type**|**Value Type**