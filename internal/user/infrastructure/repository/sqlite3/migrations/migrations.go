package migrations

type migration struct {
	list map[string][]string
}

func GetList() (map[string][]string, error) {
	migration := &migration{
		list: map[string][]string{},
	}
	if err := migration.appendAll(); err != nil {
		return nil, err
	}
	return migration.list, nil
}
