package controller

import "time"

// CanaryJob holds the reference to a canary deployment schedule
type CanaryJob struct {
	Name      string
	Namespace string
	SkipTests bool
	function  func(name string, namespace string, skipTests bool)
	done      chan bool
	ticker    *time.Ticker
}

// Start runs the canary analysis on a schedule
func (j CanaryJob) Start() {
	go func() {
		// run the infra bootstrap on job creation
		j.function(j.Name, j.Namespace, j.SkipTests)
		for {
			select {
			case <-j.ticker.C:
				j.function(j.Name, j.Namespace, j.SkipTests)
			case <-j.done:
				return
			}
		}
	}()
}

// Stop closes the job channel and stops the ticker
func (j CanaryJob) Stop() {
	close(j.done)
	j.ticker.Stop()
}
