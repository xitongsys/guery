function Dashboard(id, title) {
	dashboard = document.getElementById(id);
	this.title = title;
	this.id = id;
	this.canvas= document.getElementById(this.id + "_canvas");
	this.titleObj = document.getElementById(this.id + "_stat");


	this.Plot = function (data, number){
		this.PlotTitle(this.title, number)
		this.Resize();
		var ctx=this.canvas.getContext('2d');
		this.canvas.height = $(this.canvas).height();
		this.canvas.width = $(this.canvas).width();
		
		this.height = this.canvas.height;
		this.width = this.canvas.width;
		
		ctx.clearRect(0, 0, this.width, this.height);
		this.PlotLine(data, this.width, this.height, 0, 5);
	}
	
	this.Resize = function(){
		$(this.canvas).width($("#" + this.id).width() - $(this.titleObj).width());
		var par = $(this.canvas).parent();
		$(this.canvas).height(par.height()-7);
	}
	
	this.PlotLine = function(data, w, h, xOffset, yOffset){
		ld = data.length;
		maxd = 0;
		for(i=0; i<ld; i++){
			if(maxd<data[i]){
				maxd=data[i];
			}
		}
		dx = (w - xOffset) * 1.0 / ld;
		if(maxd<=0){
			dy=0
		}else{
			dy = (h - yOffset)*1.0 / maxd;
		}

		//this.canvas.style.backgroundColor="#fffffff";
		//this.canvas.style.backgroundColor="#292d33";

		var ctx=this.canvas.getContext('2d');
		
		/*
		ctx.strokeStyle="#424242";
		ctx.lineWidth=1;
		ctx.beginPath();
		//ctx.moveTo(xOffset,yOffset);
		//ctx.lineTo(xOffset, h);
		ctx.moveTo(xOffset, h);
		ctx.lineTo(w,h);
		ctx.stroke();
		*/
		
		ctx.strokeStyle="#AEBAED";
		ctx.fillStyle="#404752";
		ctx.lineWidth=2;
		ctx.beginPath();
		ctx.moveTo(xOffset, h);
		for(i=0; i<ld; i++){
			x = i * dx + xOffset;
			y = h - dy * data[i];
			ctx.lineTo(x, y);
		}
		ctx.lineTo(w, h); ctx.lineTo(xOffset, h); 
		ctx.stroke();
		ctx.fill();

		x = (ld-1) * dx + xOffset;
		y = h - dy * data[ld-1] - 3;
		if(x<0)x=0; if(y<0)y=0;

		ctx.beginPath();
		ctx.arc(x,y, 2, 0, 2*Math.PI, false);
		ctx.fillStyle="#00c3df";
		ctx.fill();
	}
	this.PlotTitle = function(text, number){
		res="" + number
		if(number>=1e9){
			number = number/1000000000.0
			res = ("" + number).substr(0,4) + "G";
		}else if(number>=1000000){
			number = number/1000000.0
			res = ("" + number).substr(0,4) + "M";
		}else if (number >= 1000) {
			number = number/1000.0;
			res = ("" + number).substr(0,4) + "K";
		}
		
		this.titleObj.innerHTML=res;
	}
}

