# Tyme [![License](https://img.shields.io/badge/license-MIT-4EB1BA.svg?style=flat-square)](https://github.com/6tail/tyme4go/blob/master/LICENSE)

Tyme是一个非常强大的日历工具库，可以看作 [Lunar](https://6tail.cn/calendar/api.html "https://6tail.cn/calendar/api.html") 的升级版，拥有更优的设计和扩展性，支持公历、农历、藏历、星座、干支、生肖、节气、法定假日等。

> 基于go1.19版本开发

## 示例

    package main
     
    import (
        "fmt"
        "github.com/6tail/tyme4go/tyme"
    )
     
    func main() {
        // 公历日
        solarDay, err := tyme.SolarDay{}.FromYmd(1986, 5, 29)

        // 参数有误无法创建公历日对象时，err错误信息不为空
        if err == nil {
          // 1986年5月29日
          fmt.Println(solarDay)
           
          // 农历丙寅年四月廿一
          fmt.Println(solarDay.GetLunarDay())
        }
    }

## 文档

请移步至 [https://6tail.cn/tyme.html](https://6tail.cn/tyme.html "https://6tail.cn/tyme.html")

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=6tail/tyme4go&type=Date)](https://star-history.com/#6tail/tyme4go&Date)

## 致谢
1. 感谢许剑伟老师分享的寿星天文历，本项目节气算法引自 [https://github.com/sxwnl/sxwnl](https://github.com/sxwnl/sxwnl "https://github.com/sxwnl/sxwnl")
2. 感谢[stonelf](https://github.com/stonelf "https://github.com/stonelf")，本项目藏历数据引自 [https://github.com/stonelf/zangli](https://github.com/stonelf/zangli "https://github.com/stonelf/zangli")
