package postgres

// SoftDelete ... deleting an entity; softly
func (database Database) SoftDelete(model interface{}, condition interface{}) (int64, error) {
	result := database.DB.Where(condition).Delete(model)
	if result.Error != nil {
		return -1, result.Error
	} else if result.RowsAffected == 0 {
		return 0, nil
	}
	return result.RowsAffected, nil
}

// PermaDelete terminates an entity; with extreme prejudice
func (database Database) PermaDelete(model interface{}, condition interface{}) (int64, error) {
	result := database.DB.Unscoped().Where(condition).Delete(model)
	if result.Error != nil {
		return -1, result.Error
	} else if result.RowsAffected == 0 {
		return 0, nil
	}
	return result.RowsAffected, nil
}
