package adapters

type Sign interface {
	Encode(data map[string]interface{}) (string, error)
}
