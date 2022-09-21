package register_user

import (
	"fmt"
	"time"
	"unicode/utf8"
)

type userName string
type userGender int
type userBirthDay time.Time

type registeringUser struct {
	name     userName
	gender   userGender
	birthDay userBirthDay
}

type RegisteringUser interface {
	Name() userName
	Gender() userGender
	BirthDayToTime() time.Time
}

func userNameIsEmpty(name string) bool {
	if utf8.RuneCountInString(name) == 0 {
		return true
	}
	return false
}

func userNameIsTooLong(name string) bool {
	if utf8.RuneCountInString(name) > 1000 {
		return true
	}
	return false
}

func userGenderIsNotEmpty(gender int) bool {
	if gender == 0 {
		return true
	}
	return false
}

func userGenderIsUnknown(gender int) bool {
	if gender < 0 && gender > 2 {
		return true
	}
	return false
}

func userBirthDayIsEmpty(birthDay time.Time) bool {
	if birthDay.IsZero() {
		return true
	}
	return false
}

func userBirthDayIsTooOld(birthDay time.Time) bool {
	if birthDay.Before(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return true
	}
	return false
}

func userBirthDayIsTooYoung(birthDay time.Time) bool {
	if birthDay.After(time.Now()) {
		return true
	}
	return false
}

func NewRegisteringUser(name string, gender int, birthDay string) (RegisteringUser, error) {
	if userNameIsEmpty(name) {
		return nil, fmt.Errorf("user name is empty")
	}
	if userNameIsTooLong(name) {
		return nil, fmt.Errorf("user name is too long")
	}
	userName := userName(name)

	if userGenderIsNotEmpty(gender) {
		return nil, fmt.Errorf("user gender is empty")
	}
	if userGenderIsUnknown(gender) {
		return nil, fmt.Errorf("input gender is unknown")
	}
	userGender := userGender(gender)

	timeBirthDay, err := time.Parse("2006-01-02", birthDay)
	if err != nil {
		return nil, fmt.Errorf("input birth day is invalid")
	}
	if userBirthDayIsEmpty(timeBirthDay) {
		return nil, fmt.Errorf("user birthDay is empty")
	}
	if userBirthDayIsTooOld(timeBirthDay) {
		return nil, fmt.Errorf("user birthDay is too old")
	}
	if userBirthDayIsTooYoung(timeBirthDay) {
		return nil, fmt.Errorf("user birthDay is too young")
	}
	userBirthDay := userBirthDay(timeBirthDay)

	return &registeringUser{
		name:     userName,
		gender:   userGender,
		birthDay: userBirthDay,
	}, nil
}

func (ru registeringUser) Name() userName {
	return ru.name
}

func (ru registeringUser) Gender() userGender {
	return ru.gender
}

func (ru registeringUser) BirthDayToTime() time.Time {
	return time.Time(ru.birthDay)
}
