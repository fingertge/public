// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/23 14:43:29
// * Proj: work
// * Pack: tools
// * File: psutil.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
package tools

import (
	"github.com/shirou/gopsutil/v3/process"
	"os"
)

type ProcSysInfo struct {
	Pid        int
	Name       string
	CPUPercent float64
	MemPercent float64
	UsedMem    uint64
	CreatTime  int64
}

func GetProcSysInfo() (*ProcSysInfo, error) {
	pid := os.Getegid()
	p, err := process.NewProcess(int32(pid))
	if err != nil {
		return nil, err
	}
	name, err := p.Name()
	if err != nil {
		return nil, err
	}
	cp, err := p.CPUPercent()
	if err != nil {
		return nil, err
	}
	mp, err := p.MemoryPercent()
	if err != nil {
		return nil, err
	}
	mems, err := p.MemoryInfo()
	if err != nil {
		return nil, err
	}
	tm, err := p.CreateTime()
	if err != nil {
		return nil, err
	}
	pInfo := &ProcSysInfo{
		Pid:        pid,
		Name:       name,
		CPUPercent: cp,
		MemPercent: float64(mp),
		UsedMem:    mems.RSS,
		CreatTime:  tm,
	}

	return pInfo, nil
}
