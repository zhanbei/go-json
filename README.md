# Go Customizable Encoding/Json

<!-- > 2018-04-08T19:54:12+0800 -->

The customizable encoding/json package based on [`zhanbei/golang-encoding-json`][project-upstream] -- the [`encoding/json`][package-encoding-json] package separated from [`golang/go/master`][github-golang-go].

## Installation

```bash
go get github.com/zhanbei/go-json
```

## Usage and Example

The example codes can be found in [example/main.go](example/main.go).

```go
// Example script of using Zibson with core features.
package main

import (
	"github.com/zhanbei/go-json"
	"fmt"
)

var DefaultZibson = json.NewZibson()
var OptimizedZibson = json.NewZibson()
var ZibsonToOriginalJson = json.NewZibson()
var ZibsonToIntegratedJson = json.NewZibson()

func init() {
	OptimizedZibson.SetFieldNameToJsonKeyFunc(json.FieldNameToJsonKeyFuncLowerInitialLetter)

	ZibsonToOriginalJson.SetCustomJsonTag("originalJson")
	ZibsonToOriginalJson.SetFieldNameToJsonKeyFunc(json.FieldNameToJsonKeyFuncLowerInitialLetter)

	ZibsonToIntegratedJson.SetCustomJsonTag("integratedJson")
	ZibsonToIntegratedJson.SetFieldNameToJsonKeyFunc(json.FieldNameToJsonKeyFuncLowerInitialLetter)
}

type Person struct {
	// Using `json` tag to override the JSON key.
	Id string `json:"_id"`
	// By default, the JSON key is controlled by zibson#FieldNameToJsonKeyFunc.
	Age int
	// The firstName and lastName is used internally only.
	// We expose it only when the `originalJson` tag is used.
	FirstName string `json:"-" originalJson:"firstName"`
	LastName  string `json:"-" originalJson:"lastName"`
	// The name generated in #Normalize() by connecting the FirstName with the LastName.
	// We expose it only when the `integratedJson` tag is used.
	Name string `json:"-" integratedJson:"name"`
	// The prioritised `toJson` tag makes secret not exposing to JSON when encoding.
	Secret string `json:"secret" toJson:"-"`
	// Combine signature and bio using #Normalize().
	// Expose bio only when encoding to JSON.
	Bio       string
	Signature string `toJson:"-"`
}

func (m *Person) ToIntegratedJson() string {
	return MustToJson(ZibsonToIntegratedJson, m)
}

func (m *Person) ToOriginalJson() string {
	return MustToJson(ZibsonToOriginalJson, m)
}

func (m *Person) ToOptimizedJson() string {
	return MustToJson(OptimizedZibson, m)
}

func (m *Person) ToDefaultJson() string {
	return MustToJson(DefaultZibson, m)
}

func (m *Person) Normalize() {
	m.Name = m.LastName + " " + m.LastName
	if m.Bio == "" {
		m.Bio = m.Signature
	}
}

func MustToJson(zibson *json.Zibson, v interface{}) string {
	bytes, err := zibson.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

const mPersonJson = `{"_id":"No.1","age":8,"firstName":"Tom","lastName":"Sawyer","secret":"Hola everyone, I am alive!","signature":"Bye, have a good day!"}`

func GetPerson() *Person {
	person := new(Person)
	err := ZibsonToOriginalJson.Unmarshal([]byte(mPersonJson), person)
	if err != nil {
		panic(err)
	}
	//person = &Person{FirstName:"Tom", LastName:"Sawyer"}
	person.Normalize()
	return person
}

func main() {
	person := GetPerson()

	fmt.Println("---->>")
	fmt.Println("person.ToOriginalJson():", person.ToOriginalJson())
	fmt.Println("person.ToIntegratedJson():", person.ToIntegratedJson())
	fmt.Println("person.ToOptimizedJson():", person.ToOptimizedJson())
	fmt.Println("person.ToDefaultJson():", person.ToDefaultJson())
	fmt.Println("person.Secret ----->>>:", person.Secret)

	// The exposed `json.Marshal()` uses a package-scope default instance of Zibson.
	bytes, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println("---->>")
	fmt.Println("person.ToDefaultJson() == json.Marshal(person):", person.ToDefaultJson() == string(bytes))

	fmt.Println("---->>")
	fmt.Println("ZibsonToOriginalJson:", MustToJson(DefaultZibson, ZibsonToOriginalJson))
	fmt.Println("ZibsonToOriginalJson:", MustToJson(OptimizedZibson, ZibsonToOriginalJson))
	fmt.Println("ZibsonToIntegratedJson:", MustToJson(DefaultZibson, ZibsonToIntegratedJson))
	fmt.Println("ZibsonToIntegratedJson:", MustToJson(OptimizedZibson, ZibsonToIntegratedJson))
}
```

### Example Output

```text
---->>
person.ToOriginalJson(): {"_id":"No.1","age":8,"firstName":"Tom","lastName":"Sawyer","bio":"Bye, have a good day!"}
person.ToIntegratedJson(): {"_id":"No.1","age":8,"name":"Sawyer Sawyer","bio":"Bye, have a good day!"}
person.ToOptimizedJson(): {"_id":"No.1","age":8,"bio":"Bye, have a good day!"}
person.ToDefaultJson(): {"_id":"No.1","Age":8,"Bio":"Bye, have a good day!"}
person.Secret ----->>>: Hola everyone, I am alive!
---->>
person.ToDefaultJson() == json.Marshal(person): true
---->>
ZibsonToOriginalJson: {"FromJsonTag":"fromJson","ToJsonTag":"toJson","CustomJsonTag":"originalJson","DefaultJsonTag":"json"}
ZibsonToOriginalJson: {"fromJsonTag":"fromJson","toJsonTag":"toJson","customJsonTag":"originalJson","defaultJsonTag":"json"}
ZibsonToIntegratedJson: {"FromJsonTag":"fromJson","ToJsonTag":"toJson","CustomJsonTag":"integratedJson","DefaultJsonTag":"json"}
ZibsonToIntegratedJson: {"fromJsonTag":"fromJson","toJsonTag":"toJson","customJsonTag":"integratedJson","defaultJsonTag":"json"}
```

## Features

The goal is to add more customizable features while keeping compatible with the [`encoding/json`][package-encoding-json] package.

- [x] Customizable/Different tags(`fromJson`, `toJson`) for decoding from and encoding to JSON string.
- [x] Customizable JSON struct tag(i.e. `customizedJsonTag`)
- [x] Customizable fall-back JSON struct tag(`json`).
- [x] Customizable conversion policy to JSON keys from field names of a struct.

### Differences with the [`encoding/json`][package-encoding-json] Package

- By default, the tags `fromJson` and `toJson` will be checked first respectively, over the tag `json`.

[github-golang-go]: https://github.com/golang/go "Go Source Code"
[package-encoding-json]: https://github.com/golang/go/tree/master/src/encoding/json "Go Package `encoding/json`"
[project-upstream]: https://github.com/zhanbei/golang-encoding-json "Project Upstream: Separated `encoding/json` Package"
