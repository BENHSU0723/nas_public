package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	nas "github.com/BENHSU0723/nas_public"
	"github.com/BENHSU0723/nas_public/nasType"
)

type nasTypeIdentityResponseMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeIdentityResponseMessageIdentityTable = []nasTypeIdentityResponseMessageIdentityData{
	{nas.MsgTypeIdentityResponse, nas.MsgTypeIdentityResponse},
}

func TestNasTypeNewIdentityResponseMessageIdentity(t *testing.T) {
	a := nasType.NewIdentityResponseMessageIdentity()
	assert.NotNil(t, a)
}

func TestNasTypeIdentityResponseMessageIdentity(t *testing.T) {
	a := nasType.NewIdentityResponseMessageIdentity()
	for _, table := range nasTypeIdentityResponseMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
