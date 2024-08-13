// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/BENHSU0723/nas_public/nasType"
)

type DeregistrationRequestUETerminatedDeregistration struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.DeregistrationRequestMessageIdentity
	nasType.SpareHalfOctetAndDeregistrationType
	*nasType.Cause5GMM
	*nasType.T3346Value
}

func NewDeregistrationRequestUETerminatedDeregistration(iei uint8) (deregistrationRequestUETerminatedDeregistration *DeregistrationRequestUETerminatedDeregistration) {
	deregistrationRequestUETerminatedDeregistration = &DeregistrationRequestUETerminatedDeregistration{}
	return deregistrationRequestUETerminatedDeregistration
}

const (
	DeregistrationRequestUETerminatedDeregistrationCause5GMMType  uint8 = 0x58
	DeregistrationRequestUETerminatedDeregistrationT3346ValueType uint8 = 0x5F
)

func (a *DeregistrationRequestUETerminatedDeregistration) EncodeDeregistrationRequestUETerminatedDeregistration(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationRequestUETerminatedDeregistration/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationRequestUETerminatedDeregistration/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.DeregistrationRequestMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationRequestUETerminatedDeregistration/DeregistrationRequestMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndDeregistrationType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationRequestUETerminatedDeregistration/SpareHalfOctetAndDeregistrationType): %w", err)
	}
	if a.Cause5GMM != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Cause5GMM.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (DeregistrationRequestUETerminatedDeregistration/Cause5GMM): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Cause5GMM.Octet); err != nil {
			return fmt.Errorf("NAS encode error (DeregistrationRequestUETerminatedDeregistration/Cause5GMM): %w", err)
		}
	}
	if a.T3346Value != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.T3346Value.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (DeregistrationRequestUETerminatedDeregistration/T3346Value): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.T3346Value.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (DeregistrationRequestUETerminatedDeregistration/T3346Value): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.T3346Value.Octet); err != nil {
			return fmt.Errorf("NAS encode error (DeregistrationRequestUETerminatedDeregistration/T3346Value): %w", err)
		}
	}
	return nil
}

func (a *DeregistrationRequestUETerminatedDeregistration) DecodeDeregistrationRequestUETerminatedDeregistration(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationRequestUETerminatedDeregistration/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationRequestUETerminatedDeregistration/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.DeregistrationRequestMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationRequestUETerminatedDeregistration/DeregistrationRequestMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndDeregistrationType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationRequestUETerminatedDeregistration/SpareHalfOctetAndDeregistrationType): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (DeregistrationRequestUETerminatedDeregistration/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case DeregistrationRequestUETerminatedDeregistrationCause5GMMType:
			a.Cause5GMM = nasType.NewCause5GMM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GMM.Octet); err != nil {
				return fmt.Errorf("NAS decode error (DeregistrationRequestUETerminatedDeregistration/Cause5GMM): %w", err)
			}
		case DeregistrationRequestUETerminatedDeregistrationT3346ValueType:
			a.T3346Value = nasType.NewT3346Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.T3346Value.Len); err != nil {
				return fmt.Errorf("NAS decode error (DeregistrationRequestUETerminatedDeregistration/T3346Value): %w", err)
			}
			if a.T3346Value.Len != 1 {
				return fmt.Errorf("invalid ie length (DeregistrationRequestUETerminatedDeregistration/T3346Value): %d", a.T3346Value.Len)
			}
			a.T3346Value.SetLen(a.T3346Value.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.T3346Value.Octet); err != nil {
				return fmt.Errorf("NAS decode error (DeregistrationRequestUETerminatedDeregistration/T3346Value): %w", err)
			}
		default:
		}
	}
	return nil
}
