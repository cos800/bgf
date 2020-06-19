package bgf

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type BaseC struct {
	req *ghttp.Request
}

func (b *BaseC) Init(r *ghttp.Request) {
	b.req = r
}

func (b *BaseC) Success(data interface{}, msg string, next string) {
	b.req.Response.WriteJsonExit(g.Map{
		"ok":   1,
		"msg":  msg,
		"next": next,
		"data": data,
	})
}

func (b *BaseC) Fail(msg string, next string, data interface{}) {
	b.req.Response.WriteJsonExit(g.Map{
		"ok":   0,
		"msg":  msg,
		"next": next,
		"data": data,
	})
}
