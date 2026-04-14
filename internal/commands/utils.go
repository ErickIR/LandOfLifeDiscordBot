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
	case "3", "03", "0300", "03:00":
		return domain.Hour0300, nil
	case "5", "05", "0500", "05:00":
		return domain.Hour0500, nil
	case "7", "07", "0700", "07:00":
		return domain.Hour0700, nil
	case "9", "09", "0900", "09:00":
		return domain.Hour0900, nil
	case "11", "1100", "11:00":
		return domain.Hour1100, nil
	case "13", "1300", "13:00":
		return domain.Hour1300, nil
	case "15", "1500", "15:00":
		return domain.Hour1500, nil
	case "17", "1700", "17:00":
		return domain.Hour1700, nil
	case "19", "1900", "19:00":
		return domain.Hour1900, nil
	case "21", "2100", "21:00":
		return domain.Hour2100, nil
	case "23", "2300", "23:00":
		return domain.Hour2300, nil
	}

	parsed, err := strconv.Atoi(value)
	if err == nil {
		switch parsed {
		case 1:
			return domain.Hour0100, nil
		case 3:
			return domain.Hour0300, nil
		case 5:
			return domain.Hour0500, nil
		case 7:
			return domain.Hour0700, nil
		case 9:
			return domain.Hour0900, nil
		case 11:
			return domain.Hour1100, nil
		case 13:
			return domain.Hour1300, nil
		case 15:
			return domain.Hour1500, nil
		case 17:
			return domain.Hour1700, nil
		case 19:
			return domain.Hour1900, nil
		case 21:
			return domain.Hour2100, nil
		case 23:
			return domain.Hour2300, nil
		}
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

func parseOptionalChannelOption(options map[string]any, name string) (domain.Channel, error) {
	_, ok := options[name]
	if !ok {
		return domain.UnknownChannel, nil
	}

	if parseStringOption(options, name) == "" {
		return domain.UnknownChannel, nil
	}

	return parseChannelOption(options, name)
}
