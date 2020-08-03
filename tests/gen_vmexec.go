// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ccm-chain/ccmchain/common"
	"github.com/ccm-chain/ccmchain/common/hexutil"
	"github.com/ccm-chain/ccmchain/common/math"
)

var _ = (*vmExecMarshaling)(nil)

func (v vmExec) MarshalJSON() ([]byte, error) {
	type vmExec struct {
		Address  common.UnprefixedAddress `json:"address"  gencodec:"required"`
		Caller   common.UnprefixedAddress `json:"caller"   gencodec:"required"`
		Origin   common.UnprefixedAddress `json:"origin"   gencodec:"required"`
		Code     hexutil.Bytes            `json:"code"     gencodec:"required"`
		Data     hexutil.Bytes            `json:"data"     gencodec:"required"`
		Value    *math.HexOrDecimal256    `json:"value"    gencodec:"required"`
		GasLimit math.HexOrDecimal64      `json:"gas"      gencodec:"required"`
		GasPrice *math.HexOrDecimal256    `json:"gasPrice" gencodec:"required"`
	}
	var enc vmExec
	enc.Address = common.UnprefixedAddress(v.Address)
	enc.Caller = common.UnprefixedAddress(v.Caller)
	enc.Origin = common.UnprefixedAddress(v.Origin)
	enc.Code = v.Code
	enc.Data = v.Data
	enc.Value = (*math.HexOrDecimal256)(v.Value)
	enc.GasLimit = math.HexOrDecimal64(v.GasLimit)
	enc.GasPrice = (*math.HexOrDecimal256)(v.GasPrice)
	return json.Marshal(&enc)
}

func (v *vmExec) UnmarshalJSON(input []byte) error {
	type vmExec struct {
		Address  *common.UnprefixedAddress `json:"address"  gencodec:"required"`
		Caller   *common.UnprefixedAddress `json:"caller"   gencodec:"required"`
		Origin   *common.UnprefixedAddress `json:"origin"   gencodec:"required"`
		Code     *hexutil.Bytes            `json:"code"     gencodec:"required"`
		Data     *hexutil.Bytes            `json:"data"     gencodec:"required"`
		Value    *math.HexOrDecimal256     `json:"value"    gencodec:"required"`
		GasLimit *math.HexOrDecimal64      `json:"gas"      gencodec:"required"`
		GasPrice *math.HexOrDecimal256     `json:"gasPrice" gencodec:"required"`
	}
	var dec vmExec
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Address == nil {
		return errors.New("missing required field 'address' for vmExec")
	}
	v.Address = common.Address(*dec.Address)
	if dec.Caller == nil {
		return errors.New("missing required field 'caller' for vmExec")
	}
	v.Caller = common.Address(*dec.Caller)
	if dec.Origin == nil {
		return errors.New("missing required field 'origin' for vmExec")
	}
	v.Origin = common.Address(*dec.Origin)
	if dec.Code == nil {
		return errors.New("missing required field 'code' for vmExec")
	}
	v.Code = *dec.Code
	if dec.Data == nil {
		return errors.New("missing required field 'data' for vmExec")
	}
	v.Data = *dec.Data
	if dec.Value == nil {
		return errors.New("missing required field 'value' for vmExec")
	}
	v.Value = (*big.Int)(dec.Value)
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gas' for vmExec")
	}
	v.GasLimit = uint64(*dec.GasLimit)
	if dec.GasPrice == nil {
		return errors.New("missing required field 'gasPrice' for vmExec")
	}
	v.GasPrice = (*big.Int)(dec.GasPrice)
	return nil
}
