package model

type WBHotBandItem struct {
	Num           int    `json:"num"`
	SubjectQuerys string `json:"subject_querys"`
	RawHot        int    `json:"raw_hot"`
	OnboardTime   int64  `json:"onboard_time"`
	WordScheme    string `json:"word_scheme"`
	Note          string `json:"note"`
	Category      string `json:"category"`
	Word          string `json:"word"`
	Realpos       int    `json:"realpos"`
	Rank          int    `json:"rank"`
	Url           string
	HotTag        string `json:"icon_desc"` //热搜种类
}

type WBHotGov struct {
	Note     string `json:"note"`
	IconDesc string `json:"icon_desc"`
	Word     string `json:"word"`
	Url      string `json:"url"`
	Name     string `json:"name"`
}

type WBBata struct {
	WBBandList []WBHotBandItem `json:"band_list"`
	WBHotGov   WBHotGov        `json:"hotgov"`
}

type WBHotSearch struct {
	WBDate WBBata `json:"data"`
}
