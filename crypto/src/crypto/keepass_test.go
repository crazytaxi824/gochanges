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
	XMLName xml.Name `xml:"KeyFile"` // XML 最外层 tag name
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
	// attribute
	Hash string `xml:"Hash,attr"`

	// ",chardata" (character data) 表示将该字段映射到 <Data></Data> 内部.
	// 不能使用 tag name `xml:"Text,chardata"`. invalid XML tag: cannot specify name together with option ",chardata"
	Text string `xml:",chardata"`
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

// HashHeader 检查 <Data></Data> 数据完整性, VVI: 必须使用 Sha256(Sha2)
func HashHeader(key []byte) string {
	h := sha256.New()
	h.Write(key)
	hashHeader := h.Sum(nil)
	r := strings.ToUpper(hex.EncodeToString(hashHeader))
	return r[:8]
}

func TestHash(t *testing.T) {
	// 生成 256bit (32byte) key 数据
	h := sha256.New()
	h.Write([]byte("123"))
	key := h.Sum(nil) // VVI: key 可以是任何数据

	// generate XML key file
	ss, err := XMLMarshal(&KeyFile{
		Meta: Meta{
			Version: "2.0",
		},
		Key: Key{
			Data: Data{
				// hex(sha256(key)) 生成 <Data Hash="XXXXXXXX"> 取前8位, 8 hex-string, 用于检查 <Data> 数据的完整性.
				// VVI: 必须使用 sha256
				Hash: HashHeader(key),

				// 通过 hex(key) 生成 <Data>XXXX...XXXX</Data> hex-string
				Text: strings.ToUpper(hex.EncodeToString(key)),
			},
		},
	})
	if err != nil {
		t.Log(err)
		return
	}

	fmt.Println(string(ss))
}
