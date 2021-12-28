package identity

import (
	"errors"
	"testing"
)

func Test_Identity_Parse(t *testing.T) {
	idCard := "Q7387683"
	factory, err := Factory(idCard)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	isValid := factory.IsLocalValidIDCard(idCard)
	if !isValid {
		t.Fatal("无效身份证")
		return
	}

	birthday, err := factory.Birthday(idCard)
	if err != nil && errors.Is(err, ErrUnparseBirthday) {
		t.Log("生日为空")
		return
	}
	t.Log(birthday)
	age, err := factory.Age(idCard)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	t.Log(age)

}
