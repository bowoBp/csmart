package health

type versionGetterRepository struct {
	// TODO add field like database
	//db *gorm.DB
}

type VersionGetterRepository interface {
	GetVersion() string
}

// GetVersion returns the version of the database.
// It returns an error if there was a problem executing the SQL query.
func (r versionGetterRepository) GetVersion() string {
	//TODO version of database
	var version string
	version = "this version of db"

	return version
}
