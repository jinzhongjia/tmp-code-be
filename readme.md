
## 说明

一个简易的粘代码平台，支持多种语言，自动高亮、行号显示、一键复制

## 构建

win

直接运行`build-win.sh`

linux

```bash
docker run --rm -t -v d:\project\go\tmp-code:/build gobuilder:latest
```

## 部署

**image构建**

```bash
docker build -t yunyizhiying/tmp-code:latest .
```

**docker部署**

目前仅支持内存类型的cache，如果容器需要使用redis，请自行构建

```bash
docker pull yunyizhiying/tmp-code
docker run -itd  -p 127.0.0.1:8233:8233 --name tmp-code  yunyizhiying/tmp-code
```

**命令行相关参数使用-h查看**

## demo

[http://tmp-code.jinzh.me/](http://tmp-code.jinzh.me/)
