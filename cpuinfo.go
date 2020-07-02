package cpuinfo

import (
	"bufio"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
)

// CPUInfo - is a CPU information structure
type CPUInfo struct {
	Vendor    string   `json:"vendor"`
	Family    int64    `json:"family"`
	Model     int64    `json:"model"`
	ModelName string   `json:"model_name"`
	Microcode int64    `json:"microcode"`
	CoreCount int64    `json:"core_count"`
	Features  []string `json:"features"`
}

// Get - return CPU information structure
func Get() (info *CPUInfo, err error) {
	return getCPUInfo()
}

// Sum - возвращает контрольную сумму информации о процессоре
func (info *CPUInfo) Sum() [64]byte {
	data, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	return sha512.Sum512(data)
}

func fromMap(m map[string]string) (info *CPUInfo) {
	info = &CPUInfo{}
	for key, value := range m {
		switch key {
		case "vendor":
			info.Vendor = value
		case "family":
			info.Family, _ = strconv.ParseInt(value, 10, 64)
		case "model":
			info.Model, _ = strconv.ParseInt(value, 10, 64)
		case "model_name":
			info.ModelName = value
		case "microcode":
			if strings.Contains(value, "0x") {
				info.Microcode = valueFromHex(value)
			} else {
				info.Microcode, _ = strconv.ParseInt(value, 10, 64)
			}
		case "core_count":
			info.CoreCount, _ = strconv.ParseInt(value, 10, 64)
		case "features":
			info.Features = strings.Split(strings.ToLower(value), " ")
		}
	}
	return
}

func valueFromHex(value string) (result int64) {
	b, _ := hex.DecodeString(strings.TrimPrefix(value, "0x"))
	if len(b) == 0 {
		return
	}
	switch length := len(b); length {
	case 1:
		result = int64(b[0])
	case 2:
		result = int64(binary.BigEndian.Uint16(b))
	case 4, 8:
		result = int64(binary.BigEndian.Uint64(b))
	}
	return
}

func readScanner(scanner *bufio.Scanner) (info *CPUInfo) {
	raw := make(map[string]string)

	for scanner.Scan() {
		for name, r := range regexps {
			if !r.Match(scanner.Bytes()) {
				continue
			}
			fields := r.FindStringSubmatch(scanner.Text())
			if len(fields) < 2 {
				continue
			}
			raw[name] = fields[1]
		}
	}

	info = fromMap(raw)
	return
}
