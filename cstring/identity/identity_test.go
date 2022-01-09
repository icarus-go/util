package identity

import (
	"errors"
	"testing"
)

func Test_Identity_Parse(t *testing.T) {
	idCard := "440508199606192918"
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

func Test_Identity_Mainland18(t *testing.T) {
	idCard := "440508199606192918"

	age, err := Mainland18.Age(idCard)
	if err != nil {
		return
	}

	birthday, err := Mainland18.Birthday(idCard)
	if err != nil {
		return
	}

	gender, err := Mainland18.Gender(idCard)
	if err != nil {
		return
	}

	t.Log(Mainland18.GenderCNName(gender))
	t.Log(birthday)

	t.Log(age)

}
