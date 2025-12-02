package model

import "time"

type IngestionEventType string

const (
	IngestionEventTypeTraceCreate       = "trace-create"
	IngestionEventTypeGenerationCreate  = "generation-create"
	IngestionEventTypeGenerationUpdate  = "generation-update"
	IngestionEventTypeScoreCreate       = "score-create"
	IngestionEventTypeSpanCreate        = "span-create"
	IngestionEventTypeSpanUpdate        = "span-update"
	IngestionEventTypeEventCreate       = "event-create"
	IngestionEventTypeSDKLog            = "sdk-log"
	IngestionEventTypeObservationCreate = "observation-create"
	IngestionEventTypeObservationUpdate = "observation-update"
)

type IngestionEvent struct {
	Type      IngestionEventType `json:"type"`
	ID        string             `json:"id"`
	Timestamp time.Time          `json:"timestamp"`
	Metadata  any                `json:"metadata,omitempty"`
	Body      any                `json:"body"`
}

type Trace struct {
	ID          string     `json:"id,omitempty"`
	Timestamp   *time.Time `json:"timestamp,omitempty"`
	Name        string     `json:"name,omitempty"`
	UserID      string     `json:"userId,omitempty"`
	Input       any        `json:"input,omitempty"`
	Output      any        `json:"output,omitempty"`
	SessionID   string     `json:"sessionId,omitempty"`
	Release     string     `json:"release,omitempty"`
	Version     string     `json:"version,omitempty"`
	Metadata    any        `json:"metadata,omitempty"`
	Tags        []string   `json:"tags,omitempty"`
	Environment string     `json:"environment,omitempty"`
	Public      bool       `json:"public,omitempty"`
}

type ObservationType string

const (
	ObservationTypeSpan       ObservationType = "SPAN"
	ObservationTypeGeneration ObservationType = "GENERATION"
	ObservationTypeEvent      ObservationType = "EVENT"
	ObservationTypeAgent      ObservationType = "AGENT"
	ObservationTypeTool       ObservationType = "TOOL"
	ObservationTypeChain      ObservationType = "CHAIN"
	ObservationTypeRetriever  ObservationType = "RETRIEVER"
	ObservationTypeEvaluator  ObservationType = "EVALUATOR"
	ObservationTypeEmbedding  ObservationType = "EMBEDDING"
	ObservationTypeGuardrail  ObservationType = "GUARDRAIL"
)

type ObservationLevel string

const (
	ObservationLevelDebug   ObservationLevel = "DEBUG"
	ObservationLevelDefault ObservationLevel = "DEFAULT"
	ObservationLevelWarning ObservationLevel = "WARNING"
	ObservationLevelError   ObservationLevel = "ERROR"
)

type Generation struct {
	TraceID             string                 `json:"traceId,omitempty"`
	Type                ObservationType        `json:"type,omitempty"`
	Name                string                 `json:"name,omitempty"`
	StartTime           *time.Time             `json:"startTime,omitempty"`
	EndTime             *time.Time             `json:"endTime,omitempty"`
	CompletionStartTime *time.Time             `json:"completionStartTime,omitempty"`
	Model               string                 `json:"model,omitempty"`
	ModelParameters     map[string]interface{} `json:"modelParameters,omitempty"`
	Input               any                    `json:"input,omitempty"`
	Output              any                    `json:"output,omitempty"`
	Version             string                 `json:"version,omitempty"`
	Metadata            any                    `json:"metadata,omitempty"`
	Usage               *Usage                 `json:"usage,omitempty"`
	UsageDetails        any                    `json:"usageDetails,omitempty"`
	CostDetails         map[string]float64     `json:"costDetails,omitempty"`
	Level               ObservationLevel       `json:"level,omitempty"`
	StatusMessage       string                 `json:"statusMessage,omitempty"`
	ParentObservationID string                 `json:"parentObservationId,omitempty"`
	Environment         string                 `json:"environment,omitempty"`
	PromptName          string                 `json:"promptName,omitempty"`
	PromptVersion       *int                   `json:"promptVersion,omitempty"`
	ID                  string                 `json:"id,omitempty"`
}

type Usage struct {
	Input      int            `json:"input,omitempty"`
	Output     int            `json:"output,omitempty"`
	Total      int            `json:"total,omitempty"`
	Unit       ModelUsageUnit `json:"unit,omitempty"`
	InputCost  float64        `json:"inputCost,omitempty"`
	OutputCost float64        `json:"outputCost,omitempty"`
	TotalCost  float64        `json:"totalCost,omitempty"`

	PromptTokens     int `json:"promptTokens,omitempty"`
	CompletionTokens int `json:"completionTokens,omitempty"`
	TotalTokens      int `json:"totalTokens,omitempty"`
}

