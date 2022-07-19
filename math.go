package creek

import (
	"fmt"
	"reflect"
	"strconv"
)

// The Max function returns the largest element from the stream.
func (s Stream[T]) Max() T {
	max := s.Array[0]

	for i := 1; i < len(s.Array); i++ {
		if s.Array[i] > max {
			max = s.Array[i]
		}
	}

	return max
}

// The MaxIndex function returns the index of the largest element from the stream.
func (s Stream[T]) MaxIndex() int {
	max := s.Array[0]
	maxIndex := 0

	for i := 1; i < len(s.Array); i++ {
		if s.Array[i] > max {
			max = s.Array[i]
			maxIndex = i
		}
	}

	return maxIndex
}

// The Min function returns the smallest element from the stream.
func (s Stream[T]) Min() T {
	min := s.Array[0]

	for i := 1; i < len(s.Array); i++ {
		if s.Array[i] < min {
			min = s.Array[i]
		}
	}

	return min
}

// The MinIndex function returns the index of the smallest element from the stream.
func (s Stream[T]) MinIndex() int {
	min := s.Array[0]
	minIndex := 0

	for i := 1; i < len(s.Array); i++ {
		if s.Array[i] < min {
			min = s.Array[i]
			minIndex = i
		}
	}

	return minIndex
}

// The Sum function adds up all values in a stream.
func (s Stream[T]) Sum() T {
	var sum T

	for i := 0; i < len(s.Array); i++ {
		sum += s.Array[i]
	}

	return sum
}

// The Average function calculates the average of the stream.
// This function doesn't work with strings.
func (s Stream[T]) Average() float64 {
	if reflect.TypeOf(s.Array[0]).Kind() == reflect.String {
		panic("the MinIndex function doesn't work on string types")
	}

	var sum T
	len := len(s.Array)

	for i := 0; i < len; i++ {
		sum += s.Array[i]
	}

	floatSum, _ := strconv.ParseFloat(fmt.Sprintf("%v", sum), 64)
	return floatSum / float64(len)
}

// The Max function returns the largest element from the stream.
// The first parameter is the name of the field you want to calculate by.
func (s StructStream[T]) Max(fieldName string) T {
	// getting the type of the field
	t := reflect.TypeOf(s.Array[0])
	_, found := t.FieldByName(fieldName)

	// if field not found
	if !found {
		panic(fmt.Sprintf("field with the name '%s' does not exist", fieldName))
	}

	max := fmt.Sprintf("%v", fieldValue(s.Array[0], fieldName))
	maxIndex := 0

	for i := 1; i < len(s.Array); i++ {
		if fmt.Sprintf("%v", fieldValue(s.Array[i], fieldName)) > max {
			max = fmt.Sprintf("%v", fieldValue(s.Array[i], fieldName))
			maxIndex = i
		}
	}

	return s.Array[maxIndex]
}

// The MaxIndex function returns the index of the largest element from the stream.
// The first parameter is the name of the field you want to calculate by.
func (s StructStream[T]) MaxIndex(fieldName string) int {
	// getting the type of the field
	t := reflect.TypeOf(s.Array[0])
	_, found := t.FieldByName(fieldName)

	// if field not found
	if !found {
		panic(fmt.Sprintf("field with the name '%s' does not exist", fieldName))
	}

	max := fmt.Sprintf("%v", fieldValue(s.Array[0], fieldName))
	maxIndex := 0

	for i := 1; i < len(s.Array); i++ {
		if fmt.Sprintf("%v", fieldValue(s.Array[i], fieldName)) > max {
			max = fmt.Sprintf("%v", fieldValue(s.Array[i], fieldName))
			maxIndex = i
		}
	}

	return maxIndex
}

// The Min function returns the smallest element from the stream.
// The first parameter is the name of the field you want to calculate by.
func (s StructStream[T]) Min(fieldName string) T {
	// getting the type of the field
	t := reflect.TypeOf(s.Array[0])
	_, found := t.FieldByName(fieldName)

	// if field not found
	if !found {
		panic(fmt.Sprintf("field with the name '%s' does not exist", fieldName))
	}

	min := fmt.Sprintf("%v", fieldValue(s.Array[0], fieldName))
	minIndex := 0

	for i := 1; i < len(s.Array); i++ {
		if fmt.Sprintf("%v", fieldValue(s.Array[i], fieldName)) < min {
			min = fmt.Sprintf("%v", fieldValue(s.Array[i], fieldName))
			minIndex = i
		}
	}

	return s.Array[minIndex]
}

// The MinIndex function returns the index of the smallest element from the stream.
// The first parameter is the name of the field you want to calculate by.
func (s StructStream[T]) MinIndex(fieldName string) int {
	// getting the type of the field
	t := reflect.TypeOf(s.Array[0])
	_, found := t.FieldByName(fieldName)

	// if field not found
	if !found {
		panic(fmt.Sprintf("field with the name '%s' does not exist", fieldName))
	}

	min := fmt.Sprintf("%v", fieldValue(s.Array[0], fieldName))
	minIndex := 0

	for i := 1; i < len(s.Array); i++ {
		if fmt.Sprintf("%v", fieldValue(s.Array[i], fieldName)) < min {
			min = fmt.Sprintf("%v", fieldValue(s.Array[i], fieldName))
			minIndex = i
		}
	}

	return minIndex
}

