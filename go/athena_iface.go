package athenadriver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/athena"
)

type QueryExecutionGetter interface {
	GetQueryExecution(ctx context.Context, params *athena.GetQueryExecutionInput, optFns ...func(*athena.Options)) (*athena.GetQueryExecutionOutput, error)
}

type QueryExecutionStarter interface {
	StartQueryExecution(ctx context.Context, params *athena.StartQueryExecutionInput, optFns ...func(*athena.Options)) (*athena.StartQueryExecutionOutput, error)
}

type QueryExecutionStopper interface {
	StopQueryExecution(ctx context.Context, params *athena.StopQueryExecutionInput, optFns ...func(*athena.Options)) (*athena.StopQueryExecutionOutput, error)
}

type QueryResultsGetter interface {
	GetQueryResults(ctx context.Context, params *athena.GetQueryResultsInput, optFns ...func(*athena.Options)) (*athena.GetQueryResultsOutput, error)
}

type WorkGroupCreator interface {
	CreateWorkGroup(ctx context.Context, params *athena.CreateWorkGroupInput, optFns ...func(*athena.Options)) (*athena.CreateWorkGroupOutput, error)
}

type WorkGroupGetter interface {
	GetWorkGroup(ctx context.Context, params *athena.GetWorkGroupInput, optFns ...func(*athena.Options)) (*athena.GetWorkGroupOutput, error)
}

type AthenaAPI interface {
	QueryExecutionGetter
	QueryExecutionStarter
	QueryExecutionStopper
	QueryResultsGetter
	WorkGroupGetter
	WorkGroupCreator
}
