// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: demo.proto

package demo

import strings "strings"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/infobloxopen/protoc-gen-preprocess/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (m *Demo) Preprocess() error {
	m.S = strings.TrimSpace(m.S)
	return nil
}

func (m *DemoRes) Preprocess() error {
	return nil
}
