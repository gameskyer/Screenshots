package util

import (
	"Screenshots/been"
	"fmt"
)

func MarkData(data []been.InsertSql)[]been.Router{
	names := SelectGroupByName()
	var routers []been.Router
	for _,v := range names{
		routers = append(routers,SelectByName(v))
	}
	fmt.Println(routers)
	return routers
}
