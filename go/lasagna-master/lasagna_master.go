package lasagna

// define the 'PreparationTime()' function
func PreparationTime(layers []string, averageTime int) int {
  if(averageTime == 0) {
    averageTime = 2
  }

  return len(layers) * averageTime
}

// define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64){

  for _, layer := range layers {
    switch layer {
    case "sauce":
      sauce += 0.2    
    case "noodles":
      noodles += 50
    }
  }
  return
}

// define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
  myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, potions int) []float64 {
  scaledQuantities := make([]float64, len(quantities))

  for index, quantity := range quantities {
    scaledQuantities[index] = quantity /2 * float64(potions)
  }

  return scaledQuantities
}
