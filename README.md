<p align="center">
    Creek is a Streams library written in Go.<br>
    It helps you to write functional programming code using generics.
</p>

<hr>

# Table of Contents
- [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Quick start](#quick-start)
  - [Examples](#examples)
    - [Create stream](#create-stream)
    - [Collect](#collect)
    - [Sorting](#sorting)
    - [Filter](#filter)
    - [ForEach](#foreach)
    - [Map](#map)
    - [Limit](#limit)
  - [Download and build from source](#download-and-build-from-source)
  - [Contributing](#contributing)
  - [License](#license)

<hr>

## Installation
Before installing this library, you need to have [Go](https://go.dev/dl/) installed.  
Since creek uses generics, version 1.18+ is required.  
Now, you can use this command:
```sh
go get -u github.com/0l1v3rr/creek
```

<hr>

## Quick start
```go
package main

import (
	"fmt"

	"github.com/0l1v3rr/creek"
)

func main() {
	arr := []int{1, 8, 2, 14, 22, 4, 7, 92}

	result := creek.FromArray(arr).Filter(func(item int) bool {
        return item > 3
    }).OrderBy().Collect()

	fmt.Println(result) // [4 7 8 14 22 92]
}
```

<hr>

## Examples

### Create stream
You can create a stream from almost every type of array or slice.  
The supported types are the following:
```go
string | byte | float32 | float64 | int | int16 | int32 | int64 | uint16 | uint32 | uint64
```
In order to create a Stream, use the `FromArray` function:
```go
// slice
arr := []int{1, 8, 2, 14, 22, 4, 7, 92}
intStream := creek.FromArray(arr)

// array
arr2 := [3]string{"One", "Two", "Three"}
stringStream := creek.FromArray(arr2[:])
```

### Collect
The `Collect` function returns the modified array from the streams.
```go
arr := []int{1, 8, 2, 14, 22, 4, 7, 92}
arrFromStream := creek.FromArray(arr).Collect() // [1, 8, 2, 14, 22, 4, 7, 92]
```

### Sorting
There are two ways to Sort a stream. **Ascending** and **Descending**.  
The `OrderBy` function sorts the stream in ascending order.  
The `OrderByDescending` function sorts the stream in descending order.
```go
arr := []int{1, 8, 2, 14, 22, 4, 7, 92}

ascStream := creek.FromArray(arr).OrderBy()             // [1, 2, 4, 7, 8, 14, 22, 92]
descStream := creek.FromArray(arr).OrderByDescending()  // [92, 22, 14, 8, 7, 4, 2, 1]
```

### Filter
The `Filter` function leaves only those elements in the array that make the specified condition true.
```go
arr := []string{"One", "Two", "Three"}
result := creek.FromArray(arr).Filter(func(item string) bool {
    return strings.HasPrefix(item, "T")
}) // [Two, Three]

// ----------------------------------------------

arr2 := []int{1, 4, 6, 2, 7, 3}
result2 := creek.FromArray(arr2).Filter(func(item int) bool {
    return item%2 == 0
}) // [4, 6, 2]
```

### ForEach
The `ForEach` method runs the specified method on every element in the Stream.  
Warning: this method doesn't return anything
```go
arr := []string{"One", "Two", "Three"}
creek.FromArray(arr).ForEach(func(item string) {
    fmt.Println(item)
})

// -- Output: --
// One
// Two
// Three
```

### Map
The `Map` function creates a new stream populated with the results of calling the provided function on every element.
```go
arr := []string{"One", "Two", "Three"}
result := creek.FromArray(arr).Map(func(item string) string {
    return strings.ToUpper(item)
}) // [ONE, TWO, THREE]
```

### Limit
The `Limit` function constrains the number of elements returned by the stream.
```go
result := creek.FromArray([]int{2, 7, 3, 1}).Limit(2).Collect() // [2, 7]
```

<hr>

## Download and build from source
First, clone the repository:
```sh
https://github.com/0l1v3rr/creek.git
cd creek
```
- Running the tests: `make test`
- Running the test in detailed version: `make test-detailed`
- Running the main: `make run`

<hr>

## Contributing
You can find [here](CONTRIBUTING.md) a contributing guideline.  
A star would be appreciated as well. 😉

<hr>

## License
This project is licensed under the [MIT License](LICENSE).