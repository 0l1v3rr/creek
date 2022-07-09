<div align="center">
    <img src="./images/logo.png" width="320px" alt="Creek Logo">
</div>
<div align="center">
    <a href="https://github.com/0l1v3rr/creek">Creek</a> is a tiny, fully-featured Streams library 
    for <a href="https://go.dev/">Go</a>. <br>
    Creek creates wrappers around a specific data source (array or slice), allowing us to operate with that data source and making bulk processing convenient and fast.
    Creek helps you to follow the functional programming paradigms.    
</div>
<br>
<div align="center">
    <a href="https://github.com/0l1v3rr/creek/releases">
        <img src="https://img.shields.io/github/release/0l1v3rr/creek.svg?style=flat-square" alt="Relase">
    </a>
    <a href="https://github.com/0l1v3rr/creek/actions">
        <img src="https://img.shields.io/github/workflow/status/0l1v3rr/creek/Test?style=flat-square" alt="Test Status">
    </a>
    <a href="https://github.com/0l1v3rr/creek/blob/master/LICENSE">
        <img src="https://img.shields.io/github/license/0l1v3rr/creek?style=flat-square" alt="License">
    </a>
    <a href="https://github.com/0l1v3rr/creek/blob/master/go.mod">
        <img src="https://img.shields.io/github/go-mod/go-version/0l1v3rr/creek?style=flat-square" alt="Go Version">
    </a>
    <a href="https://goreportcard.com/report/github.com/0l1v3rr/creek">
        <img src="https://goreportcard.com/badge/github.com/0l1v3rr/creek?style=flat-square" alt="Go Report">
    </a>
</div>

<hr>

