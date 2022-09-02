package helper

import (
	"errors"
	"testing"
)

func TestCheckErr(t *testing.T) {
	t.Run("CheckErr with nil value", func(t *testing.T) {
		CheckErr(nil)
	})

	t.Run("CheckErr with nil error", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		CheckErr(errors.New("test"))
	})
}
