// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/09 16:26:13
// * Proj: public
// * Pack: crypter
// * File: cryptermgr.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

package crypter

import (
	"github.com/pkg/errors"
)

type CrypterType byte

type ICrypter interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
	GetCryptId() CrypterType
}

var (
	defaultCrypterMgr *CrypterMgr
)

const (
	NoneCrypter CrypterType = 0
)

func init() {
	defaultCrypterMgr = NewCrypterMgr()
}

type CrypterMgr struct {
	crypterMap map[CrypterType]ICrypter
}

func NewCrypterMgr() *CrypterMgr {
	mgr := &CrypterMgr{
		crypterMap: make(map[CrypterType]ICrypter),
	}
	return mgr
}

func RegisterCrypter(key CrypterType, crypter ICrypter) error {
	return defaultCrypterMgr.RegisterCrypter(key, crypter)
}

func Encrypt(cryptId CrypterType, data []byte) ([]byte, error) {
	return defaultCrypterMgr.Encrypt(cryptId, data)
}

func Decrypt(cryptId CrypterType, data []byte) ([]byte, error) {
	return defaultCrypterMgr.Decrypt(cryptId, data)
}

func (c *CrypterMgr) RegisterCrypter(key CrypterType, crypter ICrypter) error {
	if _, ok := c.crypterMap[key]; ok {
		return errors.New("key already exists")
	}
	if crypter == nil || crypter.GetCryptId() == NoneCrypter {
		return errors.New("encryper id already use")
	}
	c.crypterMap[key] = crypter
	return nil
}

func (c *CrypterMgr) GetCrypter(cryptId CrypterType) (ICrypter, error) {
	if crypter, ok := c.crypterMap[cryptId]; ok {
		return crypter, nil
	}
	return nil, errors.Errorf("crypter id %d not found", cryptId)
}

func (c *CrypterMgr) Encrypt(cryptId CrypterType, data []byte) ([]byte, error) {
	if cryptId == NoneCrypter {
		return data, nil
	}
	crypter, err := c.GetCrypter(cryptId)
	if err != nil {
		return nil, err
	}
	return crypter.Encrypt(data)
}

func (c *CrypterMgr) Decrypt(cryptId CrypterType, data []byte) ([]byte, error) {
	if cryptId == NoneCrypter {
		return data, nil
	}
	crypter, err := c.GetCrypter(cryptId)
	if err != nil {
		return nil, err
	}
	return crypter.Decrypt(data)
}
