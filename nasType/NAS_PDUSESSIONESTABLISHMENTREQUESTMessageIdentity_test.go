package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	nas "github.com/BENHSU0723/nas_public"
	"github.com/BENHSU0723/nas_public/nasType"
)

func TestNasTypeNewPDUSESSIONESTABLISHMENTREQUESTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONESTABLISHMENTREQUESTMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONESTABLISHMENTREQUESTMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONESTABLISHMENTREQUESTMessageIdentityMessageTypeTable = []nasTypePDUSESSIONESTABLISHMENTREQUESTMessageIdentityMessageType{
	{nas.MsgTypePDUSessionEstablishmentRequest, nas.MsgTypePDUSessionEstablishmentRequest},
}

func TestNasTypeGetSetPDUSESSIONESTABLISHMENTREQUESTMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONESTABLISHMENTREQUESTMessageIdentity()
	for _, table := range nasTypePDUSESSIONESTABLISHMENTREQUESTMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
