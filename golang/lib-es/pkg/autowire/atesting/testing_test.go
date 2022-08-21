package atesting_test

import (
	"melody-io/lib-es/pkg/autowire"
	"melody-io/lib-es/pkg/autowire/atesting"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	fooName     = "foo"
	barName     = "bar"
	testFooName = "testFoo"
	testBarName = "testBar"
)

type FooEr interface {
	Foo()
}

// A Foo represent named struct
type Foo struct {
	Name       string
	CloseCalls int
}

// Pass method
func (f Foo) Foo() {
}

type BarEr interface {
	Baz()
}

// Foo represent named struct
type Bar struct {
	Name       string
	CloseCalls int
}

// Pass method
func (b Bar) Baz() {
}

// A FooBarUnexported represent named struct
type FooBarUnexported struct {
	foo *Foo `autowire:""`
	bar *Bar `autowire:""`
}

// A FooBar represent named struct
type FooBar struct {
	Foo *Foo `autowire:""`
	Bar *Bar `autowire:""`
}

// Baz represent named struct
type Baz struct {
	MyFoo FooEr `autowire:"Foo"`
	MyBaz BarEr `autowire:"Bar"`
}

func TestSpyUnexportedStructPtr(t *testing.T) {
	autowire.Autowire(&Foo{Name: fooName})
	autowire.Autowire(&Bar{Name: barName})
	tmpFooBar := &FooBarUnexported{}
	autowire.Autowire(tmpFooBar)
	assert.Equal(t, tmpFooBar.foo.Name, fooName)
	assert.Equal(t, tmpFooBar.bar.Name, barName)
	atesting.Spy(tmpFooBar, &Foo{Name: testFooName}, &Bar{Name: testBarName})
	assert.Equal(t, tmpFooBar.foo.Name, testFooName)
	assert.Equal(t, tmpFooBar.bar.Name, testBarName)
	assert.Equal(t, 0, len(autowire.Close()))
}

func TestSpyExportedStructPtr(t *testing.T) {
	autowire.Autowire(&Foo{Name: fooName})
	autowire.Autowire(&Bar{Name: barName})
	tmpFooBar := &FooBar{}
	autowire.Autowire(tmpFooBar)
	assert.Equal(t, tmpFooBar.Foo.Name, fooName)
	assert.Equal(t, tmpFooBar.Bar.Name, barName)
	atesting.Spy(tmpFooBar, &Foo{Name: testFooName}, &Bar{Name: testBarName})
	assert.Equal(t, tmpFooBar.Foo.Name, testFooName)
	assert.Equal(t, tmpFooBar.Bar.Name, testBarName)
	assert.Equal(t, 0, len(autowire.Close()))
}

func TestSpyInterface(t *testing.T) {
	autowire.Autowire(&Foo{Name: fooName})
	autowire.Autowire(&Bar{Name: barName})
	tmpBaz := &Baz{}
	autowire.Autowire(tmpBaz)
	assert.Equal(t, tmpBaz.MyFoo.(*Foo).Name, fooName)
	assert.Equal(t, tmpBaz.MyBaz.(*Bar).Name, barName)
	atesting.Spy(tmpBaz, &Foo{Name: testFooName}, &Bar{Name: testBarName})
	assert.Equal(t, tmpBaz.MyFoo.(*Foo).Name, testFooName)
	assert.Equal(t, tmpBaz.MyBaz.(*Bar).Name, testBarName)
	assert.Equal(t, 0, len(autowire.Close()))
}
