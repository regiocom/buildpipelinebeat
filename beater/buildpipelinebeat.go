package beater

import (
	// Custom Imports

	// Default Import
	"fmt"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/regiocom/buildpipelinebeat/config"
)

// buildpipelinebeat configuration.
type buildpipelinebeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of buildpipelinebeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &buildpipelinebeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts buildpipelinebeat.
func (bt *buildpipelinebeat) Run(b *beat.Beat) error {
	logp.Info("buildpipelinebeat is running! Hit CTRL-C to stop it.")

	// Connection Variable
	CloseTimeout := bt.config.CloseTimeout

	var err error
	bt.client, err = b.Publisher.ConnectWith(beat.ClientConfig{
		PublishMode: beat.GuaranteedSend,
		WaitClose:   CloseTimeout * time.Second,
	})
	if err != nil {
		return err
	}

	// Variables
	Status := bt.config.Status
	TeamName := bt.config.Team
	Project := bt.config.Project
	Pipeline := bt.config.Pipeline
	Error := bt.config.Error

	// Prepare the Event
	event := beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"type":     b.Info.Name,
			"team":     TeamName,
			"pipeline": Pipeline,
			"project":  Project,
			"status":   Status,
			"error":    Error,
		},
	}

	// Push the event and stop the beat
	bt.client.Publish(event)
	logp.Info("Event sent")
	defer close(bt.done)
	return nil
}

// Stop stops buildpipelinebeat.
func (bt *buildpipelinebeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
