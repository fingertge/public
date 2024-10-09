// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/09 18:37:51
// * Proj: public
// * Pack: timer
// * File: timer.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

package timer

import (
	"container/list"
	"github.com/pkg/errors"
	"time"
)

type ITimer interface {
	StartTimer(key string, delay time.Duration, isLoop bool)
	StopTimer(key string)
	StopTimeWheel()
}

type (
	Execute   func(key string)
	baseEntry struct {
		key    string
		delay  time.Duration
		isLoop bool
	}

	Task struct {
		baseEntry
		round   int
		removed bool
	}

	posEntry struct {
		pos  int
		task *Task
	}

	TimeWheel struct {
		interval    time.Duration
		ticker      *time.Ticker
		slots       []*list.List
		timers      map[string]*posEntry
		numSlots    int
		tickPos     int
		execute     Execute
		addChannel  chan *baseEntry
		stopChannel chan struct{}
		delChhannel chan string
	}
)

func NewTimeTheel(interval time.Duration, numSlots int, execute Execute) (*TimeWheel, error) {
	if interval <= 0 || numSlots <= 0 || execute == nil {
		return nil, errors.Errorf("invalid parameter, interval: %v, numSlots: %v, execute: %p", interval, numSlots, execute)
	}
	return newTimeWheel(interval, numSlots, time.NewTicker(interval), execute)
}

func newTimeWheel(interval time.Duration, numSlots int, ticker *time.Ticker, execute Execute) (*TimeWheel, error) {
	tw := &TimeWheel{
		interval:    interval,
		ticker:      ticker,
		slots:       make([]*list.List, numSlots),
		timers:      make(map[string]*posEntry),
		numSlots:    numSlots,
		tickPos:     0,
		execute:     execute,
		addChannel:  make(chan *baseEntry),
		stopChannel: make(chan struct{}),
		delChhannel: make(chan string),
	}
	tw.initSlots()
	go tw.run()
	return tw, nil
}

func (tw *TimeWheel) initSlots() {
	for i := 0; i < tw.numSlots; i++ {
		tw.slots[i] = list.New()
	}
}

func (tw *TimeWheel) run() {
	for {
		select {
		case <-tw.ticker.C:
			tw.scanAndRun()
		case base := <-tw.addChannel:
			tw.addTask(base)
		case <-tw.stopChannel:
			return
		case key := <-tw.delChhannel:
			tw.delTask(key)
		}
	}
}

func (tw *TimeWheel) addTask(base *baseEntry) {
	if base.delay < tw.interval {
		base.delay = tw.interval
	}
	round, pos := tw.getRoundAndPos(base.delay)
	task := &Task{
		baseEntry: baseEntry{key: base.key, delay: base.delay, isLoop: base.isLoop},
		round:     round,
		removed:   false,
	}
	tw.slots[pos].PushBack(task)
	tw.timers[base.key] = &posEntry{pos: pos, task: task}
}

func (tw *TimeWheel) delTask(key string) {
	if task, ok := tw.timers[key]; ok {
		task.task.removed = true
		delete(tw.timers, key)
	}
}

func (tw *TimeWheel) getRoundAndPos(d time.Duration) (int, int) {
	steps := int(d / tw.interval)
	pos := (tw.tickPos + steps) % tw.numSlots
	round := (steps - 1) / tw.numSlots
	return round, pos
}

func (tw *TimeWheel) scanAndRun() {
	tw.tickPos = (tw.tickPos + 1) % tw.numSlots
	l := tw.slots[tw.tickPos]
	var tasks []*Task
	for t := l.Front(); t != nil; {
		task := t.Value.(*Task)
		if task.removed {
			next := t.Next()
			l.Remove(t)
			t = next
			continue
		} else if task.round > 0 {
			task.round -= 1
			t = t.Next()
			continue
		}
		tasks = append(tasks, task)
		next := t.Next()
		l.Remove(t)
		delete(tw.timers, task.key)
		t = next
		if task.baseEntry.isLoop {
			tw.addTask(&baseEntry{key: task.key, delay: task.delay, isLoop: task.isLoop})
		}
	}
	tw.runTasks(tasks)
}

func (tw *TimeWheel) runTasks(tasks []*Task) {
	if len(tasks) <= 0 {
		return
	}
	go func() {
		for _, task := range tasks {
			go tw.execute(task.key)
		}
	}()
}

func (tw *TimeWheel) StartTimer(key string, delay time.Duration, isLoop bool) {
	tw.addChannel <- &baseEntry{key: key, delay: delay, isLoop: isLoop}
}

func (tw *TimeWheel) StopTimer(key string) {
	tw.delChhannel <- key
}

func (tw *TimeWheel) StopTimeWheel() {
	tw.stopChannel <- struct{}{}
}
