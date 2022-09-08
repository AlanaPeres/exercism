package cars

// CalculateWorkingCarsPerHour calcula quantos carros de trabalho são
// produzido pela linha de montagem a cada hora.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64 (productionRate) * successRate / 100
	panic("CalculateWorkingCarsPerHour not implemented")
}
// CalculateWorkingCarsPerMinute calcula quantos carros de trabalho são
// produzido pela linha de montagem a cada minuto.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
   return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60.0)
	panic("CalculateWorkingCarsPerMinute not implemented")
}
// CalculateCost calcula o custo de produção de um determinado número de carros.
func CalculateCost(carsCount int) uint {
    return uint((carsCount%10)*10000 + (carsCount/10)*95000)
	panic("CalculateCost not implemented")
}
