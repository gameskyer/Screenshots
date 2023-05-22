package service

import (
	"Screenshots/dao"
	"Screenshots/module"
	"fmt"
)

func MarkData(data []module.Screenshot) []module.Router {
	names := dao.SelectGroupByName()
	var routers []module.Router
	for _, v := range names {
		routers = append(routers, dao.SelectByName(v))
	}
	fmt.Println(routers)
	return routers
}
