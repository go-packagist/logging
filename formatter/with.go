package formatter

func WithFormat(format string) Opt {
	return func(f Formatter) {
		switch f.(type) {
		case *LineFormatter:
			f.(*LineFormatter).format = format
		default:
			panic("not support")
		}
	}
}

func WithTimeFormat(timeFormat string) Opt {
	return func(f Formatter) {
		switch f.(type) {
		case *LineFormatter:
			f.(*LineFormatter).timeFormat = timeFormat
		default:
			panic("not support")
		}
	}
}
