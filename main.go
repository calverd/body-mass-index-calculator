package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64 //0.0
	var userKg float64     //0.0
	fmt.Println("__ Калькулятор индекса массы тела __")
	fmt.Print("Введите свой рост в сантиметрах:")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес:")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	fmt.Printf("Ваш индекс массы тела: %.0f", IMT)
	advice := getAdviceStub(IMT)
	fmt.Println(",Советы по питанию и образу жизни: ")
	fmt.Println(advice)
}

func getAdviceStub(bmi float64) string {
	if bmi < 18.5 {
		return "У вас недостаточный вес. Рекомендуется увеличить калорийность рациона, включить белки и углеводы, заниматься умеренной физической активностью."
	} else if bmi >= 18.5 && bmi < 25 {
		return "Ваш вес в норме. Поддерживайте сбалансированное питание и регулярную физическую активность."
	} else if bmi >= 25 && bmi < 30 {
		return "У вас избыточный вес. Рекомендуется снизить потребление сахара и жиров, увеличить физическую активность."
	} else {
		return "Ожирение. Обратитесь к врачу-диетологу для индивидуальной программы питания и нагрузок."
	}
}
