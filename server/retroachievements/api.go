package retroachievements

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"retro/config"
	"sort"
	"strconv"

	"go.uber.org/zap"
)

const (
	baseURL = "https://retroachievements.org/API"
)

type Game struct {
	ID           string                 `json:"id"`
	Title        string                 `json:"title"`
	Icon         string                 `json:"icon"`
	TitleScreen  string                 `json:"titlescreen"`
	InGame       string                 `json:"ingame"`
	BoxArt       string                 `json:"boxart"`
	Console      string                 `json:"console"`
	Softcore     string                 `json:"softcore"`
	Hardcore     string                 `json:"hardcore"`
	Achievements map[string]Achievement `json:"achievements"`
	Order        []string               `json:"order"`
	DefaultOrder []string               `json:"defaultOrder"`
}

type Achievement struct {
	ID                 string `json:"ID"`
	Badge              string `json:"BadgeName"`
	Title              string `json:"Title"`
	Description        string `json:"Description"`
	Softcore           string `json:"NumAwarded"`
	Hardcore           string `json:"NumAwardedHardcore"`
	Points             string `json:"Points"`
	TrueRatio          string `json:"TrueRatio"`
	DisplayOrder       string `json:"DisplayOrder"`
	DateEarned         string `json:"DateEarned,omitempty"`
	DateEarnedHardcore string `json:"DateEarnedHardcore,omitempty"`
}

func GetLastGame(l *zap.Logger, username, apikey string) (string, error) {
	url := fmt.Sprintf(
		"%s/API_GetUserRecentlyPlayedGames.php?z=%s&y=%s&u=%s&c=1",
		baseURL,
		username,
		apikey,
		username,
	)
	response, err := http.Get(url)
	if code := response.StatusCode; err != nil || code < 200 || code >= 300 {
		l.Error("Error getting recently played games",
			zap.String("username", username),
			zap.Int("code", code),
			zap.Error(err),
		)
		return "", err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		l.Error("Error reading response of recently played games",
			zap.String("username", username),
			zap.Error(err),
		)
		return "", err
	}

	var container []struct {
		GameID string `json:"GameID"`
	}
	if err := json.Unmarshal(responseBody, &container); err != nil {
		l.Error("Error parsing payload of recently played games",
			zap.String("username", username),
			zap.Error(err),
		)
		return "", err
	}

	return container[0].GameID, nil
}

func GetGameInformation(l *zap.Logger, username, apikey string) (Game, error) {
	gameid, err := GetLastGame(l, username, apikey)
	if err != nil {
		l.Error("Error getting last game for user",
			zap.String("username", username),
			zap.Error(err),
		)
		return Game{}, err
	}

	url := fmt.Sprintf(
		"%s/API_GetGameInfoAndUserProgress.php?z=%s&y=%s&u=%s&g=%s",
		baseURL,
		username,
		apikey,
		username,
		gameid,
	)
	response, err := http.Get(url)
	if code := response.StatusCode; err != nil || code < 200 || code >= 300 {
		l.Error("Error getting game information for user",
			zap.String("username", username),
			zap.Int("code", code),
			zap.Error(err),
		)
		return Game{}, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		l.Error("Error reading game information for user",
			zap.String("username", username),
			zap.Error(err),
		)
		return Game{}, err
	}

	var container struct {
		Title        string                 `json:"Title"`
		Console      string                 `json:"ConsoleName"`
		Icon         string                 `json:"ImageIcon"`
		Titlescreen  string                 `json:"ImageTitle"`
		InGame       string                 `json:"ImageIngame"`
		BoxArt       string                 `json:"ImageBoxArt"`
		Softcore     string                 `json:"NumDistinctPlayersCasual"`
		Hardcore     string                 `json:"NumDistinctPlayersHardcore"`
		Achievements map[string]Achievement `json:"Achievements,omitempty"`
	}
	if err := json.Unmarshal(responseBody, &container); err != nil {
		l.Error("Error parsing game information for user",
			zap.String("username", username),
			zap.Error(err),
		)
		return Game{}, err
	}
	result := Game{
		ID:           gameid,
		Title:        container.Title,
		Console:      container.Console,
		Icon:         container.Icon,
		TitleScreen:  container.Titlescreen,
		InGame:       container.InGame,
		BoxArt:       container.BoxArt,
		Softcore:     container.Softcore,
		Hardcore:     container.Hardcore,
		Achievements: container.Achievements,
	}

	//Set default order
	type order struct {
		ID           int
		DisplayOrder int
	}

	var defaultOrder []order

	for _, achievement := range result.Achievements {
		thisID, _ := strconv.Atoi(achievement.ID)
		thisDisplayOrder, _ := strconv.Atoi(achievement.DisplayOrder)

		thisOrder := order{
			ID:           thisID,
			DisplayOrder: thisDisplayOrder,
		}

		defaultOrder = append(defaultOrder, thisOrder)
	}

	//first sort by ID
	sort.SliceStable(defaultOrder, func(i, j int) bool {
		return defaultOrder[i].ID < defaultOrder[j].ID
	})
	//then sort by display order
	sort.SliceStable(defaultOrder, func(i, j int) bool {
		return defaultOrder[i].DisplayOrder < defaultOrder[j].DisplayOrder
	})
	//save
	for _, defOrder := range defaultOrder {
		result.DefaultOrder = append(result.DefaultOrder, strconv.Itoa(defOrder.ID))
	}

	// look for a saved game order for achievements
	if order, ok := config.GetSettings().Orders[result.ID]; ok {
		result.Order = order
	} else {
		// no saved order, sort in display order
		result.Order = result.DefaultOrder
	}
	return result, nil
}