# Table of Contents
- [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Quick start](#quick-start)
  - [Create stream](#create-stream)
    - [Empty stream](#empty-stream)
    - [Stream from regular arrays and slices](#stream-from-regular-arrays-and-slices)
    - [Stream from parameter values](#stream-from-parameter-values)
  - [Functions](#functions)
    - [Collect](#collect)
    - [OrderBy](#orderby)
    - [OrderByDescending](#orderbydescending)
    - [Filter](#filter)
    - [ForEach](#foreach)
    - [Map](#map)
    - [Limit](#limit)
    - [Count](#count)
    - [Append](#append)
    - [AppendAt](#appendat)
    - [AppendIf](#appendif)
    - [Remove](#remove)
    - [RemoveAt](#removeat)
    - [RemoveIf](#removeif)
    - [RemoveWhere](#removewhere)
    - [RemoveDuplicates](#removeduplicates)
    - [IndexOf](#indexof)
    - [LastIndexOf](#lastindexof)
    - [ElementAt](#elementat)
    - [ElementAtOrElse](#elementatorelse)
    - [Find](#find)
    - [FindIndex](#findindex)
    - [FindLast](#findlast)
    - [FindLastIndex](#findlastindex)
    - [Max](#max)
    - [MaxIndex](#maxindex)
    - [Min](#min)
    - [MinIndex](#minindex)
    - [Sum](#sum)
    - [Average](#average)
    - [All](#all)
    - [Some](#some)
    - [Equals](#equals)
    - [ArrEquals](#arrequals)
    - [Reverse](#reverse)
    - [Join](#join)
    - [Contains](#contains)
    - [Distinct](#distinct)
    - [IsEmpty](#isempty)
    - [Clear](#clear)
    - [Wait](#wait)
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

## Create stream
You can create a stream from almost every type of array or slice.

### Empty stream
The `Empty` function returns an empty stream.
```go
emptyStream := creek.Empty[int]()
```

### Stream from regular arrays and slices
The supported types are the following:
```go
string | byte | float32 | float64 | int | int16 | int32 | int64 | uint16 | uint32 | uint64
```
In order to create a stream, use the `FromArray` function:
```go
// slice
arr := []int{1, 8, 2, 14, 22, 4, 7, 92}
intStream := creek.FromArray(arr)

// array
arr2 := [3]string{"One", "Two", "Three"}
stringStream := creek.FromArray(arr2[:])
```

### Stream from parameter values
```go
stream := creek.FromValues("Apple", "Strawberry", "Peach")
```

<hr>

## Functions

### Collect
The `Collect` function returns the modified array from the streams.
```go
arr := []int{1, 8, 2, 14, 22, 4, 7, 92}
arrFromStream := creek.FromArray(arr).Collect() // [1, 8, 2, 14, 22, 4, 7, 92]
```

### OrderBy
The `OrderBy` function sorts the stream in ascending order.  
```go
arr := []int{1, 8, 2, 14, 22, 4, 7, 92}
result := creek.FromArray(arr).OrderBy() // [1, 2, 4, 7, 8, 14, 22, 92]
```

### OrderByDescending
The `OrderByDescending` function sorts the stream in descending order.
```go
arr := []int{1, 8, 2, 14, 22, 4, 7, 92}
result := creek.FromArray(arr).OrderByDescending() // [92, 22, 14, 8, 7, 4, 2, 1]
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

### Count
The `Count` function returns the count of elements in the stream.
```go
arr := []int{3, 4, 1, 4, 2, 9}
result := creek.FromArray(arr).Count() // 6
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

### AppendIf
The `AppendIf` function adds an element to the stream if the second parameter is true.
```go
arr := []int{2, 7, 3, 1}
result := creek.FromArray(arr).AppendIf(8, len(arr) == 4) // [2, 7, 3, 1, 8]
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

### RemoveIf
The `RemoveIf` function removes the passed element from a stream if the second parameter is true.
```go
arr := []int{2, 7, 3, 1}
result := creek.FromArray(arr).RemoveIf(7, len(arr) == 4) // [2, 3, 1]
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
The `RemoveDuplicates` function removes every duplicate item from the stream.
```go
arr := []int{2, 7, 3, 1, 3, 9, 3}
result := creek.FromArray(arr).RemoveDuplicates() // [2, 7, 3, 1, 9]
```

### IndexOf
The `IndexOf` function returns the position of the first occurrence of the passed value in a stream.
```go
arr := []int{3, 4, 1, 4, 2, 9}
result := creek.FromArray(arr).IndexOf(4) // 1
```

### LastIndexOf
The `LastIndexOf` function returns the position of the last occurrence of the passed value in a stream.
```go
arr := []int{3, 4, 1, 4, 2, 9}
result := creek.FromArray(arr).LastIndexOf(4) // 3
```

### ElementAt
The `ElementAt` function is used to get an element from the stream at a particular index.  
If the element is not present, it throws a panic.
```go
arr := []int{3, 4, 1, 4, 2, 9}
result := creek.FromArray(arr).ElementAt(2) // 1
```

### ElementAtOrElse
The `ElementAtOrElse` function is used to get an element from the stream at a particular index.  
If the element is not present, it returns the elseValue, which is the second parameter.
```go
arr := []int{3, 4, 1, 4, 2, 9}

result := creek.FromArray(arr).ElementAtOrElse(5, 100) // 9

result2 = creek.FromArray(arr).ElementAtOrElse(6, 100) // 100
```

### Find
The `Find` function searches for an element that matches the conditions passed and returns the first occurrence within the entire stream.
```go
res := []int{2, 7, 3, 1, 4}
result := creek.FromArray(arr).Find(func(item int) bool {
    return item%2 == 0
})
// 2
```

### FindIndex
The `FindIndex` function searches for an element that matches the conditions passed and returns the index of the first occurrence within the entire stream.
```go
res := []int{2, 7, 3, 1, 4}
result := creek.FromArray(arr).FindIndex(func(item int) bool {
    return item%2 == 0
})
// 0
```

### FindLast
The `FindLast` function searches for an element that matches the conditions passed and returns the last occurrence within the entire stream.
```go
res := []int{2, 7, 3, 1, 4}
result := creek.FromArray(arr).FindLast(func(item int) bool {
    return item%2 == 0
})
// 4
```

### FindLastIndex
The `FindLastIndex` function searches for an element that matches the conditions passed and returns the index of the last occurrence within the entire stream.
```go
res := []int{2, 7, 3, 1, 4}
result := creek.FromArray(arr).FindLastIndex(func(item int) bool {
    return item%2 == 0
})
// 4
```

### Max
The `Max` function returns the largest element from the stream.
```go
arr := []int{2, 7, 3, 1, 9, 12, 5}
result := creek.FromArray(arr).Max() // 12
```

### MaxIndex
The `MaxIndex` function returns the index of the largest element from the stream.
```go
arr := []int{2, 7, 3, 1, 9, 12, 5}
result := creek.FromArray(arr).MaxIndex() // 5
```

### Min
The `Min` function returns the smallest element from the stream.
```go
arr := []int{2, 7, 3, 1, 9, 12, 5}
result := creek.FromArray(arr).Min() // 1
```

### MinIndex
The `MaxIndex` function returns the index of the smallest element from the stream.
```go
arr := []int{2, 7, 3, 1, 9, 12, 5}
result := creek.FromArray(arr).MinIndex() // 3
```

### Sum
The `Sum` function adds up all values in a stream.
```go
arr := []int{2, 7, 3, 1, 9, 12, 5}
result := creek.FromArray(arr).Sum() // 39
```

### Average
The `Average` function calculates the average of the stream.  
This function doesn't work with strings.
```go
arr := []int{2, 7, 3, 1, 9, 12, 5}
result := creek.FromArray(arr).Average() // 5.571428571428571
```

### All
The `All` function determines whether all elements of the stream satisfy the passed condition.
```go
arr := []int{2, 7, 3, 1}
result := creek.FromArray(arr).All(func(item int) bool {
    return item%2 == 0
}) 
// false
```

### Some
The `Some` function determines whether any of the elements of the stream satisfy the passed condition.
```go
arr := []int{2, 7, 3, 1}
result := creek.FromArray(arr).Some(func(item int) bool {
    return item%2 == 0
})
// true
```

### Equals
The `Equals` function compares two streams and returns true if they're equals.
```go
arr := []int{2, 7, 3, 1}
result := creek.FromArray(arr).Equals(creek.Stream[int]{
    Array: []int{2, 7, 3, 1},
})
// true
```

### ArrEquals
The `ArrEquals` function compares the stream and the passed array and returns true if they're equals.
```go
arr := []int{2, 7, 3, 1}
result = creek.FromArray(arr).ArrEquals([]int{2, 7, 3, 2})
// false
```

### Reverse
The `Reverse` function reverses the stream.
```go
arr := []int{2, 7, 3, 1}
result := creek.FromArray(arr).Reverse() // [1, 3, 7, 2]
```

### Join
The `Join` function concatenates the elements of the stream to create a single string.  
The passed parameter is placed between the elements.
```go
arr := []int{2, 7, 3, 1}
result := creek.FromArray(arr).Join(", ") // result: "2, 7, 3, 1"
```

### Contains
The `Contains` function checks whether the stream contains the passed item.
```go
arr := []int{2, 7, 3, 1}
result := creek.FromArray(arr).Contains(2) // true
```

### Distinct
The `Distinct` function filters every distinct element from the stream.
```go
arr := []int{2, 7, 3, 1, 3, 9, 3}
result := creek.FromArray(arr).Distinct() // [2, 7, 3, 1, 9]
```

### IsEmpty
The `IsEmpty` function checks whether the stream is empty.
```go
arr := []int{2, 7, 3, 1}
result = creek.FromArray(arr).IsEmpty() // false
```

### Clear
The `Clear` function clears every element from the stream.
```go
arr := []int{3, 4, 1, 4, 2, 9}
result := creek.FromArray(arr).Clear() // []
```

### Wait
The `Wait` function pauses the current stream for the duration passed.
The first and only parameter expects a value from the built-in `time.Duration` package.
```go
arr := []int{3, 4, 2, 9}
result := creek.FromArray(arr).Wait(time.Second * 5) // waits for 5 seconds
```

<hr>

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
- Running the test in detailed version: `make test_detailed`
- Test coverage: `make test_coverage`
- Go formatting: `make fmt`

<hr>

## Contributing
You can find [here](CONTRIBUTING.md) a contributing guideline.  
A star would be appreciated as well. 😉

<hr>

## License
This project is licensed under the [MIT License](LICENSE).