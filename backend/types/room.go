package types

type AllowedVisibility string

const (
	AllowedVisibilityPublic  AllowedVisibility = "public"
	AllowedVisibilityPrivate AllowedVisibility = "private"
)

type Room struct {
	Visibility      AllowedVisibility `json:"visibility"`
	RoomId          string            `json:"roomId"`
	Creator         string            `json:"creator"`
	MaxPlayers      int               `json:"maxPlayers"`
	Hints           int               `json:"hints"`
	WordChoiceCount int               `json:"wordChoiceCount"`
	GameDuration    int               `json:"gameDuration"`
}
