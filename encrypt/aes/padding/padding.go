package padding

//Padding
//  Author: Kevin·CC
//  Description: https://github.com/mardukbp/padding/blob/master/padding.go
type Padding interface {
	WipeOff([]byte, int) []byte
	Append([]byte, int) []byte
}
