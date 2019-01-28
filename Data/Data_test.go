package Data

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewData(t *testing.T) {
	key1 := Key(11)
	val1 := Value(111)
	tmp := NewData(&key1, &val1)
	fmt.Println(tmp.ToString())
}

func TestData_Swap(t *testing.T) {
	key1 := Key(11)
	val1 := Value(111)
	da := NewData(&key1, &val1)
	key2 := Key(22)
	val2 := Value(222)
	db := NewData(&key2, &val2)
	println(da.ToString())
	println(db.ToString())
	da.Swap(db)
	println(da.ToString())
	println(db.ToString())
}

func TestSort(t *testing.T) {
	keyList := []Key{
		1,
		2,
		2,
		3,
		4,
		5,
		10,
		9,
		1,
	}
	val := Value(1)
	dataList := make([]*Data, 0)
	for i, _ := range keyList {
		tmp := new(Data)
		tmp.Key = &keyList[i]
		tmp.Value = &val
		println(tmp.ToString())
		dataList = append(dataList, tmp)
	}
	for i := 0; i < 20; i++ {
		dataList = append(dataList, genData())
	}
	Sort(dataList)
	println("------------------")
	for _, item := range dataList {
		println(item.ToString())
	}
}

func genData() *Data {
	key := Key(rand.Int31n(100))
	value := Value(100 + rand.Int31n(100))
	return &Data{
		Key:   &key,
		Value: &value,
	}
}
