package data

import "encoding/json"

type Body struct {
	body map[string]interface{}
}

func (b *Body) Add(k string, v interface{}) *Body {
	b.body[k] = v
	return b
}

//NewBodyFormString
func NewBodyFormString(v string) (*Body, error) {
	m := make(map[string]interface{})

	if err := json.Unmarshal([]byte(v), &m); err != nil {
		return nil, err
	}

	instance := new(Body)

	instance.body = m

	return instance, nil
}

func NewBodyFormStruct(v interface{}) (*Body, error) {
	m := make(map[string]interface{})

	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bytes, &m); err != nil {
		return nil, err
	}

	instance := &Body{}
	instance.body = m

	return instance, nil
}

func NewBodyFormMap(v map[string]interface{}) *Body {
	instance := &Body{
		body: v,
	}
	return instance
}

//NewBody
func NewBody() *Body {
	return &Body{
		body: make(map[string]interface{}),
	}
}
