// reference: 《Go 语言从入门到进阶实战》

package simpe_json

import (
	"testing"
)

func TestMyMarshalJson(t *testing.T) {
	type Sex int
	const (
		Male Sex = iota
		Female
	)
	type PersonInfo struct {
		Name string `myjson:"name"`
		Sex  Sex `myjson:"sex"`
	}

	type StudentInfo struct {
		PersonInfo
		StuID int64
		Next *StudentInfo
	}

	stu := &StudentInfo{
		PersonInfo{"Tom", Male},
		302831,
		nil,
	}
	//stu.Next = stu
	//思考: 如何识别 ptr cycle

	s, err := MarshalJson(stu)
	if err != nil {
		t.Errorf("json Marshal failed: %v", err)
	} else {
		t.Logf("json out: %s", s)
	}
}