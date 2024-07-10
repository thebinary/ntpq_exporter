package exporter

import (
	"log"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/thebinary/ntpq_exporter/ntpq"
)

type NTPSysStatsExporter struct {
	mutex sync.Mutex

	uptime           *prometheus.Desc
	reset            *prometheus.Desc
	pktsReceived     *prometheus.Desc
	curVer           *prometheus.Desc
	oldVer           *prometheus.Desc
	badRequest       *prometheus.Desc
	authFail         *prometheus.Desc
	declined         *prometheus.Desc
	restricted       *prometheus.Desc
	rateLimited      *prometheus.Desc
	koDs             *prometheus.Desc
	processedForTime *prometheus.Desc
}

func NewSysStatsExporter() (e *NTPSysStatsExporter) {
	return &NTPSysStatsExporter{
		uptime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, sub_sysstats, "uptime"),
			"Uptime of NTP Daemon",
			nil,
			nil,
		),
		reset: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, sub_sysstats, "reset"),
			"Counter of sysstats resets",
			nil,
			nil,
		),
		pktsReceived: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, sub_sysstats, "packets_received"),
			"Number of Packets Received",
			nil,
			nil,
		),
		curVer: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, sub_sysstats, "requests_current_version"),
			"Client requests matching server protocol version",
			nil,
			nil,
		),
		oldVer: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, sub_sysstats, "requests_older_version"),
			"Older version requests from clients than the server protocol version",
			nil,
			nil,
		),
		badRequest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, sub_sysstats, "requests_bad"),
			"Requests with bad length or format",
			nil,
			nil,
		),
		authFail: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, sub_sysstats, "authentication_failed"),
			"Authenticated Failures",
			nil,
			nil,
		),
		declined: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, sub_sysstats, "requests_declined"),
			"Declined Requests",
			nil,
			nil,
		),
		restricted: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, sub_sysstats, "requests_restricted"),
			"Restricted Requests",
			nil,
			nil,
		),
		rateLimited: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, sub_sysstats, "requests_rate_limited"),
			"Rate limited Requests",
			nil,
			nil,
		),
		koDs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, sub_sysstats, "responses_kods"),
			"Kiss-O'-Death Responses",
			nil,
			nil,
		),
		processedForTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, sub_sysstats, "processed_for_time"),
			"Processed For Time",
			nil,
			nil,
		),
	}
}

func (e *NTPSysStatsExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.uptime
	ch <- e.reset
	ch <- e.pktsReceived
	ch <- e.curVer
	ch <- e.oldVer
	ch <- e.badRequest
	ch <- e.authFail
	ch <- e.declined
	ch <- e.restricted
	ch <- e.rateLimited
	ch <- e.koDs
	ch <- e.processedForTime
}

func (e *NTPSysStatsExporter) Collect(ch chan<- prometheus.Metric) {
	e.mutex.Lock() // To protect metrics from concurrent collects.
	defer e.mutex.Unlock()

	s, err := ntpq.GetNTPQSysStats()
	if err != nil {
		log.Printf("[ERROR] fetching stats: err=%s", err)
		return
	}

	ch <- prometheus.MustNewConstMetric(e.uptime, prometheus.CounterValue, float64(s.Uptime))
	ch <- prometheus.MustNewConstMetric(e.reset, prometheus.CounterValue, float64(s.SysStatsReset))
	ch <- prometheus.MustNewConstMetric(e.pktsReceived, prometheus.CounterValue, float64(s.PacketsRecieved))
	ch <- prometheus.MustNewConstMetric(e.curVer, prometheus.CounterValue, float64(s.CurrentVersion))
	ch <- prometheus.MustNewConstMetric(e.oldVer, prometheus.CounterValue, float64(s.OlderVersion))
	ch <- prometheus.MustNewConstMetric(e.badRequest, prometheus.CounterValue, float64(s.BadRequest))
	ch <- prometheus.MustNewConstMetric(e.authFail, prometheus.CounterValue, float64(s.AuthFailed))
	ch <- prometheus.MustNewConstMetric(e.declined, prometheus.CounterValue, float64(s.Declined))
	ch <- prometheus.MustNewConstMetric(e.restricted, prometheus.CounterValue, float64(s.Restricted))
	ch <- prometheus.MustNewConstMetric(e.rateLimited, prometheus.CounterValue, float64(s.RateLimited))
	ch <- prometheus.MustNewConstMetric(e.koDs, prometheus.CounterValue, float64(s.KoDResponses))
	ch <- prometheus.MustNewConstMetric(e.processedForTime, prometheus.CounterValue, float64(s.ProcessedForTime))
}
