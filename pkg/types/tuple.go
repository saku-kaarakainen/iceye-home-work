package types

// KeyValuePair is a struct that represents a key-value pair where the key and value can be of any type.
// It is a generic struct and can be used with any types that satisfy the 'comparable' constraint.
type KeyValuePair[Key any, Value any] struct {
	Key Key
	Value Value
}