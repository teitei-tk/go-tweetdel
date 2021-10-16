package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

const (
	RunMode_Detect = "detect"
	RunMode_Run    = "run"

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

	flags.StringVar(&f.ArchiveDir, "archiveDir", "archive", "your twitter data archive dir path")
	flags.StringVar(&f.RunMode, "runMode", RunMode_Detect, "The execution status of the application, which can be either 'detect' or 'run'. The initial value is detect.")

	if err := flags.Parse(args); err != nil {
		return nil, errors.Wrap(err, "failed flag parse")
	}

	return f, nil
}

func (f *CliFlags) Validate() error {
	if f.From.After(f.To) {
		return errors.Errorf("invalid time range. from: %v > to: %v", f.From, f.To)
	}

	archivePath, err := filepath.Abs(f.ArchiveDir)
	if err != nil {
		return err
	}

	_, err = os.Stat(archivePath)
	if err != nil && os.IsNotExist(err) {
		return errors.Wrap(err, "dose not archive directory")
	}

	_, err = os.Stat(filepath.Join(archivePath, "data", "tweet.js"))
	if err != nil && os.IsNotExist(err) {
		return errors.Wrap(err, "does not tweet.js")
	}

	if f.RunMode != RunMode_Detect && f.RunMode != RunMode_Run {
		return errors.New(fmt.Sprintf("%s is invalid runMode", f.RunMode))
	}

	return nil
}
