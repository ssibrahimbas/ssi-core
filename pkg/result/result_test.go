package result

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResult_Module(t *testing.T) {
	t.Run("Success function should return a success result", func(t *testing.T) {
		result := Success("Success", 200)
		assert.Equal(t, true, result.Success)
	})

	t.Run("Error function should return an error result", func(t *testing.T) {
		result := Error("Error", 500)
		assert.NotEqual(t, true, result.Success)
	})

	t.Run("SuccessData function should return a success result with data", func(t *testing.T) {
		result := SuccessData("Success", "data", 200)
		assert.Equal(t, "data", result.Data)
		assert.Equal(t, true, result.Success)
	})

	t.Run("ErrorData function should return an error result with data", func(t *testing.T) {
		result := ErrorData("Error", "data", 500)
		assert.Equal(t, "data", result.Data)
		assert.NotEqual(t, true, result.Success)
	})

	t.Run("Error function should return an error result with data", func(t *testing.T) {
		result := ErrorData("Error", "data", 500)
		assert.Equal(t, "data", result.Data)
		assert.NotEqual(t, true, result.Success)
		assert.Equal(t, "Error", result.Error())
	})

	t.Run("Error function should return an error message", func(t *testing.T) {
		result := Error("Error", 500)
		assert.Equal(t, "Error", result.Error())
	})
}
