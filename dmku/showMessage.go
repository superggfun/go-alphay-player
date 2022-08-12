package dmku

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"net"
	"net/http"
	"player/config"
	"player/dmku/database"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var m database.MyDB

func init() {
	var err error
	m, err = database.ConnectDB("ayplayer.db")
	if err != nil {
		panic(err)
	}
}

func ShowMessage(c *gin.Context) {
	ac, _ := c.GetQuery("ac")
	id, _ := c.GetQuery("id")
	switch ac {
	case "report":
		cid, _ := c.GetQuery("cid")
		ncid, _ := strconv.Atoi(cid)
		user, _ := c.GetQuery("user")
		nuser, _ := strconv.ParseInt(user, 10, 64)
		nip := 4294967295 - nuser
		str := fmt.Sprintf("%d.%d.%d.%d", byte(nip>>24), byte(nip>>16), byte(nip>>8), byte(nip))
		ty, _ := c.GetQuery("type")
		title, _ := c.GetQuery("title")
		text, _ := c.GetQuery("text")
		t, _ := c.GetQuery("_")
		nt, _ := strconv.ParseInt(t, 10, 64)
		err := m.Ins_report(database.Report{ncid, title, text, ty, nt, str})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -3,
				"name":    "report",
				"danum":   0,
				"danmuku": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    6666,
				"name":    "report",
				"danum":   0,
				"danmuku": "举报成功！感谢您为守护弹幕作出了贡献",
			})
		}

	case "get":
		l, err := m.Qy_danmaku_pool(id)
		if err != nil {
			fmt.Println(err)
		}
		var mes [][]string
		for _, v := range l {
			var word = []string{fmt.Sprintf("%f", v.Videotime), v.Typee, v.Color, strconv.Itoa(v.Cid), v.Text, enip(c.ClientIP()), time.Unix(v.Time, 0).Format("01-02 15:04"), v.Size}
			mes = append(mes, word)
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    23,
			"name":    id,
			"danum":   len(l),
			"danmuku": mes,
		})

	case "so":
		key, _ := c.GetQuery("key")
		l, err := m.Search_danmaku_pool(key)
		if err != nil {
			fmt.Println(err)
		}
		var m [][9]string
		for _, v := range l {
			m = append(m, [9]string{v.Id, fmt.Sprintf("%f", v.Videotime), v.Typee, v.Color, strconv.Itoa(v.Cid), v.Text, v.Ip, time.Unix(v.Time, 0).Format("01-02 15:04"), v.Size})
		}
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"count": len(l),
			"data":  m,
		})

	case "list":
		page, _ := c.GetQuery("page")
		limit, _ := c.GetQuery("limit")
		l, err := m.Show_danmaku_pool(page, limit)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  -1,
				"count": 0,
				"data":  err,
			})
		}
		var m [][9]string
		for _, v := range l {
			m = append(m, [9]string{v.Id, fmt.Sprintf("%f", v.Videotime), v.Typee, v.Color, strconv.Itoa(v.Cid), v.Text, v.Ip, time.Unix(v.Time, 0).Format("01-02 15:04"), v.Size})

		}
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"count": len(l),
			"data":  m,
		})
	case "reportlist":
		page, _ := c.GetQuery("page")
		limit, _ := c.GetQuery("limit")
		l, err := m.Show_report(page, limit)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  -1,
				"count": 0,
				"data":  err,
			})

		} else {
			var m [][6]string
			for _, v := range l {
				m = append(m, [6]string{v.Id, v.Type, strconv.Itoa(v.Cid), v.Text, v.Ip, time.Unix(v.Time, 0).Format("01-02 15:04")})
			}
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"count": len(l),
				"data":  m,
			})
		}

	case "dm":
		l, err := m.Qy_danmaku_pool(id)
		if err != nil {
			fmt.Println(err)
		}
		n := len(l)
		var mes [][]string
		for _, v := range l {
			var word = []string{fmt.Sprintf("%f", v.Videotime), v.Typee, v.Color, strconv.Itoa(v.Cid), v.Text, enip(c.ClientIP()), time.Unix(v.Time, 0).Format("01-02 15:04"), v.Size}
			mes = append(mes, word)
		}
		var mov string
		if n == 0 {
			mov = "一条弹幕都没有，赶紧来一发吧！"
		} else {
			mov = fmt.Sprintf("有 %v 条弹幕列队来袭~做好准备吧！", n)
		}
		config, err := config.Read()
		if err != nil {
			panic(err)
		}
		var tips = []string{"2", "right", "#fff", "", mov}
		var tips1 = []string{config.Tip.Time, "top", config.Tip.Color, "", config.Tip.Text}

		mes = append(mes, tips, tips1)
		c.JSON(http.StatusOK, gin.H{
			"code":    23,
			"name":    id,
			"danum":   n,
			"danmuku": mes,
		})

	case "del":
		ty, _ := c.GetQuery("type")
		i, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"danmuku": err,
			})
		} else {
			m.Del_danmaku(ty, i)
			c.JSON(http.StatusOK, gin.H{
				"code":    0,
				"danmuku": "true",
			})
		}

	default:

	}
}

func get_hash() string {
	rand.Seed(time.Now().Unix())
	h := md5.New()
	h.Write([]byte(strconv.Itoa(rand.Intn(999899999999) + 100000000)))
	h.Write([]byte(hex.EncodeToString(h.Sum(nil)) + "#"))
	return hex.EncodeToString(h.Sum(nil))
}

func enip(ip string) string {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return strconv.Itoa(4294967295 - int(ret.Int64()))
}
