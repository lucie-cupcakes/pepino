package pepinoservice

// PutEntry automatically loads the database/or Creates it
// ands stores the given entry Name & Value
func (svc *DatabaseService) PutEntry(dbName string, entryName string,
	entryValue []byte) error {
	dbPtr, err := svc.loadDatabaseToMemory(dbName, true)
	if err != nil {
		return err
	}

	dbPtr.Entries[entryName] = entryValue
	err = dbPtr.Save()
	if err != nil {
		return err
	}
	return nil
}
