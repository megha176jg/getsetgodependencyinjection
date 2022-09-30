package datastructures

type List []interface{}

func (l List) Exists(object interface{}) bool {
	for _, i := range l {
		if i == object {
			return true
		}
	}
	return false
}
