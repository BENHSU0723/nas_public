// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/BENHSU0723/nas_public/nasType"
)

type ConfigurationUpdateCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.ConfigurationUpdateCommandMessageIdentity
	*nasType.ConfigurationUpdateIndication
	*nasType.GUTI5G
	*nasType.TAIList
	*nasType.AllowedNSSAI
	*nasType.ServiceAreaList
	*nasType.FullNameForNetwork
	*nasType.ShortNameForNetwork
	*nasType.LocalTimeZone
	*nasType.UniversalTimeAndLocalTimeZone
	*nasType.NetworkDaylightSavingTime
	*nasType.LADNInformation
	*nasType.MICOIndication
	*nasType.NetworkSlicingIndication
	*nasType.ConfiguredNSSAI
	*nasType.RejectedNSSAI
	*nasType.OperatordefinedAccessCategoryDefinitions
	*nasType.SMSIndication
}

func NewConfigurationUpdateCommand(iei uint8) (configurationUpdateCommand *ConfigurationUpdateCommand) {
	configurationUpdateCommand = &ConfigurationUpdateCommand{}
	return configurationUpdateCommand
}

const (
	ConfigurationUpdateCommandConfigurationUpdateIndicationType            uint8 = 0x0D
	ConfigurationUpdateCommandGUTI5GType                                   uint8 = 0x77
	ConfigurationUpdateCommandTAIListType                                  uint8 = 0x54
	ConfigurationUpdateCommandAllowedNSSAIType                             uint8 = 0x15
	ConfigurationUpdateCommandServiceAreaListType                          uint8 = 0x27
	ConfigurationUpdateCommandFullNameForNetworkType                       uint8 = 0x43
	ConfigurationUpdateCommandShortNameForNetworkType                      uint8 = 0x45
	ConfigurationUpdateCommandLocalTimeZoneType                            uint8 = 0x46
	ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType            uint8 = 0x47
	ConfigurationUpdateCommandNetworkDaylightSavingTimeType                uint8 = 0x49
	ConfigurationUpdateCommandLADNInformationType                          uint8 = 0x79
	ConfigurationUpdateCommandMICOIndicationType                           uint8 = 0x0B
	ConfigurationUpdateCommandNetworkSlicingIndicationType                 uint8 = 0x09
	ConfigurationUpdateCommandConfiguredNSSAIType                          uint8 = 0x31
	ConfigurationUpdateCommandRejectedNSSAIType                            uint8 = 0x11
	ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType uint8 = 0x76
	ConfigurationUpdateCommandSMSIndicationType                            uint8 = 0x0F
)

func (a *ConfigurationUpdateCommand) EncodeConfigurationUpdateCommand(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.ConfigurationUpdateCommandMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/ConfigurationUpdateCommandMessageIdentity): %w", err)
	}
	if a.ConfigurationUpdateIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ConfigurationUpdateIndication.Octet); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/ConfigurationUpdateIndication): %w", err)
		}
	}
	if a.GUTI5G != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.GUTI5G.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/GUTI5G): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.GUTI5G.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/GUTI5G): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.GUTI5G.Octet[:]); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/GUTI5G): %w", err)
		}
	}
	if a.TAIList != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.TAIList.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/TAIList): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.TAIList.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/TAIList): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.TAIList.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/TAIList): %w", err)
		}
	}
	if a.AllowedNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AllowedNSSAI.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/AllowedNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AllowedNSSAI.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/AllowedNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AllowedNSSAI.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/AllowedNSSAI): %w", err)
		}
	}
	if a.ServiceAreaList != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ServiceAreaList.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/ServiceAreaList): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ServiceAreaList.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/ServiceAreaList): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ServiceAreaList.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/ServiceAreaList): %w", err)
		}
	}
	if a.FullNameForNetwork != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.FullNameForNetwork.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/FullNameForNetwork): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.FullNameForNetwork.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/FullNameForNetwork): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.FullNameForNetwork.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/FullNameForNetwork): %w", err)
		}
	}
	if a.ShortNameForNetwork != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ShortNameForNetwork.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/ShortNameForNetwork): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ShortNameForNetwork.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/ShortNameForNetwork): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ShortNameForNetwork.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/ShortNameForNetwork): %w", err)
		}
	}
	if a.LocalTimeZone != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.LocalTimeZone.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/LocalTimeZone): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.LocalTimeZone.Octet); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/LocalTimeZone): %w", err)
		}
	}
	if a.UniversalTimeAndLocalTimeZone != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.UniversalTimeAndLocalTimeZone.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/UniversalTimeAndLocalTimeZone): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.UniversalTimeAndLocalTimeZone.Octet[:]); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/UniversalTimeAndLocalTimeZone): %w", err)
		}
	}
	if a.NetworkDaylightSavingTime != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NetworkDaylightSavingTime.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/NetworkDaylightSavingTime): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NetworkDaylightSavingTime.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/NetworkDaylightSavingTime): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.NetworkDaylightSavingTime.Octet); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/NetworkDaylightSavingTime): %w", err)
		}
	}
	if a.LADNInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.LADNInformation.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/LADNInformation): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.LADNInformation.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/LADNInformation): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.LADNInformation.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/LADNInformation): %w", err)
		}
	}
	if a.MICOIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.MICOIndication.Octet); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/MICOIndication): %w", err)
		}
	}
	if a.NetworkSlicingIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.NetworkSlicingIndication.Octet); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/NetworkSlicingIndication): %w", err)
		}
	}
	if a.ConfiguredNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ConfiguredNSSAI.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/ConfiguredNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ConfiguredNSSAI.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/ConfiguredNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ConfiguredNSSAI.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/ConfiguredNSSAI): %w", err)
		}
	}
	if a.RejectedNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RejectedNSSAI.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/RejectedNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RejectedNSSAI.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/RejectedNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RejectedNSSAI.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/RejectedNSSAI): %w", err)
		}
	}
	if a.OperatordefinedAccessCategoryDefinitions != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/OperatordefinedAccessCategoryDefinitions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/OperatordefinedAccessCategoryDefinitions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/OperatordefinedAccessCategoryDefinitions): %w", err)
		}
	}
	if a.SMSIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SMSIndication.Octet); err != nil {
			return fmt.Errorf("NAS encode error (ConfigurationUpdateCommand/SMSIndication): %w", err)
		}
	}
	return nil
}

