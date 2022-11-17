// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CreateCategoryInput struct {
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceID"`
}

type CreateDataSourceInput struct {
	SiloDefinitionID string   `json:"siloDefinitionID"`
	Description      *string  `json:"description"`
	PropertyIDs      []string `json:"propertyIDs"`
}

type CreatePropertyInput struct {
	CategoryIDs  []string `json:"categoryIDs"`
	DataSourceID string   `json:"dataSourceID"`
	PurposeIDs   []string `json:"purposeIDs"`
}

type CreatePurposeInput struct {
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceID"`
}

type CreateSiloDefinitionInput struct {
	Description         *string  `json:"description"`
	SiloSpecificationID string   `json:"siloSpecificationID"`
	WorkspaceID         string   `json:"workspaceID"`
	SubjectIDs          []string `json:"subjectIDs"`
	SiloData            *string  `json:"siloData"`
	Name                string   `json:"name"`
}

type CreateSiloSpecificationInput struct {
	Name        string  `json:"name"`
	WorkspaceID string  `json:"workspaceID"`
	LogoURL     *string `json:"logoURL"`
	DockerImage string  `json:"dockerImage"`
	Schema      *string `json:"schema"`
}

type CreateSubjectInput struct {
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceID"`
}

type CreateUserPrimaryKeyInput struct {
	Name          string `json:"name"`
	APIIdentifier string `json:"apiIdentifier"`
	WorkspaceID   string `json:"workspaceId"`
}

type CreateWorkspaceInput struct {
	Name     string    `json:"name"`
	Settings []*KVPair `json:"settings"`
}

type DataDiscoveriesListResult struct {
	Discoveries    []*DataDiscovery `json:"discoveries"`
	NumDiscoveries int              `json:"numDiscoveries"`
}

type HandleAllDiscoveriesInput struct {
	SiloID string          `json:"siloId"`
	Action DiscoveryAction `json:"action"`
}

type HandleDiscoveryInput struct {
	DiscoveryID string          `json:"discoveryId"`
	Action      DiscoveryAction `json:"action"`
}

type JobsResult struct {
	Jobs    []*Job `json:"jobs"`
	NumJobs int    `json:"numJobs"`
}

type KVPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MonoidRecordResponse struct {
	Data        string  `json:"data"`
	SchemaGroup *string `json:"SchemaGroup"`
	SchemaName  string  `json:"SchemaName"`
}

type RequestsResult struct {
	Requests    []*Request `json:"requests"`
	NumRequests int        `json:"numRequests"`
}

type SiloScanConfigInput struct {
	SiloID string `json:"siloId"`
	// A cron string that can be used to schedule
	// the scan, or empty string if automatic scanning
	// is disabled.
	Cron string `json:"cron"`
}

type UpdateCategoryInput struct {
	Name *string `json:"name"`
}

type UpdateDataSourceInput struct {
	ID          string  `json:"id"`
	Description *string `json:"description"`
}

type UpdatePropertyInput struct {
	ID          string   `json:"id"`
	CategoryIDs []string `json:"categoryIDs"`
	PurposeIDs  []string `json:"purposeIDs"`
}

type UpdatePurposeInput struct {
	Name *string `json:"name"`
}

