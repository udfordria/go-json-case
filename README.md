# JSON case

Change the case of the keys of the maps and structures when parsing and serializing JSON data.

## Installation

```cmd
go get -u github.com/udfordria/go-json-case
```

## Usage

### Steps

1. Call the `json_case.SetNamingStrategy` function.
2. Use `jsoniter "github.com/json-iterator/go"` instead of `"encoding/json"` for all JSON-related operations.

### Example

```go
import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/stoewer/go-strcase"
	json_case "github.com/udfordria/go-json-case"
)

type User struct {
	Name    string
	Age     int
	Country string
}

func main() {
	json_case.SetNamingStrategy(strcase.SnakeCase)

	user := User{}
	jsonContent := `{"name" : "Alice", "age" : 22, "country": "DK"}`
	err := jsoniter.UnmarshalFromString(jsonContent, &user)

	if err != nil {
		panic(err)
	}

	fmt.Println(user.Name)    // Alice
	fmt.Println(user.Age)     // 22
	fmt.Println(user.Country) // DK

	output, err := jsoniter.MarshalToString(user)

	if err != nil {
		panic(err)
	}

	fmt.Println(output) // {"name":"Alice","age":22,"country":"DK"}
}
```

## Recommendation

Adding a `.golangci.yml` with the following content will make `go-golangci-lint` tool warn about the usage of the standard `"encoding/json"`:

```yml
linters-settings:
  depguard:
    rules:
      json: 
        files:
          - "**/*.go"
        deny:
          - pkg: "encoding/json"
            desc: use jsoniter instead

linters:
  enable:
    - depguard
```