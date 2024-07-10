package ntpq

import (
	"bytes"
	"os/exec"
	"strings"
)

const cmd_ntpq = "ntpq"

func GetNTPQSysStats() (stats *NTPSysStats, err error) {
	var out bytes.Buffer

	cmd := exec.Command(cmd_ntpq, "-c", "sysstats")
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	return parseNTPQSysStats(strings.NewReader(out.String()))
}
