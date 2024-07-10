package ntpq

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var KeySysStatsUptime = "uptime"
var KeySysStatsReset = "sysstats reset"
var KeySysStatsPacketsReceived = "packets received"
var KeySysStatsCurrentVersion = "current version"
var KeySysStatsOlderVersion = "older version"
var KeySysStatsBadRequest = "bad length or format"
var KeySysStatsAuthFailed = "authentication failed"
var KeySysStatsDeclined = "declined"
var KeySysStatsRestricted = "restricted"
var KeySysStatsRateLimited = "rate limited"
var KeySysStatsKoDResponses = "KoD responses"
var KeySysStatsProcessedForTime = "processed for time"

func parseNTPQSysStats(str io.Reader) (stats *NTPSysStats, err error) {
	reader := bufio.NewReader(str)

	stats = &NTPSysStats{}

	var line []byte
	var intVal int
	for {
		if line, _, err = reader.ReadLine(); err == io.EOF {
			break
		}

		split := strings.Split(string(line), ":")
		if len(split) != 2 {
			return nil, errors.New(fmt.Sprintf("malformed line: '%s", string(line)))
		}

		key := strings.TrimSpace(split[0])
		value := strings.TrimSpace(split[1])

		intVal, err = strconv.Atoi(value)
		if err != nil {
			return nil, errors.New(
				fmt.Sprintf("bad value: '%s' for key: '%s'", value, key))
		}

		switch key {
		case KeySysStatsUptime:
			stats.Uptime = intVal
		case KeySysStatsReset:
			stats.SysStatsReset = intVal
		case KeySysStatsPacketsReceived:
			stats.PacketsRecieved = intVal
		case KeySysStatsCurrentVersion:
			stats.CurrentVersion = intVal
		case KeySysStatsOlderVersion:
			stats.OlderVersion = intVal
		case KeySysStatsBadRequest:
			stats.BadRequest = intVal
		case KeySysStatsAuthFailed:
			stats.AuthFailed = intVal
		case KeySysStatsDeclined:
			stats.Declined = intVal
		case KeySysStatsRestricted:
			stats.Restricted = intVal
		case KeySysStatsRateLimited:
			stats.RateLimited = intVal
		case KeySysStatsKoDResponses:
			stats.KoDResponses = intVal
		case KeySysStatsProcessedForTime:
			stats.ProcessedForTime = intVal
		default:
			return nil, errors.New(fmt.Sprintf("unknown key: %s", key))

		}

	}

	return stats, nil
}
