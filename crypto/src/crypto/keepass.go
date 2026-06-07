package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"strings"
)

// <?xml version="1.0" encoding="UTF-8"?>
// <KeyFile>
//     <Meta>
//         <Version>2.0</Version>
//     </Meta>
//     <Key>
//         <Data Hash="{HexEncode(sha256(sha256_bytes))[:8]}">{HexEncode(sha256_bytes)}</Data>
//     </Key>
// </KeyFile>

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
func HashHeader(byts []byte) string {
	hashHeader := sha256.Sum256(byts)
	r := strings.ToUpper(hex.EncodeToString(hashHeader[:]))
	return r[:8]
}

func FormatKeepassKeyFile(src []byte) ([]byte, error) {
	if len(src) != sha256.Size {
		return nil, fmt.Errorf("src len error: %d", len(src))
	}

	// generate XML key file
	return XMLMarshal(&KeyFile{
		Meta: Meta{
			Version: "2.0",
		},
		Key: Key{
			Data: Data{
				// hex(sha256(key)) 生成 <Data Hash="XXXXXXXX"> 取前8位, 8 hex-string, 用于检查 <Data> 数据的完整性.
				// VVI: 必须使用 sha256
				Hash: HashHeader(src),

				// 通过 hex(key) 生成 <Data>XXXX...XXXX</Data> hex-string
				Text: strings.ToUpper(hex.EncodeToString(src)),
			},
		},
	})
}

func GenCompositeKey(password, keyFileHash []byte) []byte {
	pwdHash := sha256.Sum256(password)

	h := sha256.New()
	h.Write(pwdHash[:])
	h.Write(keyFileHash)
	return h.Sum(nil)
}
