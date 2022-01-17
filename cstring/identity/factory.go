package identity

import (
	"errors"
	"github.com/icarus-go/utils/cstring/constant"
)

type ParseFactory interface {
	// IsLocalValidIDCard
	//  Description: 是否是当地有效的身份证
	//  Author:  Kevin·CC
	IsLocalValidIDCard(idCard string) bool
	Birthday(idCard string) (string, error)
	Age(idCard string) (int, error)
	Gender(idCard string) (int, error)
}

func New(length constant.IdentityLength) (ParseFactory, error) {
	switch length {
	case constant.Mainland15:
		return Mainland15, nil
	case constant.Mainland18:
		return Mainland18, nil
	case constant.TaiWanAoMen:
		return TaiWanAoMen, nil
	case constant.HongKong:
		return HongKong, nil
	}
	return &mainland18{}, errors.New("错误身份证长度")
}
