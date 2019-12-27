package main

import (
	"bytes"
	"fmt"
	"github.com/ugorji/go/codec"
	"github.com/vmihailenco/msgpack"
	"testing"
)

type (
	Player struct {
		Id     int
		Level  int
		Heroes map[int]*Hero
		Equips []*Equip
	}

	Hero struct {
		Id     int
		Level  int
		Skills []*Skill
	}

	Equip struct {
		Id    int
		Level int
	}

	Skill struct {
		Id    int
		Level int
	}
)

func NewHero() *Hero {
	return &Hero{
		Id:     1,
		Level:  1,
		Skills: append([]*Skill{NewSkill()}, NewSkill(), NewSkill()),
	}
}

func NewSkill() *Skill {
	return &Skill{1, 1}
}

func NewEquip() *Equip {
	return &Equip{1, 1}
}

func NewPlayer() *Player {
	return &Player{
		Id:     1,
		Level:  1,
		Heroes: map[int]*Hero{1: NewHero(), 2: NewHero(), 3: NewHero()},
		Equips: append([]*Equip{NewEquip()}, NewEquip(), NewEquip()),
	}
}

func MsgPackDeepCopy(dst, src interface{}) error {
	b, err := msgpack.Marshal(src)
	if err != nil {
		return err
	}

	err = msgpack.Unmarshal(b, dst)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	//p1 := NewPlayer()
	//p2 := new(Player)
	//MsgPackDeepCopy(p2, p1)
	//fmt.Println(reflect.DeepEqual(p1, p2))

	mh := &codec.MsgpackHandle{RawToString: true}
	data := []interface{}{"abc", 12345, 1.2345}
	buf := &bytes.Buffer{}
	enc := codec.NewEncoder(buf, mh)
	enc.Encode(data)
	fmt.Printf("%x", buf.Bytes())

}

//output
//true

//goos: windows
//goarch: amd64
//pkg: game.lab/go-deepcopy/src/msgpack
//100000         20220 ns/op
//PASS

// 性能测试
func BenchmarkMsgPack(b *testing.B) {
	p1 := NewPlayer()
	p2 := new(Player)
	for i := 0; i < b.N; i++ {
		MsgPackDeepCopy(p2, p1)
	}
}
