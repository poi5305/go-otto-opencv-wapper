package cv2js

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type JSGenerator struct {
	scope          string
	outputPath     string
	funcPrefix     string
	cPackageStruct *PackageStruct
	debugInputs    map[string]int
	debugOutputs   map[string]int
}

func NewJSGenerator(scope, funcPrefix, output string) *JSGenerator {
	return &JSGenerator{
		scope:        scope,
		outputPath:   output,
		funcPrefix:   funcPrefix,
		debugInputs:  make(map[string]int),
		debugOutputs: make(map[string]int),
	}
}

func (j *JSGenerator) GenHeader() string {
	code := `
package %s

import (
	"fmt"
	"image"
	"image/color"

	"github.com/robertkrimen/otto"
	"gocv.io/x/gocv"
)

`
	return fmt.Sprintf(code, j.scope)
}

func (j *JSGenerator) PrintCounts() {
	fmt.Println("Function inputs\n\n")
	for key, count := range j.debugInputs {
		fmt.Println(key, count)
	}
	fmt.Println("Function outputs\n\n")
	for key, count := range j.debugOutputs {
		fmt.Println(key, count)
	}
}

func (j *JSGenerator) GenStructs(packStruct *PackageStruct) {
	j.cPackageStruct = packStruct

	for name, field := range j.cPackageStruct.TypeDefineTable {
		if !field.IsStruct {
			continue
		}
		if string(name[0]) == strings.ToLower(string(name[0])) {
			continue
		}
		fmt.Println("Struct", name)
		for _, cFields := range field.ChildFields {
			fieldName := cFields.FieldName
			if string(fieldName[0]) == strings.ToLower(string(fieldName[0])) {
				continue
			}
			fmt.Println("\t", name, cFields.FieldName, cFields.TypeName)
		}
	}
}

func (j *JSGenerator) GenFunctions(packStruct *PackageStruct) {
	code := j.GenHeader()
	registCode := "func (o *OttoFunctions) registFunctions() {\n"
	j.cPackageStruct = packStruct
	for _, fn := range j.cPackageStruct.FunctionDefines {
		// FUNCTION NAME
		functionName := j.funcPrefix
		if len(fn.ThisTypes) == 1 {
			if fn.ThisTypes[0].ChildField != nil {
				functionName += fn.ThisTypes[0].ChildField.ValueType
			} else {
				functionName += fn.ThisTypes[0].ValueType
			}
		}
		functionName += fn.Name

		registCode += "\t" + fmt.Sprintf("o.vm.Set(\"%s\", o.%s)\n", functionName, functionName)

		code += fmt.Sprintf("func (o *OttoFunctions) %s(call otto.FunctionCall) otto.Value {\n", functionName)
		// INPUT
		code += j.GenFunctionInputConverter(append(fn.ThisTypes, fn.InputTypes...))
		// CALL
		code += j.GenCallFuncCode(fn)
		// OUTPUT
		code += j.GenFunctionOutputConverter(fn.OutputTypes)
		// RETURN
		code += j.GenReturnCode(fn.OutputTypes)
		code += "}\n\n"
	}
	registCode += "}\n"
	code += registCode
	ioutil.WriteFile(j.outputPath, []byte(code), 0777)
}

func (j *JSGenerator) GenFunctionInputConverter(fields []*FieldType) string {
	code := ""

	for i, field := range fields {
		j.debugInputs[field.TypeName]++
		if i == 0 {
			code += "\t" + fmt.Sprintf("var ok bool\n")
		}
		argName := fmt.Sprintf("arg_%d_in", i)
		argNameV := fmt.Sprintf("arg_%d_out", i)
		code += "\t" + fmt.Sprintf("%s, err := call.Argument(%d).Export()\n", argName, i)
		code += genCheckNil("err")
		code += j.genInputConverterImpl(field, argName, argNameV, 1, false)
	}
	// code += "\t" + fmt.Sprintf("fmt.Println(ok)\n")
	return code
}

func (j *JSGenerator) GenFunctionOutputConverter(fields []*FieldType) string {
	code := ""
	for i, field := range fields {
		j.debugOutputs[field.TypeName]++
		// code += j.GenOutputConverter(i, field)
		in := fmt.Sprintf("res_%d", i)
		out := fmt.Sprintf("res_%d_out", i)
		outV := fmt.Sprintf("res_%dv", i)
		code += j.genOutputConverterImpl(field, in, out, 1, false)

		if field.TypeName != "error" {
			code += tab(1) + fmt.Sprintf("%s, err := o.vm.ToValue(%s)\n", outV, out)
			code += genCheckNil("err")
		}
	}
	return code
}

