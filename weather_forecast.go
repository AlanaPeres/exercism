//Package weather prevê as condições climáticas atuais nas cidades de Goblinocus.
package weather

//CurrentCondition armazena a condição do clima atual. 
var CurrentCondition string

//CurrentLocation armazena a localização da cidade que está recebendo a previsão do tempo. 
var CurrentLocation string

//Forecast função retorna a cidade e a suas condições climáticas. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
