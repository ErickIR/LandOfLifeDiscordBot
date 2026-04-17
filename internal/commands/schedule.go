package commands

import (
	"fmt"
	"strings"
	"time"
)

const (
	totalSlots     = 12
	firstStartHour = 18 // 6:00 PM
)

func ts(t time.Time) string {
	return fmt.Sprintf("<t:%d:t>", t.Unix())
}

func buildLoLScheduleMessage(day time.Time, loc *time.Location) string {
	localDay := day.In(loc)
	firstStart := time.Date(
		localDay.Year(),
		localDay.Month(),
		localDay.Day(),
		firstStartHour, 0, 0, 0,
		loc,
	)

	var b strings.Builder
	b.WriteString("**Land of Life Schedule**\n")
	b.WriteString("────────────────────────\n\n")

	for i := 0; i < totalSlots; i++ {
		start := firstStart.Add(time.Duration(i*2) * time.Hour)
		asgobas := start.Add(1 * time.Hour)
		end := start.Add(2 * time.Hour)

		b.WriteString(fmt.Sprintf("**%s LoL**\n", ordinal(i+1)))
		b.WriteString(fmt.Sprintf("Start: %s\n", ts(start)))
		b.WriteString(fmt.Sprintf("Asgobas: %s\n", ts(asgobas)))
		b.WriteString(fmt.Sprintf("End: %s\n", ts(end)))

		if i != totalSlots-1 {
			b.WriteString("\n")
		}
	}

	b.WriteString("\n────────────────────────\n")
	b.WriteString("Times are displayed in your local timezone\n")
	b.WriteString("LoL is available on all channels")

	return b.String()
}

func ordinal(n int) string {
	if n%100 >= 11 && n%100 <= 13 {
		return fmt.Sprintf("%dth", n)
	}

	switch n % 10 {
	case 1:
		return fmt.Sprintf("%dst", n)
	case 2:
		return fmt.Sprintf("%dnd", n)
	case 3:
		return fmt.Sprintf("%drd", n)
	default:
		return fmt.Sprintf("%dth", n)
	}
}
