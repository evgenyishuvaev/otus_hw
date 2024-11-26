package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func stageHandler(in In, done In, stage Stage) Out {
	var inCh Bi = make(Bi)
	go func() {
		defer close(inCh)
		for {
			select {
			case <-done:
				return
			case val, ok := <-in:
				if !ok {
					return
				}
				select {
				case inCh <- val:
				case <-done:
					return
				}
			}
		}
	}()

	outCh := stage(inCh)

	var resCh Bi = make(Bi)
	go func() {
		defer close(resCh)
		for {
			select {
			case <-done:
				return
			case val, ok := <-outCh:
				if !ok {
					return
				}
				select {
				case resCh <- val:
				case <-done:
					return
				}
			}
		}
	}()
	return resCh
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	var outCh Out
	outCh = in
	for _, stage := range stages {
		outCh = stageHandler(outCh, done, stage)
	}
	return outCh
}
