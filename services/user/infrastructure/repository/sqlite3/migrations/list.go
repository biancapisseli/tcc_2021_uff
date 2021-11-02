package migrations

func (m *migration) appendAll() error {
	if err := m.migrationInitial20211023173018(); err != nil {
		return err
	}
	return nil
}
