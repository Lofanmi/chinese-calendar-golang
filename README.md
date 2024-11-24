# chinese-calendar-golang

[![Build Status](https://github.com/Lofanmi/chinese-calendar-golang/workflows/Run%20Tests/badge.svg?branch=master)](https://github.com/Lofanmi/chinese-calendar-golang/actions?query=branch%3Amaster)
[![codecov](https://codecov.io/gh/Lofanmi/chinese-calendar-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/Lofanmi/chinese-calendar-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/Lofanmi/chinese-calendar-golang)](https://goreportcard.com/report/github.com/Lofanmi/chinese-calendar-golang)
[![Go Reference](https://pkg.go.dev/badge/github.com/Lofanmi/chinese-calendar-golang?status.svg)](https://pkg.go.dev/github.com/Lofanmi/chinese-calendar-golang?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/Lofanmi/chinese-calendar-golang/-/badge.svg)](https://sourcegraph.com/github.com/Lofanmi/chinese-calendar-golang?badge)
[![Open Source Helpers](https://www.codetriage.com/lofanmi/chinese-calendar-golang/badges/users.svg)](https://www.codetriage.com/Lofanmi/chinese-calendar-golang)

公历, 农历, 干支历转换包, 提供精确的日历转换.

使用 `Go` 编写, 觉得好用的给 `Star` 呗~

觉得**好用**, 而不是觉得**有用**. 如果不好用, 欢迎向我提 issue, 我会抽空不断改进它!

真希望有人打钱打钱打钱给我啊哈哈哈哈!!!

# 如何安装

新版本已支持 Go Module, 可以直接引入。

```bash
go mod tidy
go mod download
go build
```

如果您还在使用 GOPATH, 别忘了 `go get`:
```bash
go get -u -v github.com/Lofanmi/chinese-calendar-golang/calendar
```

# 用法

```bash
# 确保使用的是东八区(北京时间)
export TZ=PRC

# 查看时间
date -R
```

```go
import (
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/calendar"
)

func xxx() {
    t := time.Now()
    // 1. ByTimestamp
    // 时间戳
    c := calendar.ByTimestamp(t.Unix())
    // 2. BySolar
    // 公历
    c := calendar.BySolar(year, month, day, hour, minute, second)
    // 3. ByLunar
    // 农历(最后一个参数表示是否闰月)
    c := calendar.ByLunar(year, month, day, hour, minute, second, false)

    bytes, err := c.ToJSON()

    fmt.Println(string(bytes))
}
```

# 原理

1. 公历的计算比较容易, 生肖, 星座等的转换见源码, 不啰嗦哈;

2. 农历的数据和算法参考了[chinese-calendar](https://github.com/overtrue/chinese-calendar)的实现, 还有 [1900年至2100年公历、农历互转Js代码 - 晶晶的博客](http://blog.jjonline.cn/userInterFace/173.html), 节气的舍去了, 因为在干支历和生肖的计算上不够精确;

3. 干支历的转换是最繁琐的, 其计算依据是二十四节气. ~~我结合了网上的文章, 用天文方法计算了 1904-2024 的二十四节气时间.~~ 目前使用的是寿星天文历的时间, 支持 1904-3000 年, 精确到秒, 并与[香港天文台](http://data.weather.gov.hk/gts/astron2018/Solar_Term_2018_uc.htm)进行比较, 误差应该在 `1分钟` 以内, 大家如果有以往的数据, 可以看下源码比对一下.

# 使用须知

1. 干支历的时间范围是 1904-~~2024~~3000 年, ~~主要也是因为 1904 年之前的时间戳, 在32位的 `PHP` 下会识别不出(溢出), 而且再早的时间, 其实意义也不大,~~因为之前遗留的问题开放了 `NewSolarterm` 接口, 为保证兼容性, `index` 不能随便改, ~~2024之后...好像也不会用到吧. 研究二十四节气这个算法的时间花了很长很长, 反而撸代码的时间不算太多~~;

2. 实际上, 农历的时间也是可以根据天文方法计算出来的. 计算的过程, 也需要先计算二十四节气和每月的正月初一的日期(日月合朔), 才能知道闰月的信息(所以农历是阴阳历). 不过已经有了数据, 我也就懒了, 直接拿来用...所以农历的算法我还没有实现;

3. 农历的时间范围是 1900-2100 年, ~~但是这个日历只支持到 1904-2024 年.~~ 目前已经支持到 3000 年, 应该是足够使用了, 如果有需要我再加吧, ~~后续会有个 `PHP` 版本, 输出的 `JSON` 格式会保持一致~~.

# 公历(阳历) - 字段说明

```
{
    // 生肖, 以立春分界
    //     如 2018-02-04 05:28:26 立春
    //        2018-02-04 04:59:59 属鸡年
    //        2018-02-04 05:00:00 属狗年
    "animal": "鼠",

    // 星座
    "constellation": "天秤",

    // 今年是否为闰年
    "is_leep": true,

    // 年
    "year": 2020,
    // 月
    "month": 9,
    // 日
    "day": 20,

    // 时
    "hour": 5,
    // 分
    "minute": 15,
    // 秒
    "second": 26,
    // 纳秒
    "nanosecond": 0,

    // 星期日
    "week_alias": "日",
    // 星期序数, 0表示周日
    "week_number": 0,
}
```

# 农历(阴阳历) - 字段说明

```
{
    // 生肖, 以每年正月初一分界
    "animal": "鼠",

    // 年
    "year": 2020,
    // 年(汉字)
    "year_alias": "二零二零",
    // 月
    "month": 8,
    // 月(汉字)
    "month_alias": "八月",
    // 日
    "day": 4,
    // 日(汉字)
    "day_alias": "初四",

    // 是否闰年
    "is_leap": true,
    
    // 这个月是否为闰月
    "is_leap_month": false,

    // 今年闰四月
    "leap_month": 4,
}
```

# 干支历 - 字段说明

```
{
    // 生肖, 以立春分界
    //     如 2018-02-04 05:28:26 立春
    //        2018-02-03 05:28:25 属鸡年
    //        2018-02-04 05:28:26 属狗年
    "animal": "鼠",

    // 年干支
    "year": "庚子",
    // 年干支六十甲子序数
    "year_order": 37,

    // 月干支
    "month": "乙酉",
    // 月干支六十甲子序数
    "month_order": 22,

    // 日干支
    "day": "丙寅",
    // 日干支六十甲子序数
    "day_order": 3,

    // 时干支
    "hour": "辛卯",
    // 时干支六十甲子序数
    "hour_order": 28,
}
```

# 输出示例

```json
{
    "ganzhi": {
        "animal": "鼠",
        "day": "丙寅",
        "day_order": 3,
        "hour": "辛卯",
        "hour_order": 28,
        "month": "乙酉",
        "month_order": 22,
        "year": "庚子",
        "year_order": 37
    },
    "lunar": {
        "animal": "鼠",
        "day": 4,
        "day_alias": "初四",
        "is_leap": true,
        "is_leap_month": false,
        "leap_month": 4,
        "month": 8,
        "month_alias": "八月",
        "year": 2020,
        "year_alias": "二零二零"
    },
    "solar": {
        "animal": "鼠",
        "constellation": "天秤",
        "day": 20,
        "hour": 5,
        "is_leep": true,
        "minute": 15,
        "month": 9,
        "nanosecond": 0,
        "second": 26,
        "week_alias": "日",
        "week_number": 0,
        "year": 2020
    }
}
```

# TODO

1. ~~完善单元测试~~
2. ~~完善注释~~
3. ~~完善文档~~
4. ~~支持更大范围的时间~~
5. 把农历的算法实现一下
6. ~~支持 go module~~

# 参考资料

- [算法系列之十八：用天文方法计算二十四节气（上）](https://blog.csdn.net/orbit/article/details/7910220)
- [算法系列之十八：用天文方法计算二十四节气（下）](https://blog.csdn.net/orbit/article/details/7944248)
- [overtrue/chinese-calendar](https://github.com/overtrue/chinese-calendar)
- [1900年至2100年公历、农历互转Js代码 - 晶晶的博客](http://blog.jjonline.cn/userInterFace/173.html)
- [香港天文台](http://data.weather.gov.hk/)
- [五虎遁元](https://baike.baidu.com/item/%E4%BA%94%E8%99%8E%E9%81%81%E5%85%83/5471492)
- [五鼠遁元](https://baike.baidu.com/item/%E4%BA%94%E9%BC%A0%E9%81%81%E5%85%83/5471935)
- [NASA](https://eclipse.gsfc.nasa.gov/)

# License

MIT


# Stargazers over time

[![Stargazers over time](https://starchart.cc/Lofanmi/chinese-calendar-golang.svg)](https://starchart.cc/Lofanmi/chinese-calendar-golang)
