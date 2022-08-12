package database

type List struct {
	Id        string
	Cid       int
	Typee     string
	Text      string
	Color     string
	Size      string
	Videotime float32
	Ip        string
	Time      int64
}

type Ip struct {
	IP   string
	C    int8
	Time int64
}

type Report struct {
	Cid  int
	Id   string
	Text string
	Type string
	Time int64
	Ip   string
}
