package main

import (
	"fmt"
	"strings"

	"github.com/0l1v3rr/creek"
)

func main() {
	arr := []string{"One", "Two", "Three"}

	result := creek.FromArray(arr).Map(func(item string) string {
		return strings.ToUpper(item)
	}).Collect()

	fmt.Println(result)
}
