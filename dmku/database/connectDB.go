package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type MyDB interface {
	Ins_danmaku(l List) error
	Ins_frequency(ip Ip) error
	Ins_report(re Report) error
	Qy_danmaku_pool(id string) ([]List, error)
	Qy_frequency(ip string) (Ip, error)
	Search_danmaku_pool(key string) ([]List, error)
	Show_danmaku_pool(pa, li string) ([]List, error)
	Show_report(pa, li string) ([]Report, error)
	Upd_frequency(ip string, t bool) error
	Upd_danmaku(text string, color string, cid string) error
	Del_danmaku(t string, cid int) error
}

type connectDB struct {
	db            *sql.DB
	creat_list    *sql.Stmt
	creat_ip      *sql.Stmt
	creat_report  *sql.Stmt
	update_ip     *sql.Stmt
	update_list   *sql.Stmt
	update_report *sql.Stmt
	del_report    *sql.Stmt
	del_list      *sql.Stmt
	clean_ip      *sql.Stmt
}

// 数据库初始化
func ConnectDB(dbName string) (*connectDB, error) {
	// 判断数据库存不存在，如果不存在就创建
	_, err := os.Lstat(dbName)
	if err != nil {
		fp, err := os.Create(dbName)
		if err != nil {
			return nil, err
		}
		fp.Close()
	}
	var d connectDB

	// 连接数据库
	d.db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	// 创建数据表
	err = d.addTable()
	if err != nil {
		return nil, err
	}

	// 添加规则
	err = d.addRule()
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// 创建数据表
func (d *connectDB) addTable() error {
	// 创建danmaku_list表
	_, err := d.db.Exec(danmaku_list)
	if err != nil {
		return err
	}

	// 创建danmaku_ip表
	_, err = d.db.Exec(danmaku_ip)
	if err != nil {
		return err
	}

	// 创建danmaku_report表
	_, err = d.db.Exec(danmaku_report)
	if err != nil {
		return err
	}

	return nil
}

func (d *connectDB) addRule() error {
	// 插入danmaku_list规则
	stmt, err := d.db.Prepare("INSERT INTO danmaku_list(id, type, text, color, size, videotime, ip, time) values(?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	d.creat_list = stmt

	// 插入danmaku_ip规则
	d.creat_ip, err = d.db.Prepare("INSERT INTO danmaku_ip(ip, c, time) values(?,?,?)")
	if err != nil {
		return err
	}

	// 插入danmaku_report规则
	d.creat_report, err = d.db.Prepare("INSERT INTO danmaku_report(cid, id, text, type, time, ip)values(?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	// 修改danmaku_ip规则
	d.update_ip, err = d.db.Prepare(`UPDATE danmaku_ip SET c=c+1,time=? WHERE ip = ?`)
	if err != nil {
		return err
	}

	// 清零danmaku_ip规则
	d.clean_ip, err = d.db.Prepare(`UPDATE danmaku_ip SET c=1,time=? WHERE ip = ?`)
	if err != nil {
		return err
	}

	// 修改danmaku_list规则
	d.update_list, err = d.db.Prepare(`UPDATE danmaku_list SET text=?,color=? WHERE cid=?`)
	if err != nil {
		return err
	}

	// 修改danmaku_report规则
	d.update_report, err = d.db.Prepare(`UPDATE danmaku_report SET text=? WHERE cid=?`)
	if err != nil {
		return err
	}

	// 删除danmaku_report规则
	d.del_report, err = d.db.Prepare(`DELETE FROM danmaku_report WHERE cid=?`)
	if err != nil {
		return err
	}

	// 删除danmaku_list规则
	d.del_list, err = d.db.Prepare(`DELETE FROM danmaku_list WHERE cid=?`)
	if err != nil {
		return err
	}

	return nil
}

// 关闭数据库
func (d *connectDB) DbClose() {
	d.db.Close()
}
