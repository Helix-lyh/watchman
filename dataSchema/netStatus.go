package dataSchema

type NetStatus struct {
	TotalLink int64 `json:"totalLink"` // 总连接数  netstat -n | awk '/^tcp/ {n=split($(NF-1),array,":");if(n<=2)++S[array[(1)]];else++S[array[(4)]];++s[$NF];++N} END {for(a in S){printf("%-20s %s\n", a, S[a]);++I}printf("%-20s %s\n","TOTAL_IP",I);for(a in s) printf("%-20s %s\n",a, s[a]);printf("%-20s %s\n","TOTAL_LINK",N);}'
	LastAck int64  `json:"lastAck"` //正在等待处理的请求数  netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'
	Established int64  `json:"established"`  //正常数据传输状态
	TimeWait int64  `json:"timeWait"`  //处理完毕，等待超时结束的请求数
	SynRecv int64 `json:"synRecv"`  //正在等待处理的请求数
	TopLinkIP string `json:"topLinkIP"` //连接数最大的IP
	Port80Link int64 `json:"port80Link"` //80端口连接数 //netstat -nat | grep -i "80" | wc -l
}