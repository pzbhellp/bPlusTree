package Data

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().Unix()))
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

func Find(dataList []*Data, key *Key) (*Data, int) {
	return binarySearch(dataList, 0, len(dataList)-1, key)
}

func binarySearch(dataList []*Data, start, stop int, key *Key) (ret *Data, index int) {
	var mid int = 0
	index = 0
	for start <= stop {
		mid = (start + stop) / 2
		if dataList[mid].Key.Equal(key) {
			return dataList[mid], mid
		} else if dataList[mid].Key.G(key) {
			stop = mid - 1
		} else if dataList[mid].Key.L(key) {
			start = mid + 1
		}
	}
	if dataList[mid].Key.G(key) {
		index = mid - 1
	}
	ret = dataList[index]
	if index < 0 {
		ret = nil
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
func GenData() *Data {
	key := Key(r.Int31n(100))
	value := Value(100 + r.Int31n(100))
	return &Data{
		Key:   &key,
		Value: &value,
	}
}
