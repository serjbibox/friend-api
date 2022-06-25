package service

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"time"
)

const (
	EMPTY_FIELD = "пустое поле: "
	ADULT_AGE   = 18
)

func (s *UserService) ValidateFields() *UserService {
	s.ErrStatus = http.StatusBadRequest
	if s.Mode == REGISTER_MODE {
		switch {
		case s.Err != nil:
			return s
		case s.User.Phone == "":
			s.Err = errors.New(EMPTY_FIELD + "phone")
		case s.User.Name == "":
			s.Err = errors.New(EMPTY_FIELD + "name")
		case s.User.Birth == "":
			s.Err = errors.New(EMPTY_FIELD + "birth")
		case s.User.Login == "":
			s.Err = errors.New(EMPTY_FIELD + "login")
		case s.User.Password == "":
			s.Err = errors.New(EMPTY_FIELD + "password")
		}
	}
	if s.Mode == LOGIN_MODE {
		switch {
		case s.Err != nil:
			return s
		case s.User.Login == "":
			s.Err = errors.New(EMPTY_FIELD + "login")
		case s.User.Password == "":
			s.Err = errors.New(EMPTY_FIELD + "password")
		}
	}
	return s
}

func (s *UserService) ValidateData() *UserService {
	s.ErrStatus = http.StatusBadRequest
	if s.Mode == REGISTER_MODE {
		err := isAdult(s.User.Birth)
		if err != nil {
			s.Err = fmt.Errorf("%w", err)
			return s
		}
		switch {
		case s.Err != nil:
			return s
		case s.User.Email != "" && !isValid(s.User.Email):
			s.Err = errors.New("неверный email адрес")
			return s
		}
	}
	return s
}

func isValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isAdult(birth string) error {
	t, err := time.Parse("2006-01-02", birth)
	if err != nil {
		return err
	}
	now := time.Now()
	dy := now.Year() - t.Year()
	dm := now.Month() - t.Month()
	dd := now.Day() - t.Day()
	if dy > ADULT_AGE {
		return nil
	}
	if dy == ADULT_AGE && dm > 0 {
		return nil
	}
	if dy == ADULT_AGE && dm == 0 && dd >= 0 {
		return nil
	}
	return errors.New("пользователь несовершеннолетний")
}
