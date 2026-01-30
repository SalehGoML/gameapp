package entity

//
//type Category uint8
//
//const (
//	CategorySport Category = iota + 1
//	CategoryHistory
//	CategoryTech
//)
//
//func (c Category) String() string {
//	switch c {
//	case 1:
//		return "sport"
//	case 2:
//		return "history"
//	case 3:
//		return "tech"
//	}
//
//	return ""
//}

type Category struct {
	ID          uint
	Name        string
	Description string
}
