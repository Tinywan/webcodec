## 安装说明

* 系统需要安装 `Go` 环境和 `MySQL`数据库支持。

* git 方式安装

    ```bash
    $ cd $GOPATH/src
    $ git clone https://github.com/Tinywan/webcodec.git
    $ cd webcodec/
    ```
    
* 打开配置文件 `webcodec/conf/app.conf`，修改相关配置。

* 创建数据库 `webcodec`，再导入install.sql

    ```bash
    $ mysql -u username -p -D webcron < install.sql
    ```
    
* 使用 `bee run` 编译命令源码文件和运行对应的可执行文件以及依赖包的安装
    
    ```bash
    | ___ \
    | |_/ /  ___   ___
    | ___ \ / _ \ / _ \
    | |_/ /|  __/|  __/
    \____/  \___| \___| v1.9.1
    2018/03/13 17:27:25 INFO     ▶ 0001 Using 'webcodec' as 'appname'
    2018/03/13 17:27:25 INFO     ▶ 0002 Initializing watcher...
    github.com/astaxie/beego/config
    github.com/astaxie/beego/utils
    github.com/astaxie/beego/session
    github.com/astaxie/beego/context
    github.com/astaxie/beego/logs
    github.com/astaxie/beego/context/param
    github.com/astaxie/beego/grace
    github.com/astaxie/beego/toolbox
    github.com/astaxie/beego
    webcodec/libs
    github.com/astaxie/beego/orm
    github.com/go-sql-driver/mysql
    webcodec/models
    webcodec/controllers
    webcodec
    2018/03/13 17:27:32 SUCCESS  ▶ 0003 Built Successfully!
    2018/03/13 17:27:32 INFO     ▶ 0004 Restarting 'webcodec'...
    2018/03/13 17:27:32 SUCCESS  ▶ 0005 './webcodec' is running...
    [ORM]2018/03/13 17:27:32 register db Ping `default`, dial tcp 127.0.0.1:3306: getsockopt: connection refused
    2018/03/13 17:27:32.265 [I] [asm_amd64.s:2337] http server Running on http://:8080
    ```    
    
*   编译    

    ```bash
    $ cd webcron
    $ go build // 生成可执行文件 webcodec
    ```
    
*   运行
	
    ```bash
    $ ./webcodec
    //或
    $ nohup ./webcodec 2>&1 > error.log &
    //设为后台运行
    ```

*   访问： 

    * 地址栏：`http://IP:8080`

    * 帐号：`admin` 密码：`admin888`
    
    * 输入账号和密码后无法登陆，请检查数据库配置是否正确。