package stats

func Average(grades []float64)float64{
	if len(grades)==0{
		return 0.0
	}
	sum:=0.0
	for _,g:=range grades {
		sum+=g
	}
	return sum/float64(len(grades))
} 

func MaxGrade(grades []float64)float64{
	if len(grades)==0{
	return 0.0
	}
	maxgrade:=grades[0]
	for _,g:=range grades{
		if g>maxgrade{
			maxgrade=g
		}
	}
	return maxgrade
}