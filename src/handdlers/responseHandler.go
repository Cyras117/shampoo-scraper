package handdlers

type response struct {
	method     string
	errMessage string
	resp       string
}

type responsesStruct struct {
	status    string
	responses []response
}
