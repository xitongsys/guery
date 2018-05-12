function Dashboard(id, title) {
	dashboard = document.getElementById(id);
	this.title = title;
	this.id = id;
	this.canvas= document.getElementById(this.id + "_canvas");
	this.titleObj = document.getElementById(this.id + "_title");
	this.width = this.canvas.width;
	this.height = this.canvas.height;

	this.Plot = function (data, number){
		var ctx=this.canvas.getContext('2d');
		ctx.clearRect(0, 0, this.width, this.height);
		this.PlotLine(data, this.width, this.height, 0, 0);
		this.PlotTitle(this.title, number)
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
		this.canvas.style.backgroundColor="#292d33";
		var ctx=this.canvas.getContext('2d');
		ctx.strokeStyle="#404040";
		ctx.lineWidth=1;
		ctx.beginPath();
		ctx.moveTo(xOffset,yOffset);
		ctx.lineTo(xOffset, h);
		ctx.lineTo(w,h);
		ctx.stroke();
		
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
	}
	this.PlotTitle = function(text, number){
		this.titleObj.innerHTML=text + ": " + number;
		this.titleObj.style.color="#ffffff";
	}
}

