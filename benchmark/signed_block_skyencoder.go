// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package benchmark

import (
	"errors"
	"math"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/encoder"
	"github.com/skycoin/skycoin/src/coin"
)

// EncodeSizeSignedBlock computes the size of an encoded object of type SignedBlock
func EncodeSizeSignedBlock(obj *coin.SignedBlock) int {
	i0 := 0

	// obj.Block.Head.Version
	i0 += 4

	// obj.Block.Head.Time
	i0 += 8

	// obj.Block.Head.BkSeq
	i0 += 8

	// obj.Block.Head.Fee
	i0 += 8

	// obj.Block.Head.PrevHash
	i0 += 32

	// obj.Block.Head.BodyHash
	i0 += 32

	// obj.Block.Head.UxHash
	i0 += 32

	// obj.Block.Body.Transactions
	i0 += 4
	for _, x := range obj.Block.Body.Transactions {
		i1 := 0

		// x.Length
		i1 += 4

		// x.Type
		i1++

		// x.InnerHash
		i1 += 32

		// x.Sigs
		i1 += 4
		{
			i2 := 0

			// x
			i2 += 65

			i1 += len(x.Sigs) * i2
		}

		// x.In
		i1 += 4
		{
			i2 := 0

			// x
			i2 += 32

			i1 += len(x.In) * i2
		}

		// x.Out
		i1 += 4
		{
			i2 := 0

			// x.Address.Version
			i2++

			// x.Address.Key
			i2 += 20

			// x.Coins
			i2 += 8

			// x.Hours
			i2 += 8

			i1 += len(x.Out) * i2
		}

		i0 += i1
	}

	// obj.Sig
	i0 += 65

	return i0
}

// EncodeSignedBlock encodes an object of type SignedBlock to the buffer in encoder.Encoder
func EncodeSignedBlock(e *encoder.Encoder, obj *coin.SignedBlock) error {

	// obj.Block.Head.Version
	e.Uint32(obj.Block.Head.Version)

	// obj.Block.Head.Time
	e.Uint64(obj.Block.Head.Time)

	// obj.Block.Head.BkSeq
	e.Uint64(obj.Block.Head.BkSeq)

	// obj.Block.Head.Fee
	e.Uint64(obj.Block.Head.Fee)

	// obj.Block.Head.PrevHash
	e.CopyBytes(obj.Block.Head.PrevHash[:])

	// obj.Block.Head.BodyHash
	e.CopyBytes(obj.Block.Head.BodyHash[:])

	// obj.Block.Head.UxHash
	e.CopyBytes(obj.Block.Head.UxHash[:])

	// obj.Block.Body.Transactions length check
	if len(obj.Block.Body.Transactions) > math.MaxUint32 {
		return errors.New("obj.Block.Body.Transactions length exceeds math.MaxUint32")
	}

	// obj.Block.Body.Transactions length
	e.Uint32(uint32(len(obj.Block.Body.Transactions)))

	// obj.Block.Body.Transactions
	for _, x := range obj.Block.Body.Transactions {

		// x.Length
		e.Uint32(x.Length)

		// x.Type
		e.Uint8(x.Type)

		// x.InnerHash
		e.CopyBytes(x.InnerHash[:])

		// x.Sigs length check
		if len(x.Sigs) > math.MaxUint32 {
			return errors.New("x.Sigs length exceeds math.MaxUint32")
		}

		// x.Sigs length
		e.Uint32(uint32(len(x.Sigs)))

		// x.Sigs
		for _, x := range x.Sigs {

			// x
			e.CopyBytes(x[:])

		}

		// x.In length check
		if len(x.In) > math.MaxUint32 {
			return errors.New("x.In length exceeds math.MaxUint32")
		}

		// x.In length
		e.Uint32(uint32(len(x.In)))

		// x.In
		for _, x := range x.In {

			// x
			e.CopyBytes(x[:])

		}

		// x.Out length check
		if len(x.Out) > math.MaxUint32 {
			return errors.New("x.Out length exceeds math.MaxUint32")
		}

		// x.Out length
		e.Uint32(uint32(len(x.Out)))

		// x.Out
		for _, x := range x.Out {

			// x.Address.Version
			e.Uint8(x.Address.Version)

			// x.Address.Key
			e.CopyBytes(x.Address.Key[:])

			// x.Coins
			e.Uint64(x.Coins)

			// x.Hours
			e.Uint64(x.Hours)

		}

	}

	// obj.Sig
	e.CopyBytes(obj.Sig[:])

	return nil
}

