package server

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/0w0mewo/budong/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	routes     *gin.Engine
	httpServer *http.Server
	service    Service
	logger     *logrus.Entry
}

func NewRestfulServer(addr string, svr Service) *ApiServer {
	routes := gin.New()
	logger := logrus.StandardLogger().WithField("server", "rest")

	return &ApiServer{
		routes:  routes,
		service: svr,
		httpServer: &http.Server{
			Addr:    addr,
			Handler: routes,
		},
		logger: logger,
	}
}

func (r *ApiServer) Init() {
	r.routes.GET("/", r.hello)
	r.routes.GET("dofetch/:num", r.fetchsetu)
	r.routes.GET("inventory/:page/:page_size", r.inventory)
	r.routes.GET("/shetu/:id", r.givemesetu)

}

func (r *ApiServer) Run() {
	r.httpServer.ListenAndServe()
}

func (r *ApiServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	r.httpServer.Shutdown(ctx)

	r.logger.Infoln("rest server shutdown")
}

func (r *ApiServer) givemesetu(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Resp{ErrMsg: err.Error()})
		return
	}

	se, err := r.service.GetSetuFromDB(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Resp{ErrMsg: err.Error()})
		return
	}

	imgType := utils.ImageBytesFmt(se)

	c.Data(http.StatusOK, "image/"+imgType, se)
}

func (r *ApiServer) inventory(c *gin.Context) {
	resp := &Resp{}

	// get page param
	page, err := strconv.ParseUint(c.Param("page"), 10, 64)
	if err != nil {
		resp.ErrMsg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// get page size param
	size, err := strconv.ParseUint(c.Param("page_size"), 10, 64)
	if err != nil {
		resp.ErrMsg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// ensure page size is between 0 and 100
	if size > 50 || size < 1 {
		resp.ErrMsg = "invalid page size"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// ensure page is in valid range
	if page < 1 || page > r.service.Count()/size+1 {
		resp.ErrMsg = "invalid page number"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	data, err := r.service.GetInventory(page, size)
	if err != nil {
		resp.ErrMsg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.ErrMsg = "ok"
	resp.Infos = data
	resp.Count = len(data)
	c.JSON(http.StatusOK, resp)
}

func (r *ApiServer) hello(c *gin.Context) {
	img, err := r.service.RandomSetu()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Resp{ErrMsg: err.Error()})
		return
	}

	imgType := utils.ImageBytesFmt(img)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Resp{ErrMsg: err.Error()})
		return
	}

	c.Data(http.StatusOK, "image/"+imgType, img)
}

func (r *ApiServer) fetchsetu(c *gin.Context) {
	num, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Resp{ErrMsg: err.Error()})
		return
	}

	err = r.service.RequestSetu(num, false) // 不可以涩色
	if err != nil {
		c.JSON(http.StatusOK, &Resp{ErrMsg: err.Error(), Infos: nil, Count: num})
		return
	}

	c.JSON(http.StatusOK, &Resp{ErrMsg: "ok", Infos: nil, Count: num})

}
