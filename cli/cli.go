package cli

import (
	"flag"
	"time"
)

const (
	RunMode_Dry = "dry"
	RunMode_Run = "run"
)

type CliFlags struct {
	From       string
	To         string
	ArchiveDir string
	RunMode		 string
}

func ParseFlags() *CliFlags {
	f := &CliFlags{}
	flag.StringVar(&f.From, "from", time.Now().AddDate(0, -1, 0).Format("2006-05"), "year and month of from. you want to tweet delete.")
	flag.StringVar(&f.To, "to", time.Now().Format("2006-05"), "year and month of to. you want to tweet delete.")
	flag.StringVar(&f.ArchiveDir, "archiveDir", "./twitter-archives", "your twitter data archive dir path")
	flag.StringVar(&f.RunMode, "runMode", RunMode_Dry, "The execution status of the application, which can be either 'dry' or 'run'. The initial value is dry.")
	flag.Parse()
	return f
}