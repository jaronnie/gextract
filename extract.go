package extract

import (
	"os"

	"github.com/jaronnie/extract/types"

	"github.com/h2non/filetype"
)

func Extract(source string, opts ...Opt) error {
	option := &Option{}
	for _, opt := range opts {
		err := opt(option)
		if err != nil {
			return err
		}
	}

	err := defaultOption(option)
	if err != nil {
		return err
	}

	t, err := filetype.MatchFile(source)
	if err != nil {
		return err
	}
	return types.NewIExtract(t.Extension).Extract(source, option.p)
}

func defaultOption(option *Option) (err error) {
	if option.p == "" {
		option.p, err = os.Getwd()
		if err != nil {
			return err
		}
	}
	return nil
}
