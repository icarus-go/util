package optimization

import (
	"errors"
	"time"
)

type (
	Subject func() error
	_retry  struct {
		subject Subject
		count   int
		sleep   time.Duration // 秒
	}
)

//check 检查运行参数
func (r *_retry) check() error {
	if r.subject == nil {
		return errors.New("执行主体为空")
	}

	if r.count < 1 {
		return errors.New("重试次数为0")
	}

	return nil
}

//NewRetry 创建重试优化对象
func NewRetry() *_retry {
	return &_retry{
		sleep: time.Second * 3,
	}
}

//Subject 执行的方法主体
func (r *_retry) Subject(subject Subject) *_retry {
	r.subject = subject
	return r
}

//Count 重试次数
func (r *_retry) Count(count int) *_retry {
	r.count = count
	return r
}

//Sleep 睡眠间隔 ， 秒为单位
func (r *_retry) Sleep(sleep int) *_retry {
	r.sleep = time.Second * time.Duration(sleep)
	return r
}

//Run 执行
func (r *_retry) Run() error {
	if err := r.check(); err != nil {
		return err
	}

	for i := 0; i < r.count; i++ {
		if err := r.subject(); err == nil {
			return nil
		}
		time.Sleep(r.sleep)
	}
	return errors.New("始终失败")
}
