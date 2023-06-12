package types

import "github.com/pkg/errors"

type IExtract interface {
	Extract(source string, p string) error
}

func NewIExtract(t string) (IExtract, error) {
	switch t {
	case "zip":
		return &ZIP{}, nil
	case "gz":
		return &GZIP{}, nil
	}
	return nil, errors.Errorf("not support file type: [%s]", t)
}
