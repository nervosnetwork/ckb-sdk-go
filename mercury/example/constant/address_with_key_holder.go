package constant

const (
	TEST_ADDRESS0 = "ckt1qyq28wze3cw48ek9az0g4jmtfs6d8td38u4s6hp2s0"
	TEST_ADDRESS1 = "ckt1qyq27z6pccncqlaamnh8ttapwn260egnt67ss2cwvz"
	TEST_ADDRESS2 = "ckt1qyqqtg06h75ymw098r3w0l3u4xklsj04tnsqctqrmc"
	TEST_ADDRESS3 = "ckt1qyqzqfj8lmx9h8vvhk62uut8us844v0yh2hsnqvvgc"
	TEST_ADDRESS4 = "ckt1qyqg88ccqm59ksxp85788pnqg4rkejdgcg2qxcu2qf"
	TEST_KEY0     = "6fc935dad260867c749cf1ba6602d5f5ed7fb1131f1beb65be2d342e912eaafe"
	TEST_KEY1     = "9d8ca87d75d150692211fa62b0d30de4d1ee6c530d5678b40b8cedacf0750d0f"
	TEST_KEY2     = "88a09e06735d89452552e359a052315ab5130dc2e4d864ae3eed21d6505b2f67"
	TEST_KEY3     = "2d4cf0546a1dc93092ad56f2e18fbe6e41ee477d9dec0575cf43b69740ce9f74"
	TEST_KEY4     = "5e46fdbb6ffd86d232080dc71f24b60df2a119e0102ca45a7c165472de14c104"
	CEX_ADDRESS   = "ckt1qyqg03ul48cpvd3wzlqu2t5qpe80hdv3nqpq4hswge"
	CEX_KEY       = "6d88a2eab95e8546ee9b33160e941837625a40c77202cef35d9e3a1ae6f4edf1"
)

func GetKey(address string) string {
	switch address {
	case TEST_ADDRESS0:
		return TEST_KEY0
	case TEST_ADDRESS1:
		return TEST_KEY1
	case TEST_ADDRESS2:
		return TEST_KEY2
	case TEST_ADDRESS3:
		return TEST_KEY3
	case TEST_ADDRESS4:
		return TEST_KEY4
	case CEX_ADDRESS:
		return CEX_KEY
	default:
		return ""
	}
}
