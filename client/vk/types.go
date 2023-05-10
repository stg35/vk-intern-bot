package vk

type GetLongPollServerResponse struct {
	Data struct {
		Server string `json:"server"`
		Key    string `json:"key"`
		Ts     string `json:"ts"`
	} `json:"response"`
}

type Update struct {
	Ts     int `json:"ts"`
	Object struct {
		Message struct {
			Text string `json:"text"`
		} `json:"message"`
	} `json:"object"`
}

type GetUpdatesResponse struct {
	Updates []Update `json:"updates"`
}
