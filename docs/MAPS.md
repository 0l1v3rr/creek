# Regular Streams

- [Regular Streams](#regular-streams)
  - [Create stream](#create-stream)
  - [Functions](#functions)
    - [All](#all)
    - [Append](#append)
    - [AppendIf](#appendif)
    - [Collect](#collect)
    - [ElementAt](#elementat)
    - [ElementAtOrElse](#elementatorelse)
    - [First](#first)
    - [Last](#last)

<hr>

## Create stream
You can find a guide about creating streams [here](../README.md#create-stream).

<hr>

## Functions

### All
The `All` function determines whether all elements of the stream satisfy the passed condition.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}

result := creek.FromMap(arr).All(func(item creek.KeyValuePair[int, string]) bool {
    return item.Key > 0
})
// true
```

### Append
The `Append` function adds an element to the stream.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
toAppend := creek.KeyValuePair[int, string]{Key: 4, Value: "Michael"}

result := creek.FromMap(arr).Append(toAppend)
```

### AppendIf
The `AppendIf` function adds an element to the stream if the second parameter is true.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
toAppend := creek.KeyValuePair[int, string]{Key: 4, Value: "Michael"}

result := creek.FromMap(arr).AppendIf(toAppend, true)
```

### Collect
The `Collect` function returns the modified map from the streams.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).Collect() // map[int]string{1: "Mark", 2: "John", 3: "Jack"}
```

### ElementAt
The `ElementAt` function is used to get an element from the stream at a particular index.  
If the element is not present, it throws a panic.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).ElementAt(2) // 3: "Jack"
```

### ElementAtOrElse
The `ElementAtOrElse` function is used to get an element from the stream at a particular index.  
If the element is not present, it returns the elseValue, which is the second parameter.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
elseValue := creek.KeyValuePair[int, string]{Key: 4, Value: "Michael"}

result := creek.FromMap(arr).ElementAtOrElse(5, elseValue)
```

### First
The `First` method returns the first element in the stream.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).First() // 1: "Mark"
```

### Last
The `Last` method returns the last element in the stream.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).Last() // 3: "Jack"
```