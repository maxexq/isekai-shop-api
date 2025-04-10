package exception

import "fmt"

type ItemAlreadyExists struct {
	name string
}

func (e *ItemAlreadyExists) Error() string {
	return fmt.Sprintf("creating item name: %s failed. item is already exist.", e.name)
}
