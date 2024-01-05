/*
  可以引入的库和版本相关请参考 “环境说明”
  solution 函数将会在测试⽤例被调⽤

  Please refer to the "Environmental Notes" for the libraries and versions that can be introduced.
  The solution function will be called in the test case.
*/

package main

import "time"

type callJobFuncWithParams func(interface{}, ...interface{}) // 一个mock方法，用于模拟各种具体操作

type schedule struct {
	stopCh    chan struct{}
	jobFunc   interface{}
	jobParams []interface{}
	ticker    *time.Ticker
	running   bool
}

func newSchedule(period time.Duration, jobFunc interface{}, params ...interface{}) *schedule {
	return &schedule{
		stopCh:    make(chan struct{}),
		jobFunc:   jobFunc,
		jobParams: params,
		ticker:    time.NewTicker(period),
	}
}

// 请完成下面两个函数
func (s *schedule) start() {
	go func() {
		for {
			select {
			case <-s.stopCh:
				return
			case <-s.ticker.C:
				c := new(callJobFuncWithParams)
				(*c)(s.jobFunc, s.jobParams...)
			}
		}
	}()
}

func (s *schedule) stop() {
	s.running = false
	s.stopCh <- struct{}{}
	s.ticker.Stop()
}
