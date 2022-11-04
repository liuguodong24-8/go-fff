package pkgs

import (
	"bytes"
	"math/big"
)

type HexInt int

func (i *HexInt) UnmarshalJSON(data []byte) error {
	result, err := ParseInt(string(bytes.Trim(data, `"`)))
	*i = HexInt(result)
	return err
}

type HexBig big.Int

func (i *HexBig) UnmarshalJSON(data []byte) error {
	result, err := ParseBigInt(string(bytes.Trim(data, `"`)))
	*i = HexBig(result)
	return err
}

type HexBigInt int64

func (i *HexBigInt) UnmarshalJSON(data []byte) error {
	result, err := ParseBigInt(string(bytes.Trim(data, `"`)))
	*i = HexBigInt(result.Int64())
	return err
}
