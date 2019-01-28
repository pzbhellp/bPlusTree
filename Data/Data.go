package Data

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Key   *Key
	Value *Value
}

func NewData(key *Key, value *Value) (r *Data) {
	r = &Data{
		Key:   key,
		Value: value,
	}
	return
}

func Sort(dataList []*Data) []*Data {
	quick(dataList, 0, len(dataList)-1)
	return dataList
}

func quick(data []*Data, start, end int) {
	if start >= end {
		return
	}
	mid := sorting(data, start, end)
	print("mid:")
	println(mid)

	quick(data, start, mid-1)
	quick(data, mid+1, end)
}

func sorting(data []*Data, start, end int) int {
	flag := data[start]
	for start < end {
		for start < end && data[end].Key.GE(flag.Key) {
			end -= 1
		}
		for start < end && data[start].Key.LE(flag.Key) {
			start += 1
		}
		if start != end {
			data[start].Swap(data[end])
		}
	}
	data[start].Swap(flag)
	return start
}

func (data *Data) Swap(desc *Data) {
	*data, *desc = *desc, *data
}

func (data *Data) ToString() string {
	return fmt.Sprintf("key:%d,value:%v", *data.Key, *data.Value)
}

func DataListToJson(list []*Data) string {
	ret, _ := json.Marshal(list)
	return string(ret)
}
