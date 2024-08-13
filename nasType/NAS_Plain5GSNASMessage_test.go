package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/BENHSU0723/nas_public/nasType"
)

func TestNasTypeNewPlain5GSNASMessage(t *testing.T) {
	a := nasType.NewPlain5GSNASMessage()
	assert.NotNil(t, a)
}
