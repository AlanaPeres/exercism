package annalyn

var knightIsAwake = false
var archerIsAwake = true
var prisonerIsAwake = false
var petDogIsPresent = false

//CanFastAttack pode ser executado apenas quando o cavaleiro está dormindo	
func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake == true {
    	return false
	} else {
    	return true
	}
	panic("Please implement the CanFastAttack() function")
}
 	
// CanSpy pode ser executado se pelo menos um dos personagens estiver acordado
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
        if knightIsAwake == true{
            return true
		}
        if archerIsAwake == true{
    		return true
        }
        if prisonerIsAwake == true{
    		return true
        } else {
        	return false
        }
	panic("Please implement the CanSpy() function")
}
// CanSignalPrisoner pode ser executado se o prisioneiro estiver acordado e o arqueiro estiver dormindo.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
    if archerIsAwake == false && prisonerIsAwake == true {
        return true
    } else {
    	return false
    }
	panic("Please implement the CanSignalPrisoner() function")
}
// CanFreePrisoner pode ser executado se o prisioneiro estiver acordado e os outros 2 personagens estiverem dormindo ou se o cachorro de estimação de Annalyn estiver com ela e o arqueiro estiver dormindo.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool{
	return (prisonerIsAwake && !knightIsAwake && !archerIsAwake) || (!archerIsAwake && petDogIsPresent)
   
	panic("Please implement the CanFreePrisoner() function")
}
