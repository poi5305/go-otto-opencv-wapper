package cv2js

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type FieldType struct {
	IsSimple    bool
	IsBasic     bool
	IsPointer   bool
	IsArray     bool
	IsMap       bool
	IsStruct    bool
	IsInterface bool
	IsValue     bool
	Scope       string
	FieldName   string
	TypeName    string
	KeyType     string       // for map
	ValueType   string       // for map, array and pointer, value define
	ChildField  *FieldType   // for map, array, pointer, ...
	ChildFields []*FieldType // for struct
}

type FuncDef struct {
	Name        string
	ThisTypes   []*FieldType
	InputTypes  []*FieldType
	OutputTypes []*FieldType
}

func NewPackageSruct(path string, scope string) *PackageStruct {
	p := PackageStruct{
		PackagePath:     path,
		PackageScope:    scope,
		TypeDefineTable: make(map[string]*FieldType),
		FunctionDefines: make([]*FuncDef, 0, 100),
	}
	return &p
}

type PackageStruct struct {
	PackagePath     string
	PackageScope    string
	TypeDefineTable map[string]*FieldType
	FunctionDefines []*FuncDef
}

func (p *PackageStruct) Parse() {
	set := token.NewFileSet()
	packs, err := parser.ParseDir(set, p.PackagePath, nil, 0)
	if err != nil {
		fmt.Println("Failed to parse package:", err)
		os.Exit(1)
	}
	for _, pack := range packs {
		for filename, files := range pack.Files {
			if strings.Contains(filename, "test") {
				continue
			}
			fmt.Println("Parsing ...", filename)
			for _, decls := range files.Decls {
				switch decls.(type) {
				case *ast.FuncDecl:
					fn := decls.(*ast.FuncDecl)
					funcName := fn.Name.Name
					if string(funcName[0]) == strings.ToLower(string(funcName[0])) {
						continue
					}
					p.parseFuncDecl(fn)
				case *ast.GenDecl:
					decl := decls.(*ast.GenDecl)
					p.parseGenDecl(decl)
				} // switch
			} // for _, decls := range files.Decls
		}
	}
}

func (p *PackageStruct) parseFuncDecl(fn *ast.FuncDecl) {
	receiveTypes := p.parseFieldList(fn.Recv, false)
	fieldTypes := p.parseFieldList(fn.Type.Params, false)
	resultTypes := p.parseFieldList(fn.Type.Results, false)
	// ignore any C.Type
	for _, f := range fieldTypes {
		if f.Scope == "C" {
			return
		}
	}
	for _, f := range resultTypes {
		if f.Scope == "C" {
			return
		}
	}
	p.FunctionDefines = append(p.FunctionDefines, &FuncDef{
		fn.Name.Name,
		receiveTypes,
		fieldTypes,
		resultTypes,
	})
}

func (p *PackageStruct) parseGenDecl(decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		switch spec.(type) {
		case *ast.TypeSpec:
			p.parseStructType(spec.(*ast.TypeSpec))
		case *ast.ValueSpec:
			p.parseValueType(spec.(*ast.ValueSpec))
		}
	}
}

func (p *PackageStruct) parseStructType(decl *ast.TypeSpec) {
	ft := p.parseExpr(decl.Type)
	stName := decl.Name.Name
	p.TypeDefineTable[stName] = ft
}

func (p *PackageStruct) parseValueType(decl *ast.ValueSpec) {
	// TODO parse Value define

	// for _, value := range decl.Values {
	// 	valueFT := p.parseExpr(value)
	// 	for _, fieldName := range decl.Names {
	// 		ft := &FieldType{
	// 			IsSimple:  false,
	// 			IsValue:   true,
	// 			Scope:     p.PackageScope,
	// 			FieldName: fieldName.Name,
	// 			ValueType: valueFT.ValueType,
	// 		}
	// 		p.TypeDefineTable[fieldName.Name] = ft
	// 	}
	// }
}

func (p *PackageStruct) parseFieldList(fields *ast.FieldList, requirePublicName bool) []*FieldType {
	if fields == nil || fields.List == nil {
		return []*FieldType{}
	}
	types := make([]*FieldType, 0, fields.NumFields())
	for _, param := range fields.List {
		count := len(param.Names)
		if count == 0 {
			ft := p.parseExpr(param.Type)
			if requirePublicName {
				continue
			}
			types = append(types, ft)
		} else {
			for i := 0; i < count; i++ {
				ft := p.parseExpr(param.Type)
				ft.FieldName = param.Names[i].Name
				if requirePublicName && string(ft.FieldName[0]) == strings.ToLower(string(ft.FieldName[0])) {
					continue
				}
				types = append(types, ft)
			}
		}
	}
	return types
}

func (p *PackageStruct) parseExpr(expr ast.Expr) *FieldType {
	if expr == nil {
		return nil
	}
	ft := FieldType{
		IsSimple: false,
	}
	switch expr.(type) {
	case *ast.Ident:
		ft.IsSimple = true
		ft.IsBasic = true
		ft.TypeName = expr.(*ast.Ident).Name
		ft.ValueType = expr.(*ast.Ident).Name
		if !isBasicType(ft.TypeName) {
			ft.IsBasic = false
			ft.Scope = p.PackageScope
			ft.TypeName = ft.Scope + "." + ft.TypeName
		}
	case *ast.StarExpr:
		ft.ChildField = p.parseExpr(expr.(*ast.StarExpr).X)
		ft.IsPointer = true
		ft.TypeName = "*" + ft.ChildField.TypeName
		ft.ValueType = ft.ChildField.TypeName
	case *ast.ArrayType:
		ft.ChildField = p.parseExpr(expr.(*ast.ArrayType).Elt)
		ft.IsArray = true
		ft.TypeName = "[]" + ft.ChildField.TypeName
		ft.ValueType = ft.ChildField.TypeName
	case *ast.SelectorExpr:
		ft.IsStruct = true
		ft.Scope = expr.(*ast.SelectorExpr).X.(*ast.Ident).Name
		ft.TypeName = ft.Scope + "." + expr.(*ast.SelectorExpr).Sel.Name
		ft.ValueType = ft.TypeName
	case *ast.MapType:
		keyFT := p.parseExpr(expr.(*ast.MapType).Key)
		valueFT := p.parseExpr(expr.(*ast.MapType).Value)
		ft.ChildField = valueFT
		ft.IsMap = true
		ft.TypeName = fmt.Sprintf("map[%s]%s", keyFT.TypeName, valueFT.TypeName)
		ft.KeyType = keyFT.TypeName
		ft.ValueType = valueFT.TypeName
	case *ast.StructType:
		ft.IsStruct = true
		ft.ChildFields = p.parseFieldList(expr.(*ast.StructType).Fields, true)
		ft.Scope = p.PackageScope
	case *ast.InterfaceType:
		ft.IsInterface = true
	case *ast.BasicLit:
		ft.IsValue = true
		ft.ValueType = expr.(*ast.BasicLit).Value
	default:
		fmt.Printf("Unknown: %v, %T\n", expr, expr)
	}
	return &ft
}

func isBasicType(t string) bool {
	switch t {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "byte", "float32", "float64", "string", "bool", "error":
		return true
	}
	return false
}
