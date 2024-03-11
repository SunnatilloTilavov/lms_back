package check

import (
	"errors"
	"time"
	"strings"
)

func ValidateCarYear(year int) error {
	if year <= 0 || year > time.Now().Year()+1 {
		return errors.New("year is not valid")
	}
	return nil
}
func ValidateEmail(gmail string)bool{
atCount:=strings.Count(gmail,"@")
dotCount:=strings.Count(gmail,".")
return atCount==1&&dotCount==1&&(strings.HasSuffix(gmail,"@gmail.com")||strings.HasSuffix(gmail,"@mail.ru") )
}

func ValidatePhone(phone string)bool{
	plusCount:=strings.Count(phone,"+")
	return plusCount==1&&strings.HasPrefix(phone,"+998")&&len(phone)==13
	}
	
