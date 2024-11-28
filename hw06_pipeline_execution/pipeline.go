package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func garbage(ch Out) {
	for val := range ch {
		_ = val
	}
}

func stageHandler(in In, done In, stage Stage) Out {
	outCh := stage(in)

	resCh := make(Bi)
	go func() {
		defer garbage(outCh)
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
	inCh := make(Bi)
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

	outCh = inCh
	for _, stage := range stages {
		outCh = stageHandler(outCh, done, stage)
	}
	return outCh
}
