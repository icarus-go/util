package optimization

import "go.uber.org/zap"

type goroutine struct{}

var Goroutine = new(goroutine)

// Begin
//  Author: Kevin·CC
//  Description: 开启一个协程
//  Param fn 执行方法主体
func (*goroutine) Begin(fn func()) {
	go func() {
		defer func() {
			if i := recover(); i != nil {
				zap.L().Error("协程出现错误", zap.Any("panic", i))
			}
		}()

		fn()

	}()

	return
}
