package permissionVariablesSelectCyclogram

import (
	"encoding/xml"
	"log"
	"os"
	"strconv"
)

type Variable struct {
	Name    string `xml:"id,attr"`
	Content string `xml:",chardata"`
}

type PermissionVariables struct {
	XMLName  xml.Name `xml:"PermissionVariables"`
	Variable []Variable
}

type permissionVariablesSelectCyclogram struct {
	rowData []byte
	UnmarshalData PermissionVariables
	MarsalData []byte
}

func New() *permissionVariablesSelectCyclogram {
	return new(permissionVariablesSelectCyclogram)
}

// Read
func (p *permissionVariablesSelectCyclogram) Read() {
	rowData, err := os.ReadFile("PermissionControllerCommon.xml")
	if err != nil {
		log.Fatal(err)
		return
	}

	p.rowData = rowData
}

// Unmarhal
func (p *permissionVariablesSelectCyclogram) Unmarshal() {
	err := xml.Unmarshal(p.rowData, &p.UnmarshalData)
	if err != nil {
		log.Fatal(err)
		return
	}
}

// Marshal
func (p *permissionVariablesSelectCyclogram) MarshalIndent(prefix, indent string) {
	result, err := xml.MarshalIndent(p.UnmarshalData, prefix, indent)
	if err != nil {
		log.Fatal(err)
		return
	}

	p.MarsalData = result
}

// Write
func (p *permissionVariablesSelectCyclogram) Write() {
	err := os.WriteFile("permissionVariablesSelectCyclogram_result.xml", p.MarsalData, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

// Ordering
func (p *PermissionVariables) OrderId() {
	for i := 0; i < len(p.Variable); i++ {
		p.Variable[i].SetName(strconv.Itoa(i + 1))
	}
}

func (v *Variable) SetName(name string) {
	v.Name = name
}