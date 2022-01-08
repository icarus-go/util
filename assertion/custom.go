package assertion

type Unset interface {
	Unset(value interface{}) bool
}
