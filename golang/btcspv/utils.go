package btcspv

import (
	"encoding/hex"

	"github.com/btcsuite/btcutil/base58"
	"github.com/btcsuite/btcutil/bech32"
)

func strip0xPrefix(s string) string {
	if len(s) < 2 {
		return s
	}
	if s[0:2] == "0x" {
		return s[2:]
	}
	return s
}

// DecodeIfHex decodes a hex string into a byte array
func DecodeIfHex(s string) []byte {
	res, err := hex.DecodeString(strip0xPrefix(s))
	if err != nil {
		return []byte(s)
	}
	return res
}

// GetOutputType returns the name of the output type associated with the number
func GetOutputType(outputType OutputType) string {
	var typeString string
	switch outputType {
	case OutputNone:
		typeString = "Output None"
	case WPKH:
		typeString = "WPKH"
	case WSH:
		typeString = "WSH"
	case OpReturn:
		typeString = "Op Return"
	case PKH:
		typeString = "PKH"
	case SH:
		typeString = "SH"
	case Nonstandard:
		typeString = "Nonstandard"
	}
	return typeString
}

// GetInputType returns the name of the input type associated with the number
func GetInputType(inputType InputType) string {
	var typeString string
	switch inputType {
	case InputNone:
		typeString = "Input None"
	case Legacy:
		typeString = "Legacy"
	case Compatibility:
		typeString = "Compatibility"
	case Witness:
		typeString = "Witness"
	}
	return typeString
}

// EncodeP2SH turns a scripthash into an address
func EncodeP2SH(sh []byte) string {
	return base58.CheckEncode(sh, 5)
}

// EncodeP2PKH turns a pubkey hash into an address
func EncodeP2PKH(pkh []byte) string {
	return base58.CheckEncode(pkh, 0)
}

func encodeSegWit(payload []byte, version int) string {
	adj, err := bech32.ConvertBits(payload, 8, 5, true)
	if err != nil {
		return ""
	}
	combined := []byte{0x00}
	combined = append(combined, adj...)
	res, _ := bech32.Encode("bc", combined)
	return res
}

// EncodeP2WSH turns a scripthash into an address
func EncodeP2WSH(sh []byte) string {
	return encodeSegWit(sh, 0)
}

// EncodeP2WPKH turns a pubkey hash into an address
func EncodeP2WPKH(pkh []byte) string {
	return encodeSegWit(pkh, 0)
}