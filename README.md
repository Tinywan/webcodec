## 安装说明

* 系统需要安装 [Golang环境](http://www.cnblogs.com/tinywan/p/6928300.html) 和 `MySQL`数据库支持。

* `go get `方式安装

    ```bash
    $ go get -u github.com/Tinywan/webcodec
    ```
    
* 打开配置文件 `webcodec/conf/app.conf`，修改相关配置。

* 创建数据库 `webcodec`，导入 `webcodec.sql`文件

    ```bash
    $ mysql -u username -p -D webcodec < webcodec.sql
    ```
    
* 运行
	
    ```bash
    $ ./webcodec
    ```
    
* 设为后台运行
	
    ```bash
    $ nohup ./webcodec 2>&1 > error.log &
    ```
    
* 访问： 

    * 地址栏：[webcodec.tinywan.com](http://webcodec.tinywan.com/) 或者 `http://{$IP}:8080`

    * 帐号：`admin` 密码：`admin888`
    
* 可能会遇到的问题 
  
    * 输入账号和密码后无法登陆，请检查数据库配置是否正确。
    
    * 如果提示模板文件 `template` 不存在，请重新编译运行 ` cd webcodec/ && go build`