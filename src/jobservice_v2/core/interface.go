// Copyright 2018 The Harbor Authors. All rights reserved.

//Package core provides the main job operation interface and components.
package core

import (
	"github.com/vmware/harbor/src/jobservice_v2/models"
)

//Interface defines the related main methods of job operation.
type Interface interface {
	//LaunchJob is used to handle the job submission request.
	//
	//req	JobRequest    : Job request contains related required information of queuing job.
	//
	//Returns:
	//	JobStats: Job status info with ID and self link returned if job is successfully launched.
	//  error   : Error returned if failed to launch the specified job.
	LaunchJob(req models.JobRequest) (models.JobStats, error)

	//GetJob is used to handle the job stats query request.
	//
	//jobID	string: ID of job.
	//
	//Returns:
	//	JobStats: Job status info if job exists.
	//  error   : Error returned if failed to get the specified job.
	GetJob(jobID string) (models.JobStats, error)

	//StopJob is used to handle the job stopping request.
	//
	//jobID	string: ID of job.
	//
	//Return:
	//  error   : Error returned if failed to stop the specified job.
	StopJob(jobID string) error

	//RetryJob is used to handle the job retrying request.
	//
	//jobID	string        : ID of job.
	//
	//Return:
	//  error   : Error returned if failed to retry the specified job.
	RetryJob(jonID string) error

	//Cancel the job
	//
	//jobID string : ID of the enqueued job
	//
	//Returns:
	//  error           : error returned if meet any problems
	CancelJob(jobID string) error

	//CheckStatus is used to handle the job service healthy status checking request.
	CheckStatus() (models.JobPoolStats, error)

	//GetJobLogData is used to return the log text data for the specified job if exists
	GetJobLogData(jobID string) ([]byte, error)
}
