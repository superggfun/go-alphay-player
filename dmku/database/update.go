package database

import "time"

func (d *connectDB) Upd_frequency(ip string, t bool) error {
	if t {
		_, err := d.clean_ip.Exec(time.Now().Unix(), ip)
		if err != nil {
			return err
		}
	} else {
		_, err := d.update_ip.Exec(time.Now().Unix(), ip)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *connectDB) Upd_danmaku(text string, color string, cid string) error {
	_, err := d.update_list.Exec(text, color, cid)
	if err != nil {
		return err
	}

	_, err = d.update_report.Exec(text, color, cid)
	if err != nil {
		return err
	}
	return nil
}
