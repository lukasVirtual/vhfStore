package main

import (
	"sync"

	"github.com/judwhite/go-svc"
	log "github.com/sirupsen/logrus"
)

type program struct {
	wg   sync.WaitGroup
	quit chan struct{}
}

func test_main() {
	prg := &program{}

	if err := svc.Run(prg); err != nil {
		log.Fatal(err)
	}
}

// Starts before the program starts and after it
// determines if it runs in a service env or not.
func (p *program) Init(env svc.Environment) error {
	log.Infof("is win service ? %v\n", env.IsWindowsService())
	return nil
}

// Run's when progam ist started
func (p *program) Start() error {
	p.quit = make(chan struct{})

	p.wg.Add(1)
	go func() {
		log.Debugln("Starting...")
		<-p.quit
		log.Debugln("Quit signal received...")
		p.wg.Done()
	}()

	return nil
}

func (p *program) Stop() error {
	log.Debugln("Stopping..")
	close(p.quit)
	p.wg.Wait()
	log.Debugln("Stopped.")
	return nil
}
