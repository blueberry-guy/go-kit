package env

import (
	"errors"
	"strings"
)

type Type string

const (
	TypeProd  = Type("production")
	TypeLocal = Type("local")
)

func FromString(str string) (Type, error) {
	switch strings.ToLower(str) {
	case string(TypeProd):
		return TypeProd, nil
	case string(TypeLocal):
		return TypeLocal, nil
	}
	return "", errors.New("invalid type")
}
