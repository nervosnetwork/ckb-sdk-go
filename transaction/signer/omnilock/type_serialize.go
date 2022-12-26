package omnilock

import (
	"github.com/nervosnetwork/ckb-sdk-go/v2/transaction/signer/omnilock/molecule"
)

func DeserializeOmnilockWitnessLock(in []byte) (*OmnilockWitnessLock, error) {
	m, err := molecule.OmniLockWitnessLockFromSlice(in, false)
	if err != nil {
		return nil, err
	}
	return UnpackOmnilockWitnessLock(m), nil
}

func UnpackSmtProofEntry(v *molecule.SmtProofEntry) *SmtProofEntry {
	return &SmtProofEntry{
		Mask:     v.Mask().AsSlice()[0],
		SmtProof: v.Mask().AsSlice(),
	}
}

func UnpackIdentityOpt(v *molecule.IdentityOpt) *OmnilockIdentity {
	if v.IsNone() {
		return nil
	}
	var smtProofEntryVec []*SmtProofEntry
	mIdentity, _ := v.IntoIdentity()
	mSmtProofEntryVec := mIdentity.Proofs()
	for i := 0; i < int(mSmtProofEntryVec.Len()); i++ {
		smtProofEntryVec = append(smtProofEntryVec, UnpackSmtProofEntry(mSmtProofEntryVec.Get(uint(i))))
	}
	return &OmnilockIdentity{
		Identity: UnpackAuth(mIdentity.Identity()),
		Proofs:   smtProofEntryVec,
	}
}

func UnpackAuth(v *molecule.Auth) *Auth {
	b := v.AsSlice()
	return &Auth{
		Flag:        OmnilockFlag(b[0]),
		AuthContent: b[1:],
	}
}

func UnpackOmnilockWitnessLock(v *molecule.OmniLockWitnessLock) *OmnilockWitnessLock {
	o := &OmnilockWitnessLock{
		OmnilockIdentity: UnpackIdentityOpt(v.OmniIdentity()),
	}
	if v.Signature().IsSome() {
		b, _ := v.Signature().IntoBytes()
		o.Signature = b.RawData()
	}
	if v.Preimage().IsSome() {
		b, _ := v.Preimage().IntoBytes()
		o.Preimage = b.RawData()
	}
	return o
}

func (o *OmnilockWitnessLock) Serialize() []byte {
	return o.Pack().AsSlice()
}

func (o *OmnilockWitnessLock) SerializeAsPlaceholder() []byte {
	return make([]byte, len(o.Pack().AsSlice()))
}

func (o *SmtProofEntry) Pack() *molecule.SmtProofEntry {
	proofBuilder := molecule.NewSmtProofBuilder()
	for _, vv := range o.SmtProof {
		proofBuilder.Push(*packByte(vv))
	}

	builder := molecule.NewSmtProofEntryBuilder()
	builder.Mask(*packByte(o.Mask))
	builder.Proof(proofBuilder.Build())

	b := builder.Build()
	return &b
}

func (o *OmnilockIdentity) PackOpt() *molecule.IdentityOpt {
	builder := molecule.NewIdentityOptBuilder()
	if o != nil {
		builder.Set(*o.Pack())
	}
	b := builder.Build()
	return &b
}

func (o *OmnilockIdentity) Pack() *molecule.Identity {
	builder := molecule.NewIdentityBuilder()
	proofsBuilder := molecule.NewSmtProofEntryVecBuilder()
	for _, p := range o.Proofs {
		proofsBuilder.Push(*p.Pack())
	}
	builder.Proofs(proofsBuilder.Build())
	b := builder.Build()
	return &b
}

func (o *OmnilockWitnessLock) Pack() *molecule.OmniLockWitnessLock {
	builder := molecule.NewOmniLockWitnessLockBuilder()
	builder.Signature(*packBytesToOpt(o.Signature))
	builder.OmniIdentity(*o.OmnilockIdentity.PackOpt())
	builder.Preimage(*packBytesToOpt(o.Preimage))
	v := builder.Build()
	return &v
}

func packBytesToOpt(v []byte) *molecule.BytesOpt {
	builder := molecule.NewBytesOptBuilder()
	if v != nil {
		builder.Set(*packBytes(v))
	}
	b := builder.Build()
	return &b
}

func packByte(v byte) *molecule.Byte {
	b := molecule.NewByte(v)
	return &b
}

func packBytes(v []byte) *molecule.Bytes {
	builder := molecule.NewBytesBuilder()
	for _, vv := range v {
		builder.Push(*packByte(vv))
	}
	b := builder.Build()
	return &b
}
