package iface


type CacheManager interface {
	//清空所有缓存
	Clear()

	//根据Key获取value
	Get()

	//获取所有key列表
	KeySet() []interface{}

	//写入一个k/v
	Put(key,value interface{})

	//如果缓存没命中则写入
	PutIfABsent()

	//删除一个key
	Remove(key interface{})
	//删除key列表
	RemoveAll(keys []interface{})

	//替换一个k的v
	Replace(key,newValue interface{})
}
