# http-server

**安装**

- 在Go环境下安装它：

  ```bash
  git clone https://github.com/RanQiwu/http-server.git
  cd http-server
  go install
  ```

- 或者直接下载可执行文件：[http-server](https://github.com/RanQiwu/http-server/releases/download/v1.0.0/http-server.exe) [只编译了windows的]，然后把它的存放路径加 PATH 环境中。



**使用**

`http-server -h`

例子：

```bash
# 简单的使用，默认端口是 8080
http-server

# 改变端口，假如需要监听 9000 端口
http-server -p 9000

# 默认开启跨域CORS处理，可以手动关闭
http-server --cors=false

# 如果服务启动之后需要自动打开浏览器
http-server -o
```



