export version=${VERSION:=1.0}
bash build.sh
docker build -t registry.cn-shenzhen.aliyuncs.com/upa/httpencrypt:${version} .
docker push registry.cn-shenzhen.aliyuncs.com/upa/httpencrypt:${version}

