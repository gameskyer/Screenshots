package util

import (
	"Screenshots/module"
	"fmt"
)

func MarkData(data []module.InsertSql) []module.Router {
	names := SelectGroupByName()
	var routers []module.Router
	for _, v := range names {
		routers = append(routers, SelectByName(v))
	}
	fmt.Println(routers)
	return routers
}
