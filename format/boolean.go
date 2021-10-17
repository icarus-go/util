package format

type boolean struct {
	Result *bool
}

func (b *boolean) Int() int {
	isOK := 0
	if b.Result != nil && *b.Result {
		isOK = 1
	}
	return isOK
}
