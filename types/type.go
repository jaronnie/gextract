package types

type IExtract interface {
	Extract(source string, p string) error
}

func NewIExtract(t string) IExtract {
	switch t {
	case "zip":
		return &ZIP{}
	}
	return nil
}
