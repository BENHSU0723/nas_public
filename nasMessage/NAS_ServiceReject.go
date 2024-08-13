// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/BENHSU0723/nas_public/nasType"
)

type ServiceReject struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.ServiceRejectMessageIdentity
	nasType.Cause5GMM
	*nasType.PDUSessionStatus
	*nasType.T3346Value
	*nasType.EAPMessage
}

func NewServiceReject(iei uint8) (serviceReject *ServiceReject) {
	serviceReject = &ServiceReject{}
	return serviceReject
}

const (
	ServiceRejectPDUSessionStatusType uint8 = 0x50
	ServiceRejectT3346ValueType       uint8 = 0x5F
	ServiceRejectEAPMessageType       uint8 = 0x78
)

func (a *ServiceReject) EncodeServiceReject(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (ServiceReject/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (ServiceReject/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.ServiceRejectMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (ServiceReject/ServiceRejectMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.Cause5GMM.Octet); err != nil {
		return fmt.Errorf("NAS encode error (ServiceReject/Cause5GMM): %w", err)
	}
	if a.PDUSessionStatus != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ServiceReject/PDUSessionStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ServiceReject/PDUSessionStatus): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ServiceReject/PDUSessionStatus): %w", err)
		}
	}
	if a.T3346Value != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.T3346Value.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ServiceReject/T3346Value): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.T3346Value.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ServiceReject/T3346Value): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.T3346Value.Octet); err != nil {
			return fmt.Errorf("NAS encode error (ServiceReject/T3346Value): %w", err)
		}
	}
	if a.EAPMessage != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ServiceReject/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ServiceReject/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ServiceReject/EAPMessage): %w", err)
		}
	}
	return nil
}

func (a *ServiceReject) DecodeServiceReject(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (ServiceReject/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (ServiceReject/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ServiceRejectMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (ServiceReject/ServiceRejectMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GMM.Octet); err != nil {
		return fmt.Errorf("NAS decode error (ServiceReject/Cause5GMM): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (ServiceReject/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case ServiceRejectPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len); err != nil {
				return fmt.Errorf("NAS decode error (ServiceReject/PDUSessionStatus): %w", err)
			}
			if a.PDUSessionStatus.Len < 2 || a.PDUSessionStatus.Len > 32 {
				return fmt.Errorf("invalid ie length (ServiceReject/PDUSessionStatus): %d", a.PDUSessionStatus.Len)
			}
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ServiceReject/PDUSessionStatus): %w", err)
			}
		case ServiceRejectT3346ValueType:
			a.T3346Value = nasType.NewT3346Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.T3346Value.Len); err != nil {
				return fmt.Errorf("NAS decode error (ServiceReject/T3346Value): %w", err)
			}
			if a.T3346Value.Len != 1 {
				return fmt.Errorf("invalid ie length (ServiceReject/T3346Value): %d", a.T3346Value.Len)
			}
			a.T3346Value.SetLen(a.T3346Value.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.T3346Value.Octet); err != nil {
				return fmt.Errorf("NAS decode error (ServiceReject/T3346Value): %w", err)
			}
		case ServiceRejectEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return fmt.Errorf("NAS decode error (ServiceReject/EAPMessage): %w", err)
			}
			if a.EAPMessage.Len < 4 || a.EAPMessage.Len > 1500 {
				return fmt.Errorf("invalid ie length (ServiceReject/EAPMessage): %d", a.EAPMessage.Len)
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ServiceReject/EAPMessage): %w", err)
			}
		default:
		}
	}
	return nil
}
