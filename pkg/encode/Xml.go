package encode

import (
	"encoding/xml"
	"io"
)

const (
	XMLDefultTitle = `<?xml version="1.0"?>`
)

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

type StringMap map[string]string

func (m StringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}
	for k, v := range m {
		err := e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *StringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = StringMap{}
	for {
		var e xmlMapEntry
		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}

// XMLEncode 只能做一維 然後在自己組
//
// ex.
// map["a"] = "b"  =====>  <a>b</a>
func XMLEncode(m map[string]string) string {
	buf, _ := xml.Marshal(StringMap(m))
	return string(buf)
}

// XMLDecode 只能做一維string map
func XMLDecode(xmlString string) (map[string]string, error) {
	m := map[string]string{}
	err := xml.Unmarshal([]byte(xmlString), (*StringMap)(&m))
	if err != nil {
		return nil, err
	}
	return m, nil
}
