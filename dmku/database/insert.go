package database

import (
	"time"
)

func (d *connectDB) Ins_danmaku(l List) error {
	_, err := d.creat_list.Exec(l.Id, l.Typee, l.Text, l.Color, l.Size, l.Videotime, l.Ip, time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func (d *connectDB) Ins_frequency(ip Ip) error {
	_, err := d.creat_ip.Exec(ip.IP, ip.C, time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func (d *connectDB) Ins_report(re Report) error {
	_, err := d.creat_report.Exec(re.Cid, re.Id, re.Text, re.Type, time.Now().Unix(), re.Ip)
	if err != nil {
		return err
	}
	return nil
}
