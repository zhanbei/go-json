package json

import (
	"unicode"
	"sync"
)

// The default configurations of Zibson.
const (
	defaultFromJsonTag = "fromJson"

	defaultToJsonTag = "toJson"

	defaultCustomJsonTag = ""

	defaultDefaultJsonTag = "json"
)

var (
	// Lower the initial letter of the field name as JSON key using `unicode.ToLower()`.
	FieldNameToJsonKeyFuncLowerInitialLetter = func(fieldName string) string {
		if fieldName == "" {
			return fieldName
		}
		a := []rune(fieldName)
		a[0] = unicode.ToLower(a[0])
		return string(a)
	}
	// Use the original field name as the JSON key.
	FieldNameToJsonKeyFuncOrigin = func(fieldName string) string { return fieldName }
	// By default, JSON key will be exactly the field name.
	defaultFieldNameToJsonKeyFunc = FieldNameToJsonKeyFuncOrigin
)

// The default instance of Zibson, using default configurations.
// And should always be used internally.
var mDefaultZibson = NewZibson()

func init() {
	// By default, the tags `fromJson` and `toJson` will be checked first respectively, over the tag `json`.
	// This configuration makes zibson not fully compatible with the `encoding/json` package.
	// @see `#README.md#Features#Differences`.
	mDefaultZibson.SetFromAndToJsonTags(defaultFromJsonTag, defaultToJsonTag)
}

// Get the default instance of Zibson; initialize a new instance if nil is found.
func GetDefaultZibson() *Zibson {
	if mDefaultZibson == nil {
		mDefaultZibson = NewZibson()
	}
	return mDefaultZibson
}

// The JSON encoder and decoder helper with configurations.
// Named as #Zibson for nearly no reason. PS. a suggested name is recommended.
// Tag Checking Priority: FromJsonTag/ToJsonTag > CustomJsonTag > DefaultJsonTag > FieldNameToJsonKeyFunc(fieldName) > FieldName
type Zibson struct {
	// The prioritized tag for decoding from JSON, by default: "fromJson".
	FromJsonTag string
	// The prioritized tag for encoding to JSON, by default: "toJson".
	ToJsonTag string
	// The custom tag for encoding/to or decoding/from JSON, by default: "".
	CustomJsonTag string
	// The default tag for if custom tag is not found, which most often is "json", and by default is "json".
	// Set to empty string "" to disable picking up the default tag "json".
	DefaultJsonTag string
	// The conversion from field name to JSON key.
	// By default, JSON key will be exactly the field name.
	FieldNameToJsonKeyFunc func(fieldName string) string `json:"-"`
	// Caches will be cleared when fields of zibson changes.
	// @see encode#typeEncoder().
	EncodingCacheTypeToEncoderFunc sync.Map `json:"-"` // map[reflect.Type]encoderFunc
	// @see encode#getCachedTypeFieldsForEncoding().
	EncodingCacheTypeToTypeFields sync.Map `json:"-"` // map[reflect.Type][]field
	// @see encode#getCachedTypeFieldsForEncoding().
	DecodingCacheTypeToTypeFields sync.Map `json:"-"` // map[reflect.Type][]field
}

func NewZibson() *Zibson {
	return &Zibson{
		defaultFromJsonTag,
		defaultToJsonTag,
		defaultCustomJsonTag,
		defaultDefaultJsonTag,
		defaultFieldNameToJsonKeyFunc,
		sync.Map{},
		sync.Map{},
		sync.Map{},
	}
}

func (m *Zibson) ClearCaches() {
	m.EncodingCacheTypeToEncoderFunc = sync.Map{}
	m.EncodingCacheTypeToTypeFields = sync.Map{}
	m.DecodingCacheTypeToTypeFields = sync.Map{}
}

// Set the custom JSON tag for decoding from JSON and encoding to JSON.
// Set to empty string to disable the tags. The default tags are: "fromJson", and "toJson".
func (m *Zibson) SetFromAndToJsonTags(fromJsonTag, toJsonTag string) *Zibson {
	m.FromJsonTag = fromJsonTag
	m.ToJsonTag = toJsonTag
	m.ClearCaches()
	return m
}

// Set the custom JSON tag, which by default is empty.
func (m *Zibson) SetCustomJsonTag(customJsonTag string) *Zibson {
	m.CustomJsonTag = customJsonTag
	m.ClearCaches()
	return m
}

// Disable the default tag: "json".
func (m *Zibson) DisableDefaultJsonTag() *Zibson {
	m.SetDefaultJsonTag("")
	return m
}

// Set the default tag "json".
func (m *Zibson) SetDefaultJsonTag(defaultJsonTag string) *Zibson {
	m.DefaultJsonTag = defaultJsonTag
	m.ClearCaches()
	return m
}

func (m *Zibson) SetFieldNameToJsonKeyFunc(fieldNameToJsonKeyFunc func(fieldName string) string) *Zibson {
	m.FieldNameToJsonKeyFunc = fieldNameToJsonKeyFunc
	m.ClearCaches()
	return m
}
