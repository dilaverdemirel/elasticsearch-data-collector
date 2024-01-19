package scheduler

import (
	"eslasticsearchdatacollector/dao"
	"eslasticsearchdatacollector/dao/model"
	"eslasticsearchdatacollector/gormlock"
	"eslasticsearchdatacollector/service"
	"log"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

func InitializeSchedulerAndActivateJobs() {
	var filter model.Index
	filter.Valid = true
	filter.Scheduled = true
	var indices = service.FindIndices(&filter)

	var worker string = "worker-" + uuid.NewString()
	var locker, err = gormlock.NewGormLocker(dao.DB, worker)
	if err != nil {
		log.Fatal(err)
	}

	// create a scheduler
	scheduler, err := gocron.NewScheduler(gocron.WithDistributedLocker(locker))
	if err != nil {
		log.Fatal(err)
	}

	for i, index := range indices {
		// add a job to the scheduler
		j, err := scheduler.NewJob(
			gocron.CronJob(
				index.CronExpression,
				true,
			),
			gocron.NewTask(
				func() {
					Sync(index.ID)
				},
			),
		)
		if err != nil {
			log.Fatal(err)
		}
		// each job has a unique id
		log.Println("Job ID : ", j.ID())
		log.Println("At index", i, "value is", index)
	}
	// start the scheduler
	scheduler.Start()

}
