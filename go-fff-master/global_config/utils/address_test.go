package utils

import (
	"log"
	"testing"
)

func IsHexAddress(s string) bool {
	s = FFFAddressDecode(s)
	if has0xPrefix(s) {
		s = s[2:]
	}
	return len(s) == 2*20 && isHex(s)
}

// isHex validates whether each byte is valid hexadecimal string.
func isHex(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	for _, c := range []byte(str) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}

// isHexCharacter returns bool of c being a valid hexadecimal.
func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

// has0xPrefix validates str begins with '0x' or '0X'.
func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func TestFFFAddressEncode(t *testing.T) {

	//fmt.Println(IsHexAddress("FFF3Psbq3enwAmwXGa2QejWFdd9AwV1rczE6w1GPzs6WTPmJpKbmWghsLB"))
	//log.Println(FFFAddressEncode("0x43b7e972ce83812b30582E630C3D36B5889946f6"))

	//log.Println(FFFAddressDecode("FFF65YFc3TpV2FCSyJonjq5X1wPKWMpaKT2ve8jYenMP9Tf2usooGuDJu5"))

	log.Println(FFFAddressEncode("0xd2c5301f9f00aa5fa9e717b71fe95c8c3175f77f"))
	//log.Println(FFFAddressEncode("0x43b7e972ce83812b30582E630C3D36B5889946f6"))
	/*log.Println(FFFAddressEncode("0x7F37b86cdC053D70Fa5daFEbd7CC6FD9Df8afbeb"))
	log.Println(FFFAddressEncode("0x3dd9579D49f8ba528eD1ccab27A1Af5E1c6D91ab"))
	log.Println(FFFAddressEncode("0xA0bbA26b26ef4366Bb8515c7eaC3a9e5fA3F16Ed"))
	log.Println(FFFAddressEncode("0xc891CAbB5faBf3c84e45880cf683d126c4cDDC5e"))*/

	//0x7C610f6B8D468F13799c9380019eb01521507fa5, 0x7F37b86cdC053D70Fa5daFEbd7CC6FD9Df8afbeb, 0x3dd9579D49f8ba528eD1ccab27A1Af5E1c6D91ab, 0xA0bbA26b26ef4366Bb8515c7eaC3a9e5fA3F16Ed, 0xc891CAbB5faBf3c84e45880cf683d126c4cDDC5e
	//log.Println(FFFAddressDecode("FFF5ykma29oAcKrrXmQwJuaW2bfEynuaupEDzcsL6jn1JYXESVStzjzSa1"))
	/*log.Println(FFFAddressDecode("FFF5qfy2gzm4f1zu5xXGJ12vQx1hPwp8jJdF1Gyz47vXd54uCESRuhpndG"))
	log.Println(FFFAddressDecode("FFF3bSAZ9s81bFA8bcUdSQsvh8QagAStmG3swqhGUQqAh1P4XuyWCfNVMc"))
	log.Println(FFFAddressDecode("FFF3c1ars8vMkNE9ZZH91TYQrRKsMdBLParaSM745tzjJnQguSqaV8MxFk"))
	log.Println(FFFAddressDecode("FFF3PwYP6ZhYBKC5vEcELTLqMLvStPbS9N2tQHFFNk7RtJ6CQq2dNbmJAe"))

	log.Println(FFFAddressDecode("FFF3Psbq3enwAmwXGa2QejWFdd9AwV1rczE6w1GPzs6WTPmJpKbmWghsLB"))
	log.Println(FFFAddressDecode("FFF6672WbdorrmkMpavk1S5ALpoN82XpSirbMWZicxhhqqNeromt65d6TE"))*/

}

/*
{
    consensusAddr: "0x4652729e079a594E9DC6e72320ADb33ffbf13428",
    feeAddr: "0x4652729e079a594E9DC6e72320ADb33ffbf13428",
    bscFeeAddr: "0x4652729e079a594E9DC6e72320ADb33ffbf13428",
    votingPower: 0x0000000000000064
  },
  {
    consensusAddr: "0x6BE6DB3263630384F18b8825857607825843E256",
    feeAddr: "0x6BE6DB3263630384F18b8825857607825843E256",
    bscFeeAddr: "0x6BE6DB3263630384F18b8825857607825843E256",
    votingPower: 0x0000000000000064
  }, {
    consensusAddr: "0x89F77B09EE3232a2e4f91E3ffffF65150E9B89b0",
    feeAddr: "0x89F77B09EE3232a2e4f91E3ffffF65150E9B89b0",
    bscFeeAddr: "0x89F77B09EE3232a2e4f91E3ffffF65150E9B89b0",
    votingPower: 0x0000000000000064
  }, {
    consensusAddr: "0xf89c95f333ab17E584424f2104BC34aF4c3D8276",
    feeAddr: "0xf89c95f333ab17E584424f2104BC34aF4c3D8276",
    bscFeeAddr: "0xf89c95f333ab17E584424f2104BC34aF4c3D8276",
    votingPower: 0x0000000000000064
  }, {
    consensusAddr: "0x1BC78b9843544285848FCadCb2CF196dc3136B40",
    feeAddr: "0x1BC78b9843544285848FCadCb2CF196dc3136B40",
    bscFeeAddr: "0x1BC78b9843544285848FCadCb2CF196dc3136B40",
    votingPower: 0x0000000000000064
  },
*/
