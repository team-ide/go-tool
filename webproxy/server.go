package webproxy

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net"
	"net/http"
	"strings"
)

func Start() {

	gin.DefaultWriter = &nullWriter{}

	router := gin.Default()

	//fromHost := fmt.Sprintf("http%s%s", "://", "baidu.localhost.com")
	toHost := "https://www.baidu.com"

	router.MaxMultipartMemory = (1024 * 50) << 20 // 设置最大上传大小为50G

	routerGroup := router.Group("/")
	routerGroup.GET("*path", func(c *gin.Context) {
		path := c.Params.ByName("path")
		fmt.Println("path:", path)
		fmt.Println("path:", c.Params)
		client := &http.Client{}
		url := toHost + path
		fmt.Println("url:", url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}
		for k, v := range c.Request.Header {
			if strings.EqualFold(k, "Host") {
				req.Header[k] = []string{toHost}
			} else if strings.EqualFold(k, "Referer") {
				req.Header[k] = []string{url}
			} else {
				req.Header[k] = v
			}
		}
		for _, cook := range c.Request.Cookies() {
			req.AddCookie(cook)
		}

		fmt.Println("req.Header:", req.Header)
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		fmt.Println("res.Header:", res.Header)
		for k, vs := range res.Header {
			for _, v := range vs {
				c.Header(k, v)
			}
		}
		for _, cook := range res.Cookies() {
			c.SetCookie(cook.Name, cook.Value, cook.MaxAge, cook.Path, cook.Domain, cook.Secure, cook.HttpOnly)
		}
		bs, err := io.ReadAll(res.Body)
		if err == io.EOF {
			err = nil
		}
		if err != nil {
			panic(err)
		}
		text := string(bs)
		fmt.Println(text)
		c.Writer.Write(bs)

	})
	s := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: router,
	}
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return
	}
	go func() {
		err = s.Serve(ln)
		if err != nil {
			panic(err)
		}
	}()
}

type nullWriter struct{}

func (*nullWriter) Write(bs []byte) (int, error) {

	return 0, nil
}
