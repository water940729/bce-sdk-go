# CFC GO SDK文档

# 初始化

## 确认Endpoint

目前支持“华北-北京”、“华南-广州” 两个区域。北京区域：`http://cfc.bj.baidubce.com`，广州区域：`http://cfc.gz.baidubce.com` 对应信息为：

访问区域 | 对应Endpoint
---|---
BJ | cfc.bj.baidubce.com
GZ | cfc.gz.baidubce.com

## 获取密钥

要使用百度云BOS，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问BOS做签名验证。

可以通过如下步骤获得并了解您的AK/SK信息：

[注册百度云账号](https://login.bce.baidu.com/reg.html?tpl=bceplat&from=portal)

[创建AK/SK](https://console.bce.baidu.com/iam/?_=1513940574695#/iam/accesslist)

## CFC Client

CFC Client是CFC服务的客户端，为开发者与CFC服务进行交互提供了一系列的方法。

### 使用AK/SK新建CFC Client

通过AK/SK方式访问CFC，用户可以参考如下代码新建一个CFC Client：

```go
import (
	"github.com/baidubce/bce-sdk-go/services/cfc"
)

func main() {
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := <your-access-key-id>, <your-secret-access-key>

	// 用户指定的Endpoint
	ENDPOINT := <domain-name>

	// 初始化一个cfcClient
	cfcClient, err := cfc.NewClient(AK, SK, ENDPOINT)
}
```

> **注意：**`ENDPOINT`参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为`http://cfc.bj.baidubce.com`。

## CreateFunction 创建函数

```go
    zipFile, err := api.ReadBase64FromFile("your zip file path")
	args := &api.FunctionArgs{
		Code:         &api.CodeFile{ZipFile: zipFile},
		Publish:      false,
		FunctionName: FunctionName,
		Handler:      "index.handler",
		Runtime:      "python2",
		MemorySize:   129,
		Timeout:      3,
		Description:  "Description",
	}
	res, err := cfcClient.CreateFunction(args)
	if err != nil {
		panic(err)
	}
```

参数名称 | 类型 | 是否必需  | 描述
--- | --- | --- | ---
Code |Code |是|   Code信息
Description |String |否|  一个简短的说明 0-256字符
Environment |Environment |否|  环境变量
FunctionName |String |是|  您想要分配给您正在上传的函数的名称。注意，长度限制只适用于BRN。如果只指定函数名，则长度限制为64个字符。
Handler |String |是|   cfc调用的入口函数，对于node为module-name.export eg. index.handler 最大长度128
MemorySize|int|否|   内存的大小，以MB为单位，CFC使用此内存大小来推断分配给您的函数的CPU和内存数量。默认值是128mb，必须是64MB的倍数。
Runtime |String |是|  运行语言 python2 nodejs6.11 nodejs8.4 nodejs8.5
Timeout |int |是|  超时时间 1-300 最大300

## ListFunctions 函数列表

```go
	args := &api.ListFunctionsArgs{}
	args.FunctionVersion = "ALL"
	res, err := cfcClient.ListFunctions(args)
	if err != nil {
		panic(err)
	}

```

<b>请求参数<b>

参数名称 | 类型 | 是否必需  | 描述
--- | --- | --- | ---
FunctionVersion |String |否| 指定函数版本，如果没有指定返回所有函数$LATEST版本，可选有效值 ALL:将返回所有版本，包括$LATEST
Marker |int|否 |
MaxItems |int|否 |1-10000

## Invocations 调用函数

```go
	args := &api.InvocationsArgs{}
	args.InvocationType = api.InvocationTypeRequestResponse
	args.LogType = api.LogTypeTail
	res, err := cfcClient.Invocations(FunctionName, nil, args)
	if err != nil {
		panic(err)
	}
```

<b>请求参数<b>

参数名称 | 类型 | 是否必需  | 描述
--- | --- | --- | --- 
FunctionName |String |是 |您可以指定一个函数名(例如，Thumbnail)，或者您可以指定函数的BRN资源名(例如，brn:bce:cfc:bj:account-id:function:thumbnail:$LATEST)。CFC也允许您指定一个部分的BRN(例如，account-id:Thumbnail)。注意，BRN长度限制1-170。如果只指定函数名，则长度限制为64个字符。
InvocationType | String | 是  | Event(异步调用)/RequestResponse(同步调用)/DryRun(测试函数)
LogType | String | 否  | 日志类型 Tail / None 您可以将这个可选参数设置为Tail，前提是InvocationType参数必须为RequestResponse。在本例中，CFC在x-bce-log-result头中返回最后4KB的base64编码的日志数据。
Qualifier | String | 否  | 您可以使用这个可选参数来指定CFC函数版本或别名。如果指定一个函数版本，那么API将使用限定的函数BRN来调用特定的CFC函数。如果指定别名，则API使用别名BRN来调用别名指向的CFC函数版本。<br>如果您不提供此参数，那么API将调用$LATEST。

## 参数定义

### Code

状态 | 类型 | 描述
--- | --- | ---
ZipFile| String | 您要发布的zip包的 base64-encoded 注意zip包压缩目录的内容，而不是目录本身
Publish| Boolean | 是否直接发布版本



### Environment

状态 | 类型 | 描述
--- | --- | ---
Variables | map[string]string | 环境变量参数

