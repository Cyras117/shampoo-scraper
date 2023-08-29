package handdlers

/*
Respose struct
*/
type response struct {
	method     string
	errMessage string
	resp       string
}

/*
Respose struct to handle the system reponses
*/
type responsesStruct struct {
	status    string
	responses []response
}
