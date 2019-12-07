package str

import (
	"bytes"
	"strings"
	"testing"
)

func Benchmark_StrAdd(b *testing.B) {
	str := "hello world"
	for i := 0; i < b.N; i++ {
		str += "hello acbd abcd"
	}
	b.Log(len(str))
}

func Benchmark_BufferAdd(b *testing.B) {
	buffer := bytes.Buffer{}
	buffer.WriteString("hello world")
	for i := 0; i < b.N; i++ {
		buffer.WriteString("hello acbd abcd")
	}
	b.Log(buffer.Len())
}

func Benchmark_BuilderAdd(b *testing.B) {
	/***
	当buf内存不够时候,扩容方式, 成倍增长
	if cap(b.buf)-len(b.buf) < n {
	    buf := make([]byte, len(b.buf), 2*cap(b.buf)+n)
		copy(buf, b.buf)
		b.buf = buf
	}
	*/
	buffer := strings.Builder{}
	buffer.WriteString("hello world")
	for i := 0; i < b.N; i++ {
		buffer.WriteString("hello acbd abcd")
	}
	b.Log(buffer.Len())
}

func Test_sliceLenAndCap(t *testing.T) {
	buf := make([]byte, 10, 20)
	t.Log("len :", len(buf))
	t.Log("cap :", cap(buf))
}

/***
bytes.Buffer Reset 和 strings.Builder Reset区别

// Reset resets the buffer to be empty,
// but it retains the underlying storage for use by future writes.
// Reset is the same as Truncate(0).
func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
	b.off = 0
	b.lastRead = opInvalid
}

// 直接舍弃掉之前的  []byte 数组
func (b *Builder) Reset() {
	b.addr = nil
	b.buf = nil
}


返回string 区别
// String returns the accumulated string.
func (b *Builder) String() string {
	return *(*string)(unsafe.Pointer(&b.buf))
}

//返回数组的一个副本
// String returns the contents of the unread portion of the buffer
// as a string. If the Buffer is a nil pointer, it returns "<nil>".
//
// To build strings more efficiently, see the strings.Builder type.



*/

func Test_HelloWorld(t *testing.T) {

	func() {
		t.Log("测试1: 使用[]byte创建一个buffer 调用reset重置")
		buf := make([]byte, 10, 20)
		copy(buf, "0123456789x")
		t.Log("初始化slice内容 : ", buf)
		b := bytes.NewBuffer(buf)
		b.Reset()
		t.Log("reset slice : ", buf)
		b.WriteString("abcdefg")
		t.Log("WriteString slice : ", buf)
		t.Log("Buffer.String : ", b.String())
	}()

	func() {
		t.Log("\n测试2: 使用[]byte创建一个buffer 不重置")
		buf := make([]byte, 10, 20)
		copy(buf, "0123456789x")
		t.Log("初始化slice内容 : ", buf)

		b := bytes.NewBuffer(buf)

		b.WriteString("abcdefg")
		t.Log("WriteString slice : ", buf)
		t.Log("Buffer.String : ", b.String())
	}()

	func() {
		t.Log("\n测试3: 使用[]byte创建一个buffer 调用reset重置,copy(buf[4:],...) 对比测试1 ")
		buf := make([]byte, 10, 20)
		copy(buf[4:], "0123456789x")
		t.Log("初始化slice内容 : ", buf)

		b := bytes.NewBuffer(buf)
		b.Reset()
		b.WriteString("abcdefg")
		t.Log("WriteString slice : ", buf)
		t.Log("Buffer.String : ", b.String())
	}()
}

func Test_ByteSlice(t *testing.T) {
	b := make([]byte, 10)
	copy(b, []byte("hello"))

	t.Log(string(b))

	c := b
	copy(b, []byte("hello123"))

	t.Log(string(b))
	t.Log(string(c))

	c = []byte("1234567")
	t.Log(string(b))
	t.Log(string(c))

}
