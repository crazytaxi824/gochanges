// key file format <XML>:
// <Data Hash="XXXXXXXX">  必须是 <Data> 的 sha256 值前 8 位, VVI: 必须要 quote, 否则变成普通类型的文件.
// <Data>XXXX ... XXXX</Data> 可以被 hex.DecodeString(), 必须位偶数.
//
// <?xml version="1.0" encoding="UTF-8"?>
// <KeyFile>
//     <Meta>
//         <Version>2.0</Version>
//     </Meta>
//     <Key>
//         <Data Hash="{hex}">{key}</Data>
//     </Key>
// </KeyFile>

package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"strings"
	"testing"
)

type KeyFile struct {
	XMLName xml.Name `xml:"KeyFile"`
	Meta    Meta     `xml:"Meta"`
	Key     Key      `xml:"Key"`
}

type Meta struct {
	Version string `xml:"Version"`
}

type Key struct {
	Data Data `xml:"Data"`
}

type Data struct {
	Hash string `xml:"Hash,attr"` // attribute
	Text string `xml:",chardata"` // ",chardata" 表示将该字段(Text)映射到 XML 元素的内部文本内容（character data）中
}

// generate XML file
func XMLMarshal(kf *KeyFile) ([]byte, error) {
	b, err := xml.MarshalIndent(kf, "", "\t")
	if err != nil {
		return nil, err
	}

	xh := []byte(xml.Header) // `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
	xh = append(xh, b...)

	return xh, nil
}

func HashHeader(key []byte) string {
	h := sha256.New()
	h.Write(key)
	hashHeader := h.Sum(nil)
	r := strings.ToUpper(hex.EncodeToString(hashHeader))
	return r[:8]
}

func TestHash(t *testing.T) {
	// 生成 256bit (32byte) -> 64 hex-string 的 Data 数据.
	h := sha256.New()
	h.Write([]byte("123"))
	key := h.Sum(nil)

	// 生成 <Data> XXXX...XXXX </Data> (64 hex-string 256-bit)
	dataHex := strings.ToUpper(hex.EncodeToString(key))
	t.Log(dataHex)
	// t.Log(len(s))

	// 生成 <Data Hash="XXXXXXXX"> (8 hex string)
	hashHeader := HashHeader(key)

	// generate XML key file
	ss, err := XMLMarshal(&KeyFile{
		Meta: Meta{
			Version: "2.0",
		},
		Key: Key{
			Data: Data{
				Hash: hashHeader,
				Text: dataHex,
			},
		},
	})
	if err != nil {
		t.Log(err)
		return
	}

	fmt.Println(string(ss))
}
