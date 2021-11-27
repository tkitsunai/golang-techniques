package receiver_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-techniques/receiver"
	"testing"
)

func TestReceiver(t *testing.T) {
	t.Run("pointer receiver", func(t *testing.T) {
		pointerReceiver := receiver.PointerReceiver{}
		address := pointerReceiver.Address()
		assert.Equal(t, address, fmt.Sprintf("%p", &pointerReceiver))
	})

	t.Run("value receiver", func(t *testing.T) {
		valueReceiver := receiver.ValueReceiver{}
		address := valueReceiver.Address()
		assert.NotEqual(t, address, fmt.Sprintf("%p", &valueReceiver))
	})

	t.Run("immutable method", func(t *testing.T) {
		valueReceiver := receiver.ValueReceiver{
			Name: "NAME1",
		}

		valueReceiver.SetName("test")

		assert.NotEqual(t, "test", valueReceiver.Name)
		assert.Equal(t, "NAME1", valueReceiver.Name)
	})
}
