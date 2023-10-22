package main

import (
	"encoding/json"
	"fmt"
	"github.com/dave/jennifer/jen"
	"log"
)

func GenerateCode(jsonString string) (string, error) {
	var data interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	fmt.Println(data)
	if err != nil {
		return "", err
	}

	f := jen.NewFile("main")
	generateStruct("MyStruct", data, f)

	return fmt.Sprintf("%#v", f), nil
}

func generateStruct(structName string, data interface{}, f *jen.File) {
	switch value := data.(type) {
	case map[string]interface{}:
		f.Type().Id(structName).StructFunc(func(g *jen.Group) {
			for key, val := range value {
				g.Id(key).Add(generateFieldType(val))
			}
		})
	case []interface{}:
		f.Type().Id(structName).Struct(jen.Id("Items").Index().Interface())
	default:
		f.Type().Id(structName).Add(generateFieldType(value))
	}
}

func generateFieldType(value interface{}) *jen.Statement {
	switch v := value.(type) {
	case string:
		return jen.String().Tag(map[string]string{"json": v})
	case float64:
		return jen.Float64().Tag(map[string]string{"json": fmt.Sprintf("%f", v)})
	case bool:
		return jen.Bool().Tag(map[string]string{"json": fmt.Sprintf("%t", v)})
	default:
		return jen.Empty()
	}
}

func main() {
	jsonString := `{
		"field_name": "value",
		"struct_name": {
			"name": "dadad"
		},
		"numeric_field": 42,
		"boolean_field": true
	}`

	generatedCode, err := GenerateCode(jsonString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(generatedCode)
}
