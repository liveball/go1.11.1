package main

import (
	"fmt"
	"reflect"
	"testing"
)

type user1 struct {
	b byte
	i int32
	j int64
}

type user2 struct {
	b byte
	j int64
	i int32
}

type user3 struct {
	i int32
	b byte
	j int64
}

type user4 struct {
	i int32
	j int64
	b byte
}

type user5 struct {
	j int64
	b byte
	i int32
}

type user6 struct {
	j int64
	i int32
	b byte
}

// 对齐系数,GOARCH=amd64是8，GOARCH=386是4

// 在分析之前，我们先看下内存对齐的规则：
//   1、计算一个的字段的对齐，就是将当前字段的内存大小填充为对齐值的大小

//   2、struct在每个字段都对齐之后，其本身也要进行对齐。

//  以上这两条规则要好好理解，理解明白了才可以分析下面的struct结构体。
//  在这里再次提醒，对齐值也叫对齐系数、对齐倍数，对齐模数。这就是说，每个字段在内存中的偏移量是对齐值的倍数即可。

// 结构体的成员变量，第一个成员变量的偏移量为 0。
// 往后的每个成员变量的对齐值必须为编译器默认对齐长度（#pragma pack(n)）或当前成员变量类型的长度（unsafe.Sizeof），取最小值作为当前类型的对齐值。其偏移量必须为对齐值的整数倍
// 结构体本身，对齐值必须为编译器默认对齐长度（#pragma pack(n)）或结构体的所有成员变量类型中的最大长度，取最大数的最小整数倍作为对齐值
// 结合以上两点，可得知若编译器默认对齐长度（#pragma pack(n)）超过结构体内成员变量的类型最大长度时，默认对齐长度是没有任何意义的

type Part1 struct {
	a bool //axxx|bbbb|cxxx|xxxx|dddd|dddd|exxx|xxxx
	b int32
	c int8
	d int64
	e byte
}

type Part2 struct {
	e byte // ecax|bbbb|dddd|dddd
	c int8
	a bool
	b int32
	d int64
}

func Test_align(t *testing.T) {
	var u1 user1
	var u2 user2
	var u3 user3
	var u4 user4
	var u5 user5
	var u6 user6
	part1 := Part1{}
	part2 := Part2{}

	// type user1 struct {
	// 	b byte
	// 	i int32
	// 	j int64
	// }

	// 例： map[int64]int8
	// k/v/k/v/k/v/k/v
	// xxxx|xxxx|vaaa|aaaa|  ... xxxx|xxxx|vaaa|aaaa|

	// k/k/k/k/v/v/v/v
	// xxxx|xxxx|xxxx|xxxx|  ... vvvv|vvvv|vvvv|vvvv|

	showAlign(u1)    //bxxx|iiii|jjjj|jjjj
	showAlign(u2)    //bxxx|xxxx|jjjj|jjjj|iiii|xxxx
	showAlign(u3)    //iiii|bxxx|jjjj|jjjj
	showAlign(u4)    //iiii|xxxx|jjjj|jjjj|bxxx|xxxx
	showAlign(u5)    //jjjj|jjjj|bxxx|iiii
	showAlign(u6)    //jjjj|jjjj|iiii|bxxx
	showAlign(part1) // axxx|bbbb|cxxx|xxxx|dddd|dddd|exxx|xxxx
	showAlign(part2) // ecax|bbbb|dddd|dddd

	// 字段      uint8，大小： 1，对齐： 1，字段对齐： 1，偏移： 0
	// 字段      int32，大小： 4，对齐： 4，字段对齐： 4，偏移： 4
	// 字段      int64，大小： 8，对齐： 8，字段对齐： 8，偏移： 8
	// t size is (16)
	// 字段      uint8，大小： 1，对齐： 1，字段对齐： 1，偏移： 0
	// 字段      int64，大小： 8，对齐： 8，字段对齐： 8，偏移： 8
	// 字段      int32，大小： 4，对齐： 4，字段对齐： 4，偏移：16
	// t size is (24)
	// 字段      int32，大小： 4，对齐： 4，字段对齐： 4，偏移： 0
	// 字段      uint8，大小： 1，对齐： 1，字段对齐： 1，偏移： 4
	// 字段      int64，大小： 8，对齐： 8，字段对齐： 8，偏移： 8
	// t size is (16)
	// 字段      int32，大小： 4，对齐： 4，字段对齐： 4，偏移： 0
	// 字段      int64，大小： 8，对齐： 8，字段对齐： 8，偏移： 8
	// 字段      uint8，大小： 1，对齐： 1，字段对齐： 1，偏移：16
	// t size is (24)
	// 字段      int64，大小： 8，对齐： 8，字段对齐： 8，偏移： 0
	// 字段      uint8，大小： 1，对齐： 1，字段对齐： 1，偏移： 8
	// 字段      int32，大小： 4，对齐： 4，字段对齐： 4，偏移：12
	// t size is (16)
	// 字段      int64，大小： 8，对齐： 8，字段对齐： 8，偏移： 0
	// 字段      int32，大小： 4，对齐： 4，字段对齐： 4，偏移： 8
	// 字段      uint8，大小： 1，对齐： 1，字段对齐： 1，偏移：12
}

