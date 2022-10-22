package omnilock

import (
	"bytes"
	"encoding/binary"
	"fmt"
	addr "github.com/nervosnetwork/ckb-sdk-go/address"
)

type AuthFlag byte

const (
	AuthFlagCKBSecp256k1Blake160 AuthFlag = 0x0
	AuthFlagEthereum             AuthFlag = 0x1
	AuthFlagEOS                  AuthFlag = 0x2
	AuthFlagTRON                 AuthFlag = 0x3
	AuthFlagBitcoin              AuthFlag = 0x4
	AuthFlagDogcoin              AuthFlag = 0x5
	AuthFlagCKBMultiSig          AuthFlag = 0x6
	AuthFlagLockScriptHash       AuthFlag = 0xFC
	AuthFlagExec                 AuthFlag = 0xFD
	AuthFlagDynamicLinking       AuthFlag = 0xFE
)

type Authentication struct {
	Flag        AuthFlag
	AuthContent [20]byte
}

func DecodeToAuthentication(in []byte) (*Authentication, error) {
	if len(in) < 21 {
		return nil, fmt.Errorf("byte array at least should be 21 bytes")
	}
	authentication := new(Authentication)
	authentication.Flag = AuthFlag(in[0])
	copy(authentication.AuthContent[:], in[1:21])
	return authentication, nil
}

func (a Authentication) Encode() []byte {
	return append([]byte{byte(a.Flag)}, a.AuthContent[:]...)
}

type OmniConfig struct {
	Flag                     byte
	AdminListCellTypeId      [32]byte
	MinimumCKBExponentInAcp  byte
	MinimumSUDTExponentInAcp byte
	SinceForTimeLock         uint8
	TypeScriptHashForSupply  [32]byte
}

func (o OmniConfig) IsAdminModeEnabled() bool {
	return (o.Flag & 0b1) != 0
}

func (o OmniConfig) IsAnyoneCanPayModeEnabled() bool {
	return (o.Flag & 0b10) != 0
}

func (o OmniConfig) IsTimeLockModeEnabled() bool {
	return (o.Flag & 0b100) != 0
}

func (o OmniConfig) IsSupplyModeEnabled() bool {
	return (o.Flag & 0b1000) != 0
}

func DecodeToOmniConfig(in []byte) (*OmniConfig, error) {
	inLength := len(in)
	omniConfig := new(OmniConfig)
	omniConfig.Flag = in[0]
	if omniConfig.IsAdminModeEnabled() {
		if inLength < 33 {
			return nil, fmt.Errorf("byte array at least should be 33 bytes when admin mode is enabled")
		}
		copy(omniConfig.AdminListCellTypeId[:], in[1:33])
	}
	if omniConfig.IsAnyoneCanPayModeEnabled() {
		if inLength < 34 {
			return nil, fmt.Errorf("byte array at least should be 34 bytes when ACP mode is enabled")
		}
		omniConfig.MinimumCKBExponentInAcp = in[33]
		// It allows there is only minimumCKBExponentInAcp but not minimumSUDTExponentInAcp
		if len(in) >= 35 {
			omniConfig.MinimumSUDTExponentInAcp = in[34]
		}
	}
	if omniConfig.IsTimeLockModeEnabled() {
		if inLength < 43 {
			return nil, fmt.Errorf("byte array at least should be 43 bytes when time-lock mode is enabled")
		}
		buf := bytes.NewBuffer(in[35:43])
		binary.Read(buf, binary.BigEndian, &omniConfig.SinceForTimeLock)
	}
	if omniConfig.IsSupplyModeEnabled() {
		if inLength < 75 {
			return nil, fmt.Errorf("byte array at least should be 75 bytes when supply mode is enabled")
		}
		copy(omniConfig.TypeScriptHashForSupply[:], in[43:75])
	}
	return omniConfig, nil
}

func (o OmniConfig) Encode() []byte {
	var out []byte
	out = append(out, o.Flag)
	if o.IsAdminModeEnabled() {
		out = append(out, o.AdminListCellTypeId[:]...)
	} else {
		return out
	}
	if o.IsAnyoneCanPayModeEnabled() {
		out = append(out, o.MinimumCKBExponentInAcp)
		out = append(out, o.MinimumSUDTExponentInAcp)
	} else {
		return out
	}
	if o.IsTimeLockModeEnabled() {
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.BigEndian, o.SinceForTimeLock)
		out = append(out, buf.Bytes()...)
	} else {
		return out
	}
	if o.IsSupplyModeEnabled() {
		out = append(out, o.TypeScriptHashForSupply[:]...)
	}
	return out
}

type OmnilockArgs struct {
	Authentication *Authentication
	OmniConfig     *OmniConfig
}

func NewOmnilockArgsFromAddress(address string) (*OmnilockArgs, error) {
	a, err := addr.Decode(address)
	if err != nil {
		return nil, err
	}
	return NewOmnilockArgsFromAgrs(a.Script.Args)
}

func NewOmnilockArgsFromAgrs(args []byte) (*OmnilockArgs, error) {
	if len(args) < 22 {
		return nil, fmt.Errorf("byte array at least should be 22 bytes when supply mode is enabled")
	}
	authentication, err := DecodeToAuthentication(args[:21])
	if err != nil {
		return nil, err
	}
	omniConfig, err := DecodeToOmniConfig(args[21:])
	if err != nil {
		return nil, err
	}
	return &OmnilockArgs{
		Authentication: authentication,
		OmniConfig:     omniConfig,
	}, nil
}

func (o OmnilockArgs) Encode() []byte {
	out := o.Authentication.Encode()
	out = append(out, o.OmniConfig.Encode()...)
	return out
}