type UpdateSiloDefinitionInput struct {
	ID          string   `json:"id"`
	WorkspaceID string   `json:"workspaceId"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	SubjectIDs  []string `json:"subjectIDs"`
	SiloData    *string  `json:"siloData"`
}

type UpdateSiloSpecificationInput struct {
	ID          string  `json:"id"`
	DockerImage *string `json:"dockerImage"`
	Schema      *string `json:"schema"`
	Name        *string `json:"name"`
	LogoURL     *string `json:"logoUrl"`
}

type UpdateSubjectInput struct {
	Name *string `json:"name"`
}

type UpdateUserPrimaryKeyInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateWorkspaceSettingsInput struct {
	WorkspaceID string    `json:"workspaceID"`
	Settings    []*KVPair `json:"settings"`
}

type UserDataRequestInput struct {
	PrimaryKeys []*UserPrimaryKeyInput `json:"primaryKeys"`
	WorkspaceID string                 `json:"workspaceId"`
	Type        UserDataRequestType    `json:"type"`
}

type UserPrimaryKeyInput struct {
	APIIdentifier string `json:"apiIdentifier"`
	Value         string `json:"value"`
}

type DiscoveryAction string

const (
	DiscoveryActionAccept DiscoveryAction = "ACCEPT"
	DiscoveryActionReject DiscoveryAction = "REJECT"
)

var AllDiscoveryAction = []DiscoveryAction{
	DiscoveryActionAccept,
	DiscoveryActionReject,
}

func (e DiscoveryAction) IsValid() bool {
	switch e {
	case DiscoveryActionAccept, DiscoveryActionReject:
		return true
	}
	return false
}

func (e DiscoveryAction) String() string {
	return string(e)
}

func (e *DiscoveryAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DiscoveryAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DiscoveryAction", str)
	}
	return nil
}

func (e DiscoveryAction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DiscoveryStatus string

const (
	DiscoveryStatusOpen     DiscoveryStatus = "OPEN"
	DiscoveryStatusAccepted DiscoveryStatus = "ACCEPTED"
	DiscoveryStatusRejected DiscoveryStatus = "REJECTED"
)

var AllDiscoveryStatus = []DiscoveryStatus{
	DiscoveryStatusOpen,
	DiscoveryStatusAccepted,
	DiscoveryStatusRejected,
}

func (e DiscoveryStatus) IsValid() bool {
	switch e {
	case DiscoveryStatusOpen, DiscoveryStatusAccepted, DiscoveryStatusRejected:
		return true
	}
	return false
}

func (e DiscoveryStatus) String() string {
	return string(e)
}

func (e *DiscoveryStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DiscoveryStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DiscoveryStatus", str)
	}
	return nil
}

func (e DiscoveryStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DiscoveryType string

const (
	DiscoveryTypeDataSourceMissing DiscoveryType = "DATA_SOURCE_MISSING"
	DiscoveryTypeDataSourceFound   DiscoveryType = "DATA_SOURCE_FOUND"
	DiscoveryTypePropertyFound     DiscoveryType = "PROPERTY_FOUND"
	DiscoveryTypePropertyMissing   DiscoveryType = "PROPERTY_MISSING"
	DiscoveryTypeCategoryFound     DiscoveryType = "CATEGORY_FOUND"
)

var AllDiscoveryType = []DiscoveryType{
	DiscoveryTypeDataSourceMissing,
	DiscoveryTypeDataSourceFound,
	DiscoveryTypePropertyFound,
	DiscoveryTypePropertyMissing,
	DiscoveryTypeCategoryFound,
}

func (e DiscoveryType) IsValid() bool {
	switch e {
	case DiscoveryTypeDataSourceMissing, DiscoveryTypeDataSourceFound, DiscoveryTypePropertyFound, DiscoveryTypePropertyMissing, DiscoveryTypeCategoryFound:
		return true
	}
	return false
}

func (e DiscoveryType) String() string {
	return string(e)
}

func (e *DiscoveryType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DiscoveryType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DiscoveryType", str)
	}
	return nil
}

func (e DiscoveryType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type JobStatus string

const (
	JobStatusQueued    JobStatus = "QUEUED"
	JobStatusRunning   JobStatus = "RUNNING"
	JobStatusCompleted JobStatus = "COMPLETED"
	JobStatusFailed    JobStatus = "FAILED"
)

var AllJobStatus = []JobStatus{
	JobStatusQueued,
	JobStatusRunning,
	JobStatusCompleted,
	JobStatusFailed,
}

func (e JobStatus) IsValid() bool {
	switch e {
	case JobStatusQueued, JobStatusRunning, JobStatusCompleted, JobStatusFailed:
		return true
	}
	return false
}

func (e JobStatus) String() string {
	return string(e)
}

func (e *JobStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = JobStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid JobStatus", str)
	}
	return nil
}

func (e JobStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RequestStatusType string

const (
	RequestStatusTypeCreated  RequestStatusType = "CREATED"
	RequestStatusTypeExecuted RequestStatusType = "EXECUTED"
	RequestStatusTypeFailed   RequestStatusType = "FAILED"
)

var AllRequestStatusType = []RequestStatusType{
	RequestStatusTypeCreated,
	RequestStatusTypeExecuted,
	RequestStatusTypeFailed,
}

func (e RequestStatusType) IsValid() bool {
	switch e {
	case RequestStatusTypeCreated, RequestStatusTypeExecuted, RequestStatusTypeFailed:
		return true
	}
	return false
}

func (e RequestStatusType) String() string {
	return string(e)
}

func (e *RequestStatusType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RequestStatusType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RequestStatusType", str)
	}
	return nil
}

func (e RequestStatusType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserDataRequestType string

const (
	UserDataRequestTypeDelete UserDataRequestType = "DELETE"
	UserDataRequestTypeQuery  UserDataRequestType = "QUERY"
)

var AllUserDataRequestType = []UserDataRequestType{
	UserDataRequestTypeDelete,
	UserDataRequestTypeQuery,
}

func (e UserDataRequestType) IsValid() bool {
	switch e {
	case UserDataRequestTypeDelete, UserDataRequestTypeQuery:
		return true
	}
	return false
}

func (e UserDataRequestType) String() string {
	return string(e)
}

func (e *UserDataRequestType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserDataRequestType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserDataRequestType", str)
	}
	return nil
}

func (e UserDataRequestType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
