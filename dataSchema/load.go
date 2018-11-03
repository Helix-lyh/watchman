package dataSchema

type Load struct {
	OneMin float64 `json:"OneMin"`  //一分钟平均负载
	FiveMin float64 `json:"fiveMin"`  //五分钟平均负载
	FifteenMin float64 `json:"fifteenMin"`  //十五分钟平均负载
}

