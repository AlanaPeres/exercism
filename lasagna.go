package lasagna

// define the 'OvenTime' constant
const ( OvenTime  int  =  40 // todo o preparo e cozimento duram 40 minutos.
		tempoPorCamadas int = 2 //cada camada necessita de 2 minutos de preparo
		numberfLayers int = 4 // o número de fatias da 
)
// RemainingOvenTime retorna os minutos restantes com base nos minutos `reais` já no forno.
func RemainingOvenTime(actualMinutesInOven int) int {
    return OvenTime - actualMinutesInOven
	panic("RemainingOvenTime not implemented")
}
// PreparationTime calcula o tempo necessário para preparar a lasanha com base na quantidade de camadas.
func PreparationTime(numberfLayers int) int {
    return numberfLayers * tempoPorCamadas
	panic("PreparationTime not implemented")
}
//ElapsedTime calcula o tempo decorrido para cozinhar a lasanha. Este tempo inclui o tempo de preparo e o tempo que a lasanha está assando no forno.
func ElapsedTime(numberfLayers, actualMinutesInOven int) int {
 	return PreparationTime(numberfLayers) + actualMinutesInOven
	panic("ElapsedTime not implemented")