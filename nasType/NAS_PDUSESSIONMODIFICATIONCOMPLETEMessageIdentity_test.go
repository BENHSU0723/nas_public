package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	nas "github.com/BENHSU0723/nas_public"
	"github.com/BENHSU0723/nas_public/nasType"
)

func TestNasTypeNewPDUSESSIONMODIFICATIONCOMPLETEMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONCOMPLETEMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONMODIFICATIONCOMPLETEMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONMODIFICATIONCOMPLETEMessageIdentityMessageTypeTable = []nasTypePDUSESSIONMODIFICATIONCOMPLETEMessageIdentityMessageType{
	{nas.MsgTypePDUSessionModificationComplete, nas.MsgTypePDUSessionModificationComplete},
}

func TestNasTypeGetSetPDUSESSIONMODIFICATIONCOMPLETEMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONCOMPLETEMessageIdentity()
	for _, table := range nasTypePDUSESSIONMODIFICATIONCOMPLETEMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
