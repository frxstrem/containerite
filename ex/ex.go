package ex

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Autorecover(errout *error) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			*errout = err
		}
	}
}
