package cstring

// MaskAsterisk json输出星号掩码
type MaskAsterisk string

// String 获取字段串
func (m MaskAsterisk) String() string {
	return string(m)
}

// MarshalJSON  序列化
func (m MaskAsterisk) MarshalJSON() ([]byte, error) {

	if m.String() == "" {
		return []byte(""), nil
	}

	rs := []rune(m.String())
	startIndex, length := m.rule()
	asteristkString := m.asterisk(length)

	result := string(rs[:startIndex]) + asteristkString + string(rs[startIndex+length:])

	return []byte(`"` + result + `"`), nil
}

func (m MaskAsterisk) rule() (int, int) {
	rs := []rune(m.String())
	l := len(rs)
	switch l {
	case 2:
	case 3:
		return 0, 1
	case 11:
		return 3, 4
	case 18:
		return 4, 10
	default:
		if l > 3 {
			mol := l % 4
			s := l / 4
			return s, 2*s + mol
		}
		return 0, 1
	}
	return 0, 0
}

func (m MaskAsterisk) asterisk(count int) string {
	if count < 1 {
		return ""
	}

	ret := ""
	for i := 0; i < count; i++ {
		ret += "*"
	}

	return ret
}
