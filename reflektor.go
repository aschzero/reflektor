package reflektor

import (
	"compress/flate"
	"github.com/mholt/archiver"
	log "github.com/sirupsen/logrus"
	"gopkg.in/robfig/cron.v2"
	"time"
)

// Reflektor holds an instance of the cron scheduler and the collection jobs
type Reflektor struct {
	Cron *cron.Cron
	Jobs []*Job
}

// ScheduleJobs adds each job to the cron scheduler and starts the cron service
func (r *Reflektor) ScheduleJobs() {
	r.Cron = cron.New()

	for _, job := range r.Jobs {
		j := job

		id, err := r.Cron.AddFunc(job.Schedule, func() {
			r.RunJob(j)
		})

		if err != nil {
			log.WithFields(log.Fields{"job": j.Name, "error": err}).Fatal("unable to schedule job")
			continue
		}

		job.ID = id
	}

	r.Cron.Start()

	for _, job := range r.Jobs {
		entity := r.Cron.Entry(job.ID)

		log.WithFields(log.Fields{
			"job":      job.Name,
			"next_run": entity.Next,
		}).Info("job scheduled")
	}
}

// RunJob archives the given job's source directory
func (r *Reflektor) RunJob(job *Job) {
	if !job.SourceExists() {
		log.WithFields(log.Fields{
			"job":    job.Name,
			"source": job.SourcePath,
		}).Error("unable to find job source path")

		return
	}

	log.WithField("job", job.Name).Info("job running")

	start := time.Now()

	z := archiver.TarGz{
		CompressionLevel: flate.BestCompression,
		Tar: &archiver.Tar{
			OverwriteExisting:      true,
			ImplicitTopLevelFolder: true,
			MkdirAll:               true,
			ContinueOnError:        true,
		},
	}

	err := z.Archive([]string{job.SourcePath}, job.ArchivePath())
	if err != nil {
		log.WithFields(log.Fields{
			"job":   job.Name,
			"error": err,
		}).Info("job failed")

		return
	}

	elapsed := time.Since(start)
	cronEntity := r.Cron.Entry(job.ID)

	log.WithFields(log.Fields{
		"job":      job.Name,
		"elapsed":  elapsed,
		"next_run": cronEntity.Next,
	}).Info("job finished")
}
