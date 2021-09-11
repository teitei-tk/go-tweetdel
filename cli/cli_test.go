package cli_test

import (
	"flag"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/teitei-tk/go-tweetdel/cli"
)

func TestParseFlags(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		description string
		from        string
		to          string
		archiveDir  string
		runMode     string
		wantError   bool
	}{
		{
			description: "default value",
			from:        time.Now().AddDate(0, -1, 0).Format("2006-01"),
			to:          time.Now().Format("2006-01"),
			archiveDir:  "./twitter-archives",
			runMode:     cli.RunMode_Dry,
			wantError:   false,
		},
		{
			description: "from",
			from:        "invalidDate",
			to:          time.Now().Format("2006-01"),
			archiveDir:  "./twitter-archives",
			runMode:     cli.RunMode_Dry,
			wantError:   true,
		},
		{
			description: "to",
			from:        time.Now().AddDate(0, -1, 0).Format("2006-01"),
			to:          "invalidDate",
			archiveDir:  "./twitter-archives",
			runMode:     cli.RunMode_Dry,
			wantError:   true,
		},
		{
			description: "archiveDir",
			from:        time.Now().AddDate(0, -1, 0).Format("2006-01"),
			to:          time.Now().Format("2006-01"),
			archiveDir:  "./other-twitter-archives.zip",
			runMode:     cli.RunMode_Dry,
			wantError:   false,
		},
		{
			description: "runMode",
			from:        time.Now().AddDate(0, -1, 0).Format("2006-01"),
			to:          time.Now().Format("2006-01"),
			archiveDir:  "./twitter-archives",
			runMode:     "invalidMode",
			wantError:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			flag.CommandLine.Set("from", test.from)
			flag.CommandLine.Set("to", test.to)
			flag.CommandLine.Set("archiveDir", test.archiveDir)
			flag.CommandLine.Set("runMode", test.runMode)

			args := []string{"--from", test.from, "--to", test.to, "--archiveDir", test.archiveDir, "--runMode", test.runMode}
			flags, err := cli.ParseFlags("cli_test", args)
			if test.wantError {
				assert.Error(err)
				return
			}

			assert.NoError(err)
			assert.Equal(flags.From.Format("2006-01"), test.from)
			assert.Equal(flags.To.Format("2006-01"), test.to)
			assert.Equal(flags.ArchiveDir, test.archiveDir)
			assert.Equal(flags.RunMode, flags.RunMode)
		})
	}
}
