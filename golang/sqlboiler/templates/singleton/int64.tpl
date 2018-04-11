import (
	"errors"
	"strconv"
)

// Int64 represents a custom GraphQL "Int64" scalar type
// Implements graphql.Unmarshaler
type Int64 string

// ImplementsGraphQLType returns the GraphQL type name
func (Int64) ImplementsGraphQLType(name string) bool {
	return name == "Int64"
}

// UnmarshalGraphQL unmarshals the GraphQL type
func (i *Int64) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		*i = Int64(input)
	case int64:
		*i = Int64(strconv.FormatInt(int64(input), 10))
	default:
		err = errors.New("wrong type")
	}
	return err
}

// MarshalJSON implements JSON marshalling
func (i Int64) MarshalJSON() ([]byte, error) {
	return []byte(i), nil
}