func getCleanName(s string) string {
	return strings.Replace(strings.Replace(strings.Replace(strings.Replace(s, "\"", "", -1), "]", "", -1), "[", "_", -1), ".", "_", -1)
}

func tab(layer int) string {
	tab := ""
	for i := 0; i < layer; i++ {
		tab += "\t"
	}
	return tab
}

func (j *JSGenerator) genInputConverterImpl(field *FieldType, in, out string, layer int, isDefined bool) string {
	code := ""
	assign := ":"
	if isDefined {
		assign = ""
	}
	clearInNameV := getCleanName(in)
	clearOutNameV := getCleanName(out)

	if field.IsBasic {
		ottoType := mappingToOttoType(field.TypeName)
		outOttoName := fmt.Sprintf("%s_%s", clearOutNameV, ottoType)
		if ottoType == "" {
			fmt.Println("error no otto type", field)
		}
		code += genInterfaceToType(in, outOttoName, ottoType, layer, false)
		code += tab(layer) + fmt.Sprintf("%s %s= %s(%s)\n", out, assign, field.TypeName, outOttoName)
	} else if field.IsSimple {
		vField, ok := j.cPackageStruct.TypeDefineTable[field.ValueType]
		if !ok {
			fmt.Println("Error, no type found in table", field.ValueType)
			return ""
		}
		childOutNameV := fmt.Sprintf("%s_%s", clearOutNameV, field.ValueType)
		copyField := *vField
		if copyField.TypeName == "" {
			copyField.TypeName = field.TypeName
		}
		if copyField.ValueType == "" {
			copyField.ValueType = field.ValueType
		}
		code += j.genInputConverterImpl(&copyField, in, childOutNameV, layer, false)
		code += tab(layer) + fmt.Sprintf("%s %s= %s(%s)\n", out, assign, field.TypeName, childOutNameV)
	} else if field.IsStruct {
		if len(field.ChildFields) == 0 {
			vField, _ := j.cPackageStruct.TypeDefineTable[field.TypeName]
			if vField != nil && len(vField.ChildFields) != 0 {
				return j.genInputConverterImpl(vField, in, out, layer, isDefined)
			}
			code += genInterfaceToTypeValueWithPtr(in, out, field.TypeName, layer, isDefined)
			return code
		}
		// map[string]interface{}
		code += tab(layer) + fmt.Sprintf("%s %s= %s{}\n", out, assign, field.TypeName)
		childOutMapName := fmt.Sprintf("%s_map", clearOutNameV)
		code += genInterfaceToType(in, childOutMapName, "map[string]interface{}", layer, false)
		for _, cField := range field.ChildFields {
			// argNameV.xx = childArgMapNameV["xx"]
			childInName := fmt.Sprintf("%s[\"%s\"]", childOutMapName, cField.FieldName)
			childOutName := fmt.Sprintf("%s.%s", out, cField.FieldName)
			code += j.genInputConverterImpl(cField, childInName, childOutName, layer, true)
		}
	} else if field.IsArray && field.ValueType == "byte" {
		outOttoName := fmt.Sprintf("%s_string", clearOutNameV)
		code += genInterfaceToType(in, outOttoName, "string", layer, false)
		code += tab(layer) + fmt.Sprintf("%s, err %s= base64strToBytes(%s)\n", out, assign, outOttoName)
		code += genCheckNil("err")
	} else if field.IsArray {
		// []interface()
		childInArrayName := fmt.Sprintf("%s_arr", clearInNameV)
		code += genInterfaceToType(in, childInArrayName, "[]interface{}", layer, false)
		code += tab(layer) + fmt.Sprintf("%s %s= make(%s, len(%s), len(%s))\n", out, assign, field.TypeName, childInArrayName, childInArrayName)

		arrayIdxName := fmt.Sprintf("%s_i", childInArrayName)
		arrayValueName := fmt.Sprintf("%s_v", childInArrayName)
		code += tab(layer) + fmt.Sprintf("for %s, %s := range %s {\n", arrayIdxName, arrayValueName, childInArrayName)

		childOutName := fmt.Sprintf("%s[%s]", out, arrayIdxName)
		code += j.genInputConverterImpl(field.ChildField, arrayValueName, childOutName, layer+1, true)
		code += tab(layer) + fmt.Sprintf("}\n")
	} else if field.IsMap {
		// TODO have not test this
		childInMapName := fmt.Sprintf("%s_map", clearInNameV)
		code += genInterfaceToType(in, childInMapName, "map[string]interface{}", layer, false)
		code += tab(layer) + fmt.Sprintf("%s %s= make(%s)\n", out, assign, field.TypeName)

		arrayIdxName := fmt.Sprintf("%s_i", childInMapName)
		arrayValueName := fmt.Sprintf("%s_v", childInMapName)
		code += tab(layer) + fmt.Sprintf("for %s, %s := range %s {\n", arrayIdxName, arrayValueName, childInMapName)

		childOutName := fmt.Sprintf("%s[%s]", out, arrayIdxName)
		code += j.genInputConverterImpl(field.ChildField, arrayValueName, childOutName, layer+1, true)
		code += tab(layer) + fmt.Sprintf("}\n")
	} else if field.IsPointer {
		// TODO maybe tyr to de reference ptr
		code += genInterfaceToTypePtrWithValue(in, out, field.ValueType, layer, isDefined)
	} else if field.IsInterface {
		// currently only for special case: int and string
		inInt64 := fmt.Sprintf("%s_int64", clearInNameV)
		inString := fmt.Sprintf("%s_string", clearInNameV)
		code += fmt.Sprintf(`
	var %s interface{}
	%s, ok := %s.(int64)
	if ok {
		%s = int(%s)
	}
	%s, ok := %s.(string)
	if ok {
		%s = string(%s)
	}
`, out, inInt64, in, out, inInt64, inString, in, out, inString)
	}
	return code
}

