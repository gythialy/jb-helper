package bootstrap

import (
	"log"

	"github.com/fasthttp/router"
	"github.com/gythialy/jb-helper/internal/cert"
	"github.com/gythialy/jb-helper/internal/constant"
	"github.com/gythialy/jb-helper/internal/power"
	"github.com/gythialy/jb-helper/internal/product"
	router2 "github.com/gythialy/jb-helper/internal/router"
	"github.com/gythialy/jb-helper/internal/utils"
)

func Run() *router.Router {
	log.Printf("jb-helper %s @ %s\n", constant.Version, constant.BuildTime)
	log.Println("正在创建相关目录")
	err := createDirs()
	if err != nil {
		panic(err)
	}
	if utils.IsDirEmpty(utils.CertPath()) {
		log.Println("正在生成证书")
		err = cert.CreateCert()
	}
	if err != nil {
		panic(err)
	}
	log.Println("正在生成power.conf")
	err = power.GenerateEqualResult()
	if err != nil {
		panic(err)
	}
	log.Println("正在同步Jetbrains产品信息")
	err = product.SyncProduct()
	if err != nil {
		panic(err)
	}
	log.Println("正在启动服务")
	return router2.NewRouter()
}

// 创建目录
func createDirs() error {
	err := utils.CreateDir(utils.CertPath())
	if err != nil {
		return err
	}
	err = utils.CreateDir(utils.JsonPath())
	if err != nil {
		return err
	}
	err = utils.CreateDir(utils.ConfigPath())
	if err != nil {
		return err
	}
	return nil
}
