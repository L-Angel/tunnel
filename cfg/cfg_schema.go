package cfg

var schemaTables map[string][]string

func init() {
	schemaTables = make(map[string][]string)
}

func GetSchemas() []string {
	j := 0
	keys := make([]string, len(schemaTables))
	for k := range schemaTables {
		keys[j] = k
		j++
	}
	return keys
}

func GetTablesBySchema(schema string) []string {
	return schemaTables[schema]
}