func (j *JSGenerator) genOutputConverterImpl(field *FieldType, in, out string, layer int, isDefined bool) string {
	code := ""
	assign := ":"
	if isDefined {
		assign = ""
	}
	clearInNameV := getCleanName(in)
	clearOutNameV := getCleanName(out)

	if field.IsBasic && field.TypeName == "error" {
		code += genCheckNil(fmt.Sprintf("%s", in))
	} else if field.IsBasic {
		ottoType := mappingToOttoType(field.TypeName)
		if ottoType == "" {
			fmt.Println("error no otto type", field)
		}
		code += tab(layer) + fmt.Sprintf("%s %s= %s(%s)\n", out, assign, ottoType, in)
	} else if field.IsSimple {
		vField, ok := j.cPackageStruct.TypeDefineTable[field.ValueType]
		if !ok {
			fmt.Println("Error, no type found in table", field.ValueType)
			return ""
		}
		// childOutNameV := fmt.Sprintf("%s_%s", clearOutNameV, field.ValueType)
		copyField := *vField
		if copyField.TypeName == "" {
			copyField.TypeName = field.TypeName
		}
		if copyField.ValueType == "" {
			copyField.ValueType = field.ValueType
		}
		code += j.genOutputConverterImpl(&copyField, in, out, layer, isDefined)
		// fmt.Println(code)
	} else if field.IsStruct {
		if len(field.ChildFields) == 0 {
			vField, _ := j.cPackageStruct.TypeDefineTable[field.TypeName]
			if vField != nil && len(vField.ChildFields) != 0 {
				return j.genOutputConverterImpl(vField, in, out, layer, isDefined)
			}
			code += tab(layer) + fmt.Sprintf("%s %s= %s(%s)\n", out, assign, field.TypeName, in)
			return code
		}
		// map[string]interface{}
		mapOutName := fmt.Sprintf("%s_map", clearOutNameV)
		code += tab(layer) + fmt.Sprintf("%s := make(map[string]interface{})\n", mapOutName)
		for _, cField := range field.ChildFields {
			// out["xx"] = in.xx
			childInName := fmt.Sprintf("%s.%s", in, cField.FieldName)
			childOutName := fmt.Sprintf("%s[\"%s\"]", mapOutName, cField.FieldName)
			code += j.genOutputConverterImpl(cField, childInName, childOutName, layer, true)
		}
		code += tab(layer) + fmt.Sprintf("%s %s= %s\n", out, assign, mapOutName)
	} else if field.IsArray && field.ValueType == "byte" {
		code += tab(layer) + fmt.Sprintf("%s %s= bytesToBase64str(%s)\n", out, assign, in)
		code += genCheckNil("err")
	} else if field.IsArray {
		// []interface{}
		arrOutName := fmt.Sprintf("%s_arr", clearOutNameV)
		code += tab(layer) + fmt.Sprintf("%s := make([]interface{}, len(%s), len(%s))\n", arrOutName, in, in)

		arrayIdxName := fmt.Sprintf("%s_i", clearInNameV)
		arrayValueName := fmt.Sprintf("%s_v", clearInNameV)
		code += tab(layer) + fmt.Sprintf("for %s, %s := range %s {\n", arrayIdxName, arrayValueName, in)
		childOutName := fmt.Sprintf("%s[%s]", arrOutName, arrayIdxName)
		code += j.genOutputConverterImpl(field.ChildField, arrayValueName, childOutName, layer+1, true)
		code += tab(layer) + fmt.Sprintf("}\n")
		code += tab(layer) + fmt.Sprintf("%s %s= %s\n", out, assign, arrOutName)
	} else if field.IsMap {
		// code += tab(layer) + fmt.Sprintf("%s %s= (%s)(%s)\n", out, assign, field.TypeName, in)
		mapOutName := fmt.Sprintf("%s_map", clearOutNameV)
		code += tab(layer) + fmt.Sprintf("%s := make(map[string]interface{})\n", mapOutName)

		arrayIdxName := fmt.Sprintf("%s_i", clearInNameV)
		arrayValueName := fmt.Sprintf("%s_v", clearInNameV)
		code += tab(layer) + fmt.Sprintf("for %s, %s := range %s {\n", arrayIdxName, arrayValueName, in)
		childOutName := fmt.Sprintf("%s[%s]", mapOutName, arrayIdxName)
		code += j.genOutputConverterImpl(field.ChildField, arrayValueName, childOutName, layer+1, true)
		code += tab(layer) + fmt.Sprintf("}\n")
		code += tab(layer) + fmt.Sprintf("%s %s= %s\n", out, assign, mapOutName)
	} else if field.IsPointer {
		// TODO maybe dereference pointer
		code += tab(layer) + fmt.Sprintf("%s %s= (%s)(%s)\n", out, assign, field.TypeName, in)
	} else {
		code += tab(layer) + fmt.Sprintf("%s %s= (%s)(%s)\n", out, assign, field.TypeName, in)
	}
	return code
}

