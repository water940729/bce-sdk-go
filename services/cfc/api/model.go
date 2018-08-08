package api

import (
	"time"
)

type InvocationType string
type LogType string
type TriggerType string

const (
	BceFaasTriggerKey = "X-Bce-Faas-Trigger"

	InvocationTypeEvent           InvocationType = "Event"
	InvocationTypeRequestResponse InvocationType = "RequestResponse"
	InvocationTypeDryRun          InvocationType = "DryRun"

	LogTypeTail LogType = "Tail"
	LogTypeNone LogType = "None"

	TriggerTypeHTTP = "cfc-http-trigger"
	TriggerTypeBos  = "bos"
)

const (
	ParseJsonError = "Could not parse payload into json"
)

//function
type Function struct {
	Id           uint
	Uid          string
	Description  string
	FunctionBrn  string
	Region       string
	Timeout      int
	VersionDesc  string
	UpdatedAt    time.Time
	LastModified time.Time
	CodeSha256   string
	CodeSize     int32
	FunctionArn  string
	FunctionName string
	Handler      string
	Version      string
	Runtime      string
	MemorySize   int
	Environment  *Environment
	CommitID     string
	LogType      string
}

// CodeStorage
type CodeStorage struct {
	Location       string
	RepositoryType string
}

// Function Environment
type Environment struct {
	Variables map[string]string
}

type GetFunctionResult struct {
	Code          CodeStorage
	Configuration Function
}

type CodeFile struct {
	ZipFile string
	Publish bool
	DryRun  bool
}

// req body
type FunctionArgs struct {
	Code         *CodeFile
	Publish      bool
	FunctionName string
	Handler      string
	Runtime      string
	MemorySize   int
	Timeout      int
	Description  string
	Environment  *Environment
}

type InvocationsArgs struct {
	InvocationType InvocationType
	LogType        LogType
	Qualifier      string
	Trigger        TriggerType
}

type InvocationResult struct {
	Payload       string
	FunctionError string
	LogResult     string
}

type ListFunctionsArgs struct {
	FunctionVersion string
	Marker          string
	MaxItems        string
}

type ListFunctionsResult struct {
	Functions []*Function
}
