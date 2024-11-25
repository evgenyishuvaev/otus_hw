package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// log.Println("Старт пайплайна")
	var outCh Out
	var inCh Bi = make(Bi)

	go func() {
		defer close(inCh)
		for val := range in {
			// log.Printf("на очереди %d", val.(int))
			select {
			case <-done:
				return
			case inCh <- val:
				// log.Printf("Записываем в канал %d", val.(int))
			}
		}
	}()

	outCh = inCh
	// outCh = in
	for _, stage := range stages {
		outCh = stage(outCh)
	}
	return outCh
}
