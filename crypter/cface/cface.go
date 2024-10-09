// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/25 18:56:40
// * Proj: work
// * Pack: cface
// * File: cface.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
package cface

type CrypterType byte

type ICrypter interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
	GetCryptId() CrypterType
}

type IMode interface {
	Encode([]byte) ([]byte, error)
	Decode([]byte) ([]byte, error)
}
