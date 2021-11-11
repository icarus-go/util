package constant

type IdentityLength int

var (
	Any = []IdentityLength{
		Mainland18,
		Mainland15,
		TaiWanAoMen,
		HongKong,
	}
)

const (
	//Mainland15 大陆15位数身份证
	Mainland15 IdentityLength = 15
	//Mainland18 大陆18位身份证
	Mainland18 IdentityLength = 18
	//TaiWanAoMen 澳台
	TaiWanAoMen IdentityLength = 10
	// HongKong 香港
	HongKong IdentityLength = 8
)

// Value
//  Description: 值
//  Author: Kevin·CC
//  return int
func (a IdentityLength) Value() int {
	return int(a)
}
