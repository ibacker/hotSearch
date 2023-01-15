package model

type ZhiHuQuestion struct {
	Title    string `json:"title"`
	Url      string `json:"url"`
	Desc     string `json:"excerpt"`
	HotTag   string
	HotScore string
}

type Target struct {
	Title string `json:"title"`
	Id    int32  `json:"id"`
	Desc  string `json:"excerpt"`
}

type CardLabel struct {
	Type string `json:"type"`
}
type Item struct {
	CardLabel  CardLabel `json:"card_label"`
	Debut      bool      `json:"debut"`
	Target     Target    `json:"target"`
	DetailText string    `json:"detail_text" //热度文本
`
}

type HotList struct {
	Data []Item `json:"data"`
}
