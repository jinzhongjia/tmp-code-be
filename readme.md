
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

docker

```bash
docker pull yunyizhiying/tmp-code
docker run -itd  -p 127.0.0.1:8233:8233  tmp-code
```

## demo

[http://tmp-code.jinzh.me/](http://tmp-code.jinzh.me/)
