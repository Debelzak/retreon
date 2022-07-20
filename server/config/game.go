package config

import (
	"fmt"
)

func GetGameMode() string {
	return settings.Mode
}

func GetRefresh() int {
	return settings.Refresh
}

func SetGameMode(mode string) error {
	if mode != "DateEarned" && mode != "DateEarnedHardcore" {
		return fmt.Errorf("mode '%s' is not supported", mode)
	}
	settings.Mode = mode
	return save()
}

func SetRefresh(refresh int) error {
	if refresh < 0 || refresh > 6 {
		return fmt.Errorf("refresh '%d' is not within bounds 0-6", refresh)
	}
	settings.Refresh = refresh
	return save()
}

func SetGameOrder(game string, order []string) error {
	if settings.Orders == nil {
		settings.Orders = make(map[string][]string, len(order))
	}

	settings.Orders[game] = order
	return save()
}
