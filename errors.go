package mandrill

import (
	"fmt"
)

type APIError struct {
	Status  string
	Code    int
	Name    string
	Message string
}

func (a *APIError) Error() string {
	return fmt.Sprintf("%s (%d). %s. %s", a.Status, a.Code, a.Name, a.Message)
}