func (a *ConfigurationUpdateCommand) DecodeConfigurationUpdateCommand(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ConfigurationUpdateCommandMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/ConfigurationUpdateCommandMessageIdentity): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case ConfigurationUpdateCommandConfigurationUpdateIndicationType:
			a.ConfigurationUpdateIndication = nasType.NewConfigurationUpdateIndication(ieiN)
			a.ConfigurationUpdateIndication.Octet = ieiN
		case ConfigurationUpdateCommandGUTI5GType:
			a.GUTI5G = nasType.NewGUTI5G(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.GUTI5G.Len); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/GUTI5G): %w", err)
			}
			if a.GUTI5G.Len != 11 {
				return fmt.Errorf("invalid ie length (ConfigurationUpdateCommand/GUTI5G): %d", a.GUTI5G.Len)
			}
			a.GUTI5G.SetLen(a.GUTI5G.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.GUTI5G.Octet[:]); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/GUTI5G): %w", err)
			}
		case ConfigurationUpdateCommandTAIListType:
			a.TAIList = nasType.NewTAIList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.TAIList.Len); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/TAIList): %w", err)
			}
			if a.TAIList.Len < 7 || a.TAIList.Len > 112 {
				return fmt.Errorf("invalid ie length (ConfigurationUpdateCommand/TAIList): %d", a.TAIList.Len)
			}
			a.TAIList.SetLen(a.TAIList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.TAIList.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/TAIList): %w", err)
			}
		case ConfigurationUpdateCommandAllowedNSSAIType:
			a.AllowedNSSAI = nasType.NewAllowedNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AllowedNSSAI.Len); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/AllowedNSSAI): %w", err)
			}
			if a.AllowedNSSAI.Len < 2 || a.AllowedNSSAI.Len > 72 {
				return fmt.Errorf("invalid ie length (ConfigurationUpdateCommand/AllowedNSSAI): %d", a.AllowedNSSAI.Len)
			}
			a.AllowedNSSAI.SetLen(a.AllowedNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AllowedNSSAI.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/AllowedNSSAI): %w", err)
			}
		case ConfigurationUpdateCommandServiceAreaListType:
			a.ServiceAreaList = nasType.NewServiceAreaList(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ServiceAreaList.Len); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/ServiceAreaList): %w", err)
			}
			if a.ServiceAreaList.Len < 4 || a.ServiceAreaList.Len > 112 {
				return fmt.Errorf("invalid ie length (ConfigurationUpdateCommand/ServiceAreaList): %d", a.ServiceAreaList.Len)
			}
			a.ServiceAreaList.SetLen(a.ServiceAreaList.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ServiceAreaList.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/ServiceAreaList): %w", err)
			}
		case ConfigurationUpdateCommandFullNameForNetworkType:
			a.FullNameForNetwork = nasType.NewFullNameForNetwork(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.FullNameForNetwork.Len); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/FullNameForNetwork): %w", err)
			}
			if a.FullNameForNetwork.Len < 1 {
				return fmt.Errorf("invalid ie length (ConfigurationUpdateCommand/FullNameForNetwork): %d", a.FullNameForNetwork.Len)
			}
			a.FullNameForNetwork.SetLen(a.FullNameForNetwork.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.FullNameForNetwork.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/FullNameForNetwork): %w", err)
			}
		case ConfigurationUpdateCommandShortNameForNetworkType:
			a.ShortNameForNetwork = nasType.NewShortNameForNetwork(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ShortNameForNetwork.Len); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/ShortNameForNetwork): %w", err)
			}
			if a.ShortNameForNetwork.Len < 1 {
				return fmt.Errorf("invalid ie length (ConfigurationUpdateCommand/ShortNameForNetwork): %d", a.ShortNameForNetwork.Len)
			}
			a.ShortNameForNetwork.SetLen(a.ShortNameForNetwork.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ShortNameForNetwork.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/ShortNameForNetwork): %w", err)
			}
		case ConfigurationUpdateCommandLocalTimeZoneType:
			a.LocalTimeZone = nasType.NewLocalTimeZone(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.LocalTimeZone.Octet); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/LocalTimeZone): %w", err)
			}
		case ConfigurationUpdateCommandUniversalTimeAndLocalTimeZoneType:
			a.UniversalTimeAndLocalTimeZone = nasType.NewUniversalTimeAndLocalTimeZone(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, a.UniversalTimeAndLocalTimeZone.Octet[:]); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/UniversalTimeAndLocalTimeZone): %w", err)
			}
		case ConfigurationUpdateCommandNetworkDaylightSavingTimeType:
			a.NetworkDaylightSavingTime = nasType.NewNetworkDaylightSavingTime(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.NetworkDaylightSavingTime.Len); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/NetworkDaylightSavingTime): %w", err)
			}
			if a.NetworkDaylightSavingTime.Len != 1 {
				return fmt.Errorf("invalid ie length (ConfigurationUpdateCommand/NetworkDaylightSavingTime): %d", a.NetworkDaylightSavingTime.Len)
			}
			a.NetworkDaylightSavingTime.SetLen(a.NetworkDaylightSavingTime.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, &a.NetworkDaylightSavingTime.Octet); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/NetworkDaylightSavingTime): %w", err)
			}
		case ConfigurationUpdateCommandLADNInformationType:
			a.LADNInformation = nasType.NewLADNInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.LADNInformation.Len); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/LADNInformation): %w", err)
			}
			if a.LADNInformation.Len > 1712 {
				return fmt.Errorf("invalid ie length (ConfigurationUpdateCommand/LADNInformation): %d", a.LADNInformation.Len)
			}
			a.LADNInformation.SetLen(a.LADNInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.LADNInformation.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/LADNInformation): %w", err)
			}
		case ConfigurationUpdateCommandMICOIndicationType:
			a.MICOIndication = nasType.NewMICOIndication(ieiN)
			a.MICOIndication.Octet = ieiN
		case ConfigurationUpdateCommandNetworkSlicingIndicationType:
			a.NetworkSlicingIndication = nasType.NewNetworkSlicingIndication(ieiN)
			a.NetworkSlicingIndication.Octet = ieiN
		case ConfigurationUpdateCommandConfiguredNSSAIType:
			a.ConfiguredNSSAI = nasType.NewConfiguredNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ConfiguredNSSAI.Len); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/ConfiguredNSSAI): %w", err)
			}
			if a.ConfiguredNSSAI.Len < 2 {
				return fmt.Errorf("invalid ie length (ConfigurationUpdateCommand/ConfiguredNSSAI): %d", a.ConfiguredNSSAI.Len)
			}
			a.ConfiguredNSSAI.SetLen(a.ConfiguredNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ConfiguredNSSAI.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/ConfiguredNSSAI): %w", err)
			}
		case ConfigurationUpdateCommandRejectedNSSAIType:
			a.RejectedNSSAI = nasType.NewRejectedNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RejectedNSSAI.Len); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/RejectedNSSAI): %w", err)
			}
			if a.RejectedNSSAI.Len < 2 || a.RejectedNSSAI.Len > 40 {
				return fmt.Errorf("invalid ie length (ConfigurationUpdateCommand/RejectedNSSAI): %d", a.RejectedNSSAI.Len)
			}
			a.RejectedNSSAI.SetLen(a.RejectedNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.RejectedNSSAI.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/RejectedNSSAI): %w", err)
			}
		case ConfigurationUpdateCommandOperatordefinedAccessCategoryDefinitionsType:
			a.OperatordefinedAccessCategoryDefinitions = nasType.NewOperatordefinedAccessCategoryDefinitions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.OperatordefinedAccessCategoryDefinitions.Len); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/OperatordefinedAccessCategoryDefinitions): %w", err)
			}
			a.OperatordefinedAccessCategoryDefinitions.SetLen(a.OperatordefinedAccessCategoryDefinitions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.OperatordefinedAccessCategoryDefinitions.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ConfigurationUpdateCommand/OperatordefinedAccessCategoryDefinitions): %w", err)
			}
		case ConfigurationUpdateCommandSMSIndicationType:
			a.SMSIndication = nasType.NewSMSIndication(ieiN)
			a.SMSIndication.Octet = ieiN
		default:
		}
	}
	return nil
}
