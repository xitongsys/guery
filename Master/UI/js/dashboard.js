function Dashboard(id, title) {
	dashboard = document.getElementById(id);
	this.title = title;
	this.id = id;
	this.canvas= document.getElementById(this.id + "_canvas");
	this.titleObj = document.getElementById(this.id + "_stat");


	this.Plot = function (data, number){
		this.Resize();
		var ctx=this.canvas.getContext('2d');
		this.height = this.canvas.height;
		this.width = this.canvas.width;
		
		ctx.clearRect(0, 0, this.width, this.height);
		this.PlotLine(data, this.width, this.height, 0, 0);
		this.PlotTitle(this.title, number)
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
		dy = (h * 1.0 - yOffset) / maxd;

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
		
		ctx.strokeStyle="#5E6A7D";
		ctx.fillStyle="#404752";
		ctx.lineWidth=1;
		ctx.beginPath();
		ctx.moveTo(xOffset, h);
		for(i=0; i<ld; i++){
			x = i * dx + xOffset;
			y = h - dy * data[i] + yOffset;
			ctx.lineTo(x, y);
		}
		ctx.lineTo(w, h); ctx.lineTo(xOffset, h); 
		ctx.stroke();
		ctx.fill();

		x = (ld-1) * dx + xOffset;
		y = h - dy * data[ld-1] + yOffset - 3;
		if(x<0)x=0; if(y<0)y=0;
		ctx.fillStyle="#00c3df";
		ctx.fillRect(x, y, 3, 3);
	}
	this.PlotTitle = function(text, number){
		res="" + number
		if(number>=1e9){
			number = number/1000000000.0
			res = ("" + number).substr(0,3) + "G";
		}else if(number>=1000000){
			number = number/1000000.0
			res = ("" + number).substr(0,3) + "M";
		}else if (number >= 1000) {
			number = number/1000.0;
			res = ("" + number).substr(0,3) + "K";
		}
		
		this.titleObj.innerHTML=res;
	}
}

