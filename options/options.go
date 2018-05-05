package options

import (
	"errors"
	"flag"
)

type Options struct {
	Root      string
	Command   string
	Arguments []string
	IsSlave   bool
}

func (opts *Options) Parse(args []string) error {
	f := flag.NewFlagSet("", flag.ContinueOnError)

	f.BoolVar(&opts.IsSlave, "x-slave", false, "")

	err := f.Parse(args)
	if err != nil {
		return err
	}

	if f.NArg() < 2 {
		return errors.New("Too few non-option arguments")
	}

	opts.Root = f.Arg(0)
	opts.Command = f.Arg(1)
	opts.Arguments = f.Args()[2:]

	return nil
}
