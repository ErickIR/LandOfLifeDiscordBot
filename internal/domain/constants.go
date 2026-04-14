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
		"01:00", "03:00", "05:00",
		"07:00", "09:00", "11:00",
		"13:00", "15:00", "17:00",
		"19:00", "21:00", "23:00",
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
