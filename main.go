package main

import (
	// "encoding/xml"
	// "fmt"
	"log"
	// "os"
	"strconv"

	"github.com/alekslesik/otmorozova/permissionControllerCommon"
	"github.com/alekslesik/otmorozova/permissionVariablesSelectCyclogram"
	"gopkg.in/ini.v1"
)

func main() {
	//  order in ListCyclogram.ini
	ListCyclogram, err := ini.Load("ListCyclogram.ini")
	if err != nil {
		log.Fatal(err)
	}

	for i, section := range ListCyclogram.Sections() {
		if section.Name() == "DEFAULT" {
			continue
		}

		// oldName := section.GetName()
		section.SetName("Cyclo" + strconv.Itoa(i))

		key, err := section.GetKey("ID")
		if err != nil {
			log.Fatal(err)
		}
		key.SetValue(strconv.Itoa(i))
	}

	ListCyclogram.SaveTo("ListCyclogram_result.ini")

	// order in PermissionControllerCommon.xml
	permissionControllerCommon := permissionControllerCommon.New()
	permissionControllerCommon.Read()
	permissionControllerCommon.Unmarshal()
	permissionControllerCommon.UnmarshalData.Order()
	permissionControllerCommon.MarshalIndent("", "    ")
	permissionControllerCommon.Write()

	permissionVariablesSelectCyclogram := permissionVariablesSelectCyclogram.New()
	permissionVariablesSelectCyclogram.Read()
	permissionVariablesSelectCyclogram.Unmarshal()
	permissionVariablesSelectCyclogram.UnmarshalData.OrderId()
	permissionVariablesSelectCyclogram.MarshalIndent("", "    ")
	permissionVariablesSelectCyclogram.Write()

}
