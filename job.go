package reflektor

import (
	"fmt"
	"gopkg.in/robfig/cron.v2"
	"os"
)

// Job represents a scheduled archival job
type Job struct {
	ID         cron.EntryID
	Name       string
	SourcePath string
	Schedule   string
}

// NewJob builds an instance of a Job from the supplied config
func NewJob(config JobConfig) *Job {
	job := &Job{
		Name:       config.Name,
		SourcePath: config.SourcePath,
		Schedule:   config.Schedule,
	}

	return job
}

// SourceExists returns true if the job's source path is present
func (j *Job) SourceExists() bool {
	if _, err := os.Stat(j.SourcePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func (j *Job) ArchivePath() string {
	return fmt.Sprintf("/archives/%s.tar.gz", j.Name)
}
