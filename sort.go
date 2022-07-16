package creek

import (
	"fmt"
	"reflect"
	"sort"
)

// The OrderBy function sorts the stream in ascending order.
func (s Stream[T]) OrderBy() Stream[T] {
	sort.SliceStable(s.Array, func(i, j int) bool {
		return s.Array[i] < s.Array[j]
	})

	return Stream[T]{
		Array: s.Array,
	}
}

// The OrderByDescending function sorts the stream in descending order.
func (s Stream[T]) OrderByDescending() Stream[T] {
	sort.SliceStable(s.Array, func(i, j int) bool {
		return s.Array[i] > s.Array[j]
	})

	return Stream[T]{
		Array: s.Array,
	}
}

// The OrderBy function sorts the stream in ascending order.
func (s MapStream[T, V]) OrderBy(byKey bool) MapStream[T, V] {
	sort.SliceStable(s.Array, func(i, j int) bool {
		if byKey {
			return s.Array[i].Key < s.Array[j].Key
		}

		return s.Array[i].Value < s.Array[j].Value
	})

	return MapStream[T, V]{
		Array: s.Array,
	}
}

// The OrderByDescending function sorts the stream in descending order.
func (s MapStream[T, V]) OrderByDescending(byKey bool) MapStream[T, V] {
	sort.SliceStable(s.Array, func(i, j int) bool {
		if byKey {
			return s.Array[i].Key > s.Array[j].Key
		}

		return s.Array[i].Value > s.Array[j].Value
	})

	return MapStream[T, V]{
		Array: s.Array,
	}
}

// The OrderBy function sorts the stream in ascending order.
// The first parameter is the name of the field you want to sort by.
func (s StructStream[T]) OrderBy(fieldName string) StructStream[T] {
	// getting the type of the field
	t := reflect.TypeOf(s.Array[0])
	field, found := t.FieldByName(fieldName)

	// if field not found
	if !found {
		panic(fmt.Sprintf("field with the name '%s' does not exist", fieldName))
	}

	// if the field is an int
	if field.Type.Kind() == reflect.Int || field.Type.Kind() == reflect.Int16 ||
		field.Type.Kind() == reflect.Int32 || field.Type.Kind() == reflect.Int64 ||
		field.Type.Kind() == reflect.Int8 {

		sort.SliceStable(s.Array, func(i, j int) bool {
			valOne := reflect.ValueOf(s.Array[i]).FieldByName(fieldName).Int()
			valTwo := reflect.ValueOf(s.Array[j]).FieldByName(fieldName).Int()

			return valOne < valTwo
		})

		return StructStream[T]{
			Array: s.Array,
		}
	}

	// if the field is a uint
	if field.Type.Kind() == reflect.Uint || field.Type.Kind() == reflect.Uint16 ||
		field.Type.Kind() == reflect.Uint32 || field.Type.Kind() == reflect.Uint64 ||
		field.Type.Kind() == reflect.Uint8 {

		sort.SliceStable(s.Array, func(i, j int) bool {
			valOne := reflect.ValueOf(s.Array[i]).FieldByName(fieldName).Uint()
			valTwo := reflect.ValueOf(s.Array[j]).FieldByName(fieldName).Uint()

			return valOne < valTwo
		})

		return StructStream[T]{
			Array: s.Array,
		}
	}

	// if the field is a string
	if field.Type.Kind() == reflect.String {
		sort.SliceStable(s.Array, func(i, j int) bool {
			valOne := reflect.ValueOf(s.Array[i]).FieldByName(fieldName).String()
			valTwo := fmt.Sprintf("%v", reflect.ValueOf(s.Array[j]).FieldByName(fieldName))

			if !found {
				valOne = reflect.ValueOf(s.Array[i]).FieldByIndex([]int{0}).String()
				valTwo = reflect.ValueOf(s.Array[j]).FieldByIndex([]int{0}).String()
			}

			return valOne < valTwo
		})
	} else {
		panic(fmt.Sprintf("field with the type of '%v' cannot be sorten", field.Type))
	}

	return StructStream[T]{
		Array: s.Array,
	}
}

// The OrderByDescending function sorts the stream in descending order
// The first parameter is the name of the field you want to sort by.
func (s StructStream[T]) OrderByDescending(fieldName string) StructStream[T] {
	// getting the type of the field
	t := reflect.TypeOf(s.Array[0])
	field, found := t.FieldByName(fieldName)

	// if field not found
	if !found {
		panic(fmt.Sprintf("field with the name '%s' does not exist", fieldName))
	}

	// if the field is an int
	if field.Type.Kind() == reflect.Int || field.Type.Kind() == reflect.Int16 ||
		field.Type.Kind() == reflect.Int32 || field.Type.Kind() == reflect.Int64 ||
		field.Type.Kind() == reflect.Int8 {

		sort.SliceStable(s.Array, func(i, j int) bool {
			valOne := reflect.ValueOf(s.Array[i]).FieldByName(fieldName).Int()
			valTwo := reflect.ValueOf(s.Array[j]).FieldByName(fieldName).Int()

			return valOne > valTwo
		})

		return StructStream[T]{
			Array: s.Array,
		}
	}

	// if the field is a uint
	if field.Type.Kind() == reflect.Uint || field.Type.Kind() == reflect.Uint16 ||
		field.Type.Kind() == reflect.Uint32 || field.Type.Kind() == reflect.Uint64 ||
		field.Type.Kind() == reflect.Uint8 {

		sort.SliceStable(s.Array, func(i, j int) bool {
			valOne := reflect.ValueOf(s.Array[i]).FieldByName(fieldName).Uint()
			valTwo := reflect.ValueOf(s.Array[j]).FieldByName(fieldName).Uint()

			return valOne > valTwo
		})

		return StructStream[T]{
			Array: s.Array,
		}
	}

	// if the field is a string
	if field.Type.Kind() == reflect.String {
		sort.SliceStable(s.Array, func(i, j int) bool {
			valOne := reflect.ValueOf(s.Array[i]).FieldByName(fieldName).String()
			valTwo := fmt.Sprintf("%v", reflect.ValueOf(s.Array[j]).FieldByName(fieldName))

			if !found {
				valOne = reflect.ValueOf(s.Array[i]).FieldByIndex([]int{0}).String()
				valTwo = reflect.ValueOf(s.Array[j]).FieldByIndex([]int{0}).String()
			}

			return valOne > valTwo
		})
	} else {
		panic(fmt.Sprintf("field with the type of '%v' cannot be sorten", field.Type))
	}

	return StructStream[T]{
		Array: s.Array,
	}
}
