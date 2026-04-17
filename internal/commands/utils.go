package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/erickir/LandOfLifeDiscordBot/internal/domain"
)

func parseStringOption(options map[string]any, name string) string {
	if value, ok := options[name]; ok {
		switch v := value.(type) {
		case string:
			return strings.TrimSpace(v)
		case fmt.Stringer:
			return strings.TrimSpace(v.String())
		}
	}
	return ""
}

func parseHourOption(options map[string]any, name string) (domain.Hour, error) {
	raw, ok := options[name]
	if !ok {
		return 0, fmt.Errorf("missing hour")
	}

	value := fmt.Sprintf("%v", raw)
	value = strings.TrimSpace(strings.ToLower(value))

	switch value {
	case "1", "01", "0100", "01:00":
		return domain.Hour0100, nil
	case "2", "03", "0300", "03:00":
		return domain.Hour0300, nil
	case "3", "05", "0500", "05:00":
		return domain.Hour0500, nil
	case "4", "07", "0700", "07:00":
		return domain.Hour0700, nil
	case "5", "09", "0900", "09:00":
		return domain.Hour0900, nil
	case "6", "1100", "11:00":
		return domain.Hour1100, nil
	case "7", "1300", "13:00":
		return domain.Hour1300, nil
	case "8", "1500", "15:00":
		return domain.Hour1500, nil
	case "9", "1700", "17:00":
		return domain.Hour1700, nil
	case "10", "1900", "19:00":
		return domain.Hour1900, nil
	case "11", "2100", "21:00":
		return domain.Hour2100, nil
	case "12", "2300", "23:00":
		return domain.Hour2300, nil
	}

	return 0, fmt.Errorf("invalid hour value %q", raw)
}

func parseChannelOption(options map[string]any, name string) (domain.Channel, error) {
	raw, ok := options[name]
	if !ok {
		return 0, fmt.Errorf("missing channel")
	}

	value := fmt.Sprintf("%v", raw)
	parsed, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		return 0, fmt.Errorf("invalid channel value %q", raw)
	}
	if parsed < 1 || parsed > 7 {
		return 0, fmt.Errorf("channel must be between 1 and 7")
	}
	return domain.Channel(parsed), nil
}
