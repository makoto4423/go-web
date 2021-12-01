package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	_ "strings"
)

func main() {
	clearData()
	// fileName := os.Args[1]
	// bytes, err := os.ReadFile("./" + fileName + ".json")
	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }
	// base := make([]BaseJson, 0)
	// err = json.Unmarshal(bytes, &base)
	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }
	// res := recur(base, "")
	// fmt.Print(res)
}

func clearData() {
	fileName := os.Args[1]
	bytes, _ := os.ReadFile("./" + fileName + ".json")
	var m []map[string]interface{}
	json.Unmarshal(bytes, &m)
	reArr(m)
	bytes, _ = json.Marshal(m)
	fmt.Println(string(bytes))
}

func reArr(a []map[string]interface{}) {
	for _, m := range a {
		re(m)
	}
}

func re(m map[string]interface{}) {
	for k, v := range m {
		if v == nil {
			delete(m, k)
		} else if reflect.ValueOf(v).Kind() == reflect.Map {
			re(v.(map[string]interface{}))
		} else if reflect.ValueOf(v).Kind() == reflect.Slice {
			if reflect.ValueOf(v).Len() == 0 {
				delete(m, k)
			} else {
				// fmt.Println(reflect.ValueOf(v.([]interface{})[0]).Kind())
				// fmt.Println(reflect.TypeOf(v))
				// fmt.Println(reflect.TypeOf(v.([]interface{})[0]))
				for i := 0; i < reflect.ValueOf(v).Len(); i++ {
					// fmt.Println(reflect.TypeOf(v.([]interface{})[i]))
					re(v.([]interface{})[i].(map[string]interface{}))
				}
			}
		} else if reflect.ValueOf(v).Kind() == reflect.String && len(v.(string)) == 0 {
			delete(m, k)
		}
		// if v != nil {
		// 	fmt.Print(reflect.ValueOf(v).Kind())
		// }
	}

}

func recur(base []BaseJson, res string) string {
	for _, obj := range base {
		if len(obj.Choice) > 0 {
			res += "<xs:choice xml:relativeOpt=\"" + obj.Choice + "\">"
		}
		if len(obj.Name) > 0 {
			res += "<xs:element name=\"" + obj.Name + "\""
		} else {
			res += "<xs:element "
		}

		if len(obj.Type) > 0 {
			res += " type=\"xs:" + obj.Type + "\""
		}
		if len(obj.Default) > 0 {
			res += " default=\"" + obj.Default + "\""
		}

		if len(obj.Desc) > 0 {
			res += " xml:desc=\"" + obj.Desc + "\""
		}

		res += ">"

		if len(obj.Document) > 0 {
			res += "<xs:annotation><xs:documentation>" + obj.Document + "</xs:documentation></xs:annotation>"
		}

		if len(obj.Enum) > 0 {
			res += "<xs:simpleType><xs:restriction base=\"xs:" + obj.EnumType + "\">"
			for _, enum := range obj.Enum {
				res += "<xs:enumeration value=\"" + enum.Value + "\""
				if len(enum.Option) > 0 {
					res += " xml:options=\"" + enum.Option + "\""
				}
				res += "><xs:annotation><xs:documentation>" + enum.Document
				res += "</xs:documentation></xs:annotation></xs:enumeration>"
			}
			res += "</xs:restriction></xs:simpleType>"
		}

		if len(obj.Children) > 0 {
			res += "<xs:complexType><xs:sequence>"
			res = recur(obj.Children, res)
			res += "</xs:sequence></xs:complexType>"
		}

		res += "</xs:element>"
		if len(obj.Choice) > 0 {
			res += "</xs:choice>"
		}
	}
	return res
}

type BaseJson struct {
	Name     string     `json:"name"`
	Type     string     `json:"type"`
	Document string     `json:"document"`
	Children []BaseJson `json:"children"`
	Enum     []Enum     `json:"enum"`
	EnumType string     `json:"enumType"`
	Default  string     `json:"default"`
	Desc     string     `json:"desc"`
	Choice   string     `json:"choice"`
}

type Enum struct {
	Document string `json:"document"`
	Value    string `json:"value"`
	Option   string `json:"option"`
}
