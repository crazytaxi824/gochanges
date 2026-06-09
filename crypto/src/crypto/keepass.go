package crypto

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
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
	// 不能使用 tag name `xml:"Value,chardata"`. invalid XML tag: cannot specify name together with option ",chardata"
	Value string `xml:",chardata"`
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

func FormatKeepassKeyFile(val []byte) ([]byte, error) {
	if len(val) != sha256.Size {
		return nil, fmt.Errorf("value len error: %d", len(val))
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
				Hash: HashHeader(val),

				// 通过 hex(key) 生成 <Data>XXXX...XXXX</Data> hex-string
				Value: strings.ToUpper(hex.EncodeToString(val)),
			},
		},
	})
}

func expandHome(path string) (string, error) {
	if !strings.HasPrefix(path, "~/") {
		return path, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, path[2:]), nil
}

// 如果 keyfile 是一般 file, 直接 hash 整个 file
// 如果 keyfile 是 xml file, 则读取 & 验证 <Key>
func ParseKeyFile(fpath string) ([]byte, error) {
	keyfp, err := expandHome(fpath)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(keyfp)
	if err != nil {
		return nil, err
	}

	var kf KeyFile
	if err := xml.Unmarshal(content, &kf); err != nil {
		// 不是合法 XML，直接 SHA256 整个文件
		h := sha256.Sum256(content)
		return h[:], nil
	}

	version := strings.TrimSpace(kf.Meta.Version)
	dataStr := strings.TrimSpace(kf.Key.Data.Value)

	if version == "" || dataStr == "" {
		// 是 XML 但不是 KeePass KeyFile 格式
		h := sha256.Sum256(content)
		return h[:], nil
	}

	switch version {
	case "2.0":
		hexStr := strings.ReplaceAll(dataStr, " ", "")
		data, err := hex.DecodeString(hexStr)
		if err != nil {
			return nil, fmt.Errorf("invalid key data: %w", err)
		}
		// 校验 Hash
		sum := sha256.Sum256(data)
		expected := strings.ToUpper(fmt.Sprintf("%X", sum[:4]))
		if strings.ToUpper(kf.Key.Data.Hash) != expected {
			return nil, fmt.Errorf("key file hash mismatch")
		}
		return data, nil

	case "1.0":
		data, err := base64.StdEncoding.DecodeString(dataStr)
		if err != nil {
			return nil, fmt.Errorf("invalid key data: %w", err)
		}
		return data, nil

	default:
		// 不认识的版本，fallback
		h := sha256.Sum256(content)
		return h[:], nil
	}
}

func GenCompositeKey(password, keyFileHash []byte) []byte {
	pwdHash := sha256.Sum256(password)

	h := sha256.New()
	h.Write(pwdHash[:])
	h.Write(keyFileHash)
	return h.Sum(nil)
}
