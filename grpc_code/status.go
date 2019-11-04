package grpc_code

import (
	"errors"
	"fmt"
	"github.com/gw123/ghelper/reg_util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

const (
	DeviceNotFound   codes.Code = 1000
	GroupNotFound    codes.Code = 1001
	ChapterNotFound  codes.Code = 1002
	ResourceNotFound codes.Code = 1003
)

var Code2StrMap map[codes.Code]string = map[codes.Code]string{
	DeviceNotFound:   "未找到设备",
	GroupNotFound:    "未找到资源组",
	ChapterNotFound:  "未找章节",
	ResourceNotFound: "未找到资源",
}

var GrpcCodeStr2code = map[string]codes.Code{
	"Canceled":           codes.Canceled,
	"Unknown":            codes.Unknown,
	"InvalidArgument":    codes.InvalidArgument,
	"DeadlineExceeded":   codes.DeadlineExceeded,
	"NotFound":           codes.NotFound,
	"AlreadyExists":      codes.AlreadyExists,
	"PermissionDenied":   codes.PermissionDenied,
	"ResourceExhausted":  codes.ResourceExhausted,
	"FailedPrecondition": codes.FailedPrecondition,
	"Aborted":            codes.Aborted,
	"OutOfRange":         codes.OutOfRange,
	"Unimplemented":      codes.Unimplemented,
	"Internal":           codes.Internal,
	"Unavailable":        codes.Unavailable,
	"DataLoss":           codes.DataLoss,
	"Unauthenticated":    codes.Unauthenticated,
}

type Status struct {
	code    codes.Code
	message string
}

func NewStatus(code codes.Code, message string) *Status {
	if message == "" {
		if msg, ok := Code2StrMap[code]; ok {
			message = msg
		}
	}
	return &Status{
		code:    code,
		message: message,
	}
}

func (s *Status) Error() string {
	return status.Error(s.code, s.message).Error()
}

func (s *Status) GrpcError() error {
	return status.Error(s.code, s.message)
}

func (s *Status) WithMessage(msg string) {
	s.message = msg
	return
}

func (s *Status) GetMessage() string {
	return s.message
}

//
func (s *Status) String() string {

	e := status.Error(s.code, s.message)
	if e == nil {
		return "success"
	}
	return e.Error()
}

//
func (s *Status) Marshal() string {
	return s.String()
}

//
func (s *Status) UnMarshal(input string) error {
	if input == "null" {
		return nil
	}
	if s == nil {
		return fmt.Errorf("nil receiver passed to Unmarshal")
	}
	codeStr, desc, ok := reg_util.ParseGrpcError(input)
	if !ok {
		return errors.New("input format error")
	}

	if grpcCode, ok := GrpcCodeStr2code[codeStr]; ok {
		s.code = grpcCode
	} else {
		code, _ := strconv.Atoi(codeStr)
		s.code = codes.Code(code)
	}
	s.message = desc
	return nil
}

func (c *Status) MarshalJSON() ([]byte, error) {
	ret := "\"" + c.String() + "\""
	return []byte(ret), nil
}

func (c *Status) UnmarshalJSON(b []byte) error {
	return c.UnMarshal(string(b))
}

func PackGrpcError(code codes.Code) error {
	return NewStatus(code, "").GrpcError()
}