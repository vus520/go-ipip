#install

```
sh init.sh
```


#run

```
go run main.go 8.8.8.8
```

#use

```
$ ./ipip
input a ipV4 ip address or domain name:
google.com
	216.58.221.100	 中国 香港 N/A
	61.91.161.217	 泰国 泰国 N/A
baidu.com
	220.181.57.217	 中国 北京 北京
	180.149.132.47	 中国 北京 北京
	123.125.114.144	 中国 北京 北京
	111.13.101.208	 中国 北京 北京
8.8.8.8
	8.8.8.8	 GOOGLE GOOGLE N/A
xxoo.com
	138.68.109.134	 德国 黑森州 法兰克福
	138.68.25.238	 美国 加利福尼亚州 旧金山
	139.59.27.158	 印度 卡纳塔克邦 班加罗尔

```

#build on mac

### Build

```
file="ipip"_$(uname)
go build -ldflags "-w -s" -o $file
```

### Build for Linux

```
cd $(brew info go| grep Cellar | awk '{print $1"/libexec/src"}')
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./make.bash;

cd -
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ipip_linux
```