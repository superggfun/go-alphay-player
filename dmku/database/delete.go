package database

func (d *connectDB) Del_danmaku(t string, cid int) error {
	if t == "list" {
		_, err := d.del_report.Exec(cid)
		if err != nil {
			return err
		}
		_, err = d.del_list.Exec(cid)
		if err != nil {
			return err
		}
	} else if t == "report" {
		_, err := d.del_report.Exec(cid)
		if err != nil {
			return err
		}
	}
	return nil
}
