package bPlusTree

type Data struct {
	Key Key
	Value Value
}

func NewData(key Key,value Value) (r *Data){
	r = &Data{
		Key:key,
		Value:value,
	}
	return
}

func Sort(dataList []*Data){

}