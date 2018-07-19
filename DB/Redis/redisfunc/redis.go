package redisfunc

type Redisfunc struct {
	Redis
}

func (red *Redisfunc) Pipe(comm string, urls ...chan [][]interface{}) chan interface{} {

	return red.pipe(comm, urls...)

}

func (red *Redisfunc) Sign(com string, table string, key []string, val []interface{}) ([]interface{}, bool) {

	return red.mergeValue(com, table, key, val)

}
func (red *Redisfunc) Chan(value [][]interface{}) chan [][]interface{} {

	return red.buildChan(value)

}

func (red *Redisfunc) DO(com string, urls chan []interface{}) chan interface{} {

	return red.do(com, urls)

}

func (red *Redisfunc) SingChan(value []interface{}) chan []interface{} {

	return red.signbuildChan(value)

}

func (red *Redisfunc) Scanlimit(com string, value []interface{}) chan interface{} {

	return red.scanlimit(com, value)

}
