package structs

// DateTime YYYY-MM-DDThh:mm:ss.sTZD
type DateTime struct {
	Year       int
	Month      int
	Day        int
	Hour       int
	Minute     int
	Second     int
	Milisecond int
	Gmt        int
}
