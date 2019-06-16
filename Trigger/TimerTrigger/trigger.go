package MyTimerTrigger

import (
	"context"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/carlescere/scheduler"
	"strconv"
)

// Create a new logger
var log = logger.GetLogger("trigger-mytrigger")

// MyTriggerFactory My Trigger factory
type MyTriggerFactory struct {
	metadata *trigger.Metadata
}

// NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &MyTriggerFactory{metadata: md}
}

// New Creates a new trigger instance for a given id
func (t *MyTriggerFactory) New(config *trigger.Config) trigger.Trigger {
	return &MyTrigger{metadata: t.metadata, config: config}
}

// MyTrigger is a stub for your Trigger implementation
type MyTrigger struct {
	metadata *trigger.Metadata
	config   *trigger.Config
	timers   []*scheduler.Job
	handlers []*trigger.Handler
}

// Initialize implements trigger.Init.Initialize
func (t *MyTrigger) Initialize(ctx trigger.InitContext) error {
	t.handlers = ctx.GetHandlers()
	return nil
}

// Metadata implements trigger.Trigger.Metadata
func (t *MyTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Start implements trigger.Trigger.Start
func (t *MyTrigger) Start() error {
	log.Debug("Start")
	handlers := t.handlers

	log.Debug("Processing handlers")
	for _, handler := range handlers {
		t.scheduleRepeating(handler)
	}

	return nil
}

// Stop implements trigger.Trigger.Start
func (t *MyTrigger) Stop() error {
	// stop the trigger
	return nil
}

func (t *MyTrigger) scheduleRepeating(endpoint *trigger.Handler) {
	log.Info("Scheduling a repeating job")

	fn2 := func() {
		// Create a map to hold the trigger data
		triggerData := map[string]interface{}{
			"output": "Hello World from the new Timer Trigger",
		}

		_, err := endpoint.Handle(context.Background(), triggerData)
		if err != nil {
			log.Error("Error running handler: ", err.Error())
		}
	}

	t.scheduleJobEverySecond(endpoint, fn2)

}

func (t *MyTrigger) scheduleJobEverySecond(tgrHandler *trigger.Handler, fn func()) {

	var interval= 0
	seconds, _ := strconv.Atoi(tgrHandler.GetStringSetting("seconds"))
	interval = interval + seconds

	log.Debug("Repeating seconds: ", interval)
	// schedule repeating
	timerJob, err := scheduler.Every(interval).Seconds().Run(fn)
	if err != nil {
		log.Error("Error scheduleRepeating (repeat seconds) flo err: ", err.Error())
	}
	if timerJob == nil {
		log.Error("timerJob is nil")
	}

	t.timers = append(t.timers, timerJob)
}
