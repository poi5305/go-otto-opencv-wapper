
package otto_wrapper

import (
	"fmt"
	"image"
	"image/color"

	"github.com/robertkrimen/otto"
	"gocv.io/x/gocv"
)

func (o *OttoFunctions) cvIMRead(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out := string(arg_0_out_string)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_IMReadFlag_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out_IMReadFlag := int(arg_1_out_IMReadFlag_int64)
	arg_1_out := gocv.IMReadFlag(arg_1_out_IMReadFlag)

	res_0 := gocv.IMRead(arg_0_out, arg_1_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvIMWrite(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out := string(arg_0_out_string)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := gocv.IMWrite(arg_0_out, arg_1_out)

	res_0_out := bool(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvIMWriteWithParams(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out := string(arg_0_out_string)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_in_arr, ok := arg_2_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.([]interface{})"))
	}
	arg_2_out := make([]int, len(arg_2_in_arr), len(arg_2_in_arr))
	for arg_2_in_arr_i, arg_2_in_arr_v := range arg_2_in_arr {
		arg_2_out_arg_2_in_arr_i_int64, ok := arg_2_in_arr_v.(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in_arr_v.(int64)"))
		}
		arg_2_out[arg_2_in_arr_i] = int(arg_2_out_arg_2_in_arr_i_int64)
	}

	res_0 := gocv.IMWriteWithParams(arg_0_out, arg_1_out, arg_2_out)

	res_0_out := bool(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvIMEncode(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_FileExt_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out_FileExt := string(arg_0_out_FileExt_string)
	arg_0_out := gocv.FileExt(arg_0_out_FileExt)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0, res_1 := gocv.IMEncode(arg_0_out, arg_1_out)

	res_0_out := bytesToBase64str(res_0)
	if err != nil {
		return o.handleError(err)
	}
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvIMEncodeWithParams(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_FileExt_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out_FileExt := string(arg_0_out_FileExt_string)
	arg_0_out := gocv.FileExt(arg_0_out_FileExt)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_in_arr, ok := arg_2_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.([]interface{})"))
	}
	arg_2_out := make([]int, len(arg_2_in_arr), len(arg_2_in_arr))
	for arg_2_in_arr_i, arg_2_in_arr_v := range arg_2_in_arr {
		arg_2_out_arg_2_in_arr_i_int64, ok := arg_2_in_arr_v.(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in_arr_v.(int64)"))
		}
		arg_2_out[arg_2_in_arr_i] = int(arg_2_out_arg_2_in_arr_i_int64)
	}

	res_0, res_1 := gocv.IMEncodeWithParams(arg_0_out, arg_1_out, arg_2_out)

	res_0_out := bytesToBase64str(res_0)
	if err != nil {
		return o.handleError(err)
	}
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvIMDecode(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out, err := base64strToBytes(arg_0_out_string)
	if err != nil {
		return o.handleError(err)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_IMReadFlag_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out_IMReadFlag := int(arg_1_out_IMReadFlag_int64)
	arg_1_out := gocv.IMReadFlag(arg_1_out_IMReadFlag)

	res_0, res_1 := gocv.IMDecode(arg_0_out, arg_1_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvNewWindow(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out := string(arg_0_out_string)

	res_0 := gocv.NewWindow(arg_0_out)

	res_0_out := (*gocv.Window)(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvWindowClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Window)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Window)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Window)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvWindowIsOpen(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Window)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Window)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Window)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.IsOpen()

	res_0_out := bool(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvWindowGetWindowProperty(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Window)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Window)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Window)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_WindowPropertyFlag_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out_WindowPropertyFlag := int(arg_1_out_WindowPropertyFlag_int64)
	arg_1_out := gocv.WindowPropertyFlag(arg_1_out_WindowPropertyFlag)

	res_0 := arg_0_out.GetWindowProperty(arg_1_out)

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvWindowSetWindowProperty(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Window)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Window)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Window)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_WindowPropertyFlag_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out_WindowPropertyFlag := int(arg_1_out_WindowPropertyFlag_int64)
	arg_1_out := gocv.WindowPropertyFlag(arg_1_out_WindowPropertyFlag)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_WindowFlag_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out_WindowFlag := float32(arg_2_out_WindowFlag_float64)
	arg_2_out := gocv.WindowFlag(arg_2_out_WindowFlag)

	arg_0_out.SetWindowProperty(arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvWindowSetWindowTitle(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Window)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Window)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Window)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_string, ok := arg_1_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(string)"))
	}
	arg_1_out := string(arg_1_out_string)

	arg_0_out.SetWindowTitle(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvWindowIMShow(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Window)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Window)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Window)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	arg_0_out.IMShow(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvWindowWaitKey(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Window)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Window)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Window)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)

	res_0 := arg_0_out.WaitKey(arg_1_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvWindowMoveWindow(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Window)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Window)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Window)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	arg_0_out.MoveWindow(arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvWindowResizeWindow(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Window)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Window)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Window)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	arg_0_out.ResizeWindow(arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvSelectROI(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out := string(arg_0_out_string)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := gocv.SelectROI(arg_0_out, arg_1_out)

	res_0_out_map := make(map[string]interface{})
	res_0_out_map_Min_map := make(map[string]interface{})
	res_0_out_map_Min_map["X"] = int64(res_0.Min.X)
	res_0_out_map_Min_map["Y"] = int64(res_0.Min.Y)
	res_0_out_map["Min"] = res_0_out_map_Min_map
	res_0_out_map_Max_map := make(map[string]interface{})
	res_0_out_map_Max_map["X"] = int64(res_0.Max.X)
	res_0_out_map_Max_map["Y"] = int64(res_0.Max.Y)
	res_0_out_map["Max"] = res_0_out_map_Max_map
	res_0_out := res_0_out_map
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvSelectROIs(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out := string(arg_0_out_string)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := gocv.SelectROIs(arg_0_out, arg_1_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Min_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Min_map["X"] = int64(res_0_v.Min.X)
		res_0_out_arr_res_0_i_map_Min_map["Y"] = int64(res_0_v.Min.Y)
		res_0_out_arr_res_0_i_map["Min"] = res_0_out_arr_res_0_i_map_Min_map
		res_0_out_arr_res_0_i_map_Max_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Max_map["X"] = int64(res_0_v.Max.X)
		res_0_out_arr_res_0_i_map_Max_map["Y"] = int64(res_0_v.Max.Y)
		res_0_out_arr_res_0_i_map["Max"] = res_0_out_arr_res_0_i_map_Max_map
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvWaitKey(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_int64, ok := arg_0_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(int64)"))
	}
	arg_0_out := int(arg_0_out_int64)

	res_0 := gocv.WaitKey(arg_0_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvWindowCreateTrackbar(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Window)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Window)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Window)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_string, ok := arg_1_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(string)"))
	}
	arg_1_out := string(arg_1_out_string)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	res_0 := arg_0_out.CreateTrackbar(arg_1_out, arg_2_out)

	res_0_out := (*gocv.Trackbar)(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvTrackbarGetPos(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Trackbar)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Trackbar)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Trackbar)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.GetPos()

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvTrackbarSetPos(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Trackbar)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Trackbar)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Trackbar)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)

	arg_0_out.SetPos(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvTrackbarSetMin(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Trackbar)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Trackbar)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Trackbar)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)

	arg_0_out.SetMin(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvTrackbarSetMax(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Trackbar)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Trackbar)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Trackbar)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)

	arg_0_out.SetMax(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvNewMat(call otto.FunctionCall) otto.Value {

	res_0 := gocv.NewMat()

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNewMatWithSize(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_int64, ok := arg_0_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(int64)"))
	}
	arg_0_out := int(arg_0_out_int64)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_MatType_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_MatType := int(arg_2_out_MatType_int64)
	arg_2_out := gocv.MatType(arg_2_out_MatType)

	res_0 := gocv.NewMatWithSize(arg_0_out, arg_1_out, arg_2_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNewMatFromScalar(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Scalar := gocv.Scalar{}
	arg_0_out_Scalar_map, ok := arg_0_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(map[string]interface{})"))
	}
	arg_0_out_Scalar_Val1_float64, ok := arg_0_out_Scalar_map["Val1"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_Scalar_map_Val1.(float64)"))
	}
	arg_0_out_Scalar.Val1 = float64(arg_0_out_Scalar_Val1_float64)
	arg_0_out_Scalar_Val2_float64, ok := arg_0_out_Scalar_map["Val2"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_Scalar_map_Val2.(float64)"))
	}
	arg_0_out_Scalar.Val2 = float64(arg_0_out_Scalar_Val2_float64)
	arg_0_out_Scalar_Val3_float64, ok := arg_0_out_Scalar_map["Val3"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_Scalar_map_Val3.(float64)"))
	}
	arg_0_out_Scalar.Val3 = float64(arg_0_out_Scalar_Val3_float64)
	arg_0_out_Scalar_Val4_float64, ok := arg_0_out_Scalar_map["Val4"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_Scalar_map_Val4.(float64)"))
	}
	arg_0_out_Scalar.Val4 = float64(arg_0_out_Scalar_Val4_float64)
	arg_0_out := gocv.Scalar(arg_0_out_Scalar)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_MatType_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out_MatType := int(arg_1_out_MatType_int64)
	arg_1_out := gocv.MatType(arg_1_out_MatType)

	res_0 := gocv.NewMatFromScalar(arg_0_out, arg_1_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNewMatWithSizeFromScalar(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Scalar := gocv.Scalar{}
	arg_0_out_Scalar_map, ok := arg_0_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(map[string]interface{})"))
	}
	arg_0_out_Scalar_Val1_float64, ok := arg_0_out_Scalar_map["Val1"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_Scalar_map_Val1.(float64)"))
	}
	arg_0_out_Scalar.Val1 = float64(arg_0_out_Scalar_Val1_float64)
	arg_0_out_Scalar_Val2_float64, ok := arg_0_out_Scalar_map["Val2"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_Scalar_map_Val2.(float64)"))
	}
	arg_0_out_Scalar.Val2 = float64(arg_0_out_Scalar_Val2_float64)
	arg_0_out_Scalar_Val3_float64, ok := arg_0_out_Scalar_map["Val3"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_Scalar_map_Val3.(float64)"))
	}
	arg_0_out_Scalar.Val3 = float64(arg_0_out_Scalar_Val3_float64)
	arg_0_out_Scalar_Val4_float64, ok := arg_0_out_Scalar_map["Val4"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_Scalar_map_Val4.(float64)"))
	}
	arg_0_out_Scalar.Val4 = float64(arg_0_out_Scalar_Val4_float64)
	arg_0_out := gocv.Scalar(arg_0_out_Scalar)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_MatType_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out_MatType := int(arg_3_out_MatType_int64)
	arg_3_out := gocv.MatType(arg_3_out_MatType)

	res_0 := gocv.NewMatWithSizeFromScalar(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNewMatFromBytes(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_int64, ok := arg_0_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(int64)"))
	}
	arg_0_out := int(arg_0_out_int64)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_MatType_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_MatType := int(arg_2_out_MatType_int64)
	arg_2_out := gocv.MatType(arg_2_out_MatType)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_string, ok := arg_3_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(string)"))
	}
	arg_3_out, err := base64strToBytes(arg_3_out_string)
	if err != nil {
		return o.handleError(err)
	}

	res_0, res_1 := gocv.NewMatFromBytes(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatFromPtr(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_MatType_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out_MatType := int(arg_3_out_MatType_int64)
	arg_3_out := gocv.MatType(arg_3_out_MatType)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_int64, ok := arg_5_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(int64)"))
	}
	arg_5_out := int(arg_5_out_int64)

	res_0, res_1 := arg_0_out.FromPtr(arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvMatEmpty(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Empty()

	res_0_out := bool(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatClone(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Clone()

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatCopyTo(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}

	arg_0_out.CopyTo(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatCopyToWithMask(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)

	arg_0_out.CopyToWithMask(arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatConvertTo(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_MatType_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_MatType := int(arg_2_out_MatType_int64)
	arg_2_out := gocv.MatType(arg_2_out_MatType)

	arg_0_out.ConvertTo(arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatTotal(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Total()

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatSize(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Size()

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr[res_0_i] = int64(res_0_v)
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatToBytes(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.ToBytes()

	res_0_out := bytesToBase64str(res_0)
	if err != nil {
		return o.handleError(err)
	}
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatDataPtrUint8(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.DataPtrUint8()

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr[res_0_i] = int64(res_0_v)
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatDataPtrInt8(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.DataPtrInt8()

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr[res_0_i] = int64(res_0_v)
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatDataPtrUint16(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0, res_1 := arg_0_out.DataPtrUint16()

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr[res_0_i] = int64(res_0_v)
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatDataPtrInt16(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0, res_1 := arg_0_out.DataPtrInt16()

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr[res_0_i] = int64(res_0_v)
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatDataPtrFloat32(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0, res_1 := arg_0_out.DataPtrFloat32()

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr[res_0_i] = float64(res_0_v)
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatDataPtrFloat64(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0, res_1 := arg_0_out.DataPtrFloat64()

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr[res_0_i] = float64(res_0_v)
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatRegion(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out := image.Rectangle{}
	arg_1_out_map, ok := arg_1_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(map[string]interface{})"))
	}
	arg_1_out.Min = image.Point{}
	arg_1_out_Min_map, ok := arg_1_out_map["Min"].(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_map_Min.(map[string]interface{})"))
	}
	arg_1_out_Min_X_int64, ok := arg_1_out_Min_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Min_map_X.(int64)"))
	}
	arg_1_out.Min.X = int(arg_1_out_Min_X_int64)
	arg_1_out_Min_Y_int64, ok := arg_1_out_Min_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Min_map_Y.(int64)"))
	}
	arg_1_out.Min.Y = int(arg_1_out_Min_Y_int64)
	arg_1_out.Max = image.Point{}
	arg_1_out_Max_map, ok := arg_1_out_map["Max"].(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_map_Max.(map[string]interface{})"))
	}
	arg_1_out_Max_X_int64, ok := arg_1_out_Max_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Max_map_X.(int64)"))
	}
	arg_1_out.Max.X = int(arg_1_out_Max_X_int64)
	arg_1_out_Max_Y_int64, ok := arg_1_out_Max_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Max_map_Y.(int64)"))
	}
	arg_1_out.Max.Y = int(arg_1_out_Max_Y_int64)

	res_0 := arg_0_out.Region(arg_1_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatReshape(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	res_0 := arg_0_out.Reshape(arg_1_out, arg_2_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatConvertFp16(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.ConvertFp16()

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatMean(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Mean()

	res_0_out_map := make(map[string]interface{})
	res_0_out_map["Val1"] = float64(res_0.Val1)
	res_0_out_map["Val2"] = float64(res_0.Val2)
	res_0_out_map["Val3"] = float64(res_0.Val3)
	res_0_out_map["Val4"] = float64(res_0.Val4)
	res_0_out := res_0_out_map
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatSqrt(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Sqrt()

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatSum(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Sum()

	res_0_out_map := make(map[string]interface{})
	res_0_out_map["Val1"] = float64(res_0.Val1)
	res_0_out_map["Val2"] = float64(res_0.Val2)
	res_0_out_map["Val3"] = float64(res_0.Val3)
	res_0_out_map["Val4"] = float64(res_0.Val4)
	res_0_out := res_0_out_map
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatPatchNaNs(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	arg_0_out.PatchNaNs()

	return otto.Value{}
}

func (o *OttoFunctions) cvLUT(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.LUT(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatRows(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Rows()

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatCols(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Cols()

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatChannels(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Channels()

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatType(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Type()

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatStep(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Step()

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatGetUCharAt(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	res_0 := arg_0_out.GetUCharAt(arg_1_out, arg_2_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatGetUCharAt3(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)

	res_0 := arg_0_out.GetUCharAt3(arg_1_out, arg_2_out, arg_3_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatGetSCharAt(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	res_0 := arg_0_out.GetSCharAt(arg_1_out, arg_2_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatGetSCharAt3(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)

	res_0 := arg_0_out.GetSCharAt3(arg_1_out, arg_2_out, arg_3_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatGetShortAt(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	res_0 := arg_0_out.GetShortAt(arg_1_out, arg_2_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatGetShortAt3(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)

	res_0 := arg_0_out.GetShortAt3(arg_1_out, arg_2_out, arg_3_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatGetIntAt(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	res_0 := arg_0_out.GetIntAt(arg_1_out, arg_2_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatGetIntAt3(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)

	res_0 := arg_0_out.GetIntAt3(arg_1_out, arg_2_out, arg_3_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatGetFloatAt(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	res_0 := arg_0_out.GetFloatAt(arg_1_out, arg_2_out)

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatGetFloatAt3(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)

	res_0 := arg_0_out.GetFloatAt3(arg_1_out, arg_2_out, arg_3_out)

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatGetDoubleAt(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	res_0 := arg_0_out.GetDoubleAt(arg_1_out, arg_2_out)

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatGetDoubleAt3(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)

	res_0 := arg_0_out.GetDoubleAt3(arg_1_out, arg_2_out, arg_3_out)

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatSetTo(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Scalar := gocv.Scalar{}
	arg_1_out_Scalar_map, ok := arg_1_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(map[string]interface{})"))
	}
	arg_1_out_Scalar_Val1_float64, ok := arg_1_out_Scalar_map["Val1"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Scalar_map_Val1.(float64)"))
	}
	arg_1_out_Scalar.Val1 = float64(arg_1_out_Scalar_Val1_float64)
	arg_1_out_Scalar_Val2_float64, ok := arg_1_out_Scalar_map["Val2"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Scalar_map_Val2.(float64)"))
	}
	arg_1_out_Scalar.Val2 = float64(arg_1_out_Scalar_Val2_float64)
	arg_1_out_Scalar_Val3_float64, ok := arg_1_out_Scalar_map["Val3"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Scalar_map_Val3.(float64)"))
	}
	arg_1_out_Scalar.Val3 = float64(arg_1_out_Scalar_Val3_float64)
	arg_1_out_Scalar_Val4_float64, ok := arg_1_out_Scalar_map["Val4"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Scalar_map_Val4.(float64)"))
	}
	arg_1_out_Scalar.Val4 = float64(arg_1_out_Scalar_Val4_float64)
	arg_1_out := gocv.Scalar(arg_1_out_Scalar)

	arg_0_out.SetTo(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatSetUCharAt(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := uint8(arg_3_out_int64)

	arg_0_out.SetUCharAt(arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatSetUCharAt3(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := uint8(arg_4_out_int64)

	arg_0_out.SetUCharAt3(arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatSetSCharAt(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int8(arg_3_out_int64)

	arg_0_out.SetSCharAt(arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatSetSCharAt3(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int8(arg_4_out_int64)

	arg_0_out.SetSCharAt3(arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatSetShortAt(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int16(arg_3_out_int64)

	arg_0_out.SetShortAt(arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatSetShortAt3(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int16(arg_4_out_int64)

	arg_0_out.SetShortAt3(arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatSetIntAt(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int32(arg_3_out_int64)

	arg_0_out.SetIntAt(arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatSetIntAt3(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int32(arg_4_out_int64)

	arg_0_out.SetIntAt3(arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatSetFloatAt(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float32(arg_3_out_float64)

	arg_0_out.SetFloatAt(arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatSetFloatAt3(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_float64, ok := arg_4_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(float64)"))
	}
	arg_4_out := float32(arg_4_out_float64)

	arg_0_out.SetFloatAt3(arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatSetDoubleAt(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float64(arg_3_out_float64)

	arg_0_out.SetDoubleAt(arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatSetDoubleAt3(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_float64, ok := arg_4_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(float64)"))
	}
	arg_4_out := float64(arg_4_out_float64)

	arg_0_out.SetDoubleAt3(arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatAddUChar(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := uint8(arg_1_out_int64)

	arg_0_out.AddUChar(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatSubtractUChar(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := uint8(arg_1_out_int64)

	arg_0_out.SubtractUChar(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatMultiplyUChar(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := uint8(arg_1_out_int64)

	arg_0_out.MultiplyUChar(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatDivideUChar(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := uint8(arg_1_out_int64)

	arg_0_out.DivideUChar(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatAddFloat(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_float64, ok := arg_1_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(float64)"))
	}
	arg_1_out := float32(arg_1_out_float64)

	arg_0_out.AddFloat(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatSubtractFloat(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_float64, ok := arg_1_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(float64)"))
	}
	arg_1_out := float32(arg_1_out_float64)

	arg_0_out.SubtractFloat(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatMultiplyFloat(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_float64, ok := arg_1_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(float64)"))
	}
	arg_1_out := float32(arg_1_out_float64)

	arg_0_out.MultiplyFloat(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatDivideFloat(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_float64, ok := arg_1_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(float64)"))
	}
	arg_1_out := float32(arg_1_out_float64)

	arg_0_out.DivideFloat(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMatToImage(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0, res_1 := arg_0_out.ToImage()

	res_0_out := image.Image(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvImageToMatRGBA(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(image.Image)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(*image.Image)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(image.Image)"))
		}
		arg_0_out = *(arg_0_out_tmp)
	}

	res_0, res_1 := gocv.ImageToMatRGBA(arg_0_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvImageToMatRGB(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(image.Image)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(*image.Image)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(image.Image)"))
		}
		arg_0_out = *(arg_0_out_tmp)
	}

	res_0, res_1 := gocv.ImageToMatRGB(arg_0_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvImageGrayToMatGray(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*image.Gray)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(image.Gray)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(image.Gray)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0, res_1 := gocv.ImageGrayToMatGray(arg_0_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvAbsDiff(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.AbsDiff(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvAdd(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.Add(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvAddWeighted(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_float64, ok := arg_1_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(float64)"))
	}
	arg_1_out := float64(arg_1_out_float64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float64(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_float64, ok := arg_4_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(float64)"))
	}
	arg_4_out := float64(arg_4_out_float64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out, ok := arg_5_in.(*gocv.Mat)
	if !ok {
		arg_5_out_tmp, ok := arg_5_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(gocv.Mat)"))
		}
		arg_5_out = &(arg_5_out_tmp)
	}

	gocv.AddWeighted(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvBitwiseAnd(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.BitwiseAnd(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvBitwiseNot(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}

	gocv.BitwiseNot(arg_0_out, arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvBitwiseOr(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.BitwiseOr(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvBitwiseXor(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.BitwiseXor(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvBatchDistance(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_Mat, ok := arg_4_in.(gocv.Mat)
	if !ok {
		arg_4_out_Mat_tmp, ok := arg_4_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(gocv.Mat)"))
		}
		arg_4_out_Mat = *(arg_4_out_Mat_tmp)
	}
	arg_4_out := gocv.Mat(arg_4_out_Mat)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_int64, ok := arg_5_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(int64)"))
	}
	arg_5_out := int(arg_5_out_int64)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out_int64, ok := arg_6_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(int64)"))
	}
	arg_6_out := int(arg_6_out_int64)
	arg_7_in, err := call.Argument(7).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_7_out_Mat, ok := arg_7_in.(gocv.Mat)
	if !ok {
		arg_7_out_Mat_tmp, ok := arg_7_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_7_in.(gocv.Mat)"))
		}
		arg_7_out_Mat = *(arg_7_out_Mat_tmp)
	}
	arg_7_out := gocv.Mat(arg_7_out_Mat)
	arg_8_in, err := call.Argument(8).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_8_out_int64, ok := arg_8_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_8_in.(int64)"))
	}
	arg_8_out := int(arg_8_out_int64)
	arg_9_in, err := call.Argument(9).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_9_out_bool, ok := arg_9_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_9_in.(bool)"))
	}
	arg_9_out := bool(arg_9_out_bool)

	gocv.BatchDistance(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out, arg_7_out, arg_8_out, arg_9_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvBorderInterpolate(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_int64, ok := arg_0_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(int64)"))
	}
	arg_0_out := int(arg_0_out_int64)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_CovarFlags_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_CovarFlags := int(arg_2_out_CovarFlags_int64)
	arg_2_out := gocv.CovarFlags(arg_2_out_CovarFlags)

	res_0 := gocv.BorderInterpolate(arg_0_out, arg_1_out, arg_2_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvCalcCovarMatrix(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_CovarFlags_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out_CovarFlags := int(arg_3_out_CovarFlags_int64)
	arg_3_out := gocv.CovarFlags(arg_3_out_CovarFlags)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)

	gocv.CalcCovarMatrix(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvCartToPolar(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out, ok := arg_3_in.(*gocv.Mat)
	if !ok {
		arg_3_out_tmp, ok := arg_3_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(gocv.Mat)"))
		}
		arg_3_out = &(arg_3_out_tmp)
	}
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_bool, ok := arg_4_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(bool)"))
	}
	arg_4_out := bool(arg_4_out_bool)

	gocv.CartToPolar(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvCheckRange(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)

	res_0 := gocv.CheckRange(arg_0_out)

	res_0_out := bool(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvCompare(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_CompareType_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out_CompareType := int(arg_3_out_CompareType_int64)
	arg_3_out := gocv.CompareType(arg_3_out_CompareType)

	gocv.Compare(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvCountNonZero(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)

	res_0 := gocv.CountNonZero(arg_0_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvCompleteSymm(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_bool, ok := arg_1_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(bool)"))
	}
	arg_1_out := bool(arg_1_out_bool)

	gocv.CompleteSymm(arg_0_out, arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvConvertScaleAbs(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float64(arg_2_out_float64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float64(arg_3_out_float64)

	gocv.ConvertScaleAbs(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvCopyMakeBorder(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_int64, ok := arg_5_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(int64)"))
	}
	arg_5_out := int(arg_5_out_int64)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out_BorderType_int64, ok := arg_6_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(int64)"))
	}
	arg_6_out_BorderType := int(arg_6_out_BorderType_int64)
	arg_6_out := gocv.BorderType(arg_6_out_BorderType)
	arg_7_in, err := call.Argument(7).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_7_out, ok := arg_7_in.(color.RGBA)
	if !ok {
		arg_7_out_tmp, ok := arg_7_in.(*color.RGBA)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_7_in.(color.RGBA)"))
		}
		arg_7_out = *(arg_7_out_tmp)
	}

	gocv.CopyMakeBorder(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out, arg_7_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvDCT(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_DftFlags_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_DftFlags := int(arg_2_out_DftFlags_int64)
	arg_2_out := gocv.DftFlags(arg_2_out_DftFlags)

	gocv.DCT(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvDeterminant(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)

	res_0 := gocv.Determinant(arg_0_out)

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvDFT(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_DftFlags_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_DftFlags := int(arg_2_out_DftFlags_int64)
	arg_2_out := gocv.DftFlags(arg_2_out_DftFlags)

	gocv.DFT(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvDivide(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.Divide(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvEigen(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	res_0 := gocv.Eigen(arg_0_out, arg_1_out, arg_2_out)

	res_0_out := bool(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvEigenNonSymmetric(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.EigenNonSymmetric(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvExp(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}

	gocv.Exp(arg_0_out, arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvExtractChannel(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	gocv.ExtractChannel(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvFindNonZero(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}

	gocv.FindNonZero(arg_0_out, arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvFlip(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	gocv.Flip(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvGemm(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float64(arg_2_out_float64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_Mat, ok := arg_3_in.(gocv.Mat)
	if !ok {
		arg_3_out_Mat_tmp, ok := arg_3_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(gocv.Mat)"))
		}
		arg_3_out_Mat = *(arg_3_out_Mat_tmp)
	}
	arg_3_out := gocv.Mat(arg_3_out_Mat)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_float64, ok := arg_4_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(float64)"))
	}
	arg_4_out := float64(arg_4_out_float64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out, ok := arg_5_in.(*gocv.Mat)
	if !ok {
		arg_5_out_tmp, ok := arg_5_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(gocv.Mat)"))
		}
		arg_5_out = &(arg_5_out_tmp)
	}
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out_int64, ok := arg_6_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(int64)"))
	}
	arg_6_out := int(arg_6_out_int64)

	gocv.Gemm(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvGetOptimalDFTSize(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_int64, ok := arg_0_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(int64)"))
	}
	arg_0_out := int(arg_0_out_int64)

	res_0 := gocv.GetOptimalDFTSize(arg_0_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvHconcat(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.Hconcat(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvVconcat(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.Vconcat(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvRotate(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_RotateFlag_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_RotateFlag := int(arg_2_out_RotateFlag_int64)
	arg_2_out := gocv.RotateFlag(arg_2_out_RotateFlag)

	gocv.Rotate(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvIDCT(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	gocv.IDCT(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvIDFT(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)

	gocv.IDFT(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvInRange(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out, ok := arg_3_in.(*gocv.Mat)
	if !ok {
		arg_3_out_tmp, ok := arg_3_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(gocv.Mat)"))
		}
		arg_3_out = &(arg_3_out_tmp)
	}

	gocv.InRange(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvInRangeWithScalar(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Scalar := gocv.Scalar{}
	arg_1_out_Scalar_map, ok := arg_1_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(map[string]interface{})"))
	}
	arg_1_out_Scalar_Val1_float64, ok := arg_1_out_Scalar_map["Val1"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Scalar_map_Val1.(float64)"))
	}
	arg_1_out_Scalar.Val1 = float64(arg_1_out_Scalar_Val1_float64)
	arg_1_out_Scalar_Val2_float64, ok := arg_1_out_Scalar_map["Val2"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Scalar_map_Val2.(float64)"))
	}
	arg_1_out_Scalar.Val2 = float64(arg_1_out_Scalar_Val2_float64)
	arg_1_out_Scalar_Val3_float64, ok := arg_1_out_Scalar_map["Val3"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Scalar_map_Val3.(float64)"))
	}
	arg_1_out_Scalar.Val3 = float64(arg_1_out_Scalar_Val3_float64)
	arg_1_out_Scalar_Val4_float64, ok := arg_1_out_Scalar_map["Val4"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Scalar_map_Val4.(float64)"))
	}
	arg_1_out_Scalar.Val4 = float64(arg_1_out_Scalar_Val4_float64)
	arg_1_out := gocv.Scalar(arg_1_out_Scalar)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Scalar := gocv.Scalar{}
	arg_2_out_Scalar_map, ok := arg_2_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(map[string]interface{})"))
	}
	arg_2_out_Scalar_Val1_float64, ok := arg_2_out_Scalar_map["Val1"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_Scalar_map_Val1.(float64)"))
	}
	arg_2_out_Scalar.Val1 = float64(arg_2_out_Scalar_Val1_float64)
	arg_2_out_Scalar_Val2_float64, ok := arg_2_out_Scalar_map["Val2"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_Scalar_map_Val2.(float64)"))
	}
	arg_2_out_Scalar.Val2 = float64(arg_2_out_Scalar_Val2_float64)
	arg_2_out_Scalar_Val3_float64, ok := arg_2_out_Scalar_map["Val3"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_Scalar_map_Val3.(float64)"))
	}
	arg_2_out_Scalar.Val3 = float64(arg_2_out_Scalar_Val3_float64)
	arg_2_out_Scalar_Val4_float64, ok := arg_2_out_Scalar_map["Val4"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_Scalar_map_Val4.(float64)"))
	}
	arg_2_out_Scalar.Val4 = float64(arg_2_out_Scalar_Val4_float64)
	arg_2_out := gocv.Scalar(arg_2_out_Scalar)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out, ok := arg_3_in.(*gocv.Mat)
	if !ok {
		arg_3_out_tmp, ok := arg_3_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(gocv.Mat)"))
		}
		arg_3_out = &(arg_3_out_tmp)
	}

	gocv.InRangeWithScalar(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvInsertChannel(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	gocv.InsertChannel(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvInvert(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	res_0 := gocv.Invert(arg_0_out, arg_1_out, arg_2_out)

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvLog(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}

	gocv.Log(arg_0_out, arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMagnitude(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.Magnitude(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMax(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.Max(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMeanStdDev(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.MeanStdDev(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMerge(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_in_arr, ok := arg_0_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.([]interface{})"))
	}
	arg_0_out := make([]gocv.Mat, len(arg_0_in_arr), len(arg_0_in_arr))
	for arg_0_in_arr_i, arg_0_in_arr_v := range arg_0_in_arr {
		arg_0_out_arg_0_in_arr_i_Mat, ok := arg_0_in_arr_v.(gocv.Mat)
		if !ok {
			arg_0_out_arg_0_in_arr_i_Mat_tmp, ok := arg_0_in_arr_v.(*gocv.Mat)
			if !ok {
				return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in_arr_v.(gocv.Mat)"))
			}
			arg_0_out_arg_0_in_arr_i_Mat = *(arg_0_out_arg_0_in_arr_i_Mat_tmp)
		}
		arg_0_out[arg_0_in_arr_i] = gocv.Mat(arg_0_out_arg_0_in_arr_i_Mat)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}

	gocv.Merge(arg_0_out, arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMin(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.Min(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMinMaxIdx(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)

	res_0, res_1, res_2, res_3 := gocv.MinMaxIdx(arg_0_out)

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	res_1_out := float64(res_1)
	res_1v, err := o.vm.ToValue(res_1_out)
	if err != nil {
		return o.handleError(err)
	}
	res_2_out := int64(res_2)
	res_2v, err := o.vm.ToValue(res_2_out)
	if err != nil {
		return o.handleError(err)
	}
	res_3_out := int64(res_3)
	res_3v, err := o.vm.ToValue(res_3_out)
	if err != nil {
		return o.handleError(err)
	}
	results := map[string]interface{}{}
	results["r0"] = res_0v
	results["r1"] = res_1v
	results["r2"] = res_2v
	results["r3"] = res_3v
	resultsv, err := o.vm.ToValue(results)
	if err != nil {
		return o.handleError(err)
	}
	return resultsv
}

func (o *OttoFunctions) cvMinMaxLoc(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)

	res_0, res_1, res_2, res_3 := gocv.MinMaxLoc(arg_0_out)

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	res_1_out := float64(res_1)
	res_1v, err := o.vm.ToValue(res_1_out)
	if err != nil {
		return o.handleError(err)
	}
	res_2_out_map := make(map[string]interface{})
	res_2_out_map["X"] = int64(res_2.X)
	res_2_out_map["Y"] = int64(res_2.Y)
	res_2_out := res_2_out_map
	res_2v, err := o.vm.ToValue(res_2_out)
	if err != nil {
		return o.handleError(err)
	}
	res_3_out_map := make(map[string]interface{})
	res_3_out_map["X"] = int64(res_3.X)
	res_3_out_map["Y"] = int64(res_3.Y)
	res_3_out := res_3_out_map
	res_3v, err := o.vm.ToValue(res_3_out)
	if err != nil {
		return o.handleError(err)
	}
	results := map[string]interface{}{}
	results["r0"] = res_0v
	results["r1"] = res_1v
	results["r2"] = res_2v
	results["r3"] = res_3v
	resultsv, err := o.vm.ToValue(results)
	if err != nil {
		return o.handleError(err)
	}
	return resultsv
}

func (o *OttoFunctions) cvMulSpectrums(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_DftFlags_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out_DftFlags := int(arg_3_out_DftFlags_int64)
	arg_3_out := gocv.DftFlags(arg_3_out_DftFlags)

	gocv.MulSpectrums(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMultiply(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.Multiply(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvNormalize(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float64(arg_2_out_float64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float64(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_NormType_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out_NormType := int(arg_4_out_NormType_int64)
	arg_4_out := gocv.NormType(arg_4_out_NormType)

	gocv.Normalize(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvNorm(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_NormType_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out_NormType := int(arg_1_out_NormType_int64)
	arg_1_out := gocv.NormType(arg_1_out_NormType)

	res_0 := gocv.Norm(arg_0_out, arg_1_out)

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvPerspectiveTransform(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)

	gocv.PerspectiveTransform(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvSolve(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_SolveDecompositionFlags_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out_SolveDecompositionFlags := int(arg_3_out_SolveDecompositionFlags_int64)
	arg_3_out := gocv.SolveDecompositionFlags(arg_3_out_SolveDecompositionFlags)

	res_0 := gocv.Solve(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	res_0_out := bool(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvSolveCubic(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}

	res_0 := gocv.SolveCubic(arg_0_out, arg_1_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvSolvePoly(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	res_0 := gocv.SolvePoly(arg_0_out, arg_1_out, arg_2_out)

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvReduce(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_ReduceTypes_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out_ReduceTypes := int(arg_3_out_ReduceTypes_int64)
	arg_3_out := gocv.ReduceTypes(arg_3_out_ReduceTypes)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)

	gocv.Reduce(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvRepeat(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out, ok := arg_3_in.(*gocv.Mat)
	if !ok {
		arg_3_out_tmp, ok := arg_3_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(gocv.Mat)"))
		}
		arg_3_out = &(arg_3_out_tmp)
	}

	gocv.Repeat(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvSort(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_SortFlags_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_SortFlags := int(arg_2_out_SortFlags_int64)
	arg_2_out := gocv.SortFlags(arg_2_out_SortFlags)

	gocv.Sort(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvSortIdx(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_SortFlags_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_SortFlags := int(arg_2_out_SortFlags_int64)
	arg_2_out := gocv.SortFlags(arg_2_out_SortFlags)

	gocv.SortIdx(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvSplit(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)

	res_0 := gocv.Split(arg_0_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr[res_0_i] = gocv.Mat(res_0_v)
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvSubtract(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.Subtract(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvTrace(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)

	res_0 := gocv.Trace(arg_0_out)

	res_0_out_map := make(map[string]interface{})
	res_0_out_map["Val1"] = float64(res_0.Val1)
	res_0_out_map["Val2"] = float64(res_0.Val2)
	res_0_out_map["Val3"] = float64(res_0.Val3)
	res_0_out_map["Val4"] = float64(res_0.Val4)
	res_0_out := res_0_out_map
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvTransform(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)

	gocv.Transform(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvTranspose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}

	gocv.Transpose(arg_0_out, arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvPow(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_float64, ok := arg_1_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(float64)"))
	}
	arg_1_out := float64(arg_1_out_float64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.Pow(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvPolarToCart(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out, ok := arg_3_in.(*gocv.Mat)
	if !ok {
		arg_3_out_tmp, ok := arg_3_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(gocv.Mat)"))
		}
		arg_3_out = &(arg_3_out_tmp)
	}
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_bool, ok := arg_4_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(bool)"))
	}
	arg_4_out := bool(arg_4_out_bool)

	gocv.PolarToCart(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvPhase(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_bool, ok := arg_3_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(bool)"))
	}
	arg_3_out := bool(arg_3_out_bool)

	gocv.Phase(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvNewTermCriteria(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_TermCriteriaType_int64, ok := arg_0_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(int64)"))
	}
	arg_0_out_TermCriteriaType := int(arg_0_out_TermCriteriaType_int64)
	arg_0_out := gocv.TermCriteriaType(arg_0_out_TermCriteriaType)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float64(arg_2_out_float64)

	res_0 := gocv.NewTermCriteria(arg_0_out, arg_1_out, arg_2_out)

	res_0_out := gocv.TermCriteria(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNewScalar(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_float64, ok := arg_0_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(float64)"))
	}
	arg_0_out := float64(arg_0_out_float64)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_float64, ok := arg_1_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(float64)"))
	}
	arg_1_out := float64(arg_1_out_float64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float64(arg_2_out_float64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float64(arg_3_out_float64)

	res_0 := gocv.NewScalar(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	res_0_out_map := make(map[string]interface{})
	res_0_out_map["Val1"] = float64(res_0.Val1)
	res_0_out_map["Val2"] = float64(res_0.Val2)
	res_0_out_map["Val3"] = float64(res_0.Val3)
	res_0_out_map["Val4"] = float64(res_0.Val4)
	res_0_out := res_0_out_map
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatGetVecfAt(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	res_0 := arg_0_out.GetVecfAt(arg_1_out, arg_2_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr[res_0_i] = float64(res_0_v)
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatGetVeciAt(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	res_0 := arg_0_out.GetVeciAt(arg_1_out, arg_2_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr[res_0_i] = int64(res_0_v)
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvGetTickCount(call otto.FunctionCall) otto.Value {

	res_0 := gocv.GetTickCount()

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvGetTickFrequency(call otto.FunctionCall) otto.Value {

	res_0 := gocv.GetTickFrequency()

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvArcLength(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_in_arr, ok := arg_0_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.([]interface{})"))
	}
	arg_0_out := make([]image.Point, len(arg_0_in_arr), len(arg_0_in_arr))
	for arg_0_in_arr_i, arg_0_in_arr_v := range arg_0_in_arr {
		arg_0_out[arg_0_in_arr_i] = image.Point{}
		arg_0_out_arg_0_in_arr_i_map, ok := arg_0_in_arr_v.(map[string]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in_arr_v.(map[string]interface{})"))
		}
		arg_0_out_arg_0_in_arr_i_X_int64, ok := arg_0_out_arg_0_in_arr_i_map["X"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_X.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].X = int(arg_0_out_arg_0_in_arr_i_X_int64)
		arg_0_out_arg_0_in_arr_i_Y_int64, ok := arg_0_out_arg_0_in_arr_i_map["Y"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_Y.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].Y = int(arg_0_out_arg_0_in_arr_i_Y_int64)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_bool, ok := arg_1_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(bool)"))
	}
	arg_1_out := bool(arg_1_out_bool)

	res_0 := gocv.ArcLength(arg_0_out, arg_1_out)

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvApproxPolyDP(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_in_arr, ok := arg_0_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.([]interface{})"))
	}
	arg_0_out := make([]image.Point, len(arg_0_in_arr), len(arg_0_in_arr))
	for arg_0_in_arr_i, arg_0_in_arr_v := range arg_0_in_arr {
		arg_0_out[arg_0_in_arr_i] = image.Point{}
		arg_0_out_arg_0_in_arr_i_map, ok := arg_0_in_arr_v.(map[string]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in_arr_v.(map[string]interface{})"))
		}
		arg_0_out_arg_0_in_arr_i_X_int64, ok := arg_0_out_arg_0_in_arr_i_map["X"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_X.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].X = int(arg_0_out_arg_0_in_arr_i_X_int64)
		arg_0_out_arg_0_in_arr_i_Y_int64, ok := arg_0_out_arg_0_in_arr_i_map["Y"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_Y.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].Y = int(arg_0_out_arg_0_in_arr_i_Y_int64)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_float64, ok := arg_1_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(float64)"))
	}
	arg_1_out := float64(arg_1_out_float64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_bool, ok := arg_2_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(bool)"))
	}
	arg_2_out := bool(arg_2_out_bool)

	res_0 := gocv.ApproxPolyDP(arg_0_out, arg_1_out, arg_2_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map["X"] = int64(res_0_v.X)
		res_0_out_arr_res_0_i_map["Y"] = int64(res_0_v.Y)
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvConvexHull(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_in_arr, ok := arg_0_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.([]interface{})"))
	}
	arg_0_out := make([]image.Point, len(arg_0_in_arr), len(arg_0_in_arr))
	for arg_0_in_arr_i, arg_0_in_arr_v := range arg_0_in_arr {
		arg_0_out[arg_0_in_arr_i] = image.Point{}
		arg_0_out_arg_0_in_arr_i_map, ok := arg_0_in_arr_v.(map[string]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in_arr_v.(map[string]interface{})"))
		}
		arg_0_out_arg_0_in_arr_i_X_int64, ok := arg_0_out_arg_0_in_arr_i_map["X"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_X.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].X = int(arg_0_out_arg_0_in_arr_i_X_int64)
		arg_0_out_arg_0_in_arr_i_Y_int64, ok := arg_0_out_arg_0_in_arr_i_map["Y"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_Y.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].Y = int(arg_0_out_arg_0_in_arr_i_Y_int64)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_bool, ok := arg_2_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(bool)"))
	}
	arg_2_out := bool(arg_2_out_bool)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_bool, ok := arg_3_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(bool)"))
	}
	arg_3_out := bool(arg_3_out_bool)

	gocv.ConvexHull(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvConvexityDefects(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_in_arr, ok := arg_0_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.([]interface{})"))
	}
	arg_0_out := make([]image.Point, len(arg_0_in_arr), len(arg_0_in_arr))
	for arg_0_in_arr_i, arg_0_in_arr_v := range arg_0_in_arr {
		arg_0_out[arg_0_in_arr_i] = image.Point{}
		arg_0_out_arg_0_in_arr_i_map, ok := arg_0_in_arr_v.(map[string]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in_arr_v.(map[string]interface{})"))
		}
		arg_0_out_arg_0_in_arr_i_X_int64, ok := arg_0_out_arg_0_in_arr_i_map["X"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_X.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].X = int(arg_0_out_arg_0_in_arr_i_X_int64)
		arg_0_out_arg_0_in_arr_i_Y_int64, ok := arg_0_out_arg_0_in_arr_i_map["Y"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_Y.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].Y = int(arg_0_out_arg_0_in_arr_i_Y_int64)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	gocv.ConvexityDefects(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvCvtColor(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_ColorConversionCode_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_ColorConversionCode := int(arg_2_out_ColorConversionCode_int64)
	arg_2_out := gocv.ColorConversionCode(arg_2_out_ColorConversionCode)

	gocv.CvtColor(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvEqualizeHist(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}

	gocv.EqualizeHist(arg_0_out, arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvCalcHist(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_in_arr, ok := arg_0_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.([]interface{})"))
	}
	arg_0_out := make([]gocv.Mat, len(arg_0_in_arr), len(arg_0_in_arr))
	for arg_0_in_arr_i, arg_0_in_arr_v := range arg_0_in_arr {
		arg_0_out_arg_0_in_arr_i_Mat, ok := arg_0_in_arr_v.(gocv.Mat)
		if !ok {
			arg_0_out_arg_0_in_arr_i_Mat_tmp, ok := arg_0_in_arr_v.(*gocv.Mat)
			if !ok {
				return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in_arr_v.(gocv.Mat)"))
			}
			arg_0_out_arg_0_in_arr_i_Mat = *(arg_0_out_arg_0_in_arr_i_Mat_tmp)
		}
		arg_0_out[arg_0_in_arr_i] = gocv.Mat(arg_0_out_arg_0_in_arr_i_Mat)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_in_arr, ok := arg_1_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.([]interface{})"))
	}
	arg_1_out := make([]int, len(arg_1_in_arr), len(arg_1_in_arr))
	for arg_1_in_arr_i, arg_1_in_arr_v := range arg_1_in_arr {
		arg_1_out_arg_1_in_arr_i_int64, ok := arg_1_in_arr_v.(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in_arr_v.(int64)"))
		}
		arg_1_out[arg_1_in_arr_i] = int(arg_1_out_arg_1_in_arr_i_int64)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out, ok := arg_3_in.(*gocv.Mat)
	if !ok {
		arg_3_out_tmp, ok := arg_3_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(gocv.Mat)"))
		}
		arg_3_out = &(arg_3_out_tmp)
	}
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_in_arr, ok := arg_4_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.([]interface{})"))
	}
	arg_4_out := make([]int, len(arg_4_in_arr), len(arg_4_in_arr))
	for arg_4_in_arr_i, arg_4_in_arr_v := range arg_4_in_arr {
		arg_4_out_arg_4_in_arr_i_int64, ok := arg_4_in_arr_v.(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in_arr_v.(int64)"))
		}
		arg_4_out[arg_4_in_arr_i] = int(arg_4_out_arg_4_in_arr_i_int64)
	}
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_in_arr, ok := arg_5_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.([]interface{})"))
	}
	arg_5_out := make([]float64, len(arg_5_in_arr), len(arg_5_in_arr))
	for arg_5_in_arr_i, arg_5_in_arr_v := range arg_5_in_arr {
		arg_5_out_arg_5_in_arr_i_float64, ok := arg_5_in_arr_v.(float64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in_arr_v.(float64)"))
		}
		arg_5_out[arg_5_in_arr_i] = float64(arg_5_out_arg_5_in_arr_i_float64)
	}
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out_bool, ok := arg_6_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(bool)"))
	}
	arg_6_out := bool(arg_6_out_bool)

	gocv.CalcHist(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvBilateralFilter(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float64(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_float64, ok := arg_4_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(float64)"))
	}
	arg_4_out := float64(arg_4_out_float64)

	gocv.BilateralFilter(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvBlur(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out := image.Point{}
	arg_2_out_map, ok := arg_2_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(map[string]interface{})"))
	}
	arg_2_out_X_int64, ok := arg_2_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_X.(int64)"))
	}
	arg_2_out.X = int(arg_2_out_X_int64)
	arg_2_out_Y_int64, ok := arg_2_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_Y.(int64)"))
	}
	arg_2_out.Y = int(arg_2_out_Y_int64)

	gocv.Blur(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvBoxFilter(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out := image.Point{}
	arg_3_out_map, ok := arg_3_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(map[string]interface{})"))
	}
	arg_3_out_X_int64, ok := arg_3_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_map_X.(int64)"))
	}
	arg_3_out.X = int(arg_3_out_X_int64)
	arg_3_out_Y_int64, ok := arg_3_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_map_Y.(int64)"))
	}
	arg_3_out.Y = int(arg_3_out_Y_int64)

	gocv.BoxFilter(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvSqBoxFilter(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out := image.Point{}
	arg_3_out_map, ok := arg_3_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(map[string]interface{})"))
	}
	arg_3_out_X_int64, ok := arg_3_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_map_X.(int64)"))
	}
	arg_3_out.X = int(arg_3_out_X_int64)
	arg_3_out_Y_int64, ok := arg_3_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_map_Y.(int64)"))
	}
	arg_3_out.Y = int(arg_3_out_Y_int64)

	gocv.SqBoxFilter(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvDilate(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)

	gocv.Dilate(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvErode(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)

	gocv.Erode(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvBoundingRect(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_in_arr, ok := arg_0_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.([]interface{})"))
	}
	arg_0_out := make([]image.Point, len(arg_0_in_arr), len(arg_0_in_arr))
	for arg_0_in_arr_i, arg_0_in_arr_v := range arg_0_in_arr {
		arg_0_out[arg_0_in_arr_i] = image.Point{}
		arg_0_out_arg_0_in_arr_i_map, ok := arg_0_in_arr_v.(map[string]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in_arr_v.(map[string]interface{})"))
		}
		arg_0_out_arg_0_in_arr_i_X_int64, ok := arg_0_out_arg_0_in_arr_i_map["X"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_X.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].X = int(arg_0_out_arg_0_in_arr_i_X_int64)
		arg_0_out_arg_0_in_arr_i_Y_int64, ok := arg_0_out_arg_0_in_arr_i_map["Y"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_Y.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].Y = int(arg_0_out_arg_0_in_arr_i_Y_int64)
	}

	res_0 := gocv.BoundingRect(arg_0_out)

	res_0_out_map := make(map[string]interface{})
	res_0_out_map_Min_map := make(map[string]interface{})
	res_0_out_map_Min_map["X"] = int64(res_0.Min.X)
	res_0_out_map_Min_map["Y"] = int64(res_0.Min.Y)
	res_0_out_map["Min"] = res_0_out_map_Min_map
	res_0_out_map_Max_map := make(map[string]interface{})
	res_0_out_map_Max_map["X"] = int64(res_0.Max.X)
	res_0_out_map_Max_map["Y"] = int64(res_0.Max.Y)
	res_0_out_map["Max"] = res_0_out_map_Max_map
	res_0_out := res_0_out_map
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvContourArea(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_in_arr, ok := arg_0_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.([]interface{})"))
	}
	arg_0_out := make([]image.Point, len(arg_0_in_arr), len(arg_0_in_arr))
	for arg_0_in_arr_i, arg_0_in_arr_v := range arg_0_in_arr {
		arg_0_out[arg_0_in_arr_i] = image.Point{}
		arg_0_out_arg_0_in_arr_i_map, ok := arg_0_in_arr_v.(map[string]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in_arr_v.(map[string]interface{})"))
		}
		arg_0_out_arg_0_in_arr_i_X_int64, ok := arg_0_out_arg_0_in_arr_i_map["X"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_X.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].X = int(arg_0_out_arg_0_in_arr_i_X_int64)
		arg_0_out_arg_0_in_arr_i_Y_int64, ok := arg_0_out_arg_0_in_arr_i_map["Y"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_Y.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].Y = int(arg_0_out_arg_0_in_arr_i_Y_int64)
	}

	res_0 := gocv.ContourArea(arg_0_out)

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMinAreaRect(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_in_arr, ok := arg_0_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.([]interface{})"))
	}
	arg_0_out := make([]image.Point, len(arg_0_in_arr), len(arg_0_in_arr))
	for arg_0_in_arr_i, arg_0_in_arr_v := range arg_0_in_arr {
		arg_0_out[arg_0_in_arr_i] = image.Point{}
		arg_0_out_arg_0_in_arr_i_map, ok := arg_0_in_arr_v.(map[string]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in_arr_v.(map[string]interface{})"))
		}
		arg_0_out_arg_0_in_arr_i_X_int64, ok := arg_0_out_arg_0_in_arr_i_map["X"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_X.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].X = int(arg_0_out_arg_0_in_arr_i_X_int64)
		arg_0_out_arg_0_in_arr_i_Y_int64, ok := arg_0_out_arg_0_in_arr_i_map["Y"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_Y.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].Y = int(arg_0_out_arg_0_in_arr_i_Y_int64)
	}

	res_0 := gocv.MinAreaRect(arg_0_out)

	res_0_out_map := make(map[string]interface{})
	res_0_out_map_Contour_arr := make([]interface{}, len(res_0.Contour), len(res_0.Contour))
	for res_0_Contour_i, res_0_Contour_v := range res_0.Contour {
		res_0_out_map_Contour_arr_res_0_Contour_i_map := make(map[string]interface{})
		res_0_out_map_Contour_arr_res_0_Contour_i_map["X"] = int64(res_0_Contour_v.X)
		res_0_out_map_Contour_arr_res_0_Contour_i_map["Y"] = int64(res_0_Contour_v.Y)
		res_0_out_map_Contour_arr[res_0_Contour_i] = res_0_out_map_Contour_arr_res_0_Contour_i_map
	}
	res_0_out_map["Contour"] = res_0_out_map_Contour_arr
	res_0_out_map_BoundingRect_map := make(map[string]interface{})
	res_0_out_map_BoundingRect_map_Min_map := make(map[string]interface{})
	res_0_out_map_BoundingRect_map_Min_map["X"] = int64(res_0.BoundingRect.Min.X)
	res_0_out_map_BoundingRect_map_Min_map["Y"] = int64(res_0.BoundingRect.Min.Y)
	res_0_out_map_BoundingRect_map["Min"] = res_0_out_map_BoundingRect_map_Min_map
	res_0_out_map_BoundingRect_map_Max_map := make(map[string]interface{})
	res_0_out_map_BoundingRect_map_Max_map["X"] = int64(res_0.BoundingRect.Max.X)
	res_0_out_map_BoundingRect_map_Max_map["Y"] = int64(res_0.BoundingRect.Max.Y)
	res_0_out_map_BoundingRect_map["Max"] = res_0_out_map_BoundingRect_map_Max_map
	res_0_out_map["BoundingRect"] = res_0_out_map_BoundingRect_map
	res_0_out_map_Center_map := make(map[string]interface{})
	res_0_out_map_Center_map["X"] = int64(res_0.Center.X)
	res_0_out_map_Center_map["Y"] = int64(res_0.Center.Y)
	res_0_out_map["Center"] = res_0_out_map_Center_map
	res_0_out_map["Width"] = int64(res_0.Width)
	res_0_out_map["Height"] = int64(res_0.Height)
	res_0_out_map["Angle"] = float64(res_0.Angle)
	res_0_out := res_0_out_map
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvFindContours(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_RetrievalMode_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out_RetrievalMode := int(arg_1_out_RetrievalMode_int64)
	arg_1_out := gocv.RetrievalMode(arg_1_out_RetrievalMode)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_ContourApproximationMode_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_ContourApproximationMode := int(arg_2_out_ContourApproximationMode_int64)
	arg_2_out := gocv.ContourApproximationMode(arg_2_out_ContourApproximationMode)

	res_0 := gocv.FindContours(arg_0_out, arg_1_out, arg_2_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_arr := make([]interface{}, len(res_0_v), len(res_0_v))
		for res_0_v_i, res_0_v_v := range res_0_v {
			res_0_out_arr_res_0_i_arr_res_0_v_i_map := make(map[string]interface{})
			res_0_out_arr_res_0_i_arr_res_0_v_i_map["X"] = int64(res_0_v_v.X)
			res_0_out_arr_res_0_i_arr_res_0_v_i_map["Y"] = int64(res_0_v_v.Y)
			res_0_out_arr_res_0_i_arr[res_0_v_i] = res_0_out_arr_res_0_i_arr_res_0_v_i_map
		}
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_arr
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMatchTemplate(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_TemplateMatchMode_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out_TemplateMatchMode := int(arg_3_out_TemplateMatchMode_int64)
	arg_3_out := gocv.TemplateMatchMode(arg_3_out_TemplateMatchMode)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_Mat, ok := arg_4_in.(gocv.Mat)
	if !ok {
		arg_4_out_Mat_tmp, ok := arg_4_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(gocv.Mat)"))
		}
		arg_4_out_Mat = *(arg_4_out_Mat_tmp)
	}
	arg_4_out := gocv.Mat(arg_4_out_Mat)

	gocv.MatchTemplate(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMoments(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_bool, ok := arg_1_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(bool)"))
	}
	arg_1_out := bool(arg_1_out_bool)

	res_0 := gocv.Moments(arg_0_out, arg_1_out)

	res_0_out_map := make(map[string]interface{})
	for res_0_i, res_0_v := range res_0 {
		res_0_out_map[res_0_i] = float64(res_0_v)
	}
	res_0_out := res_0_out_map
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvPyrDown(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out := image.Point{}
	arg_2_out_map, ok := arg_2_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(map[string]interface{})"))
	}
	arg_2_out_X_int64, ok := arg_2_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_X.(int64)"))
	}
	arg_2_out.X = int(arg_2_out_X_int64)
	arg_2_out_Y_int64, ok := arg_2_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_Y.(int64)"))
	}
	arg_2_out.Y = int(arg_2_out_Y_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_BorderType_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out_BorderType := int(arg_3_out_BorderType_int64)
	arg_3_out := gocv.BorderType(arg_3_out_BorderType)

	gocv.PyrDown(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvPyrUp(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out := image.Point{}
	arg_2_out_map, ok := arg_2_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(map[string]interface{})"))
	}
	arg_2_out_X_int64, ok := arg_2_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_X.(int64)"))
	}
	arg_2_out.X = int(arg_2_out_X_int64)
	arg_2_out_Y_int64, ok := arg_2_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_Y.(int64)"))
	}
	arg_2_out.Y = int(arg_2_out_Y_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_BorderType_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out_BorderType := int(arg_3_out_BorderType_int64)
	arg_3_out := gocv.BorderType(arg_3_out_BorderType)

	gocv.PyrUp(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMorphologyEx(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_MorphType_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_MorphType := int(arg_2_out_MorphType_int64)
	arg_2_out := gocv.MorphType(arg_2_out_MorphType)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_Mat, ok := arg_3_in.(gocv.Mat)
	if !ok {
		arg_3_out_Mat_tmp, ok := arg_3_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(gocv.Mat)"))
		}
		arg_3_out_Mat = *(arg_3_out_Mat_tmp)
	}
	arg_3_out := gocv.Mat(arg_3_out_Mat)

	gocv.MorphologyEx(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvGetStructuringElement(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_MorphShape_int64, ok := arg_0_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(int64)"))
	}
	arg_0_out_MorphShape := int(arg_0_out_MorphShape_int64)
	arg_0_out := gocv.MorphShape(arg_0_out_MorphShape)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out := image.Point{}
	arg_1_out_map, ok := arg_1_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(map[string]interface{})"))
	}
	arg_1_out_X_int64, ok := arg_1_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_map_X.(int64)"))
	}
	arg_1_out.X = int(arg_1_out_X_int64)
	arg_1_out_Y_int64, ok := arg_1_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_map_Y.(int64)"))
	}
	arg_1_out.Y = int(arg_1_out_Y_int64)

	res_0 := gocv.GetStructuringElement(arg_0_out, arg_1_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvGaussianBlur(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out := image.Point{}
	arg_2_out_map, ok := arg_2_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(map[string]interface{})"))
	}
	arg_2_out_X_int64, ok := arg_2_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_X.(int64)"))
	}
	arg_2_out.X = int(arg_2_out_X_int64)
	arg_2_out_Y_int64, ok := arg_2_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_Y.(int64)"))
	}
	arg_2_out.Y = int(arg_2_out_Y_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float64(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_float64, ok := arg_4_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(float64)"))
	}
	arg_4_out := float64(arg_4_out_float64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_BorderType_int64, ok := arg_5_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(int64)"))
	}
	arg_5_out_BorderType := int(arg_5_out_BorderType_int64)
	arg_5_out := gocv.BorderType(arg_5_out_BorderType)

	gocv.GaussianBlur(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvSobel(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_int64, ok := arg_5_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(int64)"))
	}
	arg_5_out := int(arg_5_out_int64)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out_float64, ok := arg_6_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(float64)"))
	}
	arg_6_out := float64(arg_6_out_float64)
	arg_7_in, err := call.Argument(7).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_7_out_float64, ok := arg_7_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_7_in.(float64)"))
	}
	arg_7_out := float64(arg_7_out_float64)
	arg_8_in, err := call.Argument(8).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_8_out_BorderType_int64, ok := arg_8_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_8_in.(int64)"))
	}
	arg_8_out_BorderType := int(arg_8_out_BorderType_int64)
	arg_8_out := gocv.BorderType(arg_8_out_BorderType)

	gocv.Sobel(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out, arg_7_out, arg_8_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvSpatialGradient(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_BorderType_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out_BorderType := int(arg_4_out_BorderType_int64)
	arg_4_out := gocv.BorderType(arg_4_out_BorderType)

	gocv.SpatialGradient(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvLaplacian(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_float64, ok := arg_4_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(float64)"))
	}
	arg_4_out := float64(arg_4_out_float64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_float64, ok := arg_5_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(float64)"))
	}
	arg_5_out := float64(arg_5_out_float64)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out_BorderType_int64, ok := arg_6_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(int64)"))
	}
	arg_6_out_BorderType := int(arg_6_out_BorderType_int64)
	arg_6_out := gocv.BorderType(arg_6_out_BorderType)

	gocv.Laplacian(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvScharr(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_float64, ok := arg_5_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(float64)"))
	}
	arg_5_out := float64(arg_5_out_float64)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out_float64, ok := arg_6_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(float64)"))
	}
	arg_6_out := float64(arg_6_out_float64)
	arg_7_in, err := call.Argument(7).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_7_out_BorderType_int64, ok := arg_7_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_7_in.(int64)"))
	}
	arg_7_out_BorderType := int(arg_7_out_BorderType_int64)
	arg_7_out := gocv.BorderType(arg_7_out_BorderType)

	gocv.Scharr(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out, arg_7_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvMedianBlur(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	gocv.MedianBlur(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvCanny(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float32(arg_2_out_float64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float32(arg_3_out_float64)

	gocv.Canny(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvCornerSubPix(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out := image.Point{}
	arg_2_out_map, ok := arg_2_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(map[string]interface{})"))
	}
	arg_2_out_X_int64, ok := arg_2_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_X.(int64)"))
	}
	arg_2_out.X = int(arg_2_out_X_int64)
	arg_2_out_Y_int64, ok := arg_2_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_Y.(int64)"))
	}
	arg_2_out.Y = int(arg_2_out_Y_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out := image.Point{}
	arg_3_out_map, ok := arg_3_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(map[string]interface{})"))
	}
	arg_3_out_X_int64, ok := arg_3_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_map_X.(int64)"))
	}
	arg_3_out.X = int(arg_3_out_X_int64)
	arg_3_out_Y_int64, ok := arg_3_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_map_Y.(int64)"))
	}
	arg_3_out.Y = int(arg_3_out_Y_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_TermCriteria, ok := arg_4_in.(gocv.TermCriteria)
	if !ok {
		arg_4_out_TermCriteria_tmp, ok := arg_4_in.(*gocv.TermCriteria)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(gocv.TermCriteria)"))
		}
		arg_4_out_TermCriteria = *(arg_4_out_TermCriteria_tmp)
	}
	arg_4_out := gocv.TermCriteria(arg_4_out_TermCriteria)

	gocv.CornerSubPix(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvGoodFeaturesToTrack(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float64(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_float64, ok := arg_4_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(float64)"))
	}
	arg_4_out := float64(arg_4_out_float64)

	gocv.GoodFeaturesToTrack(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvHoughCircles(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_HoughMode_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_HoughMode := int(arg_2_out_HoughMode_int64)
	arg_2_out := gocv.HoughMode(arg_2_out_HoughMode)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float64(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_float64, ok := arg_4_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(float64)"))
	}
	arg_4_out := float64(arg_4_out_float64)

	gocv.HoughCircles(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvHoughCirclesWithParams(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_HoughMode_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_HoughMode := int(arg_2_out_HoughMode_int64)
	arg_2_out := gocv.HoughMode(arg_2_out_HoughMode)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float64(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_float64, ok := arg_4_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(float64)"))
	}
	arg_4_out := float64(arg_4_out_float64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_float64, ok := arg_5_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(float64)"))
	}
	arg_5_out := float64(arg_5_out_float64)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out_float64, ok := arg_6_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(float64)"))
	}
	arg_6_out := float64(arg_6_out_float64)
	arg_7_in, err := call.Argument(7).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_7_out_int64, ok := arg_7_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_7_in.(int64)"))
	}
	arg_7_out := int(arg_7_out_int64)
	arg_8_in, err := call.Argument(8).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_8_out_int64, ok := arg_8_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_8_in.(int64)"))
	}
	arg_8_out := int(arg_8_out_int64)

	gocv.HoughCirclesWithParams(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out, arg_7_out, arg_8_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvHoughLines(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float32(arg_2_out_float64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float32(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)

	gocv.HoughLines(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvHoughLinesP(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float32(arg_2_out_float64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float32(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)

	gocv.HoughLinesP(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvHoughLinesPWithParams(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float32(arg_2_out_float64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float32(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_float64, ok := arg_5_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(float64)"))
	}
	arg_5_out := float32(arg_5_out_float64)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out_float64, ok := arg_6_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(float64)"))
	}
	arg_6_out := float32(arg_6_out_float64)

	gocv.HoughLinesPWithParams(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvThreshold(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float32(arg_2_out_float64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float32(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_ThresholdType_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out_ThresholdType := int(arg_4_out_ThresholdType_int64)
	arg_4_out := gocv.ThresholdType(arg_4_out_ThresholdType)

	gocv.Threshold(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvAdaptiveThreshold(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float32(arg_2_out_float64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_AdaptiveThresholdType_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out_AdaptiveThresholdType := int(arg_3_out_AdaptiveThresholdType_int64)
	arg_3_out := gocv.AdaptiveThresholdType(arg_3_out_AdaptiveThresholdType)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_ThresholdType_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out_ThresholdType := int(arg_4_out_ThresholdType_int64)
	arg_4_out := gocv.ThresholdType(arg_4_out_ThresholdType)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_int64, ok := arg_5_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(int64)"))
	}
	arg_5_out := int(arg_5_out_int64)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out_float64, ok := arg_6_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(float64)"))
	}
	arg_6_out := float32(arg_6_out_float64)

	gocv.AdaptiveThreshold(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvArrowedLine(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out := image.Point{}
	arg_1_out_map, ok := arg_1_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(map[string]interface{})"))
	}
	arg_1_out_X_int64, ok := arg_1_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_map_X.(int64)"))
	}
	arg_1_out.X = int(arg_1_out_X_int64)
	arg_1_out_Y_int64, ok := arg_1_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_map_Y.(int64)"))
	}
	arg_1_out.Y = int(arg_1_out_Y_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out := image.Point{}
	arg_2_out_map, ok := arg_2_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(map[string]interface{})"))
	}
	arg_2_out_X_int64, ok := arg_2_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_X.(int64)"))
	}
	arg_2_out.X = int(arg_2_out_X_int64)
	arg_2_out_Y_int64, ok := arg_2_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_Y.(int64)"))
	}
	arg_2_out.Y = int(arg_2_out_Y_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out, ok := arg_3_in.(color.RGBA)
	if !ok {
		arg_3_out_tmp, ok := arg_3_in.(*color.RGBA)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(color.RGBA)"))
		}
		arg_3_out = *(arg_3_out_tmp)
	}
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)

	gocv.ArrowedLine(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvCircle(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out := image.Point{}
	arg_1_out_map, ok := arg_1_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(map[string]interface{})"))
	}
	arg_1_out_X_int64, ok := arg_1_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_map_X.(int64)"))
	}
	arg_1_out.X = int(arg_1_out_X_int64)
	arg_1_out_Y_int64, ok := arg_1_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_map_Y.(int64)"))
	}
	arg_1_out.Y = int(arg_1_out_Y_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out, ok := arg_3_in.(color.RGBA)
	if !ok {
		arg_3_out_tmp, ok := arg_3_in.(*color.RGBA)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(color.RGBA)"))
		}
		arg_3_out = *(arg_3_out_tmp)
	}
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)

	gocv.Circle(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvEllipse(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out := image.Point{}
	arg_1_out_map, ok := arg_1_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(map[string]interface{})"))
	}
	arg_1_out_X_int64, ok := arg_1_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_map_X.(int64)"))
	}
	arg_1_out.X = int(arg_1_out_X_int64)
	arg_1_out_Y_int64, ok := arg_1_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_map_Y.(int64)"))
	}
	arg_1_out.Y = int(arg_1_out_Y_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out := image.Point{}
	arg_2_out_map, ok := arg_2_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(map[string]interface{})"))
	}
	arg_2_out_X_int64, ok := arg_2_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_X.(int64)"))
	}
	arg_2_out.X = int(arg_2_out_X_int64)
	arg_2_out_Y_int64, ok := arg_2_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_Y.(int64)"))
	}
	arg_2_out.Y = int(arg_2_out_Y_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float64(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_float64, ok := arg_4_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(float64)"))
	}
	arg_4_out := float64(arg_4_out_float64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_float64, ok := arg_5_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(float64)"))
	}
	arg_5_out := float64(arg_5_out_float64)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out, ok := arg_6_in.(color.RGBA)
	if !ok {
		arg_6_out_tmp, ok := arg_6_in.(*color.RGBA)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(color.RGBA)"))
		}
		arg_6_out = *(arg_6_out_tmp)
	}
	arg_7_in, err := call.Argument(7).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_7_out_int64, ok := arg_7_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_7_in.(int64)"))
	}
	arg_7_out := int(arg_7_out_int64)

	gocv.Ellipse(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out, arg_7_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvLine(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out := image.Point{}
	arg_1_out_map, ok := arg_1_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(map[string]interface{})"))
	}
	arg_1_out_X_int64, ok := arg_1_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_map_X.(int64)"))
	}
	arg_1_out.X = int(arg_1_out_X_int64)
	arg_1_out_Y_int64, ok := arg_1_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_map_Y.(int64)"))
	}
	arg_1_out.Y = int(arg_1_out_Y_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out := image.Point{}
	arg_2_out_map, ok := arg_2_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(map[string]interface{})"))
	}
	arg_2_out_X_int64, ok := arg_2_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_X.(int64)"))
	}
	arg_2_out.X = int(arg_2_out_X_int64)
	arg_2_out_Y_int64, ok := arg_2_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_Y.(int64)"))
	}
	arg_2_out.Y = int(arg_2_out_Y_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out, ok := arg_3_in.(color.RGBA)
	if !ok {
		arg_3_out_tmp, ok := arg_3_in.(*color.RGBA)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(color.RGBA)"))
		}
		arg_3_out = *(arg_3_out_tmp)
	}
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)

	gocv.Line(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvRectangle(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out := image.Rectangle{}
	arg_1_out_map, ok := arg_1_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(map[string]interface{})"))
	}
	arg_1_out.Min = image.Point{}
	arg_1_out_Min_map, ok := arg_1_out_map["Min"].(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_map_Min.(map[string]interface{})"))
	}
	arg_1_out_Min_X_int64, ok := arg_1_out_Min_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Min_map_X.(int64)"))
	}
	arg_1_out.Min.X = int(arg_1_out_Min_X_int64)
	arg_1_out_Min_Y_int64, ok := arg_1_out_Min_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Min_map_Y.(int64)"))
	}
	arg_1_out.Min.Y = int(arg_1_out_Min_Y_int64)
	arg_1_out.Max = image.Point{}
	arg_1_out_Max_map, ok := arg_1_out_map["Max"].(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_map_Max.(map[string]interface{})"))
	}
	arg_1_out_Max_X_int64, ok := arg_1_out_Max_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Max_map_X.(int64)"))
	}
	arg_1_out.Max.X = int(arg_1_out_Max_X_int64)
	arg_1_out_Max_Y_int64, ok := arg_1_out_Max_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_Max_map_Y.(int64)"))
	}
	arg_1_out.Max.Y = int(arg_1_out_Max_Y_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(color.RGBA)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(*color.RGBA)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(color.RGBA)"))
		}
		arg_2_out = *(arg_2_out_tmp)
	}
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)

	gocv.Rectangle(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvFillPoly(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_in_arr, ok := arg_1_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.([]interface{})"))
	}
	arg_1_out := make([][]image.Point, len(arg_1_in_arr), len(arg_1_in_arr))
	for arg_1_in_arr_i, arg_1_in_arr_v := range arg_1_in_arr {
		arg_1_in_arr_v_arr, ok := arg_1_in_arr_v.([]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in_arr_v.([]interface{})"))
		}
		arg_1_out[arg_1_in_arr_i] = make([]image.Point, len(arg_1_in_arr_v_arr), len(arg_1_in_arr_v_arr))
		for arg_1_in_arr_v_arr_i, arg_1_in_arr_v_arr_v := range arg_1_in_arr_v_arr {
			arg_1_out[arg_1_in_arr_i][arg_1_in_arr_v_arr_i] = image.Point{}
			arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_map, ok := arg_1_in_arr_v_arr_v.(map[string]interface{})
			if !ok {
				return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in_arr_v_arr_v.(map[string]interface{})"))
			}
			arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_X_int64, ok := arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_map["X"].(int64)
			if !ok {
				return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_map_X.(int64)"))
			}
			arg_1_out[arg_1_in_arr_i][arg_1_in_arr_v_arr_i].X = int(arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_X_int64)
			arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_Y_int64, ok := arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_map["Y"].(int64)
			if !ok {
				return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_map_Y.(int64)"))
			}
			arg_1_out[arg_1_in_arr_i][arg_1_in_arr_v_arr_i].Y = int(arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_Y_int64)
		}
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(color.RGBA)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(*color.RGBA)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(color.RGBA)"))
		}
		arg_2_out = *(arg_2_out_tmp)
	}

	gocv.FillPoly(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvGetTextSize(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out := string(arg_0_out_string)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_HersheyFont_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out_HersheyFont := int(arg_1_out_HersheyFont_int64)
	arg_1_out := gocv.HersheyFont(arg_1_out_HersheyFont)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float64(arg_2_out_float64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)

	res_0 := gocv.GetTextSize(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	res_0_out_map := make(map[string]interface{})
	res_0_out_map["X"] = int64(res_0.X)
	res_0_out_map["Y"] = int64(res_0.Y)
	res_0_out := res_0_out_map
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvPutText(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_string, ok := arg_1_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(string)"))
	}
	arg_1_out := string(arg_1_out_string)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out := image.Point{}
	arg_2_out_map, ok := arg_2_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(map[string]interface{})"))
	}
	arg_2_out_X_int64, ok := arg_2_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_X.(int64)"))
	}
	arg_2_out.X = int(arg_2_out_X_int64)
	arg_2_out_Y_int64, ok := arg_2_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_Y.(int64)"))
	}
	arg_2_out.Y = int(arg_2_out_Y_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_HersheyFont_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out_HersheyFont := int(arg_3_out_HersheyFont_int64)
	arg_3_out := gocv.HersheyFont(arg_3_out_HersheyFont)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_float64, ok := arg_4_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(float64)"))
	}
	arg_4_out := float64(arg_4_out_float64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out, ok := arg_5_in.(color.RGBA)
	if !ok {
		arg_5_out_tmp, ok := arg_5_in.(*color.RGBA)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(color.RGBA)"))
		}
		arg_5_out = *(arg_5_out_tmp)
	}
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out_int64, ok := arg_6_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(int64)"))
	}
	arg_6_out := int(arg_6_out_int64)

	gocv.PutText(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvResize(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out := image.Point{}
	arg_2_out_map, ok := arg_2_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(map[string]interface{})"))
	}
	arg_2_out_X_int64, ok := arg_2_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_X.(int64)"))
	}
	arg_2_out.X = int(arg_2_out_X_int64)
	arg_2_out_Y_int64, ok := arg_2_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_Y.(int64)"))
	}
	arg_2_out.Y = int(arg_2_out_Y_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float64(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_float64, ok := arg_4_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(float64)"))
	}
	arg_4_out := float64(arg_4_out_float64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_InterpolationFlags_int64, ok := arg_5_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(int64)"))
	}
	arg_5_out_InterpolationFlags := int(arg_5_out_InterpolationFlags_int64)
	arg_5_out := gocv.InterpolationFlags(arg_5_out_InterpolationFlags)

	gocv.Resize(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvGetRotationMatrix2D(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out := image.Point{}
	arg_0_out_map, ok := arg_0_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(map[string]interface{})"))
	}
	arg_0_out_X_int64, ok := arg_0_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_map_X.(int64)"))
	}
	arg_0_out.X = int(arg_0_out_X_int64)
	arg_0_out_Y_int64, ok := arg_0_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_map_Y.(int64)"))
	}
	arg_0_out.Y = int(arg_0_out_Y_int64)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_float64, ok := arg_1_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(float64)"))
	}
	arg_1_out := float64(arg_1_out_float64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float64(arg_2_out_float64)

	res_0 := gocv.GetRotationMatrix2D(arg_0_out, arg_1_out, arg_2_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvWarpAffine(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out := image.Point{}
	arg_3_out_map, ok := arg_3_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(map[string]interface{})"))
	}
	arg_3_out_X_int64, ok := arg_3_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_map_X.(int64)"))
	}
	arg_3_out.X = int(arg_3_out_X_int64)
	arg_3_out_Y_int64, ok := arg_3_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_map_Y.(int64)"))
	}
	arg_3_out.Y = int(arg_3_out_Y_int64)

	gocv.WarpAffine(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvWarpAffineWithParams(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out := image.Point{}
	arg_3_out_map, ok := arg_3_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(map[string]interface{})"))
	}
	arg_3_out_X_int64, ok := arg_3_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_map_X.(int64)"))
	}
	arg_3_out.X = int(arg_3_out_X_int64)
	arg_3_out_Y_int64, ok := arg_3_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_map_Y.(int64)"))
	}
	arg_3_out.Y = int(arg_3_out_Y_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_InterpolationFlags_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out_InterpolationFlags := int(arg_4_out_InterpolationFlags_int64)
	arg_4_out := gocv.InterpolationFlags(arg_4_out_InterpolationFlags)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_BorderType_int64, ok := arg_5_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(int64)"))
	}
	arg_5_out_BorderType := int(arg_5_out_BorderType_int64)
	arg_5_out := gocv.BorderType(arg_5_out_BorderType)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out, ok := arg_6_in.(color.RGBA)
	if !ok {
		arg_6_out_tmp, ok := arg_6_in.(*color.RGBA)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(color.RGBA)"))
		}
		arg_6_out = *(arg_6_out_tmp)
	}

	gocv.WarpAffineWithParams(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvWarpPerspective(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out := image.Point{}
	arg_3_out_map, ok := arg_3_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(map[string]interface{})"))
	}
	arg_3_out_X_int64, ok := arg_3_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_map_X.(int64)"))
	}
	arg_3_out.X = int(arg_3_out_X_int64)
	arg_3_out_Y_int64, ok := arg_3_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_map_Y.(int64)"))
	}
	arg_3_out.Y = int(arg_3_out_Y_int64)

	gocv.WarpPerspective(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvApplyColorMap(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_ColormapTypes_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_ColormapTypes := int(arg_2_out_ColormapTypes_int64)
	arg_2_out := gocv.ColormapTypes(arg_2_out_ColormapTypes)

	gocv.ApplyColorMap(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvApplyCustomColorMap(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)

	gocv.ApplyCustomColorMap(arg_0_out, arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvGetPerspectiveTransform(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_in_arr, ok := arg_0_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.([]interface{})"))
	}
	arg_0_out := make([]image.Point, len(arg_0_in_arr), len(arg_0_in_arr))
	for arg_0_in_arr_i, arg_0_in_arr_v := range arg_0_in_arr {
		arg_0_out[arg_0_in_arr_i] = image.Point{}
		arg_0_out_arg_0_in_arr_i_map, ok := arg_0_in_arr_v.(map[string]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in_arr_v.(map[string]interface{})"))
		}
		arg_0_out_arg_0_in_arr_i_X_int64, ok := arg_0_out_arg_0_in_arr_i_map["X"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_X.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].X = int(arg_0_out_arg_0_in_arr_i_X_int64)
		arg_0_out_arg_0_in_arr_i_Y_int64, ok := arg_0_out_arg_0_in_arr_i_map["Y"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_Y.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].Y = int(arg_0_out_arg_0_in_arr_i_Y_int64)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_in_arr, ok := arg_1_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.([]interface{})"))
	}
	arg_1_out := make([]image.Point, len(arg_1_in_arr), len(arg_1_in_arr))
	for arg_1_in_arr_i, arg_1_in_arr_v := range arg_1_in_arr {
		arg_1_out[arg_1_in_arr_i] = image.Point{}
		arg_1_out_arg_1_in_arr_i_map, ok := arg_1_in_arr_v.(map[string]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in_arr_v.(map[string]interface{})"))
		}
		arg_1_out_arg_1_in_arr_i_X_int64, ok := arg_1_out_arg_1_in_arr_i_map["X"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_arg_1_in_arr_i_map_X.(int64)"))
		}
		arg_1_out[arg_1_in_arr_i].X = int(arg_1_out_arg_1_in_arr_i_X_int64)
		arg_1_out_arg_1_in_arr_i_Y_int64, ok := arg_1_out_arg_1_in_arr_i_map["Y"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_arg_1_in_arr_i_map_Y.(int64)"))
		}
		arg_1_out[arg_1_in_arr_i].Y = int(arg_1_out_arg_1_in_arr_i_Y_int64)
	}

	res_0 := gocv.GetPerspectiveTransform(arg_0_out, arg_1_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvDrawContours(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Mat)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_in_arr, ok := arg_1_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.([]interface{})"))
	}
	arg_1_out := make([][]image.Point, len(arg_1_in_arr), len(arg_1_in_arr))
	for arg_1_in_arr_i, arg_1_in_arr_v := range arg_1_in_arr {
		arg_1_in_arr_v_arr, ok := arg_1_in_arr_v.([]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in_arr_v.([]interface{})"))
		}
		arg_1_out[arg_1_in_arr_i] = make([]image.Point, len(arg_1_in_arr_v_arr), len(arg_1_in_arr_v_arr))
		for arg_1_in_arr_v_arr_i, arg_1_in_arr_v_arr_v := range arg_1_in_arr_v_arr {
			arg_1_out[arg_1_in_arr_i][arg_1_in_arr_v_arr_i] = image.Point{}
			arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_map, ok := arg_1_in_arr_v_arr_v.(map[string]interface{})
			if !ok {
				return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in_arr_v_arr_v.(map[string]interface{})"))
			}
			arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_X_int64, ok := arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_map["X"].(int64)
			if !ok {
				return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_map_X.(int64)"))
			}
			arg_1_out[arg_1_in_arr_i][arg_1_in_arr_v_arr_i].X = int(arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_X_int64)
			arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_Y_int64, ok := arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_map["Y"].(int64)
			if !ok {
				return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_map_Y.(int64)"))
			}
			arg_1_out[arg_1_in_arr_i][arg_1_in_arr_v_arr_i].Y = int(arg_1_out_arg_1_in_arr_i_arg_1_in_arr_v_arr_i_Y_int64)
		}
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out, ok := arg_3_in.(color.RGBA)
	if !ok {
		arg_3_out_tmp, ok := arg_3_in.(*color.RGBA)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(color.RGBA)"))
		}
		arg_3_out = *(arg_3_out_tmp)
	}
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)

	gocv.DrawContours(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvRemap(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out, ok := arg_3_in.(*gocv.Mat)
	if !ok {
		arg_3_out_tmp, ok := arg_3_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(gocv.Mat)"))
		}
		arg_3_out = &(arg_3_out_tmp)
	}
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_InterpolationFlags_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out_InterpolationFlags := int(arg_4_out_InterpolationFlags_int64)
	arg_4_out := gocv.InterpolationFlags(arg_4_out_InterpolationFlags)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_BorderType_int64, ok := arg_5_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(int64)"))
	}
	arg_5_out_BorderType := int(arg_5_out_BorderType_int64)
	arg_5_out := gocv.BorderType(arg_5_out_BorderType)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out, ok := arg_6_in.(color.RGBA)
	if !ok {
		arg_6_out_tmp, ok := arg_6_in.(*color.RGBA)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(color.RGBA)"))
		}
		arg_6_out = *(arg_6_out_tmp)
	}

	gocv.Remap(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvFilter2D(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_Mat, ok := arg_3_in.(gocv.Mat)
	if !ok {
		arg_3_out_Mat_tmp, ok := arg_3_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(gocv.Mat)"))
		}
		arg_3_out_Mat = *(arg_3_out_Mat_tmp)
	}
	arg_3_out := gocv.Mat(arg_3_out_Mat)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out := image.Point{}
	arg_4_out_map, ok := arg_4_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(map[string]interface{})"))
	}
	arg_4_out_X_int64, ok := arg_4_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_out_map_X.(int64)"))
	}
	arg_4_out.X = int(arg_4_out_X_int64)
	arg_4_out_Y_int64, ok := arg_4_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_out_map_Y.(int64)"))
	}
	arg_4_out.Y = int(arg_4_out_Y_int64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_float64, ok := arg_5_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(float64)"))
	}
	arg_5_out := float64(arg_5_out_float64)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out_BorderType_int64, ok := arg_6_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(int64)"))
	}
	arg_6_out_BorderType := int(arg_6_out_BorderType_int64)
	arg_6_out := gocv.BorderType(arg_6_out_BorderType)

	gocv.Filter2D(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvSepFilter2D(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_Mat, ok := arg_3_in.(gocv.Mat)
	if !ok {
		arg_3_out_Mat_tmp, ok := arg_3_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(gocv.Mat)"))
		}
		arg_3_out_Mat = *(arg_3_out_Mat_tmp)
	}
	arg_3_out := gocv.Mat(arg_3_out_Mat)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_Mat, ok := arg_4_in.(gocv.Mat)
	if !ok {
		arg_4_out_Mat_tmp, ok := arg_4_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(gocv.Mat)"))
		}
		arg_4_out_Mat = *(arg_4_out_Mat_tmp)
	}
	arg_4_out := gocv.Mat(arg_4_out_Mat)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out := image.Point{}
	arg_5_out_map, ok := arg_5_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(map[string]interface{})"))
	}
	arg_5_out_X_int64, ok := arg_5_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_out_map_X.(int64)"))
	}
	arg_5_out.X = int(arg_5_out_X_int64)
	arg_5_out_Y_int64, ok := arg_5_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_out_map_Y.(int64)"))
	}
	arg_5_out.Y = int(arg_5_out_Y_int64)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out_float64, ok := arg_6_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(float64)"))
	}
	arg_6_out := float64(arg_6_out_float64)
	arg_7_in, err := call.Argument(7).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_7_out_BorderType_int64, ok := arg_7_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_7_in.(int64)"))
	}
	arg_7_out_BorderType := int(arg_7_out_BorderType_int64)
	arg_7_out := gocv.BorderType(arg_7_out_BorderType)

	gocv.SepFilter2D(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out, arg_7_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvLogPolar(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out := image.Point{}
	arg_2_out_map, ok := arg_2_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(map[string]interface{})"))
	}
	arg_2_out_X_int64, ok := arg_2_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_X.(int64)"))
	}
	arg_2_out.X = int(arg_2_out_X_int64)
	arg_2_out_Y_int64, ok := arg_2_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_Y.(int64)"))
	}
	arg_2_out.Y = int(arg_2_out_Y_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float64(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_InterpolationFlags_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out_InterpolationFlags := int(arg_4_out_InterpolationFlags_int64)
	arg_4_out := gocv.InterpolationFlags(arg_4_out_InterpolationFlags)

	gocv.LogPolar(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvFitLine(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_in_arr, ok := arg_0_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.([]interface{})"))
	}
	arg_0_out := make([]image.Point, len(arg_0_in_arr), len(arg_0_in_arr))
	for arg_0_in_arr_i, arg_0_in_arr_v := range arg_0_in_arr {
		arg_0_out[arg_0_in_arr_i] = image.Point{}
		arg_0_out_arg_0_in_arr_i_map, ok := arg_0_in_arr_v.(map[string]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in_arr_v.(map[string]interface{})"))
		}
		arg_0_out_arg_0_in_arr_i_X_int64, ok := arg_0_out_arg_0_in_arr_i_map["X"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_X.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].X = int(arg_0_out_arg_0_in_arr_i_X_int64)
		arg_0_out_arg_0_in_arr_i_Y_int64, ok := arg_0_out_arg_0_in_arr_i_map["Y"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_Y.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].Y = int(arg_0_out_arg_0_in_arr_i_Y_int64)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_DistanceTypes_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out_DistanceTypes := int(arg_2_out_DistanceTypes_int64)
	arg_2_out := gocv.DistanceTypes(arg_2_out_DistanceTypes)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float64(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_float64, ok := arg_4_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(float64)"))
	}
	arg_4_out := float64(arg_4_out_float64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_float64, ok := arg_5_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(float64)"))
	}
	arg_5_out := float64(arg_5_out_float64)

	gocv.FitLine(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvFP16BlobFromImage(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_float64, ok := arg_1_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(float64)"))
	}
	arg_1_out := float32(arg_1_out_float64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out := image.Point{}
	arg_2_out_map, ok := arg_2_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(map[string]interface{})"))
	}
	arg_2_out_X_int64, ok := arg_2_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_X.(int64)"))
	}
	arg_2_out.X = int(arg_2_out_X_int64)
	arg_2_out_Y_int64, ok := arg_2_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_Y.(int64)"))
	}
	arg_2_out.Y = int(arg_2_out_Y_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float32(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_bool, ok := arg_4_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(bool)"))
	}
	arg_4_out := bool(arg_4_out_bool)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_bool, ok := arg_5_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(bool)"))
	}
	arg_5_out := bool(arg_5_out_bool)

	res_0 := gocv.FP16BlobFromImage(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out)

	res_0_out := bytesToBase64str(res_0)
	if err != nil {
		return o.handleError(err)
	}
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvVersion(call otto.FunctionCall) otto.Value {

	res_0 := gocv.Version()

	res_0_out := string(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvOpenCVVersion(call otto.FunctionCall) otto.Value {

	res_0 := gocv.OpenCVVersion()

	res_0_out := string(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvFisheyeUndistortImage(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_Mat, ok := arg_3_in.(gocv.Mat)
	if !ok {
		arg_3_out_Mat_tmp, ok := arg_3_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(gocv.Mat)"))
		}
		arg_3_out_Mat = *(arg_3_out_Mat_tmp)
	}
	arg_3_out := gocv.Mat(arg_3_out_Mat)

	gocv.FisheyeUndistortImage(arg_0_out, arg_1_out, arg_2_out, arg_3_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvFisheyeUndistortImageWithParams(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_Mat, ok := arg_3_in.(gocv.Mat)
	if !ok {
		arg_3_out_Mat_tmp, ok := arg_3_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(gocv.Mat)"))
		}
		arg_3_out_Mat = *(arg_3_out_Mat_tmp)
	}
	arg_3_out := gocv.Mat(arg_3_out_Mat)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_Mat, ok := arg_4_in.(gocv.Mat)
	if !ok {
		arg_4_out_Mat_tmp, ok := arg_4_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(gocv.Mat)"))
		}
		arg_4_out_Mat = *(arg_4_out_Mat_tmp)
	}
	arg_4_out := gocv.Mat(arg_4_out_Mat)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out := image.Point{}
	arg_5_out_map, ok := arg_5_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(map[string]interface{})"))
	}
	arg_5_out_X_int64, ok := arg_5_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_out_map_X.(int64)"))
	}
	arg_5_out.X = int(arg_5_out_X_int64)
	arg_5_out_Y_int64, ok := arg_5_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_out_map_Y.(int64)"))
	}
	arg_5_out.Y = int(arg_5_out_Y_int64)

	gocv.FisheyeUndistortImageWithParams(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvParseNetBackend(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out := string(arg_0_out_string)

	res_0 := gocv.ParseNetBackend(arg_0_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvParseNetTarget(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out := string(arg_0_out_string)

	res_0 := gocv.ParseNetTarget(arg_0_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNetClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Net)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Net)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Net)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvNetEmpty(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Net)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Net)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Net)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Empty()

	res_0_out := bool(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNetSetInput(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Net)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Net)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Net)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_string, ok := arg_2_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(string)"))
	}
	arg_2_out := string(arg_2_out_string)

	arg_0_out.SetInput(arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvNetForward(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Net)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Net)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Net)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_string, ok := arg_1_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(string)"))
	}
	arg_1_out := string(arg_1_out_string)

	res_0 := arg_0_out.Forward(arg_1_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNetForwardLayers(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Net)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Net)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Net)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_in_arr, ok := arg_1_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.([]interface{})"))
	}
	arg_1_out := make([]string, len(arg_1_in_arr), len(arg_1_in_arr))
	for arg_1_in_arr_i, arg_1_in_arr_v := range arg_1_in_arr {
		arg_1_out_arg_1_in_arr_i_string, ok := arg_1_in_arr_v.(string)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in_arr_v.(string)"))
		}
		arg_1_out[arg_1_in_arr_i] = string(arg_1_out_arg_1_in_arr_i_string)
	}

	res_0 := arg_0_out.ForwardLayers(arg_1_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr[res_0_i] = gocv.Mat(res_0_v)
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNetSetPreferableBackend(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Net)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Net)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Net)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_NetBackendType_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out_NetBackendType := int(arg_1_out_NetBackendType_int64)
	arg_1_out := gocv.NetBackendType(arg_1_out_NetBackendType)

	res_0 := arg_0_out.SetPreferableBackend(arg_1_out)

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvNetSetPreferableTarget(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Net)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Net)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Net)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_NetTargetType_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out_NetTargetType := int(arg_1_out_NetTargetType_int64)
	arg_1_out := gocv.NetTargetType(arg_1_out_NetTargetType)

	res_0 := arg_0_out.SetPreferableTarget(arg_1_out)

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvReadNet(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out := string(arg_0_out_string)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_string, ok := arg_1_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(string)"))
	}
	arg_1_out := string(arg_1_out_string)

	res_0 := gocv.ReadNet(arg_0_out, arg_1_out)

	res_0_out := gocv.Net(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvReadNetFromCaffe(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out := string(arg_0_out_string)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_string, ok := arg_1_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(string)"))
	}
	arg_1_out := string(arg_1_out_string)

	res_0 := gocv.ReadNetFromCaffe(arg_0_out, arg_1_out)

	res_0_out := gocv.Net(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvReadNetFromTensorflow(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out := string(arg_0_out_string)

	res_0 := gocv.ReadNetFromTensorflow(arg_0_out)

	res_0_out := gocv.Net(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvBlobFromImage(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_float64, ok := arg_1_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(float64)"))
	}
	arg_1_out := float64(arg_1_out_float64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out := image.Point{}
	arg_2_out_map, ok := arg_2_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(map[string]interface{})"))
	}
	arg_2_out_X_int64, ok := arg_2_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_X.(int64)"))
	}
	arg_2_out.X = int(arg_2_out_X_int64)
	arg_2_out_Y_int64, ok := arg_2_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_out_map_Y.(int64)"))
	}
	arg_2_out.Y = int(arg_2_out_Y_int64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_Scalar := gocv.Scalar{}
	arg_3_out_Scalar_map, ok := arg_3_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(map[string]interface{})"))
	}
	arg_3_out_Scalar_Val1_float64, ok := arg_3_out_Scalar_map["Val1"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_Scalar_map_Val1.(float64)"))
	}
	arg_3_out_Scalar.Val1 = float64(arg_3_out_Scalar_Val1_float64)
	arg_3_out_Scalar_Val2_float64, ok := arg_3_out_Scalar_map["Val2"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_Scalar_map_Val2.(float64)"))
	}
	arg_3_out_Scalar.Val2 = float64(arg_3_out_Scalar_Val2_float64)
	arg_3_out_Scalar_Val3_float64, ok := arg_3_out_Scalar_map["Val3"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_Scalar_map_Val3.(float64)"))
	}
	arg_3_out_Scalar.Val3 = float64(arg_3_out_Scalar_Val3_float64)
	arg_3_out_Scalar_Val4_float64, ok := arg_3_out_Scalar_map["Val4"].(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_Scalar_map_Val4.(float64)"))
	}
	arg_3_out_Scalar.Val4 = float64(arg_3_out_Scalar_Val4_float64)
	arg_3_out := gocv.Scalar(arg_3_out_Scalar)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_bool, ok := arg_4_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(bool)"))
	}
	arg_4_out := bool(arg_4_out_bool)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_bool, ok := arg_5_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(bool)"))
	}
	arg_5_out := bool(arg_5_out_bool)

	res_0 := gocv.BlobFromImage(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvGetBlobChannel(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_int64, ok := arg_2_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(int64)"))
	}
	arg_2_out := int(arg_2_out_int64)

	res_0 := gocv.GetBlobChannel(arg_0_out, arg_1_out, arg_2_out)

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvGetBlobSize(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)

	res_0 := gocv.GetBlobSize(arg_0_out)

	res_0_out_map := make(map[string]interface{})
	res_0_out_map["Val1"] = float64(res_0.Val1)
	res_0_out_map["Val2"] = float64(res_0.Val2)
	res_0_out_map["Val3"] = float64(res_0.Val3)
	res_0_out_map["Val4"] = float64(res_0.Val4)
	res_0_out := res_0_out_map
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNetGetLayer(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Net)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Net)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Net)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)

	res_0 := arg_0_out.GetLayer(arg_1_out)

	res_0_out := gocv.Layer(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNetGetPerfProfile(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Net)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Net)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Net)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.GetPerfProfile()

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNetGetUnconnectedOutLayers(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Net)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Net)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Net)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.GetUnconnectedOutLayers()

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr[res_0_i] = int64(res_0_v)
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvLayerClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Layer)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Layer)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Layer)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvLayerGetName(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Layer)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Layer)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Layer)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.GetName()

	res_0_out := string(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvLayerGetType(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Layer)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Layer)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Layer)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.GetType()

	res_0_out := string(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvLayerInputNameToIndex(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Layer)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Layer)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Layer)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_string, ok := arg_1_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(string)"))
	}
	arg_1_out := string(arg_1_out_string)

	res_0 := arg_0_out.InputNameToIndex(arg_1_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvLayerOutputNameToIndex(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.Layer)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.Layer)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Layer)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_string, ok := arg_1_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(string)"))
	}
	arg_1_out := string(arg_1_out_string)

	res_0 := arg_0_out.OutputNameToIndex(arg_1_out)

	res_0_out := int64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvVideoCaptureFile(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out := string(arg_0_out_string)

	res_0, res_1 := gocv.VideoCaptureFile(arg_0_out)

	res_0_out := (*gocv.VideoCapture)(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvVideoCaptureDevice(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_int64, ok := arg_0_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(int64)"))
	}
	arg_0_out := int(arg_0_out_int64)

	res_0, res_1 := gocv.VideoCaptureDevice(arg_0_out)

	res_0_out := (*gocv.VideoCapture)(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvVideoCaptureClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.VideoCapture)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.VideoCapture)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.VideoCapture)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvVideoCaptureSet(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.VideoCapture)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.VideoCapture)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.VideoCapture)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_VideoCaptureProperties_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out_VideoCaptureProperties := int(arg_1_out_VideoCaptureProperties_int64)
	arg_1_out := gocv.VideoCaptureProperties(arg_1_out_VideoCaptureProperties)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float64(arg_2_out_float64)

	arg_0_out.Set(arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvVideoCaptureGet(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_VideoCapture, ok := arg_0_in.(gocv.VideoCapture)
	if !ok {
		arg_0_out_VideoCapture_tmp, ok := arg_0_in.(*gocv.VideoCapture)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.VideoCapture)"))
		}
		arg_0_out_VideoCapture = *(arg_0_out_VideoCapture_tmp)
	}
	arg_0_out := gocv.VideoCapture(arg_0_out_VideoCapture)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_VideoCaptureProperties_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out_VideoCaptureProperties := int(arg_1_out_VideoCaptureProperties_int64)
	arg_1_out := gocv.VideoCaptureProperties(arg_1_out_VideoCaptureProperties)

	res_0 := arg_0_out.Get(arg_1_out)

	res_0_out := float64(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvVideoCaptureIsOpened(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.VideoCapture)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.VideoCapture)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.VideoCapture)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.IsOpened()

	res_0_out := bool(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvVideoCaptureRead(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.VideoCapture)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.VideoCapture)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.VideoCapture)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out, ok := arg_1_in.(*gocv.Mat)
	if !ok {
		arg_1_out_tmp, ok := arg_1_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out = &(arg_1_out_tmp)
	}

	res_0 := arg_0_out.Read(arg_1_out)

	res_0_out := bool(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvVideoCaptureGrab(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.VideoCapture)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.VideoCapture)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.VideoCapture)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)

	arg_0_out.Grab(arg_1_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvVideoCaptureCodecString(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.VideoCapture)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.VideoCapture)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.VideoCapture)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.CodecString()

	res_0_out := string(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvVideoWriterFile(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_string, ok := arg_0_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(string)"))
	}
	arg_0_out := string(arg_0_out_string)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_string, ok := arg_1_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(string)"))
	}
	arg_1_out := string(arg_1_out_string)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float64(arg_2_out_float64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_bool, ok := arg_5_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(bool)"))
	}
	arg_5_out := bool(arg_5_out_bool)

	res_0, res_1 := gocv.VideoWriterFile(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out)

	res_0_out := (*gocv.VideoWriter)(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvVideoWriterClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.VideoWriter)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.VideoWriter)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.VideoWriter)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvVideoWriterIsOpened(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.VideoWriter)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.VideoWriter)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.VideoWriter)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.IsOpened()

	res_0_out := bool(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvVideoWriterWrite(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.VideoWriter)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.VideoWriter)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.VideoWriter)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := arg_0_out.Write(arg_1_out)

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvOpenVideoCapture(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}

	var arg_0_out interface{}
	arg_0_in_int64, ok := arg_0_in.(int64)
	if ok {
		arg_0_out = int(arg_0_in_int64)
	}
	arg_0_in_string, ok := arg_0_in.(string)
	if ok {
		arg_0_out = string(arg_0_in_string)
	}

	res_0, res_1 := gocv.OpenVideoCapture(arg_0_out)

	res_0_out := (*gocv.VideoCapture)(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	if res_1 != nil {
		return o.handleError(res_1)
	}
	return res_0v
}

func (o *OttoFunctions) cvNewAKAZE(call otto.FunctionCall) otto.Value {

	res_0 := gocv.NewAKAZE()

	res_0_out := gocv.AKAZE(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvAKAZEClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.AKAZE)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.AKAZE)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.AKAZE)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvAKAZEDetect(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.AKAZE)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.AKAZE)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.AKAZE)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := arg_0_out.Detect(arg_1_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map["X"] = float64(res_0_v.X)
		res_0_out_arr_res_0_i_map["Y"] = float64(res_0_v.Y)
		res_0_out_arr_res_0_i_map["Size"] = float64(res_0_v.Size)
		res_0_out_arr_res_0_i_map["Angle"] = float64(res_0_v.Angle)
		res_0_out_arr_res_0_i_map["Response"] = float64(res_0_v.Response)
		res_0_out_arr_res_0_i_map["Octave"] = int64(res_0_v.Octave)
		res_0_out_arr_res_0_i_map["ClassID"] = int64(res_0_v.ClassID)
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvAKAZEDetectAndCompute(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.AKAZE)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.AKAZE)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.AKAZE)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)

	res_0, res_1 := arg_0_out.DetectAndCompute(arg_1_out, arg_2_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map["X"] = float64(res_0_v.X)
		res_0_out_arr_res_0_i_map["Y"] = float64(res_0_v.Y)
		res_0_out_arr_res_0_i_map["Size"] = float64(res_0_v.Size)
		res_0_out_arr_res_0_i_map["Angle"] = float64(res_0_v.Angle)
		res_0_out_arr_res_0_i_map["Response"] = float64(res_0_v.Response)
		res_0_out_arr_res_0_i_map["Octave"] = int64(res_0_v.Octave)
		res_0_out_arr_res_0_i_map["ClassID"] = int64(res_0_v.ClassID)
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	res_1_out := gocv.Mat(res_1)
	res_1v, err := o.vm.ToValue(res_1_out)
	if err != nil {
		return o.handleError(err)
	}
	results := map[string]interface{}{}
	results["r0"] = res_0v
	results["r1"] = res_1v
	resultsv, err := o.vm.ToValue(results)
	if err != nil {
		return o.handleError(err)
	}
	return resultsv
}

func (o *OttoFunctions) cvNewAgastFeatureDetector(call otto.FunctionCall) otto.Value {

	res_0 := gocv.NewAgastFeatureDetector()

	res_0_out := gocv.AgastFeatureDetector(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvAgastFeatureDetectorClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.AgastFeatureDetector)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.AgastFeatureDetector)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.AgastFeatureDetector)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvAgastFeatureDetectorDetect(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.AgastFeatureDetector)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.AgastFeatureDetector)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.AgastFeatureDetector)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := arg_0_out.Detect(arg_1_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map["X"] = float64(res_0_v.X)
		res_0_out_arr_res_0_i_map["Y"] = float64(res_0_v.Y)
		res_0_out_arr_res_0_i_map["Size"] = float64(res_0_v.Size)
		res_0_out_arr_res_0_i_map["Angle"] = float64(res_0_v.Angle)
		res_0_out_arr_res_0_i_map["Response"] = float64(res_0_v.Response)
		res_0_out_arr_res_0_i_map["Octave"] = int64(res_0_v.Octave)
		res_0_out_arr_res_0_i_map["ClassID"] = int64(res_0_v.ClassID)
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNewBRISK(call otto.FunctionCall) otto.Value {

	res_0 := gocv.NewBRISK()

	res_0_out := gocv.BRISK(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvBRISKClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.BRISK)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.BRISK)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.BRISK)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvBRISKDetect(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.BRISK)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.BRISK)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.BRISK)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := arg_0_out.Detect(arg_1_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map["X"] = float64(res_0_v.X)
		res_0_out_arr_res_0_i_map["Y"] = float64(res_0_v.Y)
		res_0_out_arr_res_0_i_map["Size"] = float64(res_0_v.Size)
		res_0_out_arr_res_0_i_map["Angle"] = float64(res_0_v.Angle)
		res_0_out_arr_res_0_i_map["Response"] = float64(res_0_v.Response)
		res_0_out_arr_res_0_i_map["Octave"] = int64(res_0_v.Octave)
		res_0_out_arr_res_0_i_map["ClassID"] = int64(res_0_v.ClassID)
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvBRISKDetectAndCompute(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.BRISK)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.BRISK)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.BRISK)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)

	res_0, res_1 := arg_0_out.DetectAndCompute(arg_1_out, arg_2_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map["X"] = float64(res_0_v.X)
		res_0_out_arr_res_0_i_map["Y"] = float64(res_0_v.Y)
		res_0_out_arr_res_0_i_map["Size"] = float64(res_0_v.Size)
		res_0_out_arr_res_0_i_map["Angle"] = float64(res_0_v.Angle)
		res_0_out_arr_res_0_i_map["Response"] = float64(res_0_v.Response)
		res_0_out_arr_res_0_i_map["Octave"] = int64(res_0_v.Octave)
		res_0_out_arr_res_0_i_map["ClassID"] = int64(res_0_v.ClassID)
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	res_1_out := gocv.Mat(res_1)
	res_1v, err := o.vm.ToValue(res_1_out)
	if err != nil {
		return o.handleError(err)
	}
	results := map[string]interface{}{}
	results["r0"] = res_0v
	results["r1"] = res_1v
	resultsv, err := o.vm.ToValue(results)
	if err != nil {
		return o.handleError(err)
	}
	return resultsv
}

func (o *OttoFunctions) cvNewFastFeatureDetector(call otto.FunctionCall) otto.Value {

	res_0 := gocv.NewFastFeatureDetector()

	res_0_out := gocv.FastFeatureDetector(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvFastFeatureDetectorClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.FastFeatureDetector)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.FastFeatureDetector)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.FastFeatureDetector)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvFastFeatureDetectorDetect(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.FastFeatureDetector)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.FastFeatureDetector)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.FastFeatureDetector)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := arg_0_out.Detect(arg_1_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map["X"] = float64(res_0_v.X)
		res_0_out_arr_res_0_i_map["Y"] = float64(res_0_v.Y)
		res_0_out_arr_res_0_i_map["Size"] = float64(res_0_v.Size)
		res_0_out_arr_res_0_i_map["Angle"] = float64(res_0_v.Angle)
		res_0_out_arr_res_0_i_map["Response"] = float64(res_0_v.Response)
		res_0_out_arr_res_0_i_map["Octave"] = int64(res_0_v.Octave)
		res_0_out_arr_res_0_i_map["ClassID"] = int64(res_0_v.ClassID)
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNewGFTTDetector(call otto.FunctionCall) otto.Value {

	res_0 := gocv.NewGFTTDetector()

	res_0_out := gocv.GFTTDetector(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvGFTTDetectorClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.GFTTDetector)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.GFTTDetector)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.GFTTDetector)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvGFTTDetectorDetect(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.GFTTDetector)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.GFTTDetector)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.GFTTDetector)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := arg_0_out.Detect(arg_1_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map["X"] = float64(res_0_v.X)
		res_0_out_arr_res_0_i_map["Y"] = float64(res_0_v.Y)
		res_0_out_arr_res_0_i_map["Size"] = float64(res_0_v.Size)
		res_0_out_arr_res_0_i_map["Angle"] = float64(res_0_v.Angle)
		res_0_out_arr_res_0_i_map["Response"] = float64(res_0_v.Response)
		res_0_out_arr_res_0_i_map["Octave"] = int64(res_0_v.Octave)
		res_0_out_arr_res_0_i_map["ClassID"] = int64(res_0_v.ClassID)
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNewKAZE(call otto.FunctionCall) otto.Value {

	res_0 := gocv.NewKAZE()

	res_0_out := gocv.KAZE(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvKAZEClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.KAZE)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.KAZE)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.KAZE)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvKAZEDetect(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.KAZE)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.KAZE)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.KAZE)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := arg_0_out.Detect(arg_1_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map["X"] = float64(res_0_v.X)
		res_0_out_arr_res_0_i_map["Y"] = float64(res_0_v.Y)
		res_0_out_arr_res_0_i_map["Size"] = float64(res_0_v.Size)
		res_0_out_arr_res_0_i_map["Angle"] = float64(res_0_v.Angle)
		res_0_out_arr_res_0_i_map["Response"] = float64(res_0_v.Response)
		res_0_out_arr_res_0_i_map["Octave"] = int64(res_0_v.Octave)
		res_0_out_arr_res_0_i_map["ClassID"] = int64(res_0_v.ClassID)
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvKAZEDetectAndCompute(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.KAZE)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.KAZE)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.KAZE)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)

	res_0, res_1 := arg_0_out.DetectAndCompute(arg_1_out, arg_2_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map["X"] = float64(res_0_v.X)
		res_0_out_arr_res_0_i_map["Y"] = float64(res_0_v.Y)
		res_0_out_arr_res_0_i_map["Size"] = float64(res_0_v.Size)
		res_0_out_arr_res_0_i_map["Angle"] = float64(res_0_v.Angle)
		res_0_out_arr_res_0_i_map["Response"] = float64(res_0_v.Response)
		res_0_out_arr_res_0_i_map["Octave"] = int64(res_0_v.Octave)
		res_0_out_arr_res_0_i_map["ClassID"] = int64(res_0_v.ClassID)
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	res_1_out := gocv.Mat(res_1)
	res_1v, err := o.vm.ToValue(res_1_out)
	if err != nil {
		return o.handleError(err)
	}
	results := map[string]interface{}{}
	results["r0"] = res_0v
	results["r1"] = res_1v
	resultsv, err := o.vm.ToValue(results)
	if err != nil {
		return o.handleError(err)
	}
	return resultsv
}

func (o *OttoFunctions) cvNewMSER(call otto.FunctionCall) otto.Value {

	res_0 := gocv.NewMSER()

	res_0_out := gocv.MSER(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvMSERClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.MSER)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.MSER)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.MSER)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvMSERDetect(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.MSER)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.MSER)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.MSER)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := arg_0_out.Detect(arg_1_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map["X"] = float64(res_0_v.X)
		res_0_out_arr_res_0_i_map["Y"] = float64(res_0_v.Y)
		res_0_out_arr_res_0_i_map["Size"] = float64(res_0_v.Size)
		res_0_out_arr_res_0_i_map["Angle"] = float64(res_0_v.Angle)
		res_0_out_arr_res_0_i_map["Response"] = float64(res_0_v.Response)
		res_0_out_arr_res_0_i_map["Octave"] = int64(res_0_v.Octave)
		res_0_out_arr_res_0_i_map["ClassID"] = int64(res_0_v.ClassID)
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNewORB(call otto.FunctionCall) otto.Value {

	res_0 := gocv.NewORB()

	res_0_out := gocv.ORB(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvORBClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.ORB)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.ORB)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.ORB)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvORBDetect(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.ORB)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.ORB)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.ORB)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := arg_0_out.Detect(arg_1_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map["X"] = float64(res_0_v.X)
		res_0_out_arr_res_0_i_map["Y"] = float64(res_0_v.Y)
		res_0_out_arr_res_0_i_map["Size"] = float64(res_0_v.Size)
		res_0_out_arr_res_0_i_map["Angle"] = float64(res_0_v.Angle)
		res_0_out_arr_res_0_i_map["Response"] = float64(res_0_v.Response)
		res_0_out_arr_res_0_i_map["Octave"] = int64(res_0_v.Octave)
		res_0_out_arr_res_0_i_map["ClassID"] = int64(res_0_v.ClassID)
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvORBDetectAndCompute(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.ORB)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.ORB)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.ORB)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)

	res_0, res_1 := arg_0_out.DetectAndCompute(arg_1_out, arg_2_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map["X"] = float64(res_0_v.X)
		res_0_out_arr_res_0_i_map["Y"] = float64(res_0_v.Y)
		res_0_out_arr_res_0_i_map["Size"] = float64(res_0_v.Size)
		res_0_out_arr_res_0_i_map["Angle"] = float64(res_0_v.Angle)
		res_0_out_arr_res_0_i_map["Response"] = float64(res_0_v.Response)
		res_0_out_arr_res_0_i_map["Octave"] = int64(res_0_v.Octave)
		res_0_out_arr_res_0_i_map["ClassID"] = int64(res_0_v.ClassID)
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	res_1_out := gocv.Mat(res_1)
	res_1v, err := o.vm.ToValue(res_1_out)
	if err != nil {
		return o.handleError(err)
	}
	results := map[string]interface{}{}
	results["r0"] = res_0v
	results["r1"] = res_1v
	resultsv, err := o.vm.ToValue(results)
	if err != nil {
		return o.handleError(err)
	}
	return resultsv
}

func (o *OttoFunctions) cvNewSimpleBlobDetector(call otto.FunctionCall) otto.Value {

	res_0 := gocv.NewSimpleBlobDetector()

	res_0_out := gocv.SimpleBlobDetector(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvSimpleBlobDetectorClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.SimpleBlobDetector)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.SimpleBlobDetector)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.SimpleBlobDetector)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvSimpleBlobDetectorDetect(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.SimpleBlobDetector)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.SimpleBlobDetector)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.SimpleBlobDetector)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := arg_0_out.Detect(arg_1_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map["X"] = float64(res_0_v.X)
		res_0_out_arr_res_0_i_map["Y"] = float64(res_0_v.Y)
		res_0_out_arr_res_0_i_map["Size"] = float64(res_0_v.Size)
		res_0_out_arr_res_0_i_map["Angle"] = float64(res_0_v.Angle)
		res_0_out_arr_res_0_i_map["Response"] = float64(res_0_v.Response)
		res_0_out_arr_res_0_i_map["Octave"] = int64(res_0_v.Octave)
		res_0_out_arr_res_0_i_map["ClassID"] = int64(res_0_v.ClassID)
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNewBFMatcher(call otto.FunctionCall) otto.Value {

	res_0 := gocv.NewBFMatcher()

	res_0_out := gocv.BFMatcher(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNewBFMatcherWithParams(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_NormType_int64, ok := arg_0_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(int64)"))
	}
	arg_0_out_NormType := int(arg_0_out_NormType_int64)
	arg_0_out := gocv.NormType(arg_0_out_NormType)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_bool, ok := arg_1_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(bool)"))
	}
	arg_1_out := bool(arg_1_out_bool)

	res_0 := gocv.NewBFMatcherWithParams(arg_0_out, arg_1_out)

	res_0_out := gocv.BFMatcher(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvBFMatcherClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.BFMatcher)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.BFMatcher)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.BFMatcher)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvBFMatcherKnnMatch(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.BFMatcher)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.BFMatcher)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.BFMatcher)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)

	res_0 := arg_0_out.KnnMatch(arg_1_out, arg_2_out, arg_3_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_arr := make([]interface{}, len(res_0_v), len(res_0_v))
		for res_0_v_i, res_0_v_v := range res_0_v {
			res_0_out_arr_res_0_i_arr_res_0_v_i_map := make(map[string]interface{})
			res_0_out_arr_res_0_i_arr_res_0_v_i_map["QueryIdx"] = int64(res_0_v_v.QueryIdx)
			res_0_out_arr_res_0_i_arr_res_0_v_i_map["TrainIdx"] = int64(res_0_v_v.TrainIdx)
			res_0_out_arr_res_0_i_arr_res_0_v_i_map["ImgIdx"] = int64(res_0_v_v.ImgIdx)
			res_0_out_arr_res_0_i_arr_res_0_v_i_map["Distance"] = float64(res_0_v_v.Distance)
			res_0_out_arr_res_0_i_arr[res_0_v_i] = res_0_out_arr_res_0_i_arr_res_0_v_i_map
		}
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_arr
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvDrawKeyPoints(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_in_arr, ok := arg_1_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.([]interface{})"))
	}
	arg_1_out := make([]gocv.KeyPoint, len(arg_1_in_arr), len(arg_1_in_arr))
	for arg_1_in_arr_i, arg_1_in_arr_v := range arg_1_in_arr {
		arg_1_out_arg_1_in_arr_i_KeyPoint := gocv.KeyPoint{}
		arg_1_out_arg_1_in_arr_i_KeyPoint_map, ok := arg_1_in_arr_v.(map[string]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in_arr_v.(map[string]interface{})"))
		}
		arg_1_out_arg_1_in_arr_i_KeyPoint_X_float64, ok := arg_1_out_arg_1_in_arr_i_KeyPoint_map["X"].(float64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_arg_1_in_arr_i_KeyPoint_map_X.(float64)"))
		}
		arg_1_out_arg_1_in_arr_i_KeyPoint.X = float64(arg_1_out_arg_1_in_arr_i_KeyPoint_X_float64)
		arg_1_out_arg_1_in_arr_i_KeyPoint_Y_float64, ok := arg_1_out_arg_1_in_arr_i_KeyPoint_map["Y"].(float64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_arg_1_in_arr_i_KeyPoint_map_Y.(float64)"))
		}
		arg_1_out_arg_1_in_arr_i_KeyPoint.Y = float64(arg_1_out_arg_1_in_arr_i_KeyPoint_Y_float64)
		arg_1_out_arg_1_in_arr_i_KeyPoint_Size_float64, ok := arg_1_out_arg_1_in_arr_i_KeyPoint_map["Size"].(float64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_arg_1_in_arr_i_KeyPoint_map_Size.(float64)"))
		}
		arg_1_out_arg_1_in_arr_i_KeyPoint.Size = float64(arg_1_out_arg_1_in_arr_i_KeyPoint_Size_float64)
		arg_1_out_arg_1_in_arr_i_KeyPoint_Angle_float64, ok := arg_1_out_arg_1_in_arr_i_KeyPoint_map["Angle"].(float64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_arg_1_in_arr_i_KeyPoint_map_Angle.(float64)"))
		}
		arg_1_out_arg_1_in_arr_i_KeyPoint.Angle = float64(arg_1_out_arg_1_in_arr_i_KeyPoint_Angle_float64)
		arg_1_out_arg_1_in_arr_i_KeyPoint_Response_float64, ok := arg_1_out_arg_1_in_arr_i_KeyPoint_map["Response"].(float64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_arg_1_in_arr_i_KeyPoint_map_Response.(float64)"))
		}
		arg_1_out_arg_1_in_arr_i_KeyPoint.Response = float64(arg_1_out_arg_1_in_arr_i_KeyPoint_Response_float64)
		arg_1_out_arg_1_in_arr_i_KeyPoint_Octave_int64, ok := arg_1_out_arg_1_in_arr_i_KeyPoint_map["Octave"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_arg_1_in_arr_i_KeyPoint_map_Octave.(int64)"))
		}
		arg_1_out_arg_1_in_arr_i_KeyPoint.Octave = int(arg_1_out_arg_1_in_arr_i_KeyPoint_Octave_int64)
		arg_1_out_arg_1_in_arr_i_KeyPoint_ClassID_int64, ok := arg_1_out_arg_1_in_arr_i_KeyPoint_map["ClassID"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_out_arg_1_in_arr_i_KeyPoint_map_ClassID.(int64)"))
		}
		arg_1_out_arg_1_in_arr_i_KeyPoint.ClassID = int(arg_1_out_arg_1_in_arr_i_KeyPoint_ClassID_int64)
		arg_1_out[arg_1_in_arr_i] = gocv.KeyPoint(arg_1_out_arg_1_in_arr_i_KeyPoint)
	}
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out, ok := arg_3_in.(color.RGBA)
	if !ok {
		arg_3_out_tmp, ok := arg_3_in.(*color.RGBA)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(color.RGBA)"))
		}
		arg_3_out = *(arg_3_out_tmp)
	}
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_DrawMatchesFlag_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out_DrawMatchesFlag := int(arg_4_out_DrawMatchesFlag_int64)
	arg_4_out := gocv.DrawMatchesFlag(arg_4_out_DrawMatchesFlag)

	gocv.DrawKeyPoints(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvNewBackgroundSubtractorMOG2(call otto.FunctionCall) otto.Value {

	res_0 := gocv.NewBackgroundSubtractorMOG2()

	res_0_out := gocv.BackgroundSubtractorMOG2(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvBackgroundSubtractorMOG2Close(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.BackgroundSubtractorMOG2)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.BackgroundSubtractorMOG2)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.BackgroundSubtractorMOG2)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvBackgroundSubtractorMOG2Apply(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.BackgroundSubtractorMOG2)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.BackgroundSubtractorMOG2)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.BackgroundSubtractorMOG2)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	arg_0_out.Apply(arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvNewBackgroundSubtractorKNN(call otto.FunctionCall) otto.Value {

	res_0 := gocv.NewBackgroundSubtractorKNN()

	res_0_out := gocv.BackgroundSubtractorKNN(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvBackgroundSubtractorKNNClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.BackgroundSubtractorKNN)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.BackgroundSubtractorKNN)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.BackgroundSubtractorKNN)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvBackgroundSubtractorKNNApply(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.BackgroundSubtractorKNN)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.BackgroundSubtractorKNN)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.BackgroundSubtractorKNN)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}

	arg_0_out.Apply(arg_1_out, arg_2_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvCalcOpticalFlowFarneback(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out, ok := arg_2_in.(*gocv.Mat)
	if !ok {
		arg_2_out_tmp, ok := arg_2_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out = &(arg_2_out_tmp)
	}
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_float64, ok := arg_3_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(float64)"))
	}
	arg_3_out := float64(arg_3_out_float64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_int64, ok := arg_5_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(int64)"))
	}
	arg_5_out := int(arg_5_out_int64)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out_int64, ok := arg_6_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(int64)"))
	}
	arg_6_out := int(arg_6_out_int64)
	arg_7_in, err := call.Argument(7).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_7_out_int64, ok := arg_7_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_7_in.(int64)"))
	}
	arg_7_out := int(arg_7_out_int64)
	arg_8_in, err := call.Argument(8).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_8_out_float64, ok := arg_8_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_8_in.(float64)"))
	}
	arg_8_out := float64(arg_8_out_float64)
	arg_9_in, err := call.Argument(9).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_9_out_int64, ok := arg_9_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_9_in.(int64)"))
	}
	arg_9_out := int(arg_9_out_int64)

	gocv.CalcOpticalFlowFarneback(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out, arg_7_out, arg_8_out, arg_9_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvCalcOpticalFlowPyrLK(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out_Mat, ok := arg_0_in.(gocv.Mat)
	if !ok {
		arg_0_out_Mat_tmp, ok := arg_0_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.Mat)"))
		}
		arg_0_out_Mat = *(arg_0_out_Mat_tmp)
	}
	arg_0_out := gocv.Mat(arg_0_out_Mat)
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_Mat, ok := arg_2_in.(gocv.Mat)
	if !ok {
		arg_2_out_Mat_tmp, ok := arg_2_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(gocv.Mat)"))
		}
		arg_2_out_Mat = *(arg_2_out_Mat_tmp)
	}
	arg_2_out := gocv.Mat(arg_2_out_Mat)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_Mat, ok := arg_3_in.(gocv.Mat)
	if !ok {
		arg_3_out_Mat_tmp, ok := arg_3_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(gocv.Mat)"))
		}
		arg_3_out_Mat = *(arg_3_out_Mat_tmp)
	}
	arg_3_out := gocv.Mat(arg_3_out_Mat)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out, ok := arg_4_in.(*gocv.Mat)
	if !ok {
		arg_4_out_tmp, ok := arg_4_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(gocv.Mat)"))
		}
		arg_4_out = &(arg_4_out_tmp)
	}
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out, ok := arg_5_in.(*gocv.Mat)
	if !ok {
		arg_5_out_tmp, ok := arg_5_in.(gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(gocv.Mat)"))
		}
		arg_5_out = &(arg_5_out_tmp)
	}

	gocv.CalcOpticalFlowPyrLK(arg_0_out, arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out)

	return otto.Value{}
}

func (o *OttoFunctions) cvNewCascadeClassifier(call otto.FunctionCall) otto.Value {

	res_0 := gocv.NewCascadeClassifier()

	res_0_out := gocv.CascadeClassifier(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvCascadeClassifierClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.CascadeClassifier)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.CascadeClassifier)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.CascadeClassifier)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvCascadeClassifierLoad(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.CascadeClassifier)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.CascadeClassifier)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.CascadeClassifier)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_string, ok := arg_1_in.(string)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(string)"))
	}
	arg_1_out := string(arg_1_out_string)

	res_0 := arg_0_out.Load(arg_1_out)

	res_0_out := bool(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvCascadeClassifierDetectMultiScale(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.CascadeClassifier)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.CascadeClassifier)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.CascadeClassifier)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := arg_0_out.DetectMultiScale(arg_1_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Min_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Min_map["X"] = int64(res_0_v.Min.X)
		res_0_out_arr_res_0_i_map_Min_map["Y"] = int64(res_0_v.Min.Y)
		res_0_out_arr_res_0_i_map["Min"] = res_0_out_arr_res_0_i_map_Min_map
		res_0_out_arr_res_0_i_map_Max_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Max_map["X"] = int64(res_0_v.Max.X)
		res_0_out_arr_res_0_i_map_Max_map["Y"] = int64(res_0_v.Max.Y)
		res_0_out_arr_res_0_i_map["Max"] = res_0_out_arr_res_0_i_map_Max_map
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvCascadeClassifierDetectMultiScaleWithParams(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.CascadeClassifier)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.CascadeClassifier)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.CascadeClassifier)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float64(arg_2_out_float64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out_int64, ok := arg_3_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(int64)"))
	}
	arg_3_out := int(arg_3_out_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out_int64, ok := arg_4_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(int64)"))
	}
	arg_4_out := int(arg_4_out_int64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out := image.Point{}
	arg_5_out_map, ok := arg_5_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(map[string]interface{})"))
	}
	arg_5_out_X_int64, ok := arg_5_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_out_map_X.(int64)"))
	}
	arg_5_out.X = int(arg_5_out_X_int64)
	arg_5_out_Y_int64, ok := arg_5_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_out_map_Y.(int64)"))
	}
	arg_5_out.Y = int(arg_5_out_Y_int64)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out := image.Point{}
	arg_6_out_map, ok := arg_6_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(map[string]interface{})"))
	}
	arg_6_out_X_int64, ok := arg_6_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_out_map_X.(int64)"))
	}
	arg_6_out.X = int(arg_6_out_X_int64)
	arg_6_out_Y_int64, ok := arg_6_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_out_map_Y.(int64)"))
	}
	arg_6_out.Y = int(arg_6_out_Y_int64)

	res_0 := arg_0_out.DetectMultiScaleWithParams(arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Min_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Min_map["X"] = int64(res_0_v.Min.X)
		res_0_out_arr_res_0_i_map_Min_map["Y"] = int64(res_0_v.Min.Y)
		res_0_out_arr_res_0_i_map["Min"] = res_0_out_arr_res_0_i_map_Min_map
		res_0_out_arr_res_0_i_map_Max_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Max_map["X"] = int64(res_0_v.Max.X)
		res_0_out_arr_res_0_i_map_Max_map["Y"] = int64(res_0_v.Max.Y)
		res_0_out_arr_res_0_i_map["Max"] = res_0_out_arr_res_0_i_map_Max_map
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvNewHOGDescriptor(call otto.FunctionCall) otto.Value {

	res_0 := gocv.NewHOGDescriptor()

	res_0_out := gocv.HOGDescriptor(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvHOGDescriptorClose(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.HOGDescriptor)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.HOGDescriptor)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.HOGDescriptor)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}

	res_0 := arg_0_out.Close()

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvHOGDescriptorDetectMultiScale(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.HOGDescriptor)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.HOGDescriptor)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.HOGDescriptor)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := arg_0_out.DetectMultiScale(arg_1_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Min_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Min_map["X"] = int64(res_0_v.Min.X)
		res_0_out_arr_res_0_i_map_Min_map["Y"] = int64(res_0_v.Min.Y)
		res_0_out_arr_res_0_i_map["Min"] = res_0_out_arr_res_0_i_map_Min_map
		res_0_out_arr_res_0_i_map_Max_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Max_map["X"] = int64(res_0_v.Max.X)
		res_0_out_arr_res_0_i_map_Max_map["Y"] = int64(res_0_v.Max.Y)
		res_0_out_arr_res_0_i_map["Max"] = res_0_out_arr_res_0_i_map_Max_map
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvHOGDescriptorDetectMultiScaleWithParams(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.HOGDescriptor)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.HOGDescriptor)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.HOGDescriptor)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float64(arg_2_out_float64)
	arg_3_in, err := call.Argument(3).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_3_out := image.Point{}
	arg_3_out_map, ok := arg_3_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_in.(map[string]interface{})"))
	}
	arg_3_out_X_int64, ok := arg_3_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_map_X.(int64)"))
	}
	arg_3_out.X = int(arg_3_out_X_int64)
	arg_3_out_Y_int64, ok := arg_3_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_3_out_map_Y.(int64)"))
	}
	arg_3_out.Y = int(arg_3_out_Y_int64)
	arg_4_in, err := call.Argument(4).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_4_out := image.Point{}
	arg_4_out_map, ok := arg_4_in.(map[string]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_in.(map[string]interface{})"))
	}
	arg_4_out_X_int64, ok := arg_4_out_map["X"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_out_map_X.(int64)"))
	}
	arg_4_out.X = int(arg_4_out_X_int64)
	arg_4_out_Y_int64, ok := arg_4_out_map["Y"].(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_4_out_map_Y.(int64)"))
	}
	arg_4_out.Y = int(arg_4_out_Y_int64)
	arg_5_in, err := call.Argument(5).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_5_out_float64, ok := arg_5_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_5_in.(float64)"))
	}
	arg_5_out := float64(arg_5_out_float64)
	arg_6_in, err := call.Argument(6).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_6_out_float64, ok := arg_6_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_6_in.(float64)"))
	}
	arg_6_out := float64(arg_6_out_float64)
	arg_7_in, err := call.Argument(7).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_7_out_bool, ok := arg_7_in.(bool)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_7_in.(bool)"))
	}
	arg_7_out := bool(arg_7_out_bool)

	res_0 := arg_0_out.DetectMultiScaleWithParams(arg_1_out, arg_2_out, arg_3_out, arg_4_out, arg_5_out, arg_6_out, arg_7_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Min_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Min_map["X"] = int64(res_0_v.Min.X)
		res_0_out_arr_res_0_i_map_Min_map["Y"] = int64(res_0_v.Min.Y)
		res_0_out_arr_res_0_i_map["Min"] = res_0_out_arr_res_0_i_map_Min_map
		res_0_out_arr_res_0_i_map_Max_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Max_map["X"] = int64(res_0_v.Max.X)
		res_0_out_arr_res_0_i_map_Max_map["Y"] = int64(res_0_v.Max.Y)
		res_0_out_arr_res_0_i_map["Max"] = res_0_out_arr_res_0_i_map_Max_map
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvHOGDefaultPeopleDetector(call otto.FunctionCall) otto.Value {

	res_0 := gocv.HOGDefaultPeopleDetector()

	res_0_out := gocv.Mat(res_0)
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) cvHOGDescriptorSetSVMDetector(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_out, ok := arg_0_in.(*gocv.HOGDescriptor)
	if !ok {
		arg_0_out_tmp, ok := arg_0_in.(gocv.HOGDescriptor)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.(gocv.HOGDescriptor)"))
		}
		arg_0_out = &(arg_0_out_tmp)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_Mat, ok := arg_1_in.(gocv.Mat)
	if !ok {
		arg_1_out_Mat_tmp, ok := arg_1_in.(*gocv.Mat)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(gocv.Mat)"))
		}
		arg_1_out_Mat = *(arg_1_out_Mat_tmp)
	}
	arg_1_out := gocv.Mat(arg_1_out_Mat)

	res_0 := arg_0_out.SetSVMDetector(arg_1_out)

	if res_0 != nil {
		return o.handleError(res_0)
	}
	return otto.Value{}
}

func (o *OttoFunctions) cvGroupRectangles(call otto.FunctionCall) otto.Value {
	var ok bool
	arg_0_in, err := call.Argument(0).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_0_in_arr, ok := arg_0_in.([]interface{})
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in.([]interface{})"))
	}
	arg_0_out := make([]image.Rectangle, len(arg_0_in_arr), len(arg_0_in_arr))
	for arg_0_in_arr_i, arg_0_in_arr_v := range arg_0_in_arr {
		arg_0_out[arg_0_in_arr_i] = image.Rectangle{}
		arg_0_out_arg_0_in_arr_i_map, ok := arg_0_in_arr_v.(map[string]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_in_arr_v.(map[string]interface{})"))
		}
		arg_0_out[arg_0_in_arr_i].Min = image.Point{}
		arg_0_out_arg_0_in_arr_i_Min_map, ok := arg_0_out_arg_0_in_arr_i_map["Min"].(map[string]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_Min.(map[string]interface{})"))
		}
		arg_0_out_arg_0_in_arr_i_Min_X_int64, ok := arg_0_out_arg_0_in_arr_i_Min_map["X"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_Min_map_X.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].Min.X = int(arg_0_out_arg_0_in_arr_i_Min_X_int64)
		arg_0_out_arg_0_in_arr_i_Min_Y_int64, ok := arg_0_out_arg_0_in_arr_i_Min_map["Y"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_Min_map_Y.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].Min.Y = int(arg_0_out_arg_0_in_arr_i_Min_Y_int64)
		arg_0_out[arg_0_in_arr_i].Max = image.Point{}
		arg_0_out_arg_0_in_arr_i_Max_map, ok := arg_0_out_arg_0_in_arr_i_map["Max"].(map[string]interface{})
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_map_Max.(map[string]interface{})"))
		}
		arg_0_out_arg_0_in_arr_i_Max_X_int64, ok := arg_0_out_arg_0_in_arr_i_Max_map["X"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_Max_map_X.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].Max.X = int(arg_0_out_arg_0_in_arr_i_Max_X_int64)
		arg_0_out_arg_0_in_arr_i_Max_Y_int64, ok := arg_0_out_arg_0_in_arr_i_Max_map["Y"].(int64)
		if !ok {
			return o.handleError(fmt.Errorf("Can not Convert To Type: arg_0_out_arg_0_in_arr_i_Max_map_Y.(int64)"))
		}
		arg_0_out[arg_0_in_arr_i].Max.Y = int(arg_0_out_arg_0_in_arr_i_Max_Y_int64)
	}
	arg_1_in, err := call.Argument(1).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_1_out_int64, ok := arg_1_in.(int64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_1_in.(int64)"))
	}
	arg_1_out := int(arg_1_out_int64)
	arg_2_in, err := call.Argument(2).Export()
	if err != nil {
		return o.handleError(err)
	}
	arg_2_out_float64, ok := arg_2_in.(float64)
	if !ok {
		return o.handleError(fmt.Errorf("Can not Convert To Type: arg_2_in.(float64)"))
	}
	arg_2_out := float64(arg_2_out_float64)

	res_0 := gocv.GroupRectangles(arg_0_out, arg_1_out, arg_2_out)

	res_0_out_arr := make([]interface{}, len(res_0), len(res_0))
	for res_0_i, res_0_v := range res_0 {
		res_0_out_arr_res_0_i_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Min_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Min_map["X"] = int64(res_0_v.Min.X)
		res_0_out_arr_res_0_i_map_Min_map["Y"] = int64(res_0_v.Min.Y)
		res_0_out_arr_res_0_i_map["Min"] = res_0_out_arr_res_0_i_map_Min_map
		res_0_out_arr_res_0_i_map_Max_map := make(map[string]interface{})
		res_0_out_arr_res_0_i_map_Max_map["X"] = int64(res_0_v.Max.X)
		res_0_out_arr_res_0_i_map_Max_map["Y"] = int64(res_0_v.Max.Y)
		res_0_out_arr_res_0_i_map["Max"] = res_0_out_arr_res_0_i_map_Max_map
		res_0_out_arr[res_0_i] = res_0_out_arr_res_0_i_map
	}
	res_0_out := res_0_out_arr
	res_0v, err := o.vm.ToValue(res_0_out)
	if err != nil {
		return o.handleError(err)
	}
	return res_0v
}

func (o *OttoFunctions) registFunctions() {
	o.vm.Set("cvIMRead", o.cvIMRead)
	o.vm.Set("cvIMWrite", o.cvIMWrite)
	o.vm.Set("cvIMWriteWithParams", o.cvIMWriteWithParams)
	o.vm.Set("cvIMEncode", o.cvIMEncode)
	o.vm.Set("cvIMEncodeWithParams", o.cvIMEncodeWithParams)
	o.vm.Set("cvIMDecode", o.cvIMDecode)
	o.vm.Set("cvNewWindow", o.cvNewWindow)
	o.vm.Set("cvWindowClose", o.cvWindowClose)
	o.vm.Set("cvWindowIsOpen", o.cvWindowIsOpen)
	o.vm.Set("cvWindowGetWindowProperty", o.cvWindowGetWindowProperty)
	o.vm.Set("cvWindowSetWindowProperty", o.cvWindowSetWindowProperty)
	o.vm.Set("cvWindowSetWindowTitle", o.cvWindowSetWindowTitle)
	o.vm.Set("cvWindowIMShow", o.cvWindowIMShow)
	o.vm.Set("cvWindowWaitKey", o.cvWindowWaitKey)
	o.vm.Set("cvWindowMoveWindow", o.cvWindowMoveWindow)
	o.vm.Set("cvWindowResizeWindow", o.cvWindowResizeWindow)
	o.vm.Set("cvSelectROI", o.cvSelectROI)
	o.vm.Set("cvSelectROIs", o.cvSelectROIs)
	o.vm.Set("cvWaitKey", o.cvWaitKey)
	o.vm.Set("cvWindowCreateTrackbar", o.cvWindowCreateTrackbar)
	o.vm.Set("cvTrackbarGetPos", o.cvTrackbarGetPos)
	o.vm.Set("cvTrackbarSetPos", o.cvTrackbarSetPos)
	o.vm.Set("cvTrackbarSetMin", o.cvTrackbarSetMin)
	o.vm.Set("cvTrackbarSetMax", o.cvTrackbarSetMax)
	o.vm.Set("cvNewMat", o.cvNewMat)
	o.vm.Set("cvNewMatWithSize", o.cvNewMatWithSize)
	o.vm.Set("cvNewMatFromScalar", o.cvNewMatFromScalar)
	o.vm.Set("cvNewMatWithSizeFromScalar", o.cvNewMatWithSizeFromScalar)
	o.vm.Set("cvNewMatFromBytes", o.cvNewMatFromBytes)
	o.vm.Set("cvMatFromPtr", o.cvMatFromPtr)
	o.vm.Set("cvMatClose", o.cvMatClose)
	o.vm.Set("cvMatEmpty", o.cvMatEmpty)
	o.vm.Set("cvMatClone", o.cvMatClone)
	o.vm.Set("cvMatCopyTo", o.cvMatCopyTo)
	o.vm.Set("cvMatCopyToWithMask", o.cvMatCopyToWithMask)
	o.vm.Set("cvMatConvertTo", o.cvMatConvertTo)
	o.vm.Set("cvMatTotal", o.cvMatTotal)
	o.vm.Set("cvMatSize", o.cvMatSize)
	o.vm.Set("cvMatToBytes", o.cvMatToBytes)
	o.vm.Set("cvMatDataPtrUint8", o.cvMatDataPtrUint8)
	o.vm.Set("cvMatDataPtrInt8", o.cvMatDataPtrInt8)
	o.vm.Set("cvMatDataPtrUint16", o.cvMatDataPtrUint16)
	o.vm.Set("cvMatDataPtrInt16", o.cvMatDataPtrInt16)
	o.vm.Set("cvMatDataPtrFloat32", o.cvMatDataPtrFloat32)
	o.vm.Set("cvMatDataPtrFloat64", o.cvMatDataPtrFloat64)
	o.vm.Set("cvMatRegion", o.cvMatRegion)
	o.vm.Set("cvMatReshape", o.cvMatReshape)
	o.vm.Set("cvMatConvertFp16", o.cvMatConvertFp16)
	o.vm.Set("cvMatMean", o.cvMatMean)
	o.vm.Set("cvMatSqrt", o.cvMatSqrt)
	o.vm.Set("cvMatSum", o.cvMatSum)
	o.vm.Set("cvMatPatchNaNs", o.cvMatPatchNaNs)
	o.vm.Set("cvLUT", o.cvLUT)
	o.vm.Set("cvMatRows", o.cvMatRows)
	o.vm.Set("cvMatCols", o.cvMatCols)
	o.vm.Set("cvMatChannels", o.cvMatChannels)
	o.vm.Set("cvMatType", o.cvMatType)
	o.vm.Set("cvMatStep", o.cvMatStep)
	o.vm.Set("cvMatGetUCharAt", o.cvMatGetUCharAt)
	o.vm.Set("cvMatGetUCharAt3", o.cvMatGetUCharAt3)
	o.vm.Set("cvMatGetSCharAt", o.cvMatGetSCharAt)
	o.vm.Set("cvMatGetSCharAt3", o.cvMatGetSCharAt3)
	o.vm.Set("cvMatGetShortAt", o.cvMatGetShortAt)
	o.vm.Set("cvMatGetShortAt3", o.cvMatGetShortAt3)
	o.vm.Set("cvMatGetIntAt", o.cvMatGetIntAt)
	o.vm.Set("cvMatGetIntAt3", o.cvMatGetIntAt3)
	o.vm.Set("cvMatGetFloatAt", o.cvMatGetFloatAt)
	o.vm.Set("cvMatGetFloatAt3", o.cvMatGetFloatAt3)
	o.vm.Set("cvMatGetDoubleAt", o.cvMatGetDoubleAt)
	o.vm.Set("cvMatGetDoubleAt3", o.cvMatGetDoubleAt3)
	o.vm.Set("cvMatSetTo", o.cvMatSetTo)
	o.vm.Set("cvMatSetUCharAt", o.cvMatSetUCharAt)
	o.vm.Set("cvMatSetUCharAt3", o.cvMatSetUCharAt3)
	o.vm.Set("cvMatSetSCharAt", o.cvMatSetSCharAt)
	o.vm.Set("cvMatSetSCharAt3", o.cvMatSetSCharAt3)
	o.vm.Set("cvMatSetShortAt", o.cvMatSetShortAt)
	o.vm.Set("cvMatSetShortAt3", o.cvMatSetShortAt3)
	o.vm.Set("cvMatSetIntAt", o.cvMatSetIntAt)
	o.vm.Set("cvMatSetIntAt3", o.cvMatSetIntAt3)
	o.vm.Set("cvMatSetFloatAt", o.cvMatSetFloatAt)
	o.vm.Set("cvMatSetFloatAt3", o.cvMatSetFloatAt3)
	o.vm.Set("cvMatSetDoubleAt", o.cvMatSetDoubleAt)
	o.vm.Set("cvMatSetDoubleAt3", o.cvMatSetDoubleAt3)
	o.vm.Set("cvMatAddUChar", o.cvMatAddUChar)
	o.vm.Set("cvMatSubtractUChar", o.cvMatSubtractUChar)
	o.vm.Set("cvMatMultiplyUChar", o.cvMatMultiplyUChar)
	o.vm.Set("cvMatDivideUChar", o.cvMatDivideUChar)
	o.vm.Set("cvMatAddFloat", o.cvMatAddFloat)
	o.vm.Set("cvMatSubtractFloat", o.cvMatSubtractFloat)
	o.vm.Set("cvMatMultiplyFloat", o.cvMatMultiplyFloat)
	o.vm.Set("cvMatDivideFloat", o.cvMatDivideFloat)
	o.vm.Set("cvMatToImage", o.cvMatToImage)
	o.vm.Set("cvImageToMatRGBA", o.cvImageToMatRGBA)
	o.vm.Set("cvImageToMatRGB", o.cvImageToMatRGB)
	o.vm.Set("cvImageGrayToMatGray", o.cvImageGrayToMatGray)
	o.vm.Set("cvAbsDiff", o.cvAbsDiff)
	o.vm.Set("cvAdd", o.cvAdd)
	o.vm.Set("cvAddWeighted", o.cvAddWeighted)
	o.vm.Set("cvBitwiseAnd", o.cvBitwiseAnd)
	o.vm.Set("cvBitwiseNot", o.cvBitwiseNot)
	o.vm.Set("cvBitwiseOr", o.cvBitwiseOr)
	o.vm.Set("cvBitwiseXor", o.cvBitwiseXor)
	o.vm.Set("cvBatchDistance", o.cvBatchDistance)
	o.vm.Set("cvBorderInterpolate", o.cvBorderInterpolate)
	o.vm.Set("cvCalcCovarMatrix", o.cvCalcCovarMatrix)
	o.vm.Set("cvCartToPolar", o.cvCartToPolar)
	o.vm.Set("cvCheckRange", o.cvCheckRange)
	o.vm.Set("cvCompare", o.cvCompare)
	o.vm.Set("cvCountNonZero", o.cvCountNonZero)
	o.vm.Set("cvCompleteSymm", o.cvCompleteSymm)
	o.vm.Set("cvConvertScaleAbs", o.cvConvertScaleAbs)
	o.vm.Set("cvCopyMakeBorder", o.cvCopyMakeBorder)
	o.vm.Set("cvDCT", o.cvDCT)
	o.vm.Set("cvDeterminant", o.cvDeterminant)
	o.vm.Set("cvDFT", o.cvDFT)
	o.vm.Set("cvDivide", o.cvDivide)
	o.vm.Set("cvEigen", o.cvEigen)
	o.vm.Set("cvEigenNonSymmetric", o.cvEigenNonSymmetric)
	o.vm.Set("cvExp", o.cvExp)
	o.vm.Set("cvExtractChannel", o.cvExtractChannel)
	o.vm.Set("cvFindNonZero", o.cvFindNonZero)
	o.vm.Set("cvFlip", o.cvFlip)
	o.vm.Set("cvGemm", o.cvGemm)
	o.vm.Set("cvGetOptimalDFTSize", o.cvGetOptimalDFTSize)
	o.vm.Set("cvHconcat", o.cvHconcat)
	o.vm.Set("cvVconcat", o.cvVconcat)
	o.vm.Set("cvRotate", o.cvRotate)
	o.vm.Set("cvIDCT", o.cvIDCT)
	o.vm.Set("cvIDFT", o.cvIDFT)
	o.vm.Set("cvInRange", o.cvInRange)
	o.vm.Set("cvInRangeWithScalar", o.cvInRangeWithScalar)
	o.vm.Set("cvInsertChannel", o.cvInsertChannel)
	o.vm.Set("cvInvert", o.cvInvert)
	o.vm.Set("cvLog", o.cvLog)
	o.vm.Set("cvMagnitude", o.cvMagnitude)
	o.vm.Set("cvMax", o.cvMax)
	o.vm.Set("cvMeanStdDev", o.cvMeanStdDev)
	o.vm.Set("cvMerge", o.cvMerge)
	o.vm.Set("cvMin", o.cvMin)
	o.vm.Set("cvMinMaxIdx", o.cvMinMaxIdx)
	o.vm.Set("cvMinMaxLoc", o.cvMinMaxLoc)
	o.vm.Set("cvMulSpectrums", o.cvMulSpectrums)
	o.vm.Set("cvMultiply", o.cvMultiply)
	o.vm.Set("cvNormalize", o.cvNormalize)
	o.vm.Set("cvNorm", o.cvNorm)
	o.vm.Set("cvPerspectiveTransform", o.cvPerspectiveTransform)
	o.vm.Set("cvSolve", o.cvSolve)
	o.vm.Set("cvSolveCubic", o.cvSolveCubic)
	o.vm.Set("cvSolvePoly", o.cvSolvePoly)
	o.vm.Set("cvReduce", o.cvReduce)
	o.vm.Set("cvRepeat", o.cvRepeat)
	o.vm.Set("cvSort", o.cvSort)
	o.vm.Set("cvSortIdx", o.cvSortIdx)
	o.vm.Set("cvSplit", o.cvSplit)
	o.vm.Set("cvSubtract", o.cvSubtract)
	o.vm.Set("cvTrace", o.cvTrace)
	o.vm.Set("cvTransform", o.cvTransform)
	o.vm.Set("cvTranspose", o.cvTranspose)
	o.vm.Set("cvPow", o.cvPow)
	o.vm.Set("cvPolarToCart", o.cvPolarToCart)
	o.vm.Set("cvPhase", o.cvPhase)
	o.vm.Set("cvNewTermCriteria", o.cvNewTermCriteria)
	o.vm.Set("cvNewScalar", o.cvNewScalar)
	o.vm.Set("cvMatGetVecfAt", o.cvMatGetVecfAt)
	o.vm.Set("cvMatGetVeciAt", o.cvMatGetVeciAt)
	o.vm.Set("cvGetTickCount", o.cvGetTickCount)
	o.vm.Set("cvGetTickFrequency", o.cvGetTickFrequency)
	o.vm.Set("cvArcLength", o.cvArcLength)
	o.vm.Set("cvApproxPolyDP", o.cvApproxPolyDP)
	o.vm.Set("cvConvexHull", o.cvConvexHull)
	o.vm.Set("cvConvexityDefects", o.cvConvexityDefects)
	o.vm.Set("cvCvtColor", o.cvCvtColor)
	o.vm.Set("cvEqualizeHist", o.cvEqualizeHist)
	o.vm.Set("cvCalcHist", o.cvCalcHist)
	o.vm.Set("cvBilateralFilter", o.cvBilateralFilter)
	o.vm.Set("cvBlur", o.cvBlur)
	o.vm.Set("cvBoxFilter", o.cvBoxFilter)
	o.vm.Set("cvSqBoxFilter", o.cvSqBoxFilter)
	o.vm.Set("cvDilate", o.cvDilate)
	o.vm.Set("cvErode", o.cvErode)
	o.vm.Set("cvBoundingRect", o.cvBoundingRect)
	o.vm.Set("cvContourArea", o.cvContourArea)
	o.vm.Set("cvMinAreaRect", o.cvMinAreaRect)
	o.vm.Set("cvFindContours", o.cvFindContours)
	o.vm.Set("cvMatchTemplate", o.cvMatchTemplate)
	o.vm.Set("cvMoments", o.cvMoments)
	o.vm.Set("cvPyrDown", o.cvPyrDown)
	o.vm.Set("cvPyrUp", o.cvPyrUp)
	o.vm.Set("cvMorphologyEx", o.cvMorphologyEx)
	o.vm.Set("cvGetStructuringElement", o.cvGetStructuringElement)
	o.vm.Set("cvGaussianBlur", o.cvGaussianBlur)
	o.vm.Set("cvSobel", o.cvSobel)
	o.vm.Set("cvSpatialGradient", o.cvSpatialGradient)
	o.vm.Set("cvLaplacian", o.cvLaplacian)
	o.vm.Set("cvScharr", o.cvScharr)
	o.vm.Set("cvMedianBlur", o.cvMedianBlur)
	o.vm.Set("cvCanny", o.cvCanny)
	o.vm.Set("cvCornerSubPix", o.cvCornerSubPix)
	o.vm.Set("cvGoodFeaturesToTrack", o.cvGoodFeaturesToTrack)
	o.vm.Set("cvHoughCircles", o.cvHoughCircles)
	o.vm.Set("cvHoughCirclesWithParams", o.cvHoughCirclesWithParams)
	o.vm.Set("cvHoughLines", o.cvHoughLines)
	o.vm.Set("cvHoughLinesP", o.cvHoughLinesP)
	o.vm.Set("cvHoughLinesPWithParams", o.cvHoughLinesPWithParams)
	o.vm.Set("cvThreshold", o.cvThreshold)
	o.vm.Set("cvAdaptiveThreshold", o.cvAdaptiveThreshold)
	o.vm.Set("cvArrowedLine", o.cvArrowedLine)
	o.vm.Set("cvCircle", o.cvCircle)
	o.vm.Set("cvEllipse", o.cvEllipse)
	o.vm.Set("cvLine", o.cvLine)
	o.vm.Set("cvRectangle", o.cvRectangle)
	o.vm.Set("cvFillPoly", o.cvFillPoly)
	o.vm.Set("cvGetTextSize", o.cvGetTextSize)
	o.vm.Set("cvPutText", o.cvPutText)
	o.vm.Set("cvResize", o.cvResize)
	o.vm.Set("cvGetRotationMatrix2D", o.cvGetRotationMatrix2D)
	o.vm.Set("cvWarpAffine", o.cvWarpAffine)
	o.vm.Set("cvWarpAffineWithParams", o.cvWarpAffineWithParams)
	o.vm.Set("cvWarpPerspective", o.cvWarpPerspective)
	o.vm.Set("cvApplyColorMap", o.cvApplyColorMap)
	o.vm.Set("cvApplyCustomColorMap", o.cvApplyCustomColorMap)
	o.vm.Set("cvGetPerspectiveTransform", o.cvGetPerspectiveTransform)
	o.vm.Set("cvDrawContours", o.cvDrawContours)
	o.vm.Set("cvRemap", o.cvRemap)
	o.vm.Set("cvFilter2D", o.cvFilter2D)
	o.vm.Set("cvSepFilter2D", o.cvSepFilter2D)
	o.vm.Set("cvLogPolar", o.cvLogPolar)
	o.vm.Set("cvFitLine", o.cvFitLine)
	o.vm.Set("cvFP16BlobFromImage", o.cvFP16BlobFromImage)
	o.vm.Set("cvVersion", o.cvVersion)
	o.vm.Set("cvOpenCVVersion", o.cvOpenCVVersion)
	o.vm.Set("cvFisheyeUndistortImage", o.cvFisheyeUndistortImage)
	o.vm.Set("cvFisheyeUndistortImageWithParams", o.cvFisheyeUndistortImageWithParams)
	o.vm.Set("cvParseNetBackend", o.cvParseNetBackend)
	o.vm.Set("cvParseNetTarget", o.cvParseNetTarget)
	o.vm.Set("cvNetClose", o.cvNetClose)
	o.vm.Set("cvNetEmpty", o.cvNetEmpty)
	o.vm.Set("cvNetSetInput", o.cvNetSetInput)
	o.vm.Set("cvNetForward", o.cvNetForward)
	o.vm.Set("cvNetForwardLayers", o.cvNetForwardLayers)
	o.vm.Set("cvNetSetPreferableBackend", o.cvNetSetPreferableBackend)
	o.vm.Set("cvNetSetPreferableTarget", o.cvNetSetPreferableTarget)
	o.vm.Set("cvReadNet", o.cvReadNet)
	o.vm.Set("cvReadNetFromCaffe", o.cvReadNetFromCaffe)
	o.vm.Set("cvReadNetFromTensorflow", o.cvReadNetFromTensorflow)
	o.vm.Set("cvBlobFromImage", o.cvBlobFromImage)
	o.vm.Set("cvGetBlobChannel", o.cvGetBlobChannel)
	o.vm.Set("cvGetBlobSize", o.cvGetBlobSize)
	o.vm.Set("cvNetGetLayer", o.cvNetGetLayer)
	o.vm.Set("cvNetGetPerfProfile", o.cvNetGetPerfProfile)
	o.vm.Set("cvNetGetUnconnectedOutLayers", o.cvNetGetUnconnectedOutLayers)
	o.vm.Set("cvLayerClose", o.cvLayerClose)
	o.vm.Set("cvLayerGetName", o.cvLayerGetName)
	o.vm.Set("cvLayerGetType", o.cvLayerGetType)
	o.vm.Set("cvLayerInputNameToIndex", o.cvLayerInputNameToIndex)
	o.vm.Set("cvLayerOutputNameToIndex", o.cvLayerOutputNameToIndex)
	o.vm.Set("cvVideoCaptureFile", o.cvVideoCaptureFile)
	o.vm.Set("cvVideoCaptureDevice", o.cvVideoCaptureDevice)
	o.vm.Set("cvVideoCaptureClose", o.cvVideoCaptureClose)
	o.vm.Set("cvVideoCaptureSet", o.cvVideoCaptureSet)
	o.vm.Set("cvVideoCaptureGet", o.cvVideoCaptureGet)
	o.vm.Set("cvVideoCaptureIsOpened", o.cvVideoCaptureIsOpened)
	o.vm.Set("cvVideoCaptureRead", o.cvVideoCaptureRead)
	o.vm.Set("cvVideoCaptureGrab", o.cvVideoCaptureGrab)
	o.vm.Set("cvVideoCaptureCodecString", o.cvVideoCaptureCodecString)
	o.vm.Set("cvVideoWriterFile", o.cvVideoWriterFile)
	o.vm.Set("cvVideoWriterClose", o.cvVideoWriterClose)
	o.vm.Set("cvVideoWriterIsOpened", o.cvVideoWriterIsOpened)
	o.vm.Set("cvVideoWriterWrite", o.cvVideoWriterWrite)
	o.vm.Set("cvOpenVideoCapture", o.cvOpenVideoCapture)
	o.vm.Set("cvNewAKAZE", o.cvNewAKAZE)
	o.vm.Set("cvAKAZEClose", o.cvAKAZEClose)
	o.vm.Set("cvAKAZEDetect", o.cvAKAZEDetect)
	o.vm.Set("cvAKAZEDetectAndCompute", o.cvAKAZEDetectAndCompute)
	o.vm.Set("cvNewAgastFeatureDetector", o.cvNewAgastFeatureDetector)
	o.vm.Set("cvAgastFeatureDetectorClose", o.cvAgastFeatureDetectorClose)
	o.vm.Set("cvAgastFeatureDetectorDetect", o.cvAgastFeatureDetectorDetect)
	o.vm.Set("cvNewBRISK", o.cvNewBRISK)
	o.vm.Set("cvBRISKClose", o.cvBRISKClose)
	o.vm.Set("cvBRISKDetect", o.cvBRISKDetect)
	o.vm.Set("cvBRISKDetectAndCompute", o.cvBRISKDetectAndCompute)
	o.vm.Set("cvNewFastFeatureDetector", o.cvNewFastFeatureDetector)
	o.vm.Set("cvFastFeatureDetectorClose", o.cvFastFeatureDetectorClose)
	o.vm.Set("cvFastFeatureDetectorDetect", o.cvFastFeatureDetectorDetect)
	o.vm.Set("cvNewGFTTDetector", o.cvNewGFTTDetector)
	o.vm.Set("cvGFTTDetectorClose", o.cvGFTTDetectorClose)
	o.vm.Set("cvGFTTDetectorDetect", o.cvGFTTDetectorDetect)
	o.vm.Set("cvNewKAZE", o.cvNewKAZE)
	o.vm.Set("cvKAZEClose", o.cvKAZEClose)
	o.vm.Set("cvKAZEDetect", o.cvKAZEDetect)
	o.vm.Set("cvKAZEDetectAndCompute", o.cvKAZEDetectAndCompute)
	o.vm.Set("cvNewMSER", o.cvNewMSER)
	o.vm.Set("cvMSERClose", o.cvMSERClose)
	o.vm.Set("cvMSERDetect", o.cvMSERDetect)
	o.vm.Set("cvNewORB", o.cvNewORB)
	o.vm.Set("cvORBClose", o.cvORBClose)
	o.vm.Set("cvORBDetect", o.cvORBDetect)
	o.vm.Set("cvORBDetectAndCompute", o.cvORBDetectAndCompute)
	o.vm.Set("cvNewSimpleBlobDetector", o.cvNewSimpleBlobDetector)
	o.vm.Set("cvSimpleBlobDetectorClose", o.cvSimpleBlobDetectorClose)
	o.vm.Set("cvSimpleBlobDetectorDetect", o.cvSimpleBlobDetectorDetect)
	o.vm.Set("cvNewBFMatcher", o.cvNewBFMatcher)
	o.vm.Set("cvNewBFMatcherWithParams", o.cvNewBFMatcherWithParams)
	o.vm.Set("cvBFMatcherClose", o.cvBFMatcherClose)
	o.vm.Set("cvBFMatcherKnnMatch", o.cvBFMatcherKnnMatch)
	o.vm.Set("cvDrawKeyPoints", o.cvDrawKeyPoints)
	o.vm.Set("cvNewBackgroundSubtractorMOG2", o.cvNewBackgroundSubtractorMOG2)
	o.vm.Set("cvBackgroundSubtractorMOG2Close", o.cvBackgroundSubtractorMOG2Close)
	o.vm.Set("cvBackgroundSubtractorMOG2Apply", o.cvBackgroundSubtractorMOG2Apply)
	o.vm.Set("cvNewBackgroundSubtractorKNN", o.cvNewBackgroundSubtractorKNN)
	o.vm.Set("cvBackgroundSubtractorKNNClose", o.cvBackgroundSubtractorKNNClose)
	o.vm.Set("cvBackgroundSubtractorKNNApply", o.cvBackgroundSubtractorKNNApply)
	o.vm.Set("cvCalcOpticalFlowFarneback", o.cvCalcOpticalFlowFarneback)
	o.vm.Set("cvCalcOpticalFlowPyrLK", o.cvCalcOpticalFlowPyrLK)
	o.vm.Set("cvNewCascadeClassifier", o.cvNewCascadeClassifier)
	o.vm.Set("cvCascadeClassifierClose", o.cvCascadeClassifierClose)
	o.vm.Set("cvCascadeClassifierLoad", o.cvCascadeClassifierLoad)
	o.vm.Set("cvCascadeClassifierDetectMultiScale", o.cvCascadeClassifierDetectMultiScale)
	o.vm.Set("cvCascadeClassifierDetectMultiScaleWithParams", o.cvCascadeClassifierDetectMultiScaleWithParams)
	o.vm.Set("cvNewHOGDescriptor", o.cvNewHOGDescriptor)
	o.vm.Set("cvHOGDescriptorClose", o.cvHOGDescriptorClose)
	o.vm.Set("cvHOGDescriptorDetectMultiScale", o.cvHOGDescriptorDetectMultiScale)
	o.vm.Set("cvHOGDescriptorDetectMultiScaleWithParams", o.cvHOGDescriptorDetectMultiScaleWithParams)
	o.vm.Set("cvHOGDefaultPeopleDetector", o.cvHOGDefaultPeopleDetector)
	o.vm.Set("cvHOGDescriptorSetSVMDetector", o.cvHOGDescriptorSetSVMDetector)
	o.vm.Set("cvGroupRectangles", o.cvGroupRectangles)
}
