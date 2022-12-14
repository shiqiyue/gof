// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package storage

import (
	"fmt"
	"strings"
)

const (
	// GormStorageMsgStatusNOTSEND is a gormStorageMsgStatus of type NOT_SEND.
	// 未发送
	GormStorageMsgStatusNOTSEND gormStorageMsgStatus = iota + 1
	// GormStorageMsgStatusSENDSUCCESS is a gormStorageMsgStatus of type SEND_SUCCESS.
	// 发送成功
	GormStorageMsgStatusSENDSUCCESS
	// GormStorageMsgStatusSENDFAIL is a gormStorageMsgStatus of type SEND_FAIL.
	// 发送失败
	GormStorageMsgStatusSENDFAIL
	// GormStorageMsgStatusCONSUMESUCCESS is a gormStorageMsgStatus of type CONSUME_SUCCESS.
	// 消费成功
	GormStorageMsgStatusCONSUMESUCCESS
	// GormStorageMsgStatusCONSUMEFAIL is a gormStorageMsgStatus of type CONSUME_FAIL.
	// 消费失败
	GormStorageMsgStatusCONSUMEFAIL
)

const _gormStorageMsgStatusName = "NOT_SENDSEND_SUCCESSSEND_FAILCONSUME_SUCCESSCONSUME_FAIL"

var _gormStorageMsgStatusNames = []string{
	_gormStorageMsgStatusName[0:8],
	_gormStorageMsgStatusName[8:20],
	_gormStorageMsgStatusName[20:29],
	_gormStorageMsgStatusName[29:44],
	_gormStorageMsgStatusName[44:56],
}

// gormStorageMsgStatusNames returns a list of possible string values of gormStorageMsgStatus.
func gormStorageMsgStatusNames() []string {
	tmp := make([]string, len(_gormStorageMsgStatusNames))
	copy(tmp, _gormStorageMsgStatusNames)
	return tmp
}

var _gormStorageMsgStatusMap = map[gormStorageMsgStatus]string{
	GormStorageMsgStatusNOTSEND:        _gormStorageMsgStatusName[0:8],
	GormStorageMsgStatusSENDSUCCESS:    _gormStorageMsgStatusName[8:20],
	GormStorageMsgStatusSENDFAIL:       _gormStorageMsgStatusName[20:29],
	GormStorageMsgStatusCONSUMESUCCESS: _gormStorageMsgStatusName[29:44],
	GormStorageMsgStatusCONSUMEFAIL:    _gormStorageMsgStatusName[44:56],
}

// String implements the Stringer interface.
func (x gormStorageMsgStatus) String() string {
	if str, ok := _gormStorageMsgStatusMap[x]; ok {
		return str
	}
	return fmt.Sprintf("gormStorageMsgStatus(%d)", x)
}

var _gormStorageMsgStatusValue = map[string]gormStorageMsgStatus{
	_gormStorageMsgStatusName[0:8]:   GormStorageMsgStatusNOTSEND,
	_gormStorageMsgStatusName[8:20]:  GormStorageMsgStatusSENDSUCCESS,
	_gormStorageMsgStatusName[20:29]: GormStorageMsgStatusSENDFAIL,
	_gormStorageMsgStatusName[29:44]: GormStorageMsgStatusCONSUMESUCCESS,
	_gormStorageMsgStatusName[44:56]: GormStorageMsgStatusCONSUMEFAIL,
}

// ParsegormStorageMsgStatus attempts to convert a string to a gormStorageMsgStatus.
func ParsegormStorageMsgStatus(name string) (gormStorageMsgStatus, error) {
	if x, ok := _gormStorageMsgStatusValue[name]; ok {
		return x, nil
	}
	return gormStorageMsgStatus(0), fmt.Errorf("%s is not a valid gormStorageMsgStatus, try [%s]", name, strings.Join(_gormStorageMsgStatusNames, ", "))
}

func (x gormStorageMsgStatus) Ptr() *gormStorageMsgStatus {
	return &x
}

// MarshalText implements the text marshaller method.
func (x gormStorageMsgStatus) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *gormStorageMsgStatus) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParsegormStorageMsgStatus(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
