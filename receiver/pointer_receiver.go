package receiver

import "fmt"

type PointerReceiver struct {
	Name string
}

func (p *PointerReceiver) Address() string {
	return fmt.Sprintf("%p", p)
}

func (p *PointerReceiver) SetName(name string) {
	p.Name = name
}
