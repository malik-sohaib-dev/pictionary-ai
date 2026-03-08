package types

type AllowedVisibility string

const (
	AllowedVisibilityPublic  AllowedVisibility = "public"
	AllowedVisibilityPrivate AllowedVisibility = "private"
)

func (v AllowedVisibility) Valid() bool {
	switch v {
	case AllowedVisibilityPrivate, AllowedVisibilityPublic:
		return true
	default:
		return false
	}
}

type Room struct {
	Visibility      AllowedVisibility `json:"visibility"`
	RoomId          string            `json:"roomId"`
	Owner           string            `json:"owner"`
	MaxPlayers      int               `json:"maxPlayers"`
	Hints           int               `json:"hints"`
	WordChoiceCount int               `json:"wordChoiceCount"`
	GameDuration    int               `json:"gameDuration"`
}

type CreateRoom struct {
	Visibility AllowedVisibility `json:"visibility"`
	Owner      string            `json:"owner"`
}

type CreateRoomResponse struct {
	RoomId string `json:"roomId"`
}

type PatchRoom struct {
	Visibility      *AllowedVisibility `json:"visibility,omitempty"`
	Owner           *string            `json:"owner,omitempty"`
	MaxPlayers      *int               `json:"maxPlayers,omitempty"`
	Hints           *int               `json:"hints,omitempty"`
	WordChoiceCount *int               `json:"wordChoiceCount,omitempty"`
	GameDuration    *int               `json:"gameDuration,omitempty"`
}
