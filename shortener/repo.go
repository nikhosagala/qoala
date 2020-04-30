package shortener

// findById return QoalaShortener from our data container
func findById(ID int64) QoalaShortener {
	for _, qs := range containers {
		if qs.id == ID {
			return qs
		}
	}
	return QoalaShortener{}
}
