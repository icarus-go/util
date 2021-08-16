package constant

type Method string

const (
	GET  Method = "GET"
	POST Method = "POST"
)

func (m Method) Value() string {
	return string(m)
}
