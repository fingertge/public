// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/09 16:10:03
// * Proj: public
// * Pack: tools
// * File: delay.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

package tools

import "time"

type Delay struct {
	maxDelay time.Duration
	duration time.Duration
}

func NewDelay(maxDelay time.Duration) *Delay {
	return &Delay{
		maxDelay: maxDelay,
	}
}

func (d *Delay) Delay() {
	d.up()
	d.do()
}

func (d *Delay) Reset() {
	d.duration = 0
}

func (d *Delay) up() {
	if d.duration == 0 {
		d.duration = 5 * time.Millisecond
		return
	}
	d.duration = 2 * d.duration
	if d.duration > d.maxDelay {
		d.duration = d.maxDelay
	}
}

func (d *Delay) do() {
	if d.duration > 0 {
		time.Sleep(d.duration)
	}
}

func (d *Delay) Duration() time.Duration {
	return d.duration
}
