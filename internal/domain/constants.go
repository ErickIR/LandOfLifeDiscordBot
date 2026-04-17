package domain

type Hour int

const (
	UnknownHour Hour = iota
	Hour0100
	Hour0300
	Hour0500
	Hour0700
	Hour0900
	Hour1100
	Hour1300
	Hour1500
	Hour1700
	Hour1900
	Hour2100
	Hour2300
)

func (h Hour) String() string {
	return [...]string{
		"Unknown",
		"<t:64060585200:t>", "<t:64060592400:t>", "<t:64060599600:t>",
		"<t:64060606800:t>", "<t:64060614000:t>", "<t:64060621200:t>",
		"<t:64060628400:t>", "<t:64060635600:t>", "<t:64060642800:t>",
		"<t:64060650000:t>", "<t:64060657200:t>", "<t:64060664400:t>",
	}[h]
}

type Channel int

const (
	UnknownChannel Channel = iota
	Channel1
	Channel2
	Channel3
	Channel4
	Channel5
	Channel6
	Channel7
)

func (c Channel) String() string {
	return [...]string{
		"Unknown",
		"Channel 1",
		"Channel 2",
		"Channel 3",
		"Channel 4",
		"Channel 5",
		"Channel 6",
		"Channel 7",
	}[c]
}

const MaxUsersPerChannel = 6
