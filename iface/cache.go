package iface

import "time"

type Cache interface {
	//替换当前cache的value
	Replace(newValue interface{})

	//设置过期时间
	SetOverTime(t time.Duration)

	//获取数据类型
	GetValueType()

	//获取值
	GetValue()

	//返回引用次数
	RetrunCount() int
}
