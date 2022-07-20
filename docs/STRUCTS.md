# Regular Streams

- [Regular Streams](#regular-streams)
  - [Create stream](#create-stream)
  - [Your struct](#your-struct)
  - [Functions](#functions)
    - [All](#all)
    - [Append](#append)
    - [AppendAt](#appendat)
    - [AppendIf](#appendif)
    - [ArrEquals](#arrequals)
    - [Average](#average)
    - [Bind](#bind)
    - [Clear](#clear)
    - [Collect](#collect)
    - [Contains](#contains)
    - [Count](#count)
    - [ElementAt](#elementat)
    - [ElementAtOrElse](#elementatorelse)
    - [Equals](#equals)
    - [Filter](#filter)
    - [Find](#find)
    - [FindIndex](#findindex)
    - [FindLast](#findlast)
    - [FindLastIndex](#findlastindex)
    - [First](#first)
    - [ForEach](#foreach)
    - [IndexOf](#indexof)
    - [IsEmpty](#isempty)
    - [IsNotEmpty](#isnotempty)
    - [Last](#last)
    - [LastIndexOf](#lastindexof)
    - [Limit](#limit)
    - [Map](#map)
    - [Max](#max)
    - [MaxIndex](#maxindex)
    - [Min](#min)
    - [MinIndex](#minindex)
    - [OrderBy](#orderby)
    - [OrderByDescending](#orderbydescending)
    - [Push](#push)
    - [PushValues](#pushvalues)
    - [Remove](#remove)
    - [RemoveAt](#removeat)
    - [RemoveIf](#removeif)
    - [RemoveWhere](#removewhere)
    - [Replace](#replace)
    - [ReplaceWhere](#replacewhere)
    - [Reverse](#reverse)
    - [Shuffle](#shuffle)
    - [Skip](#skip)
    - [Slice](#slice)
    - [Some](#some)
    - [Sum](#sum)
    - [Wait](#wait)

<hr>

## Create stream
You can find a guide about creating streams [here](../README.md#create-stream).

<hr>

## Your struct
```go
type YourStruct struct {
	Id   int64
	Name string
}
```

<hr>

## Functions

### All
The `All` function determines whether all elements of the stream satisfy the passed condition.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).All(func(item YourStruct) bool {
    return item.Id%2 == 0
})
```

### Append
The `Append` function adds an element to the stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Append(YourStruct{
    Id:   4,
    Name: "Oliver",
})
```

### AppendAt
The `AppendAt` function inserts the specified element at the specified position in the stream.  
If the index is out of range, nothing happens.  
The first parameter is the index, and the second is the item.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).AppendAt(1, YourStruct{
    Id:   4,
    Name: "Oliver",
})
```

### AppendIf
The `AppendIf` function adds an element to the stream if the second parameter is true.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
{Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).AppendIf(YourStruct{
    Id:   4,
    Name: "Oliver",
}, len(structArray) == 3)
```

### ArrEquals
The `ArrEquals` function compares the stream and the passed array and returns true if they're equals.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).ArrEquals([]YourStruct{
    {Id: 4, Name: "Oliver"},
}) // false
```

### Average
The `Average` function calculates the average of the stream.  
This function doesn't work with strings.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Average("Id")
```

### Bind
The `Bind` function binds the stream into the passed variable.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

var shouldBindHere []YourStruct

creek.FromStructs(structArray).Bind(&shouldBindHere) // shouldBindHere => [{1 John} {2 Will} {3 Mark}]
```

### Clear
The `Clear` function clears every element from the stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Clear() // []
```

### Collect
The `Collect` function returns the modified array from the streams.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

arrFromStream := creek.FromStructs(structArray).Collect()
```

### Contains
The `Contains` function checks whether the stream contains the passed item.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Contains(YourStruct{Id: 3, Name: "Mark"}) // true
```

### Count
The `Count` function returns the count of elements in the stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Count() // 3
```

### ElementAt
The `ElementAt` function is used to get an element from the stream at a particular index.  
If the element is not present, it throws a panic.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).ElementAt(2) // {Id: 3, Name: "Mark"}
```

### ElementAtOrElse
The `ElementAtOrElse` function is used to get an element from the stream at a particular index.  
If the element is not present, it returns the elseValue, which is the second parameter.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).ElementAtOrElse(6, YourStruct{Id: 4, Name: "Oliver"})
```

### Equals
The `Equals` function compares two streams and returns true if they're equals.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Equals(creek.StructStream[YourStruct]{
    Array: []YourStruct{
        {Id: 1, Name: "John"},
        {Id: 2, Name: "Will"},
        {Id: 3, Name: "Mark"},
    },
})
// true
```

### Filter
The `Filter` function leaves only those elements in the array that make the specified condition true.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Filter(func(item YourStruct) bool {
    return item.Id > 2
})
```

### Find
The `Find` function searches for an element that matches the conditions passed and returns the first occurrence within the entire stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Find(func(item YourStruct) bool {
    return item.Id > 1
})
```

### FindIndex
The `FindIndex` function searches for an element that matches the conditions passed and returns the index of the first occurrence within the entire stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).FindIndex(func(item YourStruct) bool {
    return item.Id > 1
})
```

### FindLast
The `FindLast` function searches for an element that matches the conditions passed and returns the last occurrence within the entire stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).FindLast(func(item YourStruct) bool {
    return item.Id > 1
})
```

### FindLastIndex
The `FindLastIndex` function searches for an element that matches the conditions passed and returns the index of the last occurrence within the entire stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).FindLastIndex(func(item YourStruct) bool {
    return item.Id > 1
})
```

### First
The `First` method returns the first element in the stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).First() // {Id: 1, Name: "John"}
```

### ForEach
The `ForEach` method runs the specified method on every element in the Stream.  
Warning: this method doesn't return anything
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

creek.FromStructs(structArray).ForEach(func(item YourStruct) {
    fmt.Println(item)
})
```

### IndexOf
The `IndexOf` function returns the position of the first occurrence of the passed value in a stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).IndexOf(YourStruct{Id: 1, Name: "John"}) // 0
```

### IsEmpty
The `IsEmpty` function checks whether the stream is empty.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result = creek.FromStructs(structArray).IsEmpty() // false
```

### IsNotEmpty
The `IsNotEmpty` function checks whether the stream is not empty.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result = creek.FromStructs(structArray).IsEmpty() // true
```

### Last
The `Last` method returns the last element in the stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Last() // {Id: 3, Name: "Mark"}
```

### LastIndexOf
The `LastIndexOf` function returns the position of the last occurrence of the passed value in a stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).LastIndexOf(YourStruct{Id: 1, Name: "John"}) // 0
```

### Limit
The `Limit` function constrains the number of elements returned by the stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Limit(2)
```

### Map
The `Map` function creates a new stream populated with the results of calling the provided function on every element.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Map(func(item YourStruct) YourStruct {
    return YourStruct{
        Id:   item.Id * 2,
        Name: item.Name,
    }
})
```

### Max
The `Max` function returns the largest element from the stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Max("Id")
```

### MaxIndex
The `MaxIndex` function returns the index of the largest element from the stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).MaxIndex("Id")
```

### Min
The `Min` function returns the smallest element from the stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Min("Id")
```

### MinIndex
The `MinIndex` function returns the index of the smallest element from the stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).MinIndex("Id")
```

### OrderBy
The `OrderBy` function sorts the stream in ascending order.  
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).OrderBy("Id")
```

### OrderByDescending
The `OrderByDescending` function sorts the stream in descending order.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).OrderByDescending("Id")
```

### Push
The `Push` function adds the passed array to the end of the stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

toPush := []YourStruct{
    {Id: 4, Name: "John2"},
    {Id: 5, Name: "Will2"},
    {Id: 6, Name: "Mark2"},
}

result := creek.FromStructs(structArray).Push(toPush)
```

### PushValues
The `PushValues` function adds the passed values to the end of the stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).PushValues(YourStruct{Id: 4, Name: "John2"}, YourStruct{Id: 5, Name: "Will2"})
```

### Remove
The `Remove` function removes the passed element from a stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result = creek.FromStructs(structArray).Remove(YourStruct{Id: 1, Name: "John"})
```

### RemoveAt
The `RemoveAt` function removes the item if its index matches the index passed in.  
If the index is out of range, nothing happens.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).RemoveAt(2)
```

### RemoveIf
The `RemoveIf` function removes the passed element from a stream if the second parameter is true.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).RemoveIf(YourStruct{Id: 3, Name: "Mark"}, len(structArray) == 4)
```

### RemoveWhere
The `RemoveWhere` function removes all the entries that satisfy the provided condition.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).RemoveWhere(func(item YourStruct) bool {
    return item.Id > 1
})
```

### Replace
The `Replace` function replaces every occurrence of `from` to `to`.  
The first parameter is `from`, and the second is `to`.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Replace(YourStruct{Id: 3, Name: "Mark"}, YourStruct{Id: 12, Name: "Mark12"})
```

### ReplaceWhere
The `ReplaceWhere` function replaces every element that satisfies the condition.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).ReplaceWhere(func(item YourStruct) bool {
    return item.Id%2 == 0
}, YourStruct{Id: 4, Name: "Mark"})
```

### Reverse
The `Reverse` function reverses the stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Reverse()
```

### Shuffle
The `Shuffle` function shuffles the stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Shuffle()
```

### Skip
The `Skip` function discards the first n elements of a stream, where n is the passed parameter.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Skip(2)
```

### Slice
The `Slice` function returns a copy of a portion of a stream into a new stream selected from start to end (end not included) where start and end represent the index of items in the stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Slice(1, 3)
```

### Some
The `Some` function determines whether any of the elements of the stream satisfy the passed condition.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Some(func(item YourStruct) bool {
    return item.Id%2 == 0
})
// true
```

### Sum
The `Sum` function adds up all values in a stream.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Sum("Id")
```

### Wait
The `Wait` function pauses the current stream for the duration passed.
The first and only parameter expects a value from the built-in `time.Duration` package.
```go
structArray := []YourStruct{
    {Id: 1, Name: "John"},
    {Id: 2, Name: "Will"},
    {Id: 3, Name: "Mark"},
}

result := creek.FromStructs(structArray).Wait(time.Second * 5) // waits for 5 seconds
```