package main

import (
	"fmt"

	"github.com/hardyantz/sailfish"
)

type Category struct {
	Name        string   `gql:"name"`
	Slug        string   `gql:"slug"`
	Nested      Nested   `gql:"nested"`
	NestedSlice []Nested `gql:"nestedSlice"`
}

type Nested struct {
	Child1      string   `gql:"child1"`
	Child2      string   `gql:"child2"`
	ChildNested []string `gql:"childNested"`
}

func main() {
	nested := Nested{
		Child1:      "child 1",
		Child2:      "child 2",
		ChildNested: []string{"hello 1", "hello 2"},
	}
	cat := new(Category)
	cat.Name = "category 1"
	cat.Slug = "category-1"
	cat.Nested.Child1 = "child nested 1"
	cat.Nested.Child2 = "child nested 2"
	cat.Nested.ChildNested = []string{"nested 1 ", "nested 233"}
	cat.NestedSlice = []Nested{
		nested,
		nested,
	}

	r := []string{"success", "response"}

	query := sailfish.NewQuery("mutation", "createCategory")
	s := query.Parse(cat, r)
	fmt.Println(s)
}
