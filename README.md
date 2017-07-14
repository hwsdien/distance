#### Distance
    
    计算两地球坐标点的直线距离, 相距多少米。
    
    
    
##### 示例
```go
package main

import (
    "fmt"
    "github.com/hwsdien/distance"
)

func main() {
    lng1 := 112.99266052246
    lat1 := 28.212518692017

    lng2 := 112.992158944
    lat2 := 28.2136793274

    data := distance.GetDistance(lng1, lat1, lng2, lat2)
    fmt.Println(data)
}
```
