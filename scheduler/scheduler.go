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

var SCHEDULER gocron.Scheduler

func InitializeSchedulerAndActivateJobs() {
	b := true
	var filter model.Index
	filter.Valid = &b
	filter.Scheduled = &b
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
	SCHEDULER = scheduler

	for i, index := range indices {
		// add a job to the scheduler
		add_new_job_to_scheduler_by_index(index)
		log.Println("At index", i, "value is", index)
	}
	// start the scheduler
	scheduler.Start()
}

func Add_new_job_to_scheduler_by_index_id(index_id string) {
	index := service.GetIndexById(index_id)
	add_new_job_to_scheduler_by_index(index)
}

func add_new_job_to_scheduler_by_index(index model.Index) {
	j, err := SCHEDULER.NewJob(
		gocron.CronJob(
			index.CronExpression,
			true,
		),
		gocron.NewTask(
			func() {
				Sync(index.ID)
			},
		),
		gocron.WithName(index.ID),
		gocron.WithTags(index.ID),
	)
	if err != nil {
		log.Fatal(err)
	}
	// each job has a unique id
	log.Println("Job ID : ", j.ID())
}

func Delete_job_by_index_id(index_id string) {
	SCHEDULER.RemoveByTags(index_id)
}
