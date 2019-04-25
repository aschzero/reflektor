package reflektor

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// JobCollectionConfig holds a collection of jobs from the job config yaml file
type JobCollectionConfig struct {
	Jobs []JobConfig `yaml:"jobs"`
}

// JobConfig represents a job parsed from the job config yaml file
type JobConfig struct {
	Name       string `yaml:"name"`
	SourcePath string `yaml:"source"`
	Schedule   string `yaml:"schedule"`
}

// NewJobCollection parses the supplied yaml data to return a collection of Jobs.
// Invalid jobs that are missing required fields are excluded from the collection.
func NewJobCollection(rawData []byte) ([]*Job, error) {
	config := &JobCollectionConfig{}

	err := yaml.Unmarshal(rawData, config)
	if err != nil {
		return nil, err
	}

	var jobs []*Job
	for _, job := range config.Jobs {
		err := job.Validate()
		if err != nil {
			log.WithField("job", job.Name).Error(err)
			continue
		}

		j := NewJob(job)
		jobs = append(jobs, j)
	}

	return jobs, nil
}

// Validate returns an error if any of the required fields are missing
func (j *JobConfig) Validate() error {
	if j.SourcePath == "" {
		return errors.New(fmt.Sprintf("missing source field"))
	}

	if j.Schedule == "" {
		return errors.New(fmt.Sprintf("missing schedule field"))
	}

	return nil
}
