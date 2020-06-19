package upload

import (
	"fmt"
	"github.com/cos800/bgf"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	//s.Use(func(r *ghttp.Request) {
	//	r.Exit()
	//})

	c := new(controller)
	s.BindObject("/upload", c)
}

type controller struct {
	*bgf.BaseC
}

func (c *controller) Index(r *ghttp.Request) {
	file := r.GetUploadFile("file")

	//if file.Size>1*1024*1024 {
	//
	//}

	fmt.Println(file.Size)
	fmt.Println(file.Header)
	fmt.Println(file.Filename)

	names, err := file.Save("/tmp/")
	if err != nil {
		r.Response.WriteExit(err)
	}
	r.Response.WriteExit("upload successfully: ", names)
}

// UploadShow shows uploading simgle file page.
func (c *controller) Show(r *ghttp.Request) {
	r.Response.Write(`
    <html>
    <head>
        <title>GF Upload File Demo</title>
    </head>
        <body>
            <form enctype="multipart/form-data" action="/upload" method="post">
                <input type="file" name="file" />
                <input type="submit" value="upload" />
            </form>
        </body>
    </html>
    `)
}
