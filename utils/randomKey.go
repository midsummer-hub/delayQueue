package utils

// SnowflakeID TODO
func SnowflakeID() (string, error) {
	id, err := NewWorker(1)
	if err != nil {
		return "", err
	}
	return id.GetID().String(), err
}
