package database

import (
	"fmt"
	"strconv"
)

// 查询弹幕池
func (d *connectDB) Qy_danmaku_pool(id string) ([]List, error) {
	rows, err := d.db.Query(`SELECT * FROM danmaku_list WHERE id='` + id + `' ORDER BY videotime asc`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var s []List
	for rows.Next() {
		var l List
		err = rows.Scan(&l.Id, &l.Cid, &l.Typee, &l.Text, &l.Color, &l.Size, &l.Videotime, &l.Ip, &l.Time)
		if err != nil {
			return nil, err
		}
		s = append(s, l)
	}
	return s, nil
}

// 查询弹幕频率
func (d *connectDB) Qy_frequency(ip string) (Ip, error) {
	var i Ip
	rows, err := d.db.Query(`SELECT * FROM danmaku_ip WHERE ip= '` + ip + `' LIMIT 1`)
	if err != nil {
		return i, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&i.IP, &i.C, &i.Time)
		if err != nil {
			return i, err
		}
	}
	return i, nil
}

// 搜索弹幕池
func (d *connectDB) Search_danmaku_pool(key string) ([]List, error) {
	rows, err := d.db.Query(`SELECT * FROM danmaku_list WHERE text like '` + key + `' or id like '` + key + `' ORDER BY time DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var s []List
	for rows.Next() {
		var l List
		err = rows.Scan(&l.Id, &l.Cid, &l.Typee, &l.Text, &l.Color, &l.Size, &l.Videotime, &l.Ip, &l.Time)
		if err != nil {
			return nil, err
		}
		s = append(s, l)
	}
	return s, nil
}

// 显示弹幕池
func (d *connectDB) Show_danmaku_pool(pa, li string) ([]List, error) {
	page, err := strconv.Atoi(pa)
	if err != nil {
		return nil, err
	}
	limit, err := strconv.Atoi(li)
	if err != nil {
		return nil, err
	}
	index := (page - 1) * limit
	rows, err := d.db.Query(fmt.Sprintf("SELECT * FROM danmaku_list ORDER BY time DESC limit %v,%v", index, limit))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var s []List
	for rows.Next() {
		var l List
		err = rows.Scan(&l.Id, &l.Cid, &l.Typee, &l.Text, &l.Color, &l.Size, &l.Videotime, &l.Ip, &l.Time)
		if err != nil {
			return nil, err
		}
		s = append(s, l)
	}
	return s, nil

}

// 显示举报列表
func (d *connectDB) Show_report(pa, li string) ([]Report, error) {
	page, err := strconv.Atoi(pa)
	if err != nil {
		return nil, err
	}
	limit, err := strconv.Atoi(li)
	if err != nil {
		return nil, err
	}
	index := (page - 1) * limit
	rows, err := d.db.Query(fmt.Sprintf("SELECT * FROM danmaku_report ORDER BY time DESC limit %v,%v", index, limit))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var s []Report
	for rows.Next() {
		var l Report
		err = rows.Scan(&l.Cid, &l.Id, &l.Text, &l.Type, &l.Time, &l.Ip)
		if err != nil {
			return nil, err
		}
		s = append(s, l)
	}
	return s, nil
}
