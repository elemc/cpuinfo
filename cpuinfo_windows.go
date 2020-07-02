package cpuinfo

import (
	"os/exec"
	"regexp"

	"github.com/pkg/errors"
)

var regexps = map[string]*regexp.Regexp{
	"vendor":     regexp.MustCompile("^Manufacturer\\n(?P<vendor>.*$)"),
	"family":     regexp.MustCompile("^Caption\\n.*\\sFamily\\s(?P<family>\\d).*$"),
	"model":      regexp.MustCompile("^Caption\\n.*Model\\s(?P<model>\\d+).*$"),
	"model_name": regexp.MustCompile("^Name\\n(?P<model_name>.*$)"),
	"core_count": regexp.MustCompile("^NumberOfCores\\n(?P<core_count>.*$)"),
}

func getCommandOutput(what string) (output []byte, err error) {
	cmd := exec.Command("wmic", "cpu", "get", what)
	output, err = cmd.CombinedOutput()
	if err != nil {
		err = errors.Wrapf(err, "unable to start (wmic cpu get %s)", what)
		return
	}
	return
}

func getCommandValue(what string, raw map[string]string) (err error) {
	output, err := getCommandOutput(what)
	if err != nil {
		return
	}
	setValues(output, raw)
	return
}

func setValues(data []byte, raw map[string]string) {
	for key, r := range regexps {
		if !r.Match(data) {
			return
		}
		fields := r.FindStringSubmatch(string(data))
		if len(fields) < 2 {
			continue
		}
		raw[key] = fields[1]
	}
	return
}

func getCPUInfo() (info *CPUInfo, err error) {
	raw := make(map[string]string)

	// Vendor
	err = getCommandValue("manufacturer", raw)
	if err != nil {
		err = errors.Wrap(err, "unable to get vendor")
		return
	}

	// Family Model
	err = getCommandValue("caption", raw)
	if err != nil {
		err = errors.Wrap(err, "unable to get family and model")
		return
	}

	// Model name
	err = getCommandValue("name", raw)
	if err != nil {
		err = errors.Wrap(err, "unable to get model name")
		return
	}

	// core count
	err = getCommandValue("numberofcores", raw)
	if err != nil {
		err = errors.Wrap(err, "unable to get core count")
		return
	}

	info = fromMap(raw)
	return
}
