package services

type TestService struct {
	// 某些DO对象 or PO对象
}

func (t *TestService) Test() string {
	//do something
	//获取/更新某些DO or PO对象
	return "test"
}