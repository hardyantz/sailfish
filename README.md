# Sailfish

Simple graphQL generator request body for client. 


## How to

```go
    type Category Struct {
	    Name string `gql:"name"`
    }
    cat := new(Category)
    cat.Name = "meow"
    r := []string{"success", "message"}
    query := sailfish.NewQuery("mutation", "createCategory")
	s := query.Parse(cat, r)
	fmt.Println(s)
``` 

and the output : 
```graphql
    mutation createCategory(name:"meow") {success message}
```
