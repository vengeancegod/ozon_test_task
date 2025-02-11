package generator

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestGenerateShortURL(t *testing.T) {
  shortURL := GenerateShortURL()
  assert.Equal(t, 10, len(shortURL)) 
}