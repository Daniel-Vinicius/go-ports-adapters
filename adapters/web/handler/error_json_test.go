package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(testing *testing.T) {
	msg := "Hello Json"

	expectedResult := []byte(`{"message":"Hello Json"}`)
	result := jsonError(msg)

	require.Equal(testing, expectedResult, result)
}
