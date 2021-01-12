package cerror

type DatabaseError struct {
	DBErr error
}

func (d *DatabaseError) Error() string {
	return d.DBErr.Error()
}
