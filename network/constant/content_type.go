package constant

type ContentType string

const (
	JSON        ContentType = "application/json;charset=utf-8"
	FORM        ContentType = "application/x-www-form-urlencoded"
	QueryString ContentType = "querystring"
)

func (c ContentType) Value() string {
	return string(c)
}
