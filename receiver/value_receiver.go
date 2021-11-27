package receiver

import "fmt"

type ValueReceiver struct {
	Name string
}

func (v ValueReceiver) Address() string {
	return fmt.Sprintf("%p", &v)
}

func (v ValueReceiver) SetName(name string) {
	v.Name = name
}
