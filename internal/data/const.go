package data

// 在文件开头添加自定义类型
type ContextKey string

const (
	UserTokenKey ContextKey = "user-token"
	UserName     ContextKey = "username"
	UserId       ContextKey = "userid"
)
