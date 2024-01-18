package scheduler

import (
	"eslasticsearchdatacollector/dao/model"
	"eslasticsearchdatacollector/service"
	"fmt"
	"log"

	"github.com/go-co-op/gocron/v2"
)

func InitializeSchedulerAndActivateJobs() {
	var filter model.Index
	filter.Valid = true
	filter.Scheduled = true
	var indices = service.FindIndices(&filter)

	// create a scheduler
	var scheduler, err = gocron.NewScheduler()
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
		fmt.Println(j.ID())
		fmt.Println("At index", i, "value is", index)
	}
	// start the scheduler
	scheduler.Start()

}
