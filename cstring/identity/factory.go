package identity

import (
	"errors"
	"pmo-test4.yz-intelligence.com/base/utils/cstring/constant"
)

type AreaFactory interface {
	// IsLocalValidIDCard
	//  Description: 是否是当地有效的身份证
	//  Author:  Kevin·CC
	IsLocalValidIDCard(idCard string) bool
}

func New(length constant.IdentityLength) (AreaFactory, error) {
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
