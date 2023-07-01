package delivery

import "encoding/json"

type Period struct {
	Day  string `json:"day"`
	From string `json:"from"`
	To   string `json:"to"`
}

type Area struct {
	Latitude  string `json:"latitude" `
	Longitude string `json:"longitude" `
}

type Request struct {
	IsActive bool     `json:"isActive" `
	Periods  []Period `json:"periods" `
	Areas    []Area   `json:"areas" `
}

type Response struct {
	IsActive bool     `json:"isActive" `
	Periods  []Period `json:"periods" `
	Areas    []Area   `json:"areas" `
}

func ParseFromEntity(data *Entity) (res *Response) {
	if data == nil {
		return
	}

	res = &Response{
		IsActive: data.IsActive,
	}
	json.Unmarshal(data.Periods, &res.Periods)
	json.Unmarshal(data.Areas, &res.Areas)

	return
}
