package player

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sentJson struct {
	Code int                    `json:"code"`
	Data map[string]interface{} `json:"data"`
}

func Api(c *gin.Context) {
	f, err := ioutil.ReadFile("admin.json")
	if err != nil {
		fmt.Print(err)
	}

	var d = make(map[string]interface{})
	err = json.Unmarshal(f, &d)
	if err != nil {
		fmt.Println(err)
	}
	c.AsciiJSON(http.StatusOK, sentJson{
		Code: 1,
		Data: d,
	})
}
