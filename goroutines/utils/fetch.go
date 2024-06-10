package utils

func FetchData(ch chan<- []DbResult, service DbService) {
	ch <- service.GetDbResult()
}
