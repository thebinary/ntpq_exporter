package ntpq

type NTPSysStats struct {
	Uptime           int
	SysStatsReset    int
	PacketsRecieved  int
	CurrentVersion   int
	OlderVersion     int
	BadRequest       int
	AuthFailed       int
	Declined         int
	Restricted       int
	RateLimited      int
	KoDResponses     int
	ProcessedForTime int
}
