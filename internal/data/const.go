package data

// 在文件开头添加自定义类型
type ContextKey string

const (
	CtxUserTokenKey ContextKey = "user-token"
	CtxUserName     ContextKey = "username"
	CtxUserId       ContextKey = "userid"
)
