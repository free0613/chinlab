package errcode

var (
	Success = NewError(0, "成功")
	Fail    = NewError(10000000, "内部错误")
)
