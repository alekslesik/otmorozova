package permissionControllerCommon

import (
	"encoding/xml"
	"log"
	"os"
	"strconv"
)

type Variable struct {
	Name    string `xml:"name,attr"`
	Content string `xml:",chardata"`
}

type PermissionVariables struct {
	XMLName  xml.Name `xml:"PermissionVariables"`
	Variable []Variable
}

type PermissionControllerCommon struct {
	rowData []byte
	UnmarshalData PermissionVariables
	MarsalData []byte
}

func New() *PermissionControllerCommon {
	return new(PermissionControllerCommon)
}

// Read
func (p *PermissionControllerCommon) Read() {
	rowData, err := os.ReadFile("PermissionControllerCommon.xml")
	if err != nil {
		log.Fatal(err)
		return
	}

	p.rowData = rowData
}

// Unmarhal
func (p *PermissionControllerCommon) Unmarshal() {
	err := xml.Unmarshal(p.rowData, &p.UnmarshalData)
	if err != nil {
		log.Fatal(err)
		return
	}
}

// Marshal
func (p *PermissionControllerCommon) MarshalIndent(prefix, indent string) {
	result, err := xml.MarshalIndent(p.UnmarshalData, prefix, indent)
	if err != nil {
		log.Fatal(err)
		return
	}

	p.MarsalData = result
}



// Write
func (p *PermissionControllerCommon) Write() {
	err := os.WriteFile("PermissionControllerCommon_result.xml", p.MarsalData, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

// Ordering
func (p *PermissionVariables) Order() {
	for i := 0; i < len(p.Variable); i++ {
		p.Variable[i].SetName("bCyclo" + strconv.Itoa(i + 1))
	}
}

func (v *Variable) SetName(name string) {
	v.Name = name
}