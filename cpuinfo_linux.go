package cpuinfo

import (
	"bufio"
	"os"
	"regexp"

	"github.com/pkg/errors"
)

const sourceFileName = "/proc/cpuinfo"

var regexps = map[string]*regexp.Regexp{
	"vendor":     regexp.MustCompile("^vendor_id\\s+:\\s(?P<vendor>.*$)"),
	"family":     regexp.MustCompile("^cpu family\\s+:\\s(?P<family>.*$)"),
	"model":      regexp.MustCompile("^model\\s+:\\s(?P<model>.*$)"),
	"model_name": regexp.MustCompile("^model name\\s+:\\s(?P<model_name>.*$)"),
	"microcode":  regexp.MustCompile("^microcode\\s+:\\s(?P<microcode>.*$)"),
	"core_count": regexp.MustCompile("^cpu cores\\s+:\\s(?P<core_count>.*$)"),
	"features":   regexp.MustCompile("^flags\\s+:\\s(?P<features>.*$)"),
}

func getCPUInfo() (info *CPUInfo, err error) {
	f, err := os.OpenFile(sourceFileName, os.O_RDONLY, 0644)
	if err != nil {
		err = errors.Wrapf(err, "unable to open file (%s)", sourceFileName)
		return
	}
	defer f.Close()

	info = readScanner(bufio.NewScanner(f))
	return
}
