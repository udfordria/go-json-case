package __tests__

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stoewer/go-strcase"
	json_case "github.com/udfordria/go-json-case"
)

type User struct {
	Name    string
	Age     int
	Country string
}

func TestJSONCase(t *testing.T) {
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
