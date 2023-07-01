package schedule

import "encoding/json"

type Period struct {
	Day  string `json:"day"`
	From string `json:"from"`
	To   string `json:"to"`
}

type Request struct {
	IsActive bool     `json:"isActive" `
	Periods  []Period `json:"periods" `
}

type Response struct {
	IsActive bool     `json:"isActive" `
	Periods  []Period `json:"periods" `
}

func ParseFromEntity(data *Entity) (res *Response) {
	if data == nil {
		return
	}

	res = &Response{
		IsActive: data.IsActive,
	}
	json.Unmarshal(data.Periods, &res.Periods)

	return
}
