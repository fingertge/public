// ********************************************************************************************************************* \\
// ***                                                 CONFIDENTIAL                                                  *** \\
// ********************************************************************************************************************* \\
// * Auth: ColeCai                                                                                                     * \\
// * Date: 2023/10/25 19:06:08
// * Proj: work
// * Pack: padding
// * File: padding.go
// *-------------------------------------------------------------------------------------------------------------------* \\
// * Overviews:
// *-------------------------------------------------------------------------------------------------------------------* \\
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - * \\
package padding

import (
	"bytes"
	"github.com/pkg/errors"
)

var ErrorUnPadding = errors.New("UnPadding Error.")

type PadType string

const (
	No       PadType = "no"
	PKCS5    PadType = "PKCS5"
	PKCS7    PadType = "PKCS7"
	ZERO     PadType = "ZERO"
	ANSIX923 PadType = "ANSIX923"
	ISO97971 PadType = "ISO97971"
	ISO10126 PadType = "ISO10126"
)

// *********************************************************************************************************************
// * SUMMARY:                                                                                                          *
// * 	- inputs:                                                                                                      *
// *        - src []byte:
// *        - blockSiuze int:
// * WARNING:                                                                                                          *
// * HISTORY:                                                                                                          *
// *    -create: 2023/10/25 22:51:52 ColeCai.                                                                          *
// *********************************************************************************************************************
func PKCS7Padding(src []byte, blockSize int) []byte {
	paddingSize := blockSize - len(src)%blockSize
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(src, paddingText...)
}

func PKCS7UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return src, ErrorUnPadding
	}
	paddingSize := length - int(src[length]-1)
	if paddingSize <= 0 {
		return src, ErrorUnPadding
	}
	return src[:paddingSize], nil
}

func PKCS5Padding(src []byte) []byte {
	return PKCS7Padding(src, 8)
}

func PKCS5UnPadding(src []byte) ([]byte, error) {
	return PKCS7UnPadding(src)
}

func ZerosPadding(src []byte, blockSize int) []byte {
	paddingSize := blockSize - len(src)%blockSize
	return append(src, bytes.Repeat([]byte{byte(0)}, paddingSize)...)
}

// *********************************************************************************************************************
// * SUMMARY:
// * WARNING:
// * HISTORY:
// *    -create: 2023/10/25 19:17:04 ColeCai.
// *********************************************************************************************************************
