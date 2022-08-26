func (t *Test) Shoot() bool {
	if(t.On && t.Ammo > 0){
		t.Ammo--
		return true
	}
	return false
}

func (t *Test) RideBike() bool {
	if(t.On && t.Power > 0){
		t.Power--
		return true
	}
	return false
}

type Test struct {
	On bool
	Ammo, Power int
}

func main() {

testStruct := &Test{}
/*
 * Экземпляр созданной вами структуры необходимо передать в качестве
 * аргумента функции testStruct, которая выполнит проверку соблюдения
 * всех условий задания/
 */
// }