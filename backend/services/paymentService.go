package services

import (
	"errors"
	"flight/models"
	"strconv"
	"time"
	"unicode"
)

func CheckPaymentValidation(payment *models.PaymentModel) (bool, error) {
	if !validate(payment.CardNumber) {
		return false, errors.New("invalid card number")
	}
	if len(payment.CVV) != 3 {
		return false, errors.New("invalid CVV")
	}
	if len(payment.ExpiryDate) != 5 || payment.ExpiryDate[2] != '/' {
		return false, errors.New("invalid expiry date format")
	}
	month, err := strconv.Atoi(payment.ExpiryDate[:2])
	if err != nil || month < 1 || month > 12 {
		return false, errors.New("invalid expiry month")
	}
	year, err := strconv.Atoi(payment.ExpiryDate[3:])
	if err != nil || year < time.Now().Year()%100 {
		return false, errors.New("invalid expiry year")
	}
	return true, nil
}

func validate(cn string) bool {
	var ss string
	for _, r := range cn {
		if !unicode.IsSpace(r) {
			ss += string(r)
		}
	}
	var sum int64 = 0
	parity := len(ss) % 2

	cardNumWithoutChecksum := ss[:len(ss)-1]

	for i, v := range cardNumWithoutChecksum {

		item, err := strconv.Atoi(string(v))

		if err != nil {
			return false
		}
		if int64(i)%2 != int64(parity) {
			sum += int64(item)
		} else if item > 4 {
			sum += int64(2*item - 9)
		} else {
			sum += int64(2 * item)
		}

	}

	checkDigit, err := strconv.Atoi(ss[len(ss)-1:])

	if err != nil {
		return false
	}

	SumMod := sum % 10

	if SumMod == int64(0) {
		return SumMod == int64(checkDigit)
	}
	return int64(10)-SumMod == int64(checkDigit)
}
