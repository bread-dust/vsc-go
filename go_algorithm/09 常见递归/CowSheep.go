package main


func Pick(){
	pick := 1
	for pick <=rest{
		if Pick(rest-pick,wait,turn)==turn{
			return turn 
		}
		if pick<=rest/4{

			pick *=pick
		}else{
			break
		}
	}
	return wait
}