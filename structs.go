package main

type ServerStats struct {
	RAMUsedPercent  float64
	CPUPercent 		float64
	DiskUsedPercent float64
	TempCelsius  	float64
}

type QuizDB struct {
	question 	string
	options		[]string
	answer_id 	int
}