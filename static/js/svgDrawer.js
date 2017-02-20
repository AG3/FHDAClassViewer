var WIDTH
var HEIGHT
var s
var COL_WIDTH
var ROW_HEIGHT
var ClassEleSet = Snap.set()

function drawTable(){
	for(i=0;i<6;i++){
		s.rect(i*COL_WIDTH,0 , 1,HEIGHT).attr({fill:"#AAAAAA"});
	}
	ROW_HEIGHT = (HEIGHT-1)/14
	for(i=0;i<15;i++){
		s.rect(0, i*ROW_HEIGHT, WIDTH, 1).attr({fill:"#AAAAAA"});
	}
}

function AddClass(data){
	weekday=[]
	week = data["days"]
	for(i=0;i<week.length;i++){
		if(week[i]=='M')
			weekday.push(0)
		if(week[i]=='T' && week[i+1]!='h')
			weekday.push(1)
		if(week[i]=='W')
			weekday.push(2)
		if((week[i]=='T'&& week[i+1]=='h') || week[i]=='R')
			weekday.push(3)
		if(week[i]=='F')
			weekday.push(4)
	}
	TOTAL_MINUTES = 14*60

	s_time = data["start_time"]
	e_time = data["end_time"]

	s_h = s_time.split(':')[0]
	s_m = s_time.split(':')[1]
	if(s_h=="24")
		s_h="12"
	s_h = parseInt(s_h)
	s_m = parseInt(s_m)

	s_tm = (s_h-7)*60+s_m

	e_h = e_time.split(':')[0]
	e_m = e_time.split(':')[1]
	if(e_h=="24")
		e_h="12"
	e_h = parseInt(e_h)
	e_m = parseInt(e_m)
	e_tm = (e_h-7)*60+e_m

	s_y = (s_tm/TOTAL_MINUTES)*HEIGHT
	e_y = (e_tm/TOTAL_MINUTES)*HEIGHT
	console.log(s_time, e_time, s_tm, e_tm, s_y, e_y)
	for(i=0;i<weekday.length;i++){
		t = s.rect(weekday[i]*COL_WIDTH,s_y,COL_WIDTH,e_y-s_y).attr({fill:"red", opacity:"0.2"})
		t.click(function(){
			this.test = 1
			console.log(this.test)
		})
		ClassEleSet.push(t)
	}
}

function Init() {
	s = Snap("svg")
	WIDTH = $('svg').width()
	HEIGHT = $('svg').height()
	console.log(WIDTH,HEIGHT)
	COL_WIDTH = (WIDTH-1) / 5
	drawTable()
}
