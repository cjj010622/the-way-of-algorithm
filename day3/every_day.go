package day3

// 计算列车到站时间
func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}
