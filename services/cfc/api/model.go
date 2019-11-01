/*
 * Copyright 2017 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

package api

import (
	"time"
)

type InvocationType string
type LogType string
type SourceType string
type TriggerType string

const (
	InvocationTypeEvent           InvocationType = "Event"
	InvocationTypeRequestResponse InvocationType = "RequestResponse"
	InvocationTypeDryRun          InvocationType = "DryRun"

	LogTypeTail LogType = "Tail"
	LogTypeNone LogType = "None"

	SourceTypeDuerOS SourceType = "dueros"
	SourceTypeBOS    SourceType = "bos/your-bucket-name"
	SourceTypeDuEdge SourceType = "duedge"
	SourceTypeHTTP   SourceType = "cfc-http-trigger/v1/CFCAPI"

	TriggerTypeHTTP    TriggerType = "cfc-http-trigger"
	TriggerTypeGeneric TriggerType = "generic"
)

type Function struct {
	Uid          string       `json:"Uid"`
	Description  string       `json:"Description"`
	FunctionBrn  string       `json:"FunctionBrn"`
	Region       string       `json:"Region"`
	Timeout      int          `json:"Timeout"`
	VersionDesc  string       `json:"VersionDesc"`
	UpdatedAt    time.Time    `json:"UpdatedAt"`
	LastModified time.Time    `json:"LastModified"`
	CodeSha256   string       `json:"CodeSha256"`
	CodeSize     int32        `json:"CodeSize"`
	FunctionArn  string       `json:"FunctionArn"`
	FunctionName string       `json:"FunctionName"`
	Handler      string       `json:"Handler"`
	Version      string       `json:"Version"`
	Runtime      string       `json:"Runtime"`
	MemorySize   int          `json:"MemorySize"`
	Environment  *Environment `json:"Environment"`
	CommitID     string       `json:"CommitID"`
	Role         string       `json:"Role"`
	LogType      string       `json:"LogType"`
	LogBosDir    string       `json:"LogBosDir"`
}

//functionInfo
type FunctionInfo struct {
	Code          *CodeStorage `json:"Code"`
	Configuration *Function    `json:"Configuration"`
}

type Alias struct {
	AliasBrn        string    `json:"AliasBrn"`
	AliasArn        string    `json:"AliasArn"`
	FunctionName    string    `json:"FunctionName"`
	FunctionVersion string    `json:"FunctionVersion"`
	Name            string    `json:"Name"`
	Description     string    `json:"Description"`
	Uid             string    `json:"Uid"`
	UpdatedAt       time.Time `json:"UpdatedAt"`
	CreatedAt       time.Time `json:"CreatedAt"`
}

type RelationInfo struct {
	RelationId string      `json:"RelationId"`
	Sid        string      `json:"Sid"`
	Source     string      `json:"Source"`
	Target     string      `json:"Target"`
	Data       interface{} `json:"Data"`
}

type CodeStorage struct {
	Location       string `json:"Location"`
	RepositoryType string `json:"RepositoryType"`
}

type Environment struct {
	Variables map[string]string
}

type CodeFile struct {
	ZipFile string
	Publish bool
	DryRun  bool
}

type InvocationsArgs struct {
	FunctionName   string
	InvocationType InvocationType
	LogType        LogType
	Qualifier      string
	Payload        interface{}
}

type InvocationsResult struct {
	Payload       string
	FunctionError string
	LogResult     string
}

type GetFunctionArgs struct {
	FunctionName string
	Qualifier    string
}

type GetFunctionResult struct {
	Code          CodeStorage
	Configuration Function
}

type ListFunctionsArgs struct {
	FunctionVersion string
	Marker          int
	MaxItems        int
}

type ListFunctionsResult struct {
	Functions []*Function
}

type CreateFunctionArgs struct {
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

type CreateFunctionResult Function

type DeleteFunctionArgs struct {
	FunctionName string
	Qualifier    string
}

type GetFunctionConfigurationArgs struct {
	FunctionName string
	Qualifier    string
}

type GetFunctionConfigurationResult Function

type UpdateFunctionConfigurationArgs struct {
	FunctionName string
	Timeout      int
	Description  string
	Handler      string
	Runtime      string
	Environment  *Environment
}

type UpdateFunctionConfigurationResult Function

type UpdateFunctionCodeArgs struct {
	FunctionName string
	ZipFile      string
	Publish      bool
	DryRun       bool
}

type UpdateFunctionCodeResult Function

type ListVersionsByFunctionArgs struct {
	FunctionName string
	Marker       int
	MaxItems     int
}
type ListVersionsByFunctionResult struct {
	Versions []*Function
}
type PublishVersionArgs struct {
	FunctionName string
	Description  string
	CodeSha256   string
}
type PublishVersionResult Function

type ListAliasesArgs struct {
	FunctionName    string
	FunctionVersion string
	Marker          int
	MaxItems        int
}

type ListAliasesResult struct {
	Aliases []*Alias
}

type GetAliasArgs struct {
	FunctionName string
	AliasName    string
}

type GetAliasResult Alias

type CreateAliasArgs struct {
	FunctionName    string
	FunctionVersion string
	Name            string
	Description     string
}

type CreateAliasResult Alias

type UpdateAliasArgs struct {
	FunctionName    string
	AliasName       string
	FunctionVersion string
	Description     string
}
type UpdateAliasResult Alias

type DeleteAliasArgs struct {
	FunctionName string
	AliasName    string
}

type ListTriggersArgs struct {
	FunctionBrn string
}
type ListTriggersResult struct {
	Relation []*RelationInfo
}

type CreateTriggerArgs struct {
	Target string
	Source SourceType
	Data   interface{}
}
type CreateTriggerResult struct {
	Relation *RelationInfo
}

type UpdateTriggerArgs struct {
	RelationId string
	Target     string
	Source     SourceType
	Data       interface{}
}
type UpdateTriggerResult struct {
	Relation *RelationInfo
}

type DeleteTriggerArgs struct {
	RelationId string
	Target     string
	Source     SourceType
}
