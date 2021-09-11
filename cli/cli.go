package cli

import (
	"flag"
	"time"

	"github.com/pkg/errors"
)

const (
	RunMode_Dry = "dry"
	RunMode_Run = "run"

	dateLayout = "2006-01"
)

type CliFlags struct {
	From       time.Time
	To         time.Time
	ArchiveDir string
	RunMode    string
}

func ParseFlags(prgName string, args []string) (*CliFlags, error) {
	flags := flag.NewFlagSet(prgName, flag.ContinueOnError)

	f := &CliFlags{}
	flags.Func("from", "year and month of from. you want to tweet delete.", func(s string) error {
		if s == "" {
			f.From = time.Now().AddDate(0, -1, 0)
			return nil
		}

		from, err := time.Parse(dateLayout, s)
		if err != nil {
			return errors.Wrap(err, "invalid from time value")
		}

		f.From = from
		return nil
	})

	flags.Func("to", "year and month of to. you want to tweet delete.", func(s string) error {
		if s == "" {
			f.To = time.Now()
			return nil
		}

		to, err := time.Parse(dateLayout, s)
		if err != nil {
			return errors.Wrap(err, "invalid to time value")
		}

		f.To = to
		return nil
	})

	flags.StringVar(&f.ArchiveDir, "archiveDir", "./twitter-archives", "your twitter data archive dir path")
	flags.StringVar(&f.RunMode, "runMode", RunMode_Dry, "The execution status of the application, which can be either 'dry' or 'run'. The initial value is dry.")

	if err := flags.Parse(args); err != nil {
		return nil, errors.Wrap(err, "failed flag parse")
	}

	return f, nil
}
