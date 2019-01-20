package cv2js

import (
	"fmt"
	"image"
	"testing"
	"time"

	"github.com/robertkrimen/otto"
)

type TestFloat []float32
type TestFloat2 float32
type TestStruct struct {
	X int64
	Y int64
}

func TestRun(t *testing.T) {
	// GenOttoFuncCode()
	// parseAll()
	// PrintDefines()
}

func TestInterface(t *testing.T) {
	// v := TestFloat2(1)
	// f := func(i interface{}) {
	// 	vv := TestFloat2(2)
	// 	ri := reflect.ValueOf(&vv)
	// 	rj := reflect.ValueOf(i)

	// 	rj.Set(ri)
	// 	fmt.Println("a", i)
	// }
	// f(&v)
	// fmt.Println("b", v)
}

func TestVMGetLine(t *testing.T) {
	vm := otto.New()
	vm.Set("testP1", func(call otto.FunctionCall) otto.Value {
		f0, _ := call.Argument(0).Export()
		fmt.Printf("%T %v\n", f0, f0)
		f1, err := call.Argument(1).ToFloat()
		fmt.Println(f1, err)

		v := &TestStruct{10, 20}
		fmt.Printf("%T %p\n", v, v)

		var f TestFloat2 = 10.33
		v1, _ := vm.ToValue(float32(f))
		v2, _ := vm.ToValue("hello")
		v3, _ := vm.ToValue([]byte{'H', 'e', 'l', 'l', 'o'})
		v4, _ := vm.ToValue(v)
		results := map[string]interface{}{}
		results["float"] = v1
		results["string"] = v2
		results["bytes"] = v3
		results["struct"] = v4
		results["rect"] = []image.Rectangle{image.Rectangle{}}
		r, _ := vm.ToValue(results)
		return r
	})
	vm.Set("testP2", func(call otto.FunctionCall) otto.Value {
		f0, _ := call.Argument(0).Export()
		fmt.Printf("%T %p\n", f0, f0)
		f0v := f0.(*TestStruct)
		f0v.X = 30
		fmt.Println(f0v)
		r, _ := vm.ToValue(0)
		return r
	})
	_, e := vm.Run(`
		function a() {
			var obj1 = testP1({'a': 10, 'b': 20}, 1.2);
			obj1.bytes[0] = 255;
			var obj2 = testP1(2);

			// obj2.struct.X = 20;
			console.log(JSON.stringify(obj2));
			testP2(obj2.struct);
			console.log(JSON.stringify(obj2));
		}
		a();
	`)
	if e != nil {
		fmt.Println(e)
	}
	time.Sleep(50 * time.Millisecond)
}
