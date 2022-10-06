package main

import (
	"fmt"
	"time"
)

func main()  {
	// 初始化时间轮
	// 第一个参数为tick刻度, 即时间轮多久转动一次
	// 第二个参数为时间轮槽slot数量
	tw := New(1 * time.Second, 3600)

	// 启动时间轮
	tw.Start()

	// 添加定时器
	// 第一个参数为延迟时间
	// 第二个参数为定时器唯一标识, 删除定时器需传递此参数
	//第三个参数为传入的自定义函数, 此参数将会传递给回调函数, 类型为func(data interface{})
	// 第四个参数为用户自定义数据, 将参数传给上面的方法, 类型为interface{}
	tw.AddTimer(1 * time.Second, 1, func(data interface{}) {
		v, ok := data.(map[string]int)
		if !ok {
			// Can't assert, handle error.
		}
		for _, s := range v {
			fmt.Printf("Value: %v\n", s)
		}
	},map[string]int{"uid" : 1056})
	tw.AddTimer(4 * time.Second, 2, func(data interface{}) {
		v, ok := data.(map[string]int)//解析data->map
		if !ok {
			// Can't assert, handle error.
		}
		for _, s := range v {
			fmt.Printf("Value: %v\n", s)
		}
	},map[string]int{"uid" : 626})
	// 删除定时器, 参数为添加定时器传递的唯一标识
	tw.RemoveTimer(2)

	// 停止时间轮
	tw.Stop()

	select{}
}