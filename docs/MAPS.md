# Regular Streams

- [Regular Streams](#regular-streams)
  - [Create stream](#create-stream)
  - [Functions](#functions)
    - [All](#all)
    - [Append](#append)
    - [AppendIf](#appendif)
    - [Clear](#clear)
    - [Collect](#collect)
    - [Count](#count)
    - [ContainsKey](#containskey)
    - [ElementAt](#elementat)
    - [ElementAtOrElse](#elementatorelse)
    - [Filter](#filter)
    - [Find](#find)
    - [FindIndex](#findindex)
    - [FindLast](#findlast)
    - [FindLastIndex](#findlastindex)
    - [First](#first)
    - [ForEach](#foreach)
    - [IsEmpty](#isempty)
    - [IsNotEmpty](#isnotempty)
    - [Last](#last)
    - [Map](#map)
    - [OrderBy](#orderby)
    - [OrderByDescending](#orderbydescending)
    - [Shuffle](#shuffle)
    - [Some](#some)
    - [Wait](#wait)

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

### Clear
The `Clear` function clears every element from the stream.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).Clear() // []
```

### Collect
The `Collect` function returns the modified map from the streams.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).Collect() // map[int]string{1: "Mark", 2: "John", 3: "Jack"}
```

### Count
The `Count` function returns the count of elements in the stream.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).Count() // 3
```

### ContainsKey
The `ContainsKey` function checks whether the stream contains an item with the passed key
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).ContainsKey(2) // true
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

### Filter
The `Filter` function leaves only those elements in the array that make the specified condition true.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).Filter(func(kvp creek.KeyValuePair[int, string]) bool {
    return kvp.Key > 1
})
```

### Find
The `Find` function searches for an element that matches the conditions passed and returns the first occurrence within the entire stream.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).Find(func(kvp creek.KeyValuePair[int, string]) bool {
    return kvp.Key%2 == 0
})
```

### FindIndex
The `FindIndex` function searches for an element that matches the conditions passed and returns the index of the first occurrence within the entire stream.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).FindIndex(func(kvp creek.KeyValuePair[int, string]) bool {
    return kvp.Key%2 == 0
})
```

### FindLast
The `FindLast` function searches for an element that matches the conditions passed and returns the last occurrence within the entire stream.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).FindLast(func(kvp creek.KeyValuePair[int, string]) bool {
    return kvp.Key%2 == 0
})
```

### FindLastIndex
The `FindLastIndex` function searches for an element that matches the conditions passed and returns the index of the last occurrence within the entire stream.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).FindLastIndex(func(kvp creek.KeyValuePair[int, string]) bool {
    return kvp.Key%2 == 0
})
```

### First
The `First` method returns the first element in the stream.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).First() // 1: "Mark"
```

### ForEach
The `ForEach` method runs the specified method on every element in the Stream.  
Warning: this method doesn't return anything
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
creek.FromMap(arr).ForEach(func(kvp creek.KeyValuePair[int, string]) {
    fmt.Println(kvp.Key, kvp.Value)
})

// -- Output: --
// 1 Mark
// 2 John
// 3 Jack
```

### IsEmpty
The `IsEmpty` function checks whether the stream is empty.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result = creek.FromMap(arr).IsEmpty() // false
```

### IsNotEmpty
The `IsNotEmpty` function checks whether the stream is not empty.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result = creek.FromMap(arr).IsEmpty() // true
```

### Last
The `Last` method returns the last element in the stream.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).Last() // 3: "Jack"
```

### Map
The `Map` function creates a new stream populated with the results of calling the provided function on every element.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).Map(func(item creek.KeyValuePair[int, string]) creek.KeyValuePair[int, string] {
    return creek.KeyValuePair[int, string]{
        Key:   item.Key * 2,
        Value: strings.ToUpper(item.Value),
    }
}) // [{2 MARK} {4 JOHN} {6 JACK}]
```

### OrderBy
The `OrderBy` function sorts the stream in ascending order.  
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).OrderBy(creek.ByKey) // [{1 Mark} {2 John} {3 Jack}]
```

### OrderByDescending
The `OrderByDescending` function sorts the stream in descending order.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).OrderByDescending(creek.ByValue) // [{1 Mark} {2 John} {3 Jack}]
```

### Shuffle
The `Shuffle` function shuffles the stream.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).Shuffle()
```

### Some
The `Some` function determines whether any of the elements of the stream satisfy the passed condition.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}

result := creek.FromMap(arr).Some(func(item creek.KeyValuePair[int, string]) bool {
    return item.Key > 0
})
// true
```

### Wait
The `Wait` function pauses the current stream for the duration passed.
The first and only parameter expects a value from the built-in `time.Duration` package.
```go
arr := map[int]string{1: "Mark", 2: "John", 3: "Jack"}
result := creek.FromMap(arr).Wait(time.Second * 5) // waits for 5 seconds
```