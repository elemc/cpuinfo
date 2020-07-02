package cpuinfo

import (
	"bufio"
	"bytes"
	"os/exec"
	"regexp"

	"github.com/pkg/errors"
)

var regexps = map[string]*regexp.Regexp{
	"vendor":     regexp.MustCompile("^machdep.cpu.vendor:\\s(?P<vendor>.*$)"),
	"family":     regexp.MustCompile("^machdep.cpu.family:\\s(?P<family>.*$)"),
	"model":      regexp.MustCompile("^machdep.cpu.model:\\s(?P<model>.*$)"),
	"model_name": regexp.MustCompile("^machdep.cpu.brand_string:\\s(?P<model_name>.*$)"),
	"microcode":  regexp.MustCompile("^machdep.cpu.microcode_version:\\s(?P<microcode>.*$)"),
	"core_count": regexp.MustCompile("^machdep.cpu.core_count:\\s(?P<core_count>.*$)"),
	"features":   regexp.MustCompile("^machdep.cpu.features:\\s(?P<features>.*$)"),
}

func getCPUInfo() (info *CPUInfo, err error) {
	cmd := exec.Command("sysctl", "-a")
	data, err := cmd.CombinedOutput()
	if err != nil {
		err = errors.Wrap(err, "unable to run command `sysctl -a`")
		return
	}
	info = readScanner(bufio.NewScanner(bytes.NewReader(data)))
	return
}
