package identity

import "errors"

var (
	ErrUnparseGender   = errors.New("无法解析性别")
	ErrUnparseAge      = errors.New("无法解析年龄")
	ErrUnparseBirthday = errors.New("无法解析生日")
)
