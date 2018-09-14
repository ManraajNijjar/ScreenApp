
var canvas = document.getElementById("draw");


var ctxCanvas = canvas.getContext("2d");
resize();


function resize() {
  ctxCanvas.canvas.width = window.innerWidth;
  ctxCanvas.canvas.height = window.innerHeight;
  drawGrid();
}


window.addEventListener("resize", resize);
document.addEventListener("mousemove", draw);
document.addEventListener("mousedown", setPosition);
document.addEventListener("mouseenter", setPosition);

var pos = { x: 0, y: 0 };


function setPosition(e) {
  pos.x = e.clientX;
  pos.y = e.clientY;
}

function draw(e) {
  if (e.buttons !== 1) return;

  var color = document.getElementById("hex").value;
  console.log(color);
  ctxCanvas.beginPath();

  ctxCanvas.lineWidth = 5;
  ctxCanvas.lineCap = "round";
  ctxCanvas.strokeStyle = color;

  ctxCanvas.moveTo(pos.x, pos.y);
  setPosition(e);
  ctxCanvas.lineTo(pos.x, pos.y);

  ctxCanvas.stroke();
}

function drawGrid() {
    var bw = canvas.width
    var bh = canvas.height
    var p = 10;
    var gridSize = 40
    
    for (var x = 0; x <= bw; x += gridSize) {
        ctxCanvas.moveTo(0.5 + x + p, p);
        ctxCanvas.lineTo(0.5 + x + p, bh + p);
    }

    for (var x = 0; x <= bh; x += gridSize) {
        ctxCanvas.moveTo(p, 0.5 + x + p);
        ctxCanvas.lineTo(bw + p, 0.5 + x + p);
    }

    ctxCanvas.strokeStyle = "black";
    ctxCanvas.stroke();
}