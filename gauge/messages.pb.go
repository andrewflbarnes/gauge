// Code generated by protoc-gen-go.
// source: messages.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	KillProcessRequest
	ExecutionStatusResponse
	ExecutionStartingRequest
	SpecExecutionStartingRequest
	SpecExecutionEndingRequest
	ScenarioExecutionStartingRequest
	ScenarioExecutionEndingRequest
	StepExecutionStartingRequest
	StepExecutionEndingRequest
	ExecutionInfo
	SpecInfo
	ScenarioInfo
	StepInfo
	ExecuteStepRequest
	StepValidateRequest
	StepValidateResponse
	ExecutionEndingRequest
	SuiteExecutionResult
	StepNamesRequest
	StepNamesResponse
	Message
*/
package main

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Message_MessageType int32

const (
	Message_ExecutionStarting         Message_MessageType = 0
	Message_SpecExecutionStarting     Message_MessageType = 1
	Message_SpecExecutionEnding       Message_MessageType = 2
	Message_ScenarioExecutionStarting Message_MessageType = 3
	Message_ScenarioExecutionEnding   Message_MessageType = 4
	Message_StepExecutionStarting     Message_MessageType = 5
	Message_StepExecutionEnding       Message_MessageType = 6
	Message_ExecuteStep               Message_MessageType = 7
	Message_ExecutionEnding           Message_MessageType = 8
	Message_StepValidateRequest       Message_MessageType = 9
	Message_StepValidateResponse      Message_MessageType = 10
	Message_ExecutionStatusResponse   Message_MessageType = 11
	Message_StepNamesRequest          Message_MessageType = 12
	Message_StepNamesResponse         Message_MessageType = 13
	Message_KillProcessRequest        Message_MessageType = 14
	Message_SuiteExecutionResult      Message_MessageType = 15
)

var Message_MessageType_name = map[int32]string{
	0:  "ExecutionStarting",
	1:  "SpecExecutionStarting",
	2:  "SpecExecutionEnding",
	3:  "ScenarioExecutionStarting",
	4:  "ScenarioExecutionEnding",
	5:  "StepExecutionStarting",
	6:  "StepExecutionEnding",
	7:  "ExecuteStep",
	8:  "ExecutionEnding",
	9:  "StepValidateRequest",
	10: "StepValidateResponse",
	11: "ExecutionStatusResponse",
	12: "StepNamesRequest",
	13: "StepNamesResponse",
	14: "KillProcessRequest",
	15: "SuiteExecutionResult",
}
var Message_MessageType_value = map[string]int32{
	"ExecutionStarting":         0,
	"SpecExecutionStarting":     1,
	"SpecExecutionEnding":       2,
	"ScenarioExecutionStarting": 3,
	"ScenarioExecutionEnding":   4,
	"StepExecutionStarting":     5,
	"StepExecutionEnding":       6,
	"ExecuteStep":               7,
	"ExecutionEnding":           8,
	"StepValidateRequest":       9,
	"StepValidateResponse":      10,
	"ExecutionStatusResponse":   11,
	"StepNamesRequest":          12,
	"StepNamesResponse":         13,
	"KillProcessRequest":        14,
	"SuiteExecutionResult":      15,
}

func (x Message_MessageType) Enum() *Message_MessageType {
	p := new(Message_MessageType)
	*p = x
	return p
}
func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (x *Message_MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Message_MessageType_value, data, "Message_MessageType")
	if err != nil {
		return err
	}
	*x = Message_MessageType(value)
	return nil
}

type KillProcessRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *KillProcessRequest) Reset()         { *m = KillProcessRequest{} }
func (m *KillProcessRequest) String() string { return proto.CompactTextString(m) }
func (*KillProcessRequest) ProtoMessage()    {}

// Sends to any request which needs a execution status as response
// usually step execution, hooks etc will return this
type ExecutionStatusResponse struct {
	ExecutionResult  *ProtoExecutionResult `protobuf:"bytes,1,req,name=executionResult" json:"executionResult,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *ExecutionStatusResponse) Reset()         { *m = ExecutionStatusResponse{} }
func (m *ExecutionStatusResponse) String() string { return proto.CompactTextString(m) }
func (*ExecutionStatusResponse) ProtoMessage()    {}

func (m *ExecutionStatusResponse) GetExecutionResult() *ProtoExecutionResult {
	if m != nil {
		return m.ExecutionResult
	}
	return nil
}

type ExecutionStartingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *ExecutionStartingRequest) Reset()         { *m = ExecutionStartingRequest{} }
func (m *ExecutionStartingRequest) String() string { return proto.CompactTextString(m) }
func (*ExecutionStartingRequest) ProtoMessage()    {}

func (m *ExecutionStartingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

type SpecExecutionStartingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *SpecExecutionStartingRequest) Reset()         { *m = SpecExecutionStartingRequest{} }
func (m *SpecExecutionStartingRequest) String() string { return proto.CompactTextString(m) }
func (*SpecExecutionStartingRequest) ProtoMessage()    {}

func (m *SpecExecutionStartingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

type SpecExecutionEndingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *SpecExecutionEndingRequest) Reset()         { *m = SpecExecutionEndingRequest{} }
func (m *SpecExecutionEndingRequest) String() string { return proto.CompactTextString(m) }
func (*SpecExecutionEndingRequest) ProtoMessage()    {}

func (m *SpecExecutionEndingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

type ScenarioExecutionStartingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *ScenarioExecutionStartingRequest) Reset()         { *m = ScenarioExecutionStartingRequest{} }
func (m *ScenarioExecutionStartingRequest) String() string { return proto.CompactTextString(m) }
func (*ScenarioExecutionStartingRequest) ProtoMessage()    {}

func (m *ScenarioExecutionStartingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

type ScenarioExecutionEndingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *ScenarioExecutionEndingRequest) Reset()         { *m = ScenarioExecutionEndingRequest{} }
func (m *ScenarioExecutionEndingRequest) String() string { return proto.CompactTextString(m) }
func (*ScenarioExecutionEndingRequest) ProtoMessage()    {}

func (m *ScenarioExecutionEndingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

type StepExecutionStartingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *StepExecutionStartingRequest) Reset()         { *m = StepExecutionStartingRequest{} }
func (m *StepExecutionStartingRequest) String() string { return proto.CompactTextString(m) }
func (*StepExecutionStartingRequest) ProtoMessage()    {}

func (m *StepExecutionStartingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

type StepExecutionEndingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *StepExecutionEndingRequest) Reset()         { *m = StepExecutionEndingRequest{} }
func (m *StepExecutionEndingRequest) String() string { return proto.CompactTextString(m) }
func (*StepExecutionEndingRequest) ProtoMessage()    {}

func (m *StepExecutionEndingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

type ExecutionInfo struct {
	CurrentSpec      *SpecInfo     `protobuf:"bytes,1,opt,name=currentSpec" json:"currentSpec,omitempty"`
	CurrentScenario  *ScenarioInfo `protobuf:"bytes,2,opt,name=currentScenario" json:"currentScenario,omitempty"`
	CurrentStep      *StepInfo     `protobuf:"bytes,3,opt,name=currentStep" json:"currentStep,omitempty"`
	Stacktrace       *string       `protobuf:"bytes,4,opt,name=stacktrace" json:"stacktrace,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *ExecutionInfo) Reset()         { *m = ExecutionInfo{} }
func (m *ExecutionInfo) String() string { return proto.CompactTextString(m) }
func (*ExecutionInfo) ProtoMessage()    {}

func (m *ExecutionInfo) GetCurrentSpec() *SpecInfo {
	if m != nil {
		return m.CurrentSpec
	}
	return nil
}

func (m *ExecutionInfo) GetCurrentScenario() *ScenarioInfo {
	if m != nil {
		return m.CurrentScenario
	}
	return nil
}

func (m *ExecutionInfo) GetCurrentStep() *StepInfo {
	if m != nil {
		return m.CurrentStep
	}
	return nil
}

func (m *ExecutionInfo) GetStacktrace() string {
	if m != nil && m.Stacktrace != nil {
		return *m.Stacktrace
	}
	return ""
}

type SpecInfo struct {
	Name             *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	FileName         *string  `protobuf:"bytes,2,req,name=fileName" json:"fileName,omitempty"`
	IsFailed         *bool    `protobuf:"varint,3,req,name=isFailed" json:"isFailed,omitempty"`
	Tags             []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SpecInfo) Reset()         { *m = SpecInfo{} }
func (m *SpecInfo) String() string { return proto.CompactTextString(m) }
func (*SpecInfo) ProtoMessage()    {}

func (m *SpecInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SpecInfo) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *SpecInfo) GetIsFailed() bool {
	if m != nil && m.IsFailed != nil {
		return *m.IsFailed
	}
	return false
}

