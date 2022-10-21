package omnilock

type SmtProofEntry struct {
	Mask     byte
	SmtProof []byte
}

type OmnilockFlag byte

const (
	OmnilockFlagCKBSecp256k1Blake160 OmnilockFlag = 0x0
	OmnilockFlagLockScriptHash       OmnilockFlag = 0xfc
)

type Auth struct {
	Flag        OmnilockFlag
	AuthContent []byte
}

func (a Auth) encode() []byte {
	var out []byte
	out = []byte{byte(a.Flag)}
	out = append(out, a.AuthContent...)
	return out
}

type OmnilockIdentity struct {
	Identity *Auth
	Proofs   []*SmtProofEntry
}

type OmnilockWitnessLock struct {
	Signature        []byte
	OmnilockIdentity *OmnilockIdentity
	Preimage         []byte
}
