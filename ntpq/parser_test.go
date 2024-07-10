package ntpq

import (
	"strings"
	"testing"
)

func TestParseNTPQSysStats(t *testing.T) {

	t.Run("expected values", func(t *testing.T) {

		cmdout := `uptime: 1
sysstats reset:         2
packets received:       3
current version:        4
older version:          5
bad length or format:   6
authentication failed:  7
declined:               8
restricted:             9
rate limited:           10
KoD responses:          11
processed for time:     12`

		expected := &NTPSysStats{
			Uptime:           1,
			SysStatsReset:    2,
			PacketsRecieved:  3,
			CurrentVersion:   4,
			OlderVersion:     5,
			BadRequest:       6,
			AuthFailed:       7,
			Declined:         8,
			Restricted:       9,
			RateLimited:      10,
			KoDResponses:     11,
			ProcessedForTime: 12,
		}

		stats, err := parseNTPQSysStats(strings.NewReader(cmdout))
		if err != nil {
			t.Errorf("unexpected error returned: %s", err)
		}

		if stats.Uptime != expected.Uptime {
			t.Errorf("uptime: want=%d; got=%d", expected.Uptime, stats.Uptime)
		}
		if stats.SysStatsReset != expected.SysStatsReset {
			t.Errorf("sysstats reset: want=%d; got=%d", expected.SysStatsReset, stats.SysStatsReset)
		}
		if stats.PacketsRecieved != expected.PacketsRecieved {
			t.Errorf("packets received: want=%d; got=%d", expected.PacketsRecieved, stats.PacketsRecieved)
		}
		if stats.CurrentVersion != expected.CurrentVersion {
			t.Errorf("current version: want=%d; got=%d", expected.CurrentVersion, stats.CurrentVersion)
		}
		if stats.OlderVersion != expected.OlderVersion {
			t.Errorf("older version: want=%d; got=%d", expected.OlderVersion, stats.OlderVersion)
		}
		if stats.BadRequest != expected.BadRequest {
			t.Errorf("bad length or format: want=%d; got=%d", expected.BadRequest, stats.BadRequest)
		}
		if stats.AuthFailed != expected.AuthFailed {
			t.Errorf("authentication failed: want=%d; got=%d", expected.AuthFailed, stats.AuthFailed)
		}
		if stats.Declined != expected.Declined {
			t.Errorf("declined: want=%d; got=%d", expected.Declined, stats.Declined)
		}
		if stats.Restricted != expected.Restricted {
			t.Errorf("restricted: want=%d; got=%d", expected.Restricted, stats.Restricted)
		}
		if stats.RateLimited != expected.RateLimited {
			t.Errorf("rate limited: want=%d; got=%d", expected.RateLimited, stats.RateLimited)
		}
		if stats.KoDResponses != expected.KoDResponses {
			t.Errorf("KoD responses: want=%d; got=%d", expected.KoDResponses, stats.KoDResponses)
		}
		if stats.ProcessedForTime != expected.ProcessedForTime {
			t.Errorf("processed for time: want=%d; got=%d", expected.ProcessedForTime, stats.ProcessedForTime)
		}
	})

	t.Run("detect malformed line", func(t *testing.T) {
		cmdout := `uptime: 1: 2`
		_, err := parseNTPQSysStats(strings.NewReader(cmdout))
		if err == nil {
			t.Error("failed to detect malformed line")
		}
	})

	t.Run("detect unknown stat key", func(t *testing.T) {
		cmdout := `nostat: 2`
		_, err := parseNTPQSysStats(strings.NewReader(cmdout))
		if err == nil {
			t.Error("failed to detect unknown key: nostat")
		}
	})

	t.Run("detect invalid value", func(t *testing.T) {
		cmdout := `uptime: 1
nostat: nan`
		_, err := parseNTPQSysStats(strings.NewReader(cmdout))
		if err == nil {
			t.Error("failed to detect invalid value for key: nostat")
		}
	})
}
