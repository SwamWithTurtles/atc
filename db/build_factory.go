package db

import (
	"database/sql"

	"github.com/lib/pq"
)

func newBuildFactory(conn Conn, bus *notificationsBus) *buildFactory {
	return &buildFactory{
		conn: conn,
		bus:  bus,
	}
}

type buildFactory struct {
	conn Conn
	bus  *notificationsBus
}

func (f *buildFactory) ScanBuild(row scannable) (Build, bool, error) {
	var id int
	var name string
	var jobID, pipelineID, teamID sql.NullInt64
	var status string
	var scheduled bool
	var engine, engineMetadata, jobName, pipelineName sql.NullString
	var startTime pq.NullTime
	var endTime pq.NullTime
	var reapTime pq.NullTime
	var teamName string

	err := row.Scan(&id, &name, &jobID, &teamID, &status, &scheduled, &engine, &engineMetadata, &startTime, &endTime, &reapTime, &jobName, &pipelineID, &pipelineName, &teamName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}

		return nil, false, err
	}

	build := &build{
		conn: f.conn,
		bus:  f.bus,

		id:        id,
		name:      name,
		status:    Status(status),
		scheduled: scheduled,

		engine:         engine.String,
		engineMetadata: engineMetadata.String,

		startTime: startTime.Time,
		endTime:   endTime.Time,
		reapTime:  reapTime.Time,

		teamName: teamName,
	}

	if jobID.Valid {
		build.jobName = jobName.String
		build.pipelineName = pipelineName.String
		build.pipelineID = int(pipelineID.Int64)
	}

	if teamID.Valid {
		build.teamID = int(teamID.Int64)
	}

	return build, true, nil
}
