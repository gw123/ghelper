package grpc_code

import (
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestNewStatus(t *testing.T) {
	tests := []codes.Code{
		codes.Canceled,
		codes.AlreadyExists,
		codes.PermissionDenied,
		DeviceNotFound,
		GroupNotFound,
	}

	for _, testItem := range tests {
		status1 := NewStatus(testItem, "")
		tempStatus := Status{}
		err := tempStatus.UnMarshal(status1.Error())
		if err != nil {
			t.Fatal(err)
		}

		if status1.String() != tempStatus.String() {
			t.Fail()
		}
	}
}

func TestStatus_MarshalAndUnmarshal(t *testing.T) {
	tests := []codes.Code{
		codes.Canceled,
		codes.AlreadyExists,
		codes.PermissionDenied,
		DeviceNotFound,
		GroupNotFound,
	}

	for _, testItem := range tests {
		status1 := NewStatus(testItem, "")
		tempStatus := Status{}
		err := tempStatus.UnMarshal(status1.Marshal())
		if err != nil {
			t.Fatal(err)
		}

		if status1.String() != tempStatus.String() {
			t.Fail()
		}
	}
}

func TestStatus_WithMessage(t *testing.T) {
	tests := map[string]codes.Code{
		"用户取消":       codes.Canceled,
		"已经存在":       codes.AlreadyExists,
		"权限不够":       codes.PermissionDenied,
		"设备notfound": DeviceNotFound,
	}

	for msg, testItem := range tests {
		status1 := NewStatus(testItem, "")
		status1.WithMessage(msg)
		tempStatus := Status{}
		err := tempStatus.UnMarshal(status1.Marshal())
		if err != nil {
			t.Fatal(err)
		}

		if status1.String() != tempStatus.String() {
			t.Fail()
		}
		t.Log(tempStatus.GetMessage(), tempStatus.String())
		if msg != tempStatus.GetMessage() {
			t.Fail()
		}
	}
}

func TestStatus_MarshalJSON(t *testing.T) {
	tests := map[string]codes.Code{
		"用户取消":       codes.Canceled,
		"已经存在":       codes.AlreadyExists,
		"权限不够":       codes.PermissionDenied,
		"设备notfound": DeviceNotFound,
	}

	for msg, testItem := range tests {
		status1 := NewStatus(testItem, msg)
		tempStatus := Status{}

		data, err := json.Marshal(status1)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(data))
		err = json.Unmarshal(data, &tempStatus)
		if err != nil {
			t.Fatal(err)
		}

		if status1.String() != tempStatus.String() {
			t.Fail()
		}
		t.Log(tempStatus.GetMessage(), " == ", tempStatus.String())
		if msg != tempStatus.GetMessage() {
			t.Fail()
		}
	}
}

func TestPackGrpcError(t *testing.T) {
	tests := map[string]codes.Code{
		"用户取消":       codes.Canceled,
		"已经存在":       codes.AlreadyExists,
		"权限不够":       codes.PermissionDenied,
	}

	for _, testItem := range tests {
		err := PackGrpcError(testItem)
		err2 := status.Error(testItem, "")

		if err.Error() != err2.Error() {
			t.Log(err, err2)
			t.Fatal()
		}
	}
}
