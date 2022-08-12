package admin

type dataForm struct {
	Danmuon       string `form:"danmuon"`
	Color         string `form:"color"`
	Logo          string `form:"logo"`
	Trytime       string `form:"trytime"`
	Waittime      string `form:"waittime"`
	Sendtime      string `form:"sendtime"`
	Dmrule        string `form:"dmrule"`
	Pbgjz         string `form:"pbgjz"`
	AdsState      string `form:"adsState"`
	AdsSetState   string `form:"adsSetState"`
	AdsSetGroup   string `form:"adsSetGroup"`
	AdsSetPicTime string `form:"adsSetPicTime"`
	AdsSetPicImg  string `form:"adsSetPicImg"`
	AdsSetPicLink string `form:"adsSetPicLink"`
	AdsSetVodUrl  string `form:"adsSetVodUrl"`
	AdsSetVodLink string `form:"adsSetVodLink"`
	AdsPauseState string `form:"adsPauseState"`
	AdsPausePic   string `form:"adsPausePic"`
	AdsPauseLink  string `form:"adsPauseLink"`
}
