package admin

type data struct {
	Danmuon  string `json:"danmuon"`
	Color    string `json:"color"`
	Logo     string `json:"logo"`
	Trytime  string `json:"trytime"`
	Waittime string `json:"waittime"`
	Sendtime string `json:"sendtime"`
	Dmrule   string `json:"dmrule"`
	Pbgjz    string `json:"pbgjz"`
	Ads      ads    `json:"ads"`
}

type ads struct {
	State string `json:"state"`
	Set   set    `json:"set"`
	Pause pause  `json:"pause"`
}

type set struct {
	State string `json:"state"`
	Group string `json:"group"`
	Pic   pic    `json:"pic"`
	Vod   vod    `json:"vod"`
}

type pic struct {
	Time string `json:"time"`
	Img  string `json:"img"`
	Link string `json:"link"`
}

type vod struct {
	Url  string `json:"url"`
	Link string `json:"link"`
}

type pause struct {
	State string `json:"state"`
	Pic   string `json:"pic"`
	Link  string `json:"link"`
}