func (m *SpecInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ScenarioInfo struct {
	Name             *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	IsFailed         *bool    `protobuf:"varint,2,req,name=isFailed" json:"isFailed,omitempty"`
	Tags             []string `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ScenarioInfo) Reset()         { *m = ScenarioInfo{} }
func (m *ScenarioInfo) String() string { return proto.CompactTextString(m) }
func (*ScenarioInfo) ProtoMessage()    {}

func (m *ScenarioInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ScenarioInfo) GetIsFailed() bool {
	if m != nil && m.IsFailed != nil {
		return *m.IsFailed
	}
	return false
}

func (m *ScenarioInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type StepInfo struct {
	Step             *ExecuteStepRequest `protobuf:"bytes,1,req,name=step" json:"step,omitempty"`
	IsFailed         *bool               `protobuf:"varint,2,req,name=isFailed" json:"isFailed,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *StepInfo) Reset()         { *m = StepInfo{} }
func (m *StepInfo) String() string { return proto.CompactTextString(m) }
func (*StepInfo) ProtoMessage()    {}

func (m *StepInfo) GetStep() *ExecuteStepRequest {
	if m != nil {
		return m.Step
	}
	return nil
}

func (m *StepInfo) GetIsFailed() bool {
	if m != nil && m.IsFailed != nil {
		return *m.IsFailed
	}
	return false
}

type ExecuteStepRequest struct {
	ActualStepText   *string     `protobuf:"bytes,1,req,name=actualStepText" json:"actualStepText,omitempty"`
	ParsedStepText   *string     `protobuf:"bytes,2,req,name=parsedStepText" json:"parsedStepText,omitempty"`
	ScenarioFailing  *bool       `protobuf:"varint,3,opt,name=scenarioFailing" json:"scenarioFailing,omitempty"`
	Args             []*Argument `protobuf:"bytes,4,rep,name=args" json:"args,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ExecuteStepRequest) Reset()         { *m = ExecuteStepRequest{} }
func (m *ExecuteStepRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteStepRequest) ProtoMessage()    {}

func (m *ExecuteStepRequest) GetActualStepText() string {
	if m != nil && m.ActualStepText != nil {
		return *m.ActualStepText
	}
	return ""
}

func (m *ExecuteStepRequest) GetParsedStepText() string {
	if m != nil && m.ParsedStepText != nil {
		return *m.ParsedStepText
	}
	return ""
}

func (m *ExecuteStepRequest) GetScenarioFailing() bool {
	if m != nil && m.ScenarioFailing != nil {
		return *m.ScenarioFailing
	}
	return false
}

func (m *ExecuteStepRequest) GetArgs() []*Argument {
	if m != nil {
		return m.Args
	}
	return nil
}

type StepValidateRequest struct {
	StepText          *string `protobuf:"bytes,1,req,name=stepText" json:"stepText,omitempty"`
	NumberOfArguments *int32  `protobuf:"varint,2,req,name=numberOfArguments" json:"numberOfArguments,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *StepValidateRequest) Reset()         { *m = StepValidateRequest{} }
func (m *StepValidateRequest) String() string { return proto.CompactTextString(m) }
func (*StepValidateRequest) ProtoMessage()    {}

func (m *StepValidateRequest) GetStepText() string {
	if m != nil && m.StepText != nil {
		return *m.StepText
	}
	return ""
}

func (m *StepValidateRequest) GetNumberOfArguments() int32 {
	if m != nil && m.NumberOfArguments != nil {
		return *m.NumberOfArguments
	}
	return 0
}

type StepValidateResponse struct {
	IsValid          *bool   `protobuf:"varint,1,req,name=isValid" json:"isValid,omitempty"`
	ErrorMessage     *string `protobuf:"bytes,2,opt,name=errorMessage" json:"errorMessage,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StepValidateResponse) Reset()         { *m = StepValidateResponse{} }
func (m *StepValidateResponse) String() string { return proto.CompactTextString(m) }
func (*StepValidateResponse) ProtoMessage()    {}

func (m *StepValidateResponse) GetIsValid() bool {
	if m != nil && m.IsValid != nil {
		return *m.IsValid
	}
	return false
}

func (m *StepValidateResponse) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

type ExecutionEndingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *ExecutionEndingRequest) Reset()         { *m = ExecutionEndingRequest{} }
func (m *ExecutionEndingRequest) String() string { return proto.CompactTextString(m) }
func (*ExecutionEndingRequest) ProtoMessage()    {}

func (m *ExecutionEndingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

type SuiteExecutionResult struct {
	SuiteResult      *ProtoSuiteResult `protobuf:"bytes,1,req,name=suiteResult" json:"suiteResult,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *SuiteExecutionResult) Reset()         { *m = SuiteExecutionResult{} }
func (m *SuiteExecutionResult) String() string { return proto.CompactTextString(m) }
func (*SuiteExecutionResult) ProtoMessage()    {}

func (m *SuiteExecutionResult) GetSuiteResult() *ProtoSuiteResult {
	if m != nil {
		return m.SuiteResult
	}
	return nil
}

type StepNamesRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *StepNamesRequest) Reset()         { *m = StepNamesRequest{} }
func (m *StepNamesRequest) String() string { return proto.CompactTextString(m) }
func (*StepNamesRequest) ProtoMessage()    {}

type StepNamesResponse struct {
	Steps            []string `protobuf:"bytes,1,rep,name=steps" json:"steps,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *StepNamesResponse) Reset()         { *m = StepNamesResponse{} }
func (m *StepNamesResponse) String() string { return proto.CompactTextString(m) }
func (*StepNamesResponse) ProtoMessage()    {}

func (m *StepNamesResponse) GetSteps() []string {
	if m != nil {
		return m.Steps
	}
	return nil
}

// This is the message which gets transferred all the time
// with proper message type set
type Message struct {
	MessageType *Message_MessageType `protobuf:"varint,1,req,name=messageType,enum=main.Message_MessageType" json:"messageType,omitempty"`
	// A unique id to represent this message. A response to the message should copy over this value
	// this is used to synchronize messages & responses
	MessageId *int64 `protobuf:"varint,2,req,name=messageId" json:"messageId,omitempty"`
	// One of the following will have a value
	ExecutionStartingRequest         *ExecutionStartingRequest         `protobuf:"bytes,3,opt,name=executionStartingRequest" json:"executionStartingRequest,omitempty"`
	SpecExecutionStartingRequest     *SpecExecutionStartingRequest     `protobuf:"bytes,4,opt,name=specExecutionStartingRequest" json:"specExecutionStartingRequest,omitempty"`
	SpecExecutionEndingRequest       *SpecExecutionEndingRequest       `protobuf:"bytes,5,opt,name=specExecutionEndingRequest" json:"specExecutionEndingRequest,omitempty"`
	ScenarioExecutionStartingRequest *ScenarioExecutionStartingRequest `protobuf:"bytes,6,opt,name=scenarioExecutionStartingRequest" json:"scenarioExecutionStartingRequest,omitempty"`
	ScenarioExecutionEndingRequest   *ScenarioExecutionEndingRequest   `protobuf:"bytes,7,opt,name=scenarioExecutionEndingRequest" json:"scenarioExecutionEndingRequest,omitempty"`
	StepExecutionStartingRequest     *StepExecutionStartingRequest     `protobuf:"bytes,8,opt,name=stepExecutionStartingRequest" json:"stepExecutionStartingRequest,omitempty"`
	StepExecutionEndingRequest       *StepExecutionEndingRequest       `protobuf:"bytes,9,opt,name=stepExecutionEndingRequest" json:"stepExecutionEndingRequest,omitempty"`
	ExecuteStepRequest               *ExecuteStepRequest               `protobuf:"bytes,10,opt,name=executeStepRequest" json:"executeStepRequest,omitempty"`
	ExecutionEndingRequest           *ExecutionEndingRequest           `protobuf:"bytes,11,opt,name=executionEndingRequest" json:"executionEndingRequest,omitempty"`
	StepValidateRequest              *StepValidateRequest              `protobuf:"bytes,12,opt,name=stepValidateRequest" json:"stepValidateRequest,omitempty"`
	StepValidateResponse             *StepValidateResponse             `protobuf:"bytes,13,opt,name=stepValidateResponse" json:"stepValidateResponse,omitempty"`
	ExecutionStatusResponse          *ExecutionStatusResponse          `protobuf:"bytes,14,opt,name=executionStatusResponse" json:"executionStatusResponse,omitempty"`
	StepNamesRequest                 *StepNamesRequest                 `protobuf:"bytes,15,opt,name=stepNamesRequest" json:"stepNamesRequest,omitempty"`
	StepNamesResponse                *StepNamesResponse                `protobuf:"bytes,16,opt,name=stepNamesResponse" json:"stepNamesResponse,omitempty"`
	SuiteExecutionResult             *SuiteExecutionResult             `protobuf:"bytes,17,opt,name=suiteExecutionResult" json:"suiteExecutionResult,omitempty"`
	KillProcessRequest               *KillProcessRequest               `protobuf:"bytes,18,opt,name=killProcessRequest" json:"killProcessRequest,omitempty"`
	XXX_unrecognized                 []byte                            `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

func (m *Message) GetMessageType() Message_MessageType {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return Message_ExecutionStarting
}

func (m *Message) GetMessageId() int64 {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return 0
}

func (m *Message) GetExecutionStartingRequest() *ExecutionStartingRequest {
	if m != nil {
		return m.ExecutionStartingRequest
	}
	return nil
}

func (m *Message) GetSpecExecutionStartingRequest() *SpecExecutionStartingRequest {
	if m != nil {
		return m.SpecExecutionStartingRequest
	}
	return nil
}

func (m *Message) GetSpecExecutionEndingRequest() *SpecExecutionEndingRequest {
	if m != nil {
		return m.SpecExecutionEndingRequest
	}
	return nil
}

func (m *Message) GetScenarioExecutionStartingRequest() *ScenarioExecutionStartingRequest {
	if m != nil {
		return m.ScenarioExecutionStartingRequest
	}
	return nil
}

func (m *Message) GetScenarioExecutionEndingRequest() *ScenarioExecutionEndingRequest {
	if m != nil {
		return m.ScenarioExecutionEndingRequest
	}
	return nil
}

func (m *Message) GetStepExecutionStartingRequest() *StepExecutionStartingRequest {
	if m != nil {
		return m.StepExecutionStartingRequest
	}
	return nil
}

func (m *Message) GetStepExecutionEndingRequest() *StepExecutionEndingRequest {
	if m != nil {
		return m.StepExecutionEndingRequest
	}
	return nil
}

func (m *Message) GetExecuteStepRequest() *ExecuteStepRequest {
	if m != nil {
		return m.ExecuteStepRequest
	}
	return nil
}

func (m *Message) GetExecutionEndingRequest() *ExecutionEndingRequest {
	if m != nil {
		return m.ExecutionEndingRequest
	}
	return nil
}

func (m *Message) GetStepValidateRequest() *StepValidateRequest {
	if m != nil {
		return m.StepValidateRequest
	}
	return nil
}

func (m *Message) GetStepValidateResponse() *StepValidateResponse {
	if m != nil {
		return m.StepValidateResponse
	}
	return nil
}

func (m *Message) GetExecutionStatusResponse() *ExecutionStatusResponse {
	if m != nil {
		return m.ExecutionStatusResponse
	}
	return nil
}

func (m *Message) GetStepNamesRequest() *StepNamesRequest {
	if m != nil {
		return m.StepNamesRequest
	}
	return nil
}

func (m *Message) GetStepNamesResponse() *StepNamesResponse {
	if m != nil {
		return m.StepNamesResponse
	}
	return nil
}

func (m *Message) GetSuiteExecutionResult() *SuiteExecutionResult {
	if m != nil {
		return m.SuiteExecutionResult
	}
	return nil
}

func (m *Message) GetKillProcessRequest() *KillProcessRequest {
	if m != nil {
		return m.KillProcessRequest
	}
	return nil
}

func init() {
	proto.RegisterEnum("main.Message_MessageType", Message_MessageType_name, Message_MessageType_value)
}
