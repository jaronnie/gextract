package extract

type Opt func(*Option) error

type Option struct {
	p string // output path, default is `pwd`
}

func WithOutputPath(p string) Opt {
	return func(opt *Option) error {
		opt.p = p
		return nil
	}
}
