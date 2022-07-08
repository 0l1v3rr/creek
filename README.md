# Creek

Creek is a tiny Streams library written in Go.<br>
It helps you to write functional programming code using generics.
<br>
<img src="https://img.shields.io/github/release/0l1v3rr/creek.svg?style=flat-square">
<img src="https://img.shields.io/github/workflow/status/0l1v3rr/creek/Test?style=flat-square">
<img src="https://img.shields.io/github/repo-size/0l1v3rr/creek?style=flat-square">
<img src="https://img.shields.io/github/license/0l1v3rr/creek?style=flat-square">
<img src="https://img.shields.io/github/go-mod/go-version/0l1v3rr/creek?style=flat-square">

<hr>

# Table of Contents
- [Creek](#creek)
- [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Quick start](#quick-start)
  - [Examples](#examples)
    - [Create stream](#create-stream)
    - [Collect](#collect)
    - [OrderBy, OrderByDescending](#orderby-orderbydescending)
    - [Filter](#filter)
    - [ForEach](#foreach)
    - [Map](#map)
    - [Limit](#limit)
    - [Append](#append)
    - [AppendAt](#appendat)
    - [Remove](#remove)
    - [RemoveAt](#removeat)
    - [RemoveWhere](#removewhere)
    - [RemoveDuplicates](#removeduplicates)
  - [Method chaining](#method-chaining)
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

### OrderBy, OrderByDescending
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
arr := []int{2, 7, 3, 1}
result := creek.FromArray(arr).Limit(2) // [2, 7]
```

### Append
The `Append` function adds an element to the stream.
```go
arr := []int{2, 7, 3, 1}
result := creek.FromArray(arr).Append(2) // [2, 7, 3, 1, 2]
```

### AppendAt
The `AppendAt` function inserts the specified element at the specified position in the stream.  
If the index is out of range, nothing happens.  
The first parameter is the index, and the second is the item.
```go
arr := []int{2, 7, 3, 1}
result := creek.FromArray(arr).AppendAt(1, 14) // [2, 14, 7, 3, 1]
```

### Remove
The `Remove` function removes the passed element from a stream.
```go
arr := []int{2, 7, 3, 1, 3}
result = creek.FromArray(arr).Remove(3) // [2, 7, 1]
```

### RemoveAt
The `RemoveAt` function removes the item if its index matches the index passed in.  
If the index is out of range, nothing happens.
```go
arr := []int{2, 7, 3, 1}
result := creek.FromArray(arr).RemoveAt(2) // [2, 7, 1]
```

### RemoveWhere
The `RemoveWhere` function removes all the entries that satisfy the provided condition.
```go
arr := []int{2, 7, 3, 1}
result := creek.FromArray(arr).RemoveWhere(func(item int) bool {
    return item > 2
}) // [2, 1]
```

### RemoveDuplicates
The RemoveDuplicates function removes every duplicate item from the stream.
```go
arr := []int{2, 7, 3, 1, 3, 9, 3}
result := creek.FromArray(arr).RemoveDuplicates() // [2, 7, 3, 1, 9]
```

## Method chaining
With **creek**, method chaining is very straightforward and easy to read.  
This code below: 
- filters the array, so only the even numbers stay
- multiplies every number by 3
- sorts the array in descending order
- constrains the array to the length of 5
- collects the array from the stream
```go
arr := []int{2, 7, 3, 1, 12, 6, 82, 101, 23, 24, 72, 13, 7}

result := creek.FromArray(arr).Filter(func(item int) bool {
    return item%2 == 0
}).Map(func(item int) int {
    return item * 3
}).OrderByDescending().Limit(5).Collect()

// result: [246, 216, 72, 36, 18]
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