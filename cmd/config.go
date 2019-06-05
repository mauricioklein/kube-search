package cmd

type config struct {
	namespace  string
	resource   string
	nRecords   uint16
	printScore bool
}

func (c config) printer() printable {
	if c.printScore {
		return scorePrinter{}
	}

	return simplePrinter{}
}
