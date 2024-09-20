package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

type PhoneNumber struct {
  area string
  numberBegin string
  numberEnd string
}

func (p PhoneNumber) String() string {
  return fmt.Sprintf("(%s) %s-%s", p.area, p.numberBegin, p.numberEnd)
}

func NewPhoneNumber(phoneNumber string) (PhoneNumber, error) {
  re := regexp.MustCompile(`^\+?(1)?\D*\(?([2-9]\d{2})\)?\D*([2-9]\d{2})\D*(\d{4})\D*$`)
  res := re.FindStringSubmatch(phoneNumber)
  
  if res == nil {
    return PhoneNumber{}, errors.New("Invalid phoneNumber")
  }
  if len(res) == 5 {
    return PhoneNumber{area: res[2], numberBegin: res[3], numberEnd: res[4]}, nil
  }
  return PhoneNumber{area: res[1], numberBegin: res[2], numberEnd: res[3]}, nil
}

func Number(phoneNumber string) (string, error) {
  p, err := NewPhoneNumber(phoneNumber)
  if err != nil {
    return "", err
  }
  return fmt.Sprintf("%s%s%s", p.area, p.numberBegin, p.numberEnd), nil
}

func AreaCode(phoneNumber string) (string, error) {
  p, err := NewPhoneNumber(phoneNumber)
  if err != nil {
    return "", err
  }
  return p.area, nil
}

func Format(phoneNumber string) (string, error) {
  p, err := NewPhoneNumber(phoneNumber)
  if err != nil {
    return "", err
  }
  return p.String(), nil
}