// The Sum function adds up all values in a stream.
// The first parameter is the name of the field you want to calculate by.
func (s StructStream[T]) Sum(fieldName string) float64 {
	// getting the type of the field
	t := reflect.TypeOf(s.Array[0])
	_, found := t.FieldByName(fieldName)

	// if field not found
	if !found {
		panic(fmt.Sprintf("field with the name '%s' does not exist", fieldName))
	}

	// if the field cannot be converted to float
	_, err := strconv.ParseFloat(fmt.Sprintf("%v", fieldValue(s.Array[0], fieldName)), 64)
	if err != nil {
		panic("you can't sum this field")
	}

	var sum float64

	for i := 0; i < len(s.Array); i++ {
		converted, err := strconv.ParseFloat(fmt.Sprintf("%v", fieldValue(s.Array[i], fieldName)), 64)
		if err != nil {
			panic("you can't sum this field")
		}
		sum += converted
	}

	return sum
}

// The Average function calculates the average of the stream.
// This function doesn't work with strings.
// The first parameter is the name of the field you want to calculate by.
func (s StructStream[T]) Average(fieldName string) float64 {
	return s.Sum(fieldName) / float64(len(s.Array))
}

func fieldValue[T any](item T, fieldName string) interface{} {
	// getting the type of the field
	t := reflect.TypeOf(item)
	field, found := t.FieldByName(fieldName)

	// if field not found
	if !found {
		panic(fmt.Sprintf("field with the name '%s' does not exist", fieldName))
	}

	// if the field is an int
	if field.Type.Kind() == reflect.Int || field.Type.Kind() == reflect.Int16 ||
		field.Type.Kind() == reflect.Int32 || field.Type.Kind() == reflect.Int64 ||
		field.Type.Kind() == reflect.Int8 {

		return reflect.ValueOf(item).FieldByName(fieldName).Int()
	}

	// if the field is a uint
	if field.Type.Kind() == reflect.Uint || field.Type.Kind() == reflect.Uint16 ||
		field.Type.Kind() == reflect.Uint32 || field.Type.Kind() == reflect.Uint64 ||
		field.Type.Kind() == reflect.Uint8 {

		return reflect.ValueOf(item).FieldByName(fieldName).Uint()
	}

	// if the field is not a string
	if field.Type.Kind() != reflect.String {
		panic(fmt.Sprintf("wrong field type: '%v'", field.Type))
	}

	return reflect.ValueOf(item).FieldByName(fieldName).String()
}

// The Max function returns the largest element from the stream.
func (s MapStream[T, V]) Max(byKey bool) KeyValuePair[T, V] {
	max := s.Array[0]

	for i := 1; i < len(s.Array); i++ {
		if byKey {
			if s.Array[i].Key > max.Key {
				max = s.Array[i]
			}

			continue
		}

		if s.Array[i].Value > max.Value {
			max = s.Array[i]
		}
	}

	return max
}

// The MaxIndex function returns the index of the largest element from the stream.
func (s MapStream[T, V]) MaxIndex(byKey bool) int {
	max := s.Array[0]
	maxIndex := 0

	for i := 1; i < len(s.Array); i++ {
		if byKey {
			if s.Array[i].Key > max.Key {
				max = s.Array[i]
				maxIndex = i
			}

			continue
		}

		if s.Array[i].Value > max.Value {
			max = s.Array[i]
			maxIndex = i
		}
	}

	return maxIndex
}

// The Min function returns the smallest element from the stream.
func (s MapStream[T, V]) Min(byKey bool) KeyValuePair[T, V] {
	min := s.Array[0]

	for i := 1; i < len(s.Array); i++ {
		if byKey {
			if s.Array[i].Key < min.Key {
				min = s.Array[i]
			}

			continue
		}

		if s.Array[i].Value < min.Value {
			min = s.Array[i]
		}
	}

	return min
}

// The MinIndex function returns the index of the smallest element from the stream.
func (s MapStream[T, V]) MinIndex(byKey bool) int {
	min := s.Array[0]
	minIndex := 0

	for i := 1; i < len(s.Array); i++ {
		if byKey {
			if s.Array[i].Key < min.Key {
				min = s.Array[i]
				minIndex = i
			}

			continue
		}

		if s.Array[i].Value < min.Value {
			min = s.Array[i]
			minIndex = i
		}
	}

	return minIndex
}

// The Sum function adds up all values in a stream.
func (s MapStream[T, V]) Sum(byKey bool) interface{} {
	if byKey {
		var sum T

		for i := 0; i < len(s.Array); i++ {
			sum += s.Array[i].Key
		}

		return sum
	}

	var sum V

	for i := 0; i < len(s.Array); i++ {
		sum += s.Array[i].Value
	}

	return sum
}

// The Average function calculates the average of the stream.
// This function doesn't work with strings.
func (s MapStream[T, V]) Average(byKey bool) float64 {
	sum := s.Sum(byKey)

	if reflect.TypeOf(sum).Kind() == reflect.String {
		panic("the MinIndex function doesn't work on string types")
	}

	floatSum, _ := strconv.ParseFloat(fmt.Sprintf("%v", sum), 64)
	return floatSum / float64(len(s.Array))
}