// DecodeSignedBlock decodes an object of type SignedBlock from the buffer in encoder.Decoder
func DecodeSignedBlock(d *encoder.Decoder, obj *coin.SignedBlock) error {
	{
		// obj.Block.Head.Version
		i, err := d.Uint32()
		if err != nil {
			return err
		}
		obj.Block.Head.Version = i
	}

	{
		// obj.Block.Head.Time
		i, err := d.Uint64()
		if err != nil {
			return err
		}
		obj.Block.Head.Time = i
	}

	{
		// obj.Block.Head.BkSeq
		i, err := d.Uint64()
		if err != nil {
			return err
		}
		obj.Block.Head.BkSeq = i
	}

	{
		// obj.Block.Head.Fee
		i, err := d.Uint64()
		if err != nil {
			return err
		}
		obj.Block.Head.Fee = i
	}

	{
		// obj.Block.Head.PrevHash
		if len(d.Buffer) < len(obj.Block.Head.PrevHash) {
			return encoder.ErrBufferUnderflow
		}
		copy(obj.Block.Head.PrevHash[:], d.Buffer[:len(obj.Block.Head.PrevHash)])
		d.Buffer = d.Buffer[len(obj.Block.Head.PrevHash):]
	}

	{
		// obj.Block.Head.BodyHash
		if len(d.Buffer) < len(obj.Block.Head.BodyHash) {
			return encoder.ErrBufferUnderflow
		}
		copy(obj.Block.Head.BodyHash[:], d.Buffer[:len(obj.Block.Head.BodyHash)])
		d.Buffer = d.Buffer[len(obj.Block.Head.BodyHash):]
	}

	{
		// obj.Block.Head.UxHash
		if len(d.Buffer) < len(obj.Block.Head.UxHash) {
			return encoder.ErrBufferUnderflow
		}
		copy(obj.Block.Head.UxHash[:], d.Buffer[:len(obj.Block.Head.UxHash)])
		d.Buffer = d.Buffer[len(obj.Block.Head.UxHash):]
	}

	{
		// obj.Block.Body.Transactions

		if len(d.Buffer) < 4 {
			return encoder.ErrBufferUnderflow
		}

		ul, err := d.Uint32()
		if err != nil {
			return err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return encoder.ErrBufferUnderflow
		}

		obj.Block.Body.Transactions = make([]coin.Transaction, length)

		for z3 := range obj.Block.Body.Transactions {
			{
				// obj.Block.Body.Transactions[z3].Length
				i, err := d.Uint32()
				if err != nil {
					return err
				}
				obj.Block.Body.Transactions[z3].Length = i
			}

			{
				// obj.Block.Body.Transactions[z3].Type
				i, err := d.Uint8()
				if err != nil {
					return err
				}
				obj.Block.Body.Transactions[z3].Type = i
			}

			{
				// obj.Block.Body.Transactions[z3].InnerHash
				if len(d.Buffer) < len(obj.Block.Body.Transactions[z3].InnerHash) {
					return encoder.ErrBufferUnderflow
				}
				copy(obj.Block.Body.Transactions[z3].InnerHash[:], d.Buffer[:len(obj.Block.Body.Transactions[z3].InnerHash)])
				d.Buffer = d.Buffer[len(obj.Block.Body.Transactions[z3].InnerHash):]
			}

			{
				// obj.Block.Body.Transactions[z3].Sigs

				if len(d.Buffer) < 4 {
					return encoder.ErrBufferUnderflow
				}

				ul, err := d.Uint32()
				if err != nil {
					return err
				}

				length := int(ul)
				if length < 0 || length > len(d.Buffer) {
					return encoder.ErrBufferUnderflow
				}

				obj.Block.Body.Transactions[z3].Sigs = make([]cipher.Sig, length)

				for z5 := range obj.Block.Body.Transactions[z3].Sigs {
					{
						// obj.Block.Body.Transactions[z3].Sigs[z5]
						if len(d.Buffer) < len(obj.Block.Body.Transactions[z3].Sigs[z5]) {
							return encoder.ErrBufferUnderflow
						}
						copy(obj.Block.Body.Transactions[z3].Sigs[z5][:], d.Buffer[:len(obj.Block.Body.Transactions[z3].Sigs[z5])])
						d.Buffer = d.Buffer[len(obj.Block.Body.Transactions[z3].Sigs[z5]):]
					}

				}

			}

			{
				// obj.Block.Body.Transactions[z3].In

				if len(d.Buffer) < 4 {
					return encoder.ErrBufferUnderflow
				}

				ul, err := d.Uint32()
				if err != nil {
					return err
				}

				length := int(ul)
				if length < 0 || length > len(d.Buffer) {
					return encoder.ErrBufferUnderflow
				}

				obj.Block.Body.Transactions[z3].In = make([]cipher.SHA256, length)

				for z5 := range obj.Block.Body.Transactions[z3].In {
					{
						// obj.Block.Body.Transactions[z3].In[z5]
						if len(d.Buffer) < len(obj.Block.Body.Transactions[z3].In[z5]) {
							return encoder.ErrBufferUnderflow
						}
						copy(obj.Block.Body.Transactions[z3].In[z5][:], d.Buffer[:len(obj.Block.Body.Transactions[z3].In[z5])])
						d.Buffer = d.Buffer[len(obj.Block.Body.Transactions[z3].In[z5]):]
					}

				}

			}

			{
				// obj.Block.Body.Transactions[z3].Out

				if len(d.Buffer) < 4 {
					return encoder.ErrBufferUnderflow
				}

				ul, err := d.Uint32()
				if err != nil {
					return err
				}

				length := int(ul)
				if length < 0 || length > len(d.Buffer) {
					return encoder.ErrBufferUnderflow
				}

				obj.Block.Body.Transactions[z3].Out = make([]coin.TransactionOutput, length)

				for z5 := range obj.Block.Body.Transactions[z3].Out {
					{
						// obj.Block.Body.Transactions[z3].Out[z5].Address.Version
						i, err := d.Uint8()
						if err != nil {
							return err
						}
						obj.Block.Body.Transactions[z3].Out[z5].Address.Version = i
					}

					{
						// obj.Block.Body.Transactions[z3].Out[z5].Address.Key
						if len(d.Buffer) < len(obj.Block.Body.Transactions[z3].Out[z5].Address.Key) {
							return encoder.ErrBufferUnderflow
						}
						copy(obj.Block.Body.Transactions[z3].Out[z5].Address.Key[:], d.Buffer[:len(obj.Block.Body.Transactions[z3].Out[z5].Address.Key)])
						d.Buffer = d.Buffer[len(obj.Block.Body.Transactions[z3].Out[z5].Address.Key):]
					}

					{
						// obj.Block.Body.Transactions[z3].Out[z5].Coins
						i, err := d.Uint64()
						if err != nil {
							return err
						}
						obj.Block.Body.Transactions[z3].Out[z5].Coins = i
					}

					{
						// obj.Block.Body.Transactions[z3].Out[z5].Hours
						i, err := d.Uint64()
						if err != nil {
							return err
						}
						obj.Block.Body.Transactions[z3].Out[z5].Hours = i
					}

				}

			}
		}

	}

	{
		// obj.Sig
		if len(d.Buffer) < len(obj.Sig) {
			return encoder.ErrBufferUnderflow
		}
		copy(obj.Sig[:], d.Buffer[:len(obj.Sig)])
		d.Buffer = d.Buffer[len(obj.Sig):]
	}

	if len(d.Buffer) != 0 {
		return encoder.ErrRemainingBytes
	}

	return nil
}
