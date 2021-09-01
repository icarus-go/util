package date

type Format string

const (
	//YMDHms 年-月-日 时:分:秒
	YMDHms Format = "2006-01-02 15:04:06"
	//Ymd 年-月-日
	Ymd Format = "2006-01-02"
	//YMDHmsSlash 年/月/日 时:分:秒
	YMDHmsSlash Format = "2006/01/02 15:04:06"
	//YmdSlash 年/月/日
	YmdSlash Format = "2006/01/02"
)

//Value 值
func (d Format) Value() string {
	return string(d)
}
