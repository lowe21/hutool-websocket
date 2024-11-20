# websocket

## 目录结构
| 目录                                  | 描述        |
|:------------------------------------|:----------|
| config                              | 配置目录      |
| hack                                | gf工具配置目录  |
| internal                            | 内部目录      |
| &emsp;│─ api                        | 接口目录      |
| &emsp;│─ bo                         | 业务对象目录    |
| &emsp;│─ cmd                        | 命令引导目录    |
| &emsp;│─ consts                     | 常量目录      |
| &emsp;│─ controller                 | 控制器目录     |
| &emsp;│─ dao                        | 数据对象目录    |
| &emsp;│&emsp;&emsp;│─ do            | 数据模型目录    |
| &emsp;│&emsp;&emsp;│─ internal      | 数据对象内部目录  |
| &emsp;│&emsp;&emsp;│─ po            | 数据实例目录    |
| &emsp;│&emsp;&emsp;└─ tpl           | 数据对象模板目录  |
| &emsp;│─ imports                    | 引入目录      |
| &emsp;│─ logic                      | 逻辑目录      |
| &emsp;│─ middleware                 | 中间件目录     |
| &emsp;│─ pkg                        | 包目录       |
| &emsp;│&emsp;&emsp;│─ config_center | 配置中心目录    |
| &emsp;│&emsp;&emsp;│─ lbs           | 位置服务目录    |
| &emsp;│&emsp;&emsp;│─ log_center    | 日志中心目录    |
| &emsp;│&emsp;&emsp;│─ mongodb       | 文件存储数据库目录 |
| &emsp;│&emsp;&emsp;│─ rpc           | 远程调用目录    |
| &emsp;│&emsp;&emsp;│─ thirdparty    | 三方服务目录    |
| &emsp;│&emsp;&emsp;└─ tracing       | 链路追踪目录    |
| &emsp;│─ rpc                        | rpc目录     |
| &emsp;│&emsp;&emsp;│─ consumer      | rpc消费者目录  |
| &emsp;│&emsp;&emsp;└─ provider      | rpc提供者目录  |
| &emsp;│─ service                    | 服务目录      |
| &emsp;└─ util                       | 工具目录      |
| static                              | 静态文件目录    |
| template                            | 模版文件目录    |
| .gitattributes                      | git属性文件   |
| .gitignore                          | git忽略文件   |
| Dockerfile                          | 镜像文件      |
| go.mod                              | 包管理文件     |
| main.go                             | 主文件       |
| Makefile                            | 构建文件      |
| README.md                           | 自述文件      |

## 启动服务
设置 dubbo go 配置文件的环境变量

`DUBBO_GO_CONFIG_PATH=config/dubbo.yaml`

## 构建项目
`make build` 编译项目

`make image` 生成镜像

`make image-push` 生成镜像并推送到远程仓库
