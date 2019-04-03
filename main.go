package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v2"
	"httpencrypt/http"
	"httpencrypt/rsa"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {
	var port string
	var token string
	var privateKey string
	var publicKey string
	var isClient bool
	var targetUrl string

	app := &cli.App{
		Name:    "http的rsa加密通讯中间服务",
		Version: "1.0",
		Usage:   "分为两个端进行互相的加密通讯",
		Authors: []*cli.Author{{Name: "lhn", Email: "550124023@qq.com"}},
		/*Flags: []cli.Flag{
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
				Name:        "privateKey",
				Usage:       "ras 密钥对中的一个",
				Destination: &privateKey,
				EnvVars:     []string{"privateKey"},
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
		},*/
		Commands: []*cli.Command{
			&cli.Command{
				Name:    "server",
				Aliases: []string{"s"},
				Usage:   "做服务端服务",
				Action: func(context *cli.Context) error {
					if isClient {
						fmt.Printf("启动客户端\n")
					} else {
						fmt.Printf("启动服务端\n")
					}
					var err error
					privateKey, err = checkKey(privateKey)
					if err != nil {
						return err
					}
					publicKey, err = checkKey(publicKey)
					if err != nil {
						return err
					}
					err = http.Start(port, publicKey, privateKey, token, targetUrl, isClient)
					if err != nil {
						fmt.Printf("启动失败:%s", err.Error())
					}
					return err
				},
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
						Name:        "publicKey",
						Usage:       "ras 密钥对中的对方的公钥",
						Destination: &publicKey,
						Aliases:     []string{"public"},
						EnvVars:     []string{"publickKey"},
					},
					&cli.StringFlag{
						Name:        "privateKey",
						Usage:       "ras 密钥对中的自己密钥",
						Destination: &privateKey,
						Aliases:     []string{"private"},
						EnvVars:     []string{"privateKey"},
					},
					/*&cli.StringFlag{
						Name:        "keyPath",
						Usage:       "ras 密钥对中的一个，这里写密钥路径",
						Destination: &keyPath,
						EnvVars:     []string{"keyPath"},
					},*/
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
			},
			&cli.Command{
				Name:    "genkey",
				Aliases: []string{"g"},
				Usage:   "生成密钥对",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "size",
						Aliases: []string{"s"},
						Usage:   "密钥大小 8的2次方",
						Value:   2048,
					},
					&cli.StringFlag{
						Name:    "path",
						Aliases: []string{"p"},
						Usage:   "保存的文件夹,会在该文件夹下面生成rsa.public和rsa.private文件,默认保存到当前目录",
						Value:   "./",
					},
				},
				Action: func(context *cli.Context) error {
					size := context.Int("size")
					ph := context.String("path")
					fmt.Printf("生成密钥大小%d,位置%s\n", size, path.Join(ph))
					err := rsa.Genkey(ph, size)
					if err == nil {
						fmt.Printf("生成密钥成功，位置：%s\n", path.Join(ph))
					} else {
						fmt.Printf("生成密钥失败,%s\n", err.Error())
					}
					return err
				},
			},
		},
	}
	app.Run(os.Args)
}
func checkKey(key string) (string, error) {
	if strings.HasPrefix(key, "-----") {
		return key, nil
	}
	bys, err := ioutil.ReadFile(key)
	if err != nil {
		fmt.Printf("读取密钥失败:%s\n", err.Error())
		return "", err
	}
	return string(bys), nil
}
