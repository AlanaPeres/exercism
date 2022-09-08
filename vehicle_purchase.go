package purchase
// NeedsLicense determina se uma licença é necessária para dirigir um tipo de veículo. Apenas "carro" e "caminhão" exigem uma licença.
func NeedsLicense(kind string) bool {
    return kind == "car" || kind == "truck"
	panic("NeedsLicense not implemented")
}
// ChooseVehicle recomenda um veículo para seleção. Recomenda sempre o veículo que vem primeiro na ordem lexicográfica.
func ChooseVehicle(option1, option2 string) string {
    text := " is clearly the better choice."
    if option1 < option2 {
        return option1 + text
    } else {
        return option2 + text
 }
	panic("ChooseVehicle not implemented")
}
// CalculateResellPrice calcula quanto um veículo pode revender em uma determinada idade.
func CalculateResellPrice(originalPrice, age float64) float64 {
    if age < 3 {
        return originalPrice * 0.8
    } 
	if age < 10 {
        return originalPrice * 0.7
    }
    	return originalPrice * 0.5 
	panic("CalculateResellPrice not implemented")
}