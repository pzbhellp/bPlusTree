package Data

//主键值

type Key int

func (k *Key) G(desc *Key) bool {
	return *k > *desc
}

func (k *Key) GE(desc *Key) bool {
	return *k >= *desc
}

func (k *Key) L(desc *Key) bool {
	return *k < *desc
}

func (k *Key) LE(desc *Key) bool {
	return *k <= *desc
}

func (k *Key) Equal(desc *Key) bool {
	return *desc == *k
}

func CompareKey(src *Key, desc *Key) int {
	if src.G(desc) {
		return 1
	} else if src.L(desc) {
		return -1
	} else {
		return 0
	}
}
