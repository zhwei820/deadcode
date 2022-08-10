package bb

type T1 struct{} // used
type t2 struct { // used
	t3 // used
}
type t3 struct{} // used

func (t3) Foo() {} // used
func (t3) foo() {} // unused

type T5 struct{} // used

func (T5) Foo() {} // used

func init() { // used
	_ = T1{}
	t := t2{}
	t.Foo()
}
