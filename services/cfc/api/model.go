package api

import (
	"time"
)

type InvocationType string
type LogType string

const (
	InvocationTypeEvent           InvocationType = "Event"
	InvocationTypeRequestResponse InvocationType = "RequestResponse"
	InvocationTypeDryRun          InvocationType = "DryRun"
)

const (
	LogTypeTail LogType = "Tail"
	LogTypeNone LogType = "None"
)

const (
	ParseJsonError = "Could not parse payload into json"
)

//function
type Function struct {
	Id             uint
	Uid            string
	Description    string
	FunctionBrn    string
	Region         string
	Timeout        int
	EnvironmentStr string `json:"-"`
	VersionDesc    string
	UpdatedAt      time.Time
	LastModified   time.Time
	FunctionConfig
}

// CodeStorage
type CodeStorage struct {
	Location       string
	RepositoryType string
}

// FunctionConfig
type FunctionConfig struct {
	CodeSha256   string
	CodeSize     int32
	FunctionArn  string
	FunctionName string `valid:"optional,matches(^[a-zA-Z0-9-_]+$),runelength(1|64)"`
	Handler      string `valid:"optional,matches(^([a-zA-Z0-9-_]+\\.[a-zA-Z0-9-_]+)$),runelength(1|32)"`
	Version      string `valid:"optional,matches(^\\$LATEST|([0-9]+)$),runelength(1|32)"`
	Runtime      string `valid:"optional,matches(^(nodejs6\\.11|nodejs8\\.4|nodejs8\\.5|python2|python3)$)"`
	MemorySize   int
	Environment  *Environment
	CommitID     string
	LogType      string
}

// Function Environment
type Environment struct {
	Variables map[string]string
}

type InvocationsArgs struct {
	InvocationType InvocationType
	LogType        LogType
	Qualifier      string
}

type InvocationResult struct {
	Payload       string
	FunctionError string
	LogResult     string
}

type ListFunctionsArgs struct {
	FunctionVersion string `valid:"optional,matches(^\\$LATEST|([0-9]+)|ALL$),runelength(1|32)"`
	Marker          string `valid:"int"`
	MaxItems        string `valid:"int"`
}

type ListFunctionsResult struct {
	Functions []*Function
}

type GetFunctionResult struct {
	Code          CodeStorage
	Configuration Function
}

type CodeFile struct {
	ZipFile string `valid:"base64,required"`
	Publish bool
	DryRun  bool
}

// req body
type FunctionArgs struct {
	Code         *CodeFile
	Publish      bool
	FunctionName string `valid:"required,matches(^[a-zA-Z0-9-_]+$),runelength(1|64)"`
	Handler      string `valid:"required,matches(^([a-zA-Z0-9-_]+\\.[a-zA-Z0-9-_]+)$),runelength(1|32)"`
	Runtime      string `valid:"required,matches(^(nodejs6\\.11|nodejs8\\.4|nodejs8\\.5|python2|python3)$)"`
	MemorySize   int
	Timeout      int
	Description  string       `valid:"optional,"`
	Environment  *Environment `valid:"-"`
}
