// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/09 18:25:42
// * Proj: public
// * Pack: logs
// * File: rotator.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

package logs

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Rotator struct {
	pattern  string
	fileName string
	maxSize  uint64
	curSize  uint64
	isRotate bool
	fd       *os.File
	sync.RWMutex
}

type OptionRotator func(*Rotator)

func WithPattern(pattern string) OptionRotator {
	return func(r *Rotator) {
		r.pattern = pattern
	}
}

func WithMaxSize(maxSize uint64) OptionRotator {
	return func(r *Rotator) {
		r.maxSize = maxSize
	}
}

func NewRotator(opts ...OptionRotator) *Rotator {
	r := &Rotator{
		pattern:  "./rotator.log",
		fileName: "logs",
		maxSize:  0,
		curSize:  0,
		isRotate: true,
		fd:       nil,
	}
	for _, opt := range opts {
		opt(r)
	}

	name := r.genFileName()
	if strings.Compare(name, r.pattern) == 0 {
		r.isRotate = false
	}

	if !r.isSplit() {
		fd, _ := r.rotate(r.pattern)
		r.fd = fd
	}

	return r
}

func (r *Rotator) genFileName() string {
	if !r.isRotate {
		return r.pattern
	}
	return time.Now().Format(r.pattern)
}

func (r *Rotator) isSplit() bool {
	return r.maxSize > 0 || r.isRotate
}

func (r *Rotator) rotate(fullName string) (*os.File, error) {
	dir := filepath.Dir(fullName)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, err
	}
	idx := 0
	for {
		if idx >= 1000 {
			return nil, errors.WithStack(fmt.Errorf("over max log file cnt"))
		}
		file := fullName + "_" + strconv.Itoa(idx)
		_, err := os.Stat(file)
		if !os.IsNotExist(err) {
			idx++
			continue
		}
		return os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	}
}

func (r *Rotator) sizeCheck(size uint64) bool {
	if r.maxSize <= 0 {
		return false
	}
	if r.curSize+size > r.maxSize {
		return true
	}
	return false
}

func (r *Rotator) fileCheck(name string) bool {
	return strings.Compare(name, r.fileName) != 0
}

// *********************************************************************************************************************
// * SUMMARY: none.                                                                                                    *
// * WARNING: none.                                                                                                    *
// * HISTORY:                                                                                                          *
// *    -create: 2023/10/25 23:13:52 ColeCai.                                                                          *
// *********************************************************************************************************************
func (r *Rotator) write(buff []byte) (int, error) {
	if r.fd == nil {
		return 0, errors.WithStack(fmt.Errorf("f.fd is nil"))
	}
	n, err := r.fd.Write(buff)
	if r.maxSize > 0 {
		r.curSize += uint64(n)
	}
	return n, err
}

// *********************************************************************************************************************
// * SUMMARY: implement io.Write interface. do split log file, then write log to Writer.                               *
// * WARNING: none.                                                                                                    *
// * HISTORY:                                                                                                          *
// *    -create: 2023/10/25 23:11:07 ColeCai.                                                                          *
// *********************************************************************************************************************
func (r *Rotator) Write(buff []byte) (int, error) {
	if !r.isSplit() {
		return r.write(buff)
	}
	name := r.genFileName()
	size := uint64(len(buff))

	r.Lock()
	defer r.Unlock()

	if !r.sizeCheck(size) && !r.fileCheck(name) {
		return r.write(buff)
	}

	fd, err := r.rotate(name)
	if err != nil {
		return r.write(buff)
	}
	if r.fd != nil {
		r.fd.Close()
	}
	r.fd = fd
	r.fileName = name
	r.curSize = 0
	return r.write(buff)
}
