package errors

type invalidArgument struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
