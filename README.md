# hash
golang简化加密类库

## 依赖要求
没有

## 安装
使用`go`命令获取类库

```bash
go get github.com/flxxyz/hash
```

## 例子
```go
package main

import (
    "fmt"
    "github.com/flxxyz/hash"
)

func main() {
    //sha1编码
    hashValue := hash.SHA1("0")
    fmt.Printf("[sha1] 0 => %s", hashValue)

    //sha256编码
    hashValue = hash.SHA256("0")
    fmt.Printf("[sha256] 0 => %s", hashValue)

    //sha512编码
    hashValue = hash.SHA512("0")
    fmt.Printf("[sha512] 0 => %s", hashValue)

    //md5编码
    hashValue = hash.MD5("0")
    fmt.Printf("[md5] 0 => %s", hashValue)

    //hmac
    {
        data := "0"
        key := "1234567890"

        //以md5规则生成
        hashValue = hash.HMAC("md5", key, data)
        fmt.Printf("[md5] \"%s\" => %s", data, hashValue)

        //以sha1规则生成
        hashValue = hash.HMAC("sha1", key, data)
        fmt.Printf("[sha1] \"%s\" => %s", data, hashValue)

        //以sha256规则生成
        hashValue = hash.HMAC("sha256", key, data)
        fmt.Printf("[sha256] \"%s\" => %s", data, hashValue)

        //以sha512规则生成
        hashValue = hash.HMAC("sha512", key, data)
        fmt.Printf("[sha512] \"%s\" => %s", data, hashValue)
    }

    //base64
    {
        //base64编码
        encodeStr := "hello world"
        hashValue = hash.Base64Encode(encodeStr)
        fmt.Printf("[base64] encode \"%s\" => %s", encodeStr, hashValue)

        //base64解码
        decodeStr := "aGVsbG8gd29ybGQ="
        hashValue = hash.Base64Decode(decodeStr)
        fmt.Printf("[sha512] \"%s\" => %s", decodeStr, hashValue)
    }
}
```

## 文档
[文档点这里](http://godoc.org/github.com/flxxyz/hash)

## 版权
hash包在MIT License下发布。有关详细信息，请参阅LICENSE。