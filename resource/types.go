package resource

type Source struct {
	Id    string  `json:"id"`
	Token *string `json:"token,omitempty"`
}

type Version map[string]string
