package model

type BDHotSearchItem struct {
	Word     string `json:"word"`
	Url      string `json:"url"`
	HotScore string `json:"hotScore"`
	Desc     string `json:"desc"`
}

type BDCards struct {
	Component          string            `json:"component"`
	BDHotSearchContent []BDHotSearchItem `json:"content"`
}

type BDData struct {
	BDCards []BDCards `json:"cards"`
}

type BDHotSearch struct {
	BDDate BDData `json:"data"`
}
