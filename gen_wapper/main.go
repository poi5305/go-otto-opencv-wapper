package main

import (
	"os"

	cv2js "github.com/poi5305/go-otto-opencv-wapper"
)

func main() {
	gocvPath := os.Getenv("GOPATH") + "/src/gocv.io/x/gocv"

	p := cv2js.NewPackageSruct(gocvPath, "gocv")
	p.Parse()
	p.TypeDefineTable["image.Point"] = &cv2js.FieldType{
		IsStruct: true,
		TypeName: "image.Point",
		ChildFields: []*cv2js.FieldType{
			&cv2js.FieldType{
				IsBasic:   true,
				FieldName: "X",
				TypeName:  "int",
			},
			&cv2js.FieldType{
				IsBasic:   true,
				FieldName: "Y",
				TypeName:  "int",
			},
		},
	}

	p.TypeDefineTable["image.Rectangle"] = &cv2js.FieldType{
		IsStruct: true,
		TypeName: "image.Rectangle",
		ChildFields: []*cv2js.FieldType{
			&cv2js.FieldType{
				IsStruct:    true,
				FieldName:   "Min",
				TypeName:    "image.Point",
				ChildFields: p.TypeDefineTable["image.Point"].ChildFields,
			},
			&cv2js.FieldType{
				IsStruct:    true,
				FieldName:   "Max",
				TypeName:    "image.Point",
				ChildFields: p.TypeDefineTable["image.Point"].ChildFields,
			},
		},
	}

	// *** change this path for output ***
	jsg := cv2js.NewJSGenerator("otto_wrapper", "cv", "../otto_wrapper/otto_gocv.go")
	jsg.GenFunctions(p)
}
