// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/09 17:25:42
// * Proj: public
// * Pack: tools
// * File: bytepool.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

package tools

import "sync"

type BytePool struct {
	p sync.Pool
}

func NewBytePool(size, cap int) *BytePool {
	if size > cap {
		panic("size must be less then cap")
	}
	p := &BytePool{}
	p.p.New = func() interface{} {
		return make([]byte, size, cap)
	}
	return p
}

func (p *BytePool) Get() []byte {
	return p.p.Get().([]byte)
}

func (p *BytePool) Put(b interface{}) {

	p.p.Put(b)
}

func (p *BytePool) PutWithReset(b []byte) {
	// 重置已用大小
	b = b[:0]
	p.p.Put(b)
}
