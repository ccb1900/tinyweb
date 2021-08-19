package api

import "github.com/gin-gonic/gin"

type PageMeta struct {
	Total uint `json:"total"`
}
type Pagination struct {
	Data interface{} `json:"data"`
	Meta PageMeta `json:"meta"`
}

func Page(ctx *gin.Context,data interface{},total uint)  {
	Success(ctx,Pagination{
		Data: data,
		Meta: PageMeta{
			total,
		},
	})
}
