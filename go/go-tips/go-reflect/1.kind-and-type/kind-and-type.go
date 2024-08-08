package main

import (
	"fmt"
	"reflect"
)

type CLC string

const (
	CLC_TYPE_A CLC = "马克思主义、列宁主义、毛泽东思想、邓小平理论"
	CLC_TYPE_B CLC = "哲学、宗教"
	CLC_TYPE_C CLC = "社会科学总论"
)

type Goods struct {
	Name  string
	Price int64
}

type Book struct {
	Goods
	Author string
	CLC    CLC `cn:"中图分类号"`
}

func KindAndType() {
	// Type 代码 Go 语言的类型对象：
	// Kind 是Type 的具体类型:
	//	const (
	//		Invalid Kind = iota
	//		Bool
	//		Int
	//		Int8
	//		Int16
	//		Int32
	//		Int64
	//		Uint
	//		Uint8
	//		Uint16
	//		Uint32
	//		Uint64
	//		Uintptr
	//		Float32
	//		Float64
	//		Complex64
	//		Complex128
	//		Array
	//		Chan
	//		Func
	//		Interface
	//		Map
	//		Ptr
	//		Slice
	//		String
	//		Struct
	//		UnsafePointer
	//	)

	typeOfRead := reflect.TypeOf(CLC_TYPE_A)
	kind := typeOfRead.Kind()
	name := typeOfRead.Name()
	fmt.Printf("Type Of CLC_TYPE_A: Kind(%s), Type(%s)\n", kind, name)
	// Type Of CLC_TYPE_A: Kind(string), Type(CLC)

	var book Book
	typeOfBook := reflect.TypeOf(book)
	fmt.Printf("Number of Field: %d\n", typeOfBook.NumField())
	// Number of Field: 3

	// find StructField
	// By Index 0
	field := typeOfBook.Field(0)
	fmt.Printf("Index[0]-> StructField: %v\n", field)
	// Index[0]-> StructField: {Goods  main.Goods  0 [0] true}

	// By Name
	field, _ = typeOfBook.FieldByName("Author")
	fmt.Printf("Name[Author] -> StructField: %v\n", field)
	// Name[Author] -> StructField: {Author  string  24 [1] false}

	// By NameFunc
	field, _ = typeOfBook.FieldByNameFunc(func(s string) bool {
		if s == "Author" { return true}
		return false
	})
	fmt.Printf("Func[Author] --> StructField: %v\n", field)
	// Func[Author] --> StructField: {Author  string  24 [1] false}

	// By Indexes
	field = typeOfBook.FieldByIndex([]int{0, 1})
	fmt.Printf("Indexes[0,1] --> StructField: %v\n", field)
	// Indexes[0,1] --> StructField: {Price  int64  16 [1] false}

	// Tag
	tag := typeOfBook.Field(2).Tag
	fmt.Printf("Index[0] Tag By `cn`: %s\n", tag.Get("cn"))
}

func main()  {
	KindAndType()
}
