package vo

//公共的结构体
type PublicVO struct {
	user User
	info info
}
type User struct {
	name string
	age  int
}

//小写是私有的结构体
type info struct {
	id     int64
	IDCARD string
}
