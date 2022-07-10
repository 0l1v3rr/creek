package test

type TestStruct struct {
	Id   int
	Name string
}

func GetTestStructArray() []TestStruct {
	return []TestStruct{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Will"},
		{Id: 3, Name: "Mark"},
	}
}

func GetOtherStructArray() []TestStruct {
	return []TestStruct{
		{Id: 12, Name: "Ian"},
		{Id: 14, Name: "Paul"},
		{Id: 13, Name: "Josh"},
	}
}

func GetTestStructItem() TestStruct {
	return TestStruct{Id: 12, Name: "Ian"}
}

func GetOtherStructItem() TestStruct {
	return TestStruct{Id: 13, Name: "Josh"}
}
