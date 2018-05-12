function Dashboard(id, title) {
	dashboard = document.getElementById(id);
	this.title = title;
	this.id = id;
	this.canvas= document.getElementById(this.id + "_canvas");
	this.titleObj = document.getElementById(this.id + "_title");
	this.width = this.canvas.width;
	this.height = this.canvas.height;
	this.charWidth = 50;

	this.Plot = function (data, number){
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
		dy = (h * 1.0 - yOffset - 10.0) / maxd;

		this.canvas.style.backgroundColor="#fffffff";
		//canvas.style.backgroundColor="#292C33";			 
		var ctx=this.canvas.getContext('2d');
		ctx.strokeStyle="#505762";
		ctx.fillStyle="#404752";
		ctx.lineWidth=1;
		ctx.beginPath();
		ctx.moveTo(xOffset, h);
		for(i=0; i<ld; i++){
			x = i * dx + xOffset;
			y = h - dy * data[i] - 5 + yOffset;
			ctx.lineTo(x, y);
		}
		ctx.lineTo(w, h); ctx.lineTo(xOffset, h); 
		ctx.stroke();
		ctx.fill();
	}
	this.PlotTitle = function(text, number){
		this.titleObj.innerHTML=text + ": " + number;
	}
}