func (j *JSGenerator) GenCallFuncCode(fn *FuncDef) string {
	code := "\n\t"
	// gen output
	for i := range fn.OutputTypes {
		if i == len(fn.OutputTypes)-1 {
			code += fmt.Sprintf("res_%d := ", i)
		} else {
			code += fmt.Sprintf("res_%d, ", i)
		}
	}
	// gen func this
	inputIdx := 0
	if len(fn.ThisTypes) == 1 {
		code += fmt.Sprintf("arg_%d_out.", inputIdx)
		inputIdx++
	} else {
		code += j.cPackageStruct.PackageScope + "."
	}
	code += fmt.Sprintf("%s(", fn.Name)
	for i := range fn.InputTypes {
		if i == len(fn.InputTypes)-1 {
			code += fmt.Sprintf("arg_%d_out", inputIdx)
		} else {
			code += fmt.Sprintf("arg_%d_out, ", inputIdx)
		}
		inputIdx++
	}
	code += fmt.Sprintf(")\n\n")
	return code
}

func (j *JSGenerator) GenReturnCode(fields []*FieldType) string {
	code := ""
	nonErrorIdx := 0
	returnCount := 0
	for i, field := range fields {
		if field.TypeName != "error" {
			returnCount++
			nonErrorIdx = i
		}
	}
	if returnCount == 0 {
		code += "\t" + fmt.Sprintf("return otto.Value{}\n")
	} else if returnCount == 1 {
		code += "\t" + fmt.Sprintf("return res_%dv\n", nonErrorIdx)
	} else {
		// convert to map
		code += "\t" + fmt.Sprintf("results := map[string]interface{}{}\n")
		for i, field := range fields {
			if field.TypeName == "error" {
				continue
			}
			code += "\t" + fmt.Sprintf("results[\"r%d\"] = res_%dv\n", i, i)
		}
		code += "\t" + fmt.Sprintf("resultsv, err := o.vm.ToValue(results)\n")
		code += "\t" + fmt.Sprintf("if err != nil {\n")
		code += "\t\t" + fmt.Sprintf("return o.handleError(err)\n")
		code += "\t" + fmt.Sprintf("}\n")
		code += "\t" + fmt.Sprintf("return resultsv\n")
	}
	return code
}

func genCheckNil(varName string) string {
	code := "\t" + fmt.Sprintf("if %s != nil {\n", varName)
	code += "\t\t" + fmt.Sprintf("return o.handleError(%s)\n", varName)
	code += "\t" + fmt.Sprintf("}\n")
	return code
}

func genInterfaceToTypeValueWithPtr(in, out, toType string, layer int, isDefined bool) string {
	assign := ":"
	if isDefined {
		assign = ""
	}
	clearOutNameV := getCleanName(out)
	outTmp := clearOutNameV + "_tmp"
	code := ""
	code += tab(layer) + fmt.Sprintf("%s, ok %s= %s.(%s)\n", out, assign, in, toType)
	code += tab(layer) + fmt.Sprintf("if !ok {\n")
	code += tab(layer+1) + fmt.Sprintf("%s, ok := %s.(*%s)\n", outTmp, in, toType)
	code += tab(layer+1) + fmt.Sprintf("if !ok {\n")
	code += tab(layer+2) + fmt.Sprintf("return o.handleError(fmt.Errorf(\"Can not Convert To Type: %s.(%s)\"))\n", getCleanName(in), toType)
	code += tab(layer+1) + fmt.Sprintf("}\n")
	code += tab(layer+1) + fmt.Sprintf("%s = *(%s)\n", out, outTmp)
	code += tab(layer) + fmt.Sprintf("}\n")
	return code
}

func genInterfaceToTypePtrWithValue(in, out, toType string, layer int, isDefined bool) string {
	assign := ":"
	if isDefined {
		assign = ""
	}
	clearOutNameV := getCleanName(out)
	outTmp := clearOutNameV + "_tmp"
	code := ""
	code += tab(layer) + fmt.Sprintf("%s, ok %s= %s.(*%s)\n", out, assign, in, toType)
	code += tab(layer) + fmt.Sprintf("if !ok {\n")
	code += tab(layer+1) + fmt.Sprintf("%s, ok := %s.(%s)\n", outTmp, in, toType)
	code += tab(layer+1) + fmt.Sprintf("if !ok {\n")
	code += tab(layer+2) + fmt.Sprintf("return o.handleError(fmt.Errorf(\"Can not Convert To Type: %s.(%s)\"))\n", getCleanName(in), toType)
	code += tab(layer+1) + fmt.Sprintf("}\n")
	code += tab(layer+1) + fmt.Sprintf("%s = &(%s)\n", out, outTmp)
	code += tab(layer) + fmt.Sprintf("}\n")
	return code
}

func genInterfaceToType(in, out, toType string, layer int, isDefined bool) string {
	assign := ":"
	if isDefined {
		assign = ""
	}
	code := tab(layer) + fmt.Sprintf("%s, ok %s= %s.(%s)\n", out, assign, in, toType)
	code += tab(layer) + fmt.Sprintf("if !ok {\n")
	code += tab(layer+1) + fmt.Sprintf("return o.handleError(fmt.Errorf(\"Can not Convert To Type: %s.(%s)\"))\n", getCleanName(in), toType)
	code += tab(layer) + fmt.Sprintf("}\n")
	return code
}

// int -> int64, float32 -> float64, Mat -> ""
func mappingToOttoType(t string) string {
	switch t {
	case "int", "int8", "int16", "int32", "int64":
		return "int64"
	case "uint", "uint8", "uint16", "uint32", "uint64", "byte":
		return "int64"
	case "float32", "float64":
		return "float64"
	case "string":
		return "string"
	case "bool":
		return "bool"
	}
	return ""
}

func getOttoArgumentCall(ottoType string) string {
	switch ottoType {
	case "int64":
		return "ToInteger"
	case "float64":
		return "ToFloat"
	case "bool":
		return "ToBoolean"
	case "string":
		return "ToString"
	}
	return "Export"
}
