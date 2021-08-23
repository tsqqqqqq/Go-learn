package infar

type RetrieveFactory struct {
}

const (
	CODERETRIEVE = iota
	TESTRETRIEVE
)

/**
创造类型枚举，传递类型的枚举，从而返回相应的接口实现。
这个方式有点繁杂没有断言来的块。
*/
func (rf *RetrieveFactory) GetRetrieve(retrieve int) Retrieve {
	var res Retrieve
	switch retrieve {
	case 0:
		res = new(CodeRetrieve)
	case 1:
		res = new(TestRetrieve)
	default:
		panic("not implements Retrieve Interface")
	}
	return res
}

/**
使用go语言的TypeAssertion 类型断言来判断当前传入工厂的类型，进而直接获取实现，
特点是需要传递当前类型的指针进入工厂。若是值传递则无效
*/
func (rf *RetrieveFactory) GetTypeAssertionRetrieve(r Retrieve) Retrieve {
	switch r.(type) {
	case CodeRetrieve:
		return new(CodeRetrieve)
	case TestRetrieve:
		return new(TestRetrieve)
	default:
		panic("not implements Retrieve Interface")
	}
}
