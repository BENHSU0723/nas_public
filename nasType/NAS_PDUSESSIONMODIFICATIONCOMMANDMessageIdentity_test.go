package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	nas "github.com/BENHSU0723/nas_public"
	"github.com/BENHSU0723/nas_public/nasType"
)

func TestNasTypeNewPDUSESSIONMODIFICATIONCOMMANDMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONCOMMANDMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONMODIFICATIONCOMMANDMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONMODIFICATIONCOMMANDMessageIdentityMessageTypeTable = []nasTypePDUSESSIONMODIFICATIONCOMMANDMessageIdentityMessageType{
	{nas.MsgTypePDUSessionModificationCommand, nas.MsgTypePDUSessionModificationCommand},
}

func TestNasTypeGetSetPDUSESSIONMODIFICATIONCOMMANDMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONCOMMANDMessageIdentity()
	for _, table := range nasTypePDUSESSIONMODIFICATIONCOMMANDMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
