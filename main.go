package main

import (
	"gopkg.in/urfave/cli.v2"
	"httpencrypt/http"
	"os"
)

func main() {
	var port string
	var token string
	var key string
	var isClient bool
	var targetUrl string
	app := &cli.App{
		Name:    "http的rsa加密通讯中间服务",
		Version: "1.0",
		Usage:   "分为两个端进行互相的加密通讯",
		Authors: []*cli.Author{{Name: "lhn", Email: "550124023@qq.com"}},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Usage:       "本地服务器监听的段都,注意前面添加符号\":\"",
				Destination: &port,
				EnvVars:     []string{"port", "p"},
			},
			&cli.StringFlag{
				Name:        "token",
				Usage:       "两个中间件互相认证的token",
				Destination: &token,
				EnvVars:     []string{"token"},
			},
			&cli.StringFlag{
				Name:        "key",
				Usage:       "ras 密钥对中的一个",
				Destination: &key,
				EnvVars:     []string{"key"},
			},
			&cli.StringFlag{
				Name:        "targetUrl",
				Aliases:     []string{"t", "target"},
				Usage:       "目标地址",
				Destination: &targetUrl,
				EnvVars:     []string{"t", "target", "targetUrl"},
			},
			&cli.BoolFlag{
				Name:        "client",
				Usage:       "是否是客户端，如果不是客户端就是服务端",
				Destination: &isClient,
				Aliases:     []string{"c"},
				EnvVars:     []string{"client", "c"},
				Value:       false,
			},
		},
		Action: func(*cli.Context) error {
			return http.Start(port, key, token, targetUrl, isClient)
		},
	}
	app.Run(os.Args)
}
