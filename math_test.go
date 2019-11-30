package fortnight_test

import (
	"testing"

	fortnight "github.com/bickyeric/fuzzy-octo-fortnight"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := fortnight.Add(5, 9)
	assert.Equal(t, 14, result)
}