func showAlign(t interface{}) {
	v := reflect.TypeOf(t)
	n := v.NumField()
	for i := 0; i < n; i++ {
		sf := v.Field(i)
		fmt.Printf("字段 %10v，大小：%2v，对齐：%2v，字段对齐：%2v，偏移：%2v\n",
			sf.Type.Kind(),
			sf.Type.Size(),
			sf.Type.Align(),
			sf.Type.FieldAlign(),
			sf.Offset,
		)
	}
	fmt.Printf("t size is (%2v) align is (%2v)\n", v.Size(), v.Align())
}

// So(unsafe.Sizeof(true), ShouldEqual, 1)
// So(unsafe.Sizeof(int8(0)), ShouldEqual, 1)
// So(unsafe.Sizeof(int16(0)), ShouldEqual, 2)
// So(unsafe.Sizeof(int32(0)), ShouldEqual, 4)
// So(unsafe.Sizeof(int64(0)), ShouldEqual, 8)
// So(unsafe.Sizeof(int(0)), ShouldEqual, 8)
// So(unsafe.Sizeof(float32(0)), ShouldEqual, 4)
// So(unsafe.Sizeof(float64(0)), ShouldEqual, 8)
// So(unsafe.Sizeof(""), ShouldEqual, 16)
// So(unsafe.Sizeof("hello world"), ShouldEqual, 16)
// So(unsafe.Sizeof([]int{}), ShouldEqual, 24)
// So(unsafe.Sizeof([]int{1, 2, 3}), ShouldEqual, 24)
// So(unsafe.Sizeof([3]int{1, 2, 3}), ShouldEqual, 24)
// So(unsafe.Sizeof(map[string]string{}), ShouldEqual, 8)
// So(unsafe.Sizeof(map[string]string{"1": "one", "2": "two"}), ShouldEqual, 8)
// So(unsafe.Sizeof(struct{}{}), ShouldEqual, 0)

// bool 类型虽然只有一位，但也需要占用1个字节，因为计算机是以字节为单位
// 64为的机器，一个 int 占8个字节
// string 类型占16个字节，内部包含一个指向数据的指针（8个字节）和一个 int 的长度（8个字节）
// slice 类型占24个字节，内部包含一个指向数据的指针（8个字节）和一个 int 的长度（8个字节）和一个 int 的容量（8个字节）
// map 类型占8个字节，是一个指向 map 结构的指针
// 可以用 struct{} 表示空类型，这个类型不占用任何空间，用这个作为 map 的 value，可以讲 map 当做 set 来用