type ModelUsageUnit string

const (
	ModelUsageUnitCharacters   ModelUsageUnit = "CHARACTERS"
	ModelUsageUnitTokens       ModelUsageUnit = "TOKENS"
	ModelUsageUnitMilliseconds ModelUsageUnit = "MILLISECONDS"
	ModelUsageUnitSeconds      ModelUsageUnit = "SECONDS"
	ModelUsageUnitImages       ModelUsageUnit = "IMAGES"
)

// UsageUnit is kept for backward compatibility with previous versions of the SDK.
//
//nolint:revive
type UsageUnit = ModelUsageUnit

type Score struct {
	ID            string        `json:"id,omitempty"`
	TraceID       string        `json:"traceId,omitempty"`
	SessionID     string        `json:"sessionId,omitempty"`
	ObservationID string        `json:"observationId,omitempty"`
	DatasetRunID  string        `json:"datasetRunId,omitempty"`
	Name          string        `json:"name,omitempty"`
	Environment   string        `json:"environment,omitempty"`
	QueueID       string        `json:"queueId,omitempty"`
	Value         any           `json:"value,omitempty"`
	Comment       string        `json:"comment,omitempty"`
	Metadata      any           `json:"metadata,omitempty"`
	DataType      ScoreDataType `json:"dataType,omitempty"`
	ConfigID      string        `json:"configId,omitempty"`
}

type Span struct {
	TraceID             string           `json:"traceId,omitempty"`
	Type                ObservationType  `json:"type,omitempty"`
	Name                string           `json:"name,omitempty"`
	StartTime           *time.Time       `json:"startTime,omitempty"`
	EndTime             *time.Time       `json:"endTime,omitempty"`
	Metadata            any              `json:"metadata,omitempty"`
	Input               any              `json:"input,omitempty"`
	Output              any              `json:"output,omitempty"`
	Level               ObservationLevel `json:"level,omitempty"`
	StatusMessage       string           `json:"statusMessage,omitempty"`
	ParentObservationID string           `json:"parentObservationId,omitempty"`
	Version             string           `json:"version,omitempty"`
	Environment         string           `json:"environment,omitempty"`
	ID                  string           `json:"id,omitempty"`
}

type Event struct {
	TraceID             string           `json:"traceId,omitempty"`
	Type                ObservationType  `json:"type,omitempty"`
	Name                string           `json:"name,omitempty"`
	StartTime           *time.Time       `json:"startTime,omitempty"`
	EndTime             *time.Time       `json:"endTime,omitempty"`
	Metadata            any              `json:"metadata,omitempty"`
	Input               any              `json:"input,omitempty"`
	Output              any              `json:"output,omitempty"`
	Level               ObservationLevel `json:"level,omitempty"`
	StatusMessage       string           `json:"statusMessage,omitempty"`
	ParentObservationID string           `json:"parentObservationId,omitempty"`
	Version             string           `json:"version,omitempty"`
	Environment         string           `json:"environment,omitempty"`
	ID                  string           `json:"id,omitempty"`
}

type Observation struct {
	TraceID             string             `json:"traceId,omitempty"`
	Type                ObservationType    `json:"type,omitempty"`
	Name                string             `json:"name,omitempty"`
	StartTime           *time.Time         `json:"startTime,omitempty"`
	EndTime             *time.Time         `json:"endTime,omitempty"`
	CompletionStartTime *time.Time         `json:"completionStartTime,omitempty"`
	Model               string             `json:"model,omitempty"`
	ModelParameters     map[string]any     `json:"modelParameters,omitempty"`
	Input               any                `json:"input,omitempty"`
	Version             string             `json:"version,omitempty"`
	Metadata            any                `json:"metadata,omitempty"`
	Output              any                `json:"output,omitempty"`
	Usage               *Usage             `json:"usage,omitempty"`
	UsageDetails        any                `json:"usageDetails,omitempty"`
	CostDetails         map[string]float64 `json:"costDetails,omitempty"`
	Level               ObservationLevel   `json:"level,omitempty"`
	StatusMessage       string             `json:"statusMessage,omitempty"`
	ParentObservationID string             `json:"parentObservationId,omitempty"`
	Environment         string             `json:"environment,omitempty"`
	ID                  string             `json:"id,omitempty"`
}

type SDKLog struct {
	Log any `json:"log"`
}

type ScoreDataType string

const (
	ScoreDataTypeNumeric     ScoreDataType = "NUMERIC"
	ScoreDataTypeBoolean     ScoreDataType = "BOOLEAN"
	ScoreDataTypeCategorical ScoreDataType = "CATEGORICAL"
)

type M map[string]interface{}
