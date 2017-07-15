
document.addEventListener("DOMContentLoaded", function(e){
  var websocket = new WebSocket("ws://localhost:3000/api/websocket")
  websocket.onopen = function(e) {
    isOpen = true
  }

  var viewport = document.getElementById("viewport")
  var point = [], mouse = {}, dif= {}, isOpen = false, isDrawn = false
  var body = HTMLElement, html = HTMLElement
  var height = 0
  // Event listener for click on plain
  viewport.addEventListener("click", function(e){
    // var bounds = getCoords(e.target)
    point = [e.pageX-this.offsetLeft, e.pageY - this.offsetTop]
    // body = document.body
    // html = document.documentElement
    // height = Math.max( body.scrollHeight, body.offsetHeight, html.clientHeight, html.scrollHeight, html.offsetHeight )
    // dif.Y = html.offsetHeight - viewport.offsetHeight
    // dif.X =  html.offsetWidth - viewport.offsetWidth
    // mouse.Y = e.clientY
    // mouse.X = e.clientX
    // console.log("Dif of point mouse.X - dif.X: ", parseFloat(mouse.X) - parseFloat(dif.X), "\nDif of point mouse.Y - dif.Y: ", parseFloat(mouse.Y) - parseFloat(dif.Y), "\nHeight: ", height, "\nDif.X: ", dif.X, "\nDif.Y: ", dif.Y, "\nMouse.X: ", mouse.X, "\nMouse.Y", mouse.Y)
    // console.log("Diff y", parseFloat(dif.Y))
    // point = [parseFloat(mouse.X - dif.X), (parseFloat(mouse.Y % dif.Y))]
    console.log(point[0], point[1])
    if (isOpen) {
      websocket.send(JSON.stringify({"api":"neuron", "data":{"x":point[0], "y": point[1]},}))
    }
  })
  
  // canvas stuff
  var canvas = loadCanvas(viewport)
  viewport.appendChild(canvas)
  ctx = canvas.getContext('2d')
  // websocket information train

  websocket.onmessage = function(e) {
    var data = JSON.parse(e.data)
    console.log(data)
    points = [
      [0, data.B],
      [canvas.width, data.M*canvas.width+data.B]
    ]
    if (!isDrawn) {
      ctx.beginPath()
      isDrawn = true
      ctx.moveTo(points[0][0],points[0][1])
      ctx.lineTo(points[1][0],points[1][1])
      ctx.stroke()
    }
    console.log("raw data.output: ", data.output)

    if (parseFloat(data.output) == 1 ) {
      console.log("red")
      ctx.fillStyle = "#FF0000"
    }
    if (parseFloat(data.output) == 0 ) {
      console.log("green")
      ctx.fillStyle = "#00FF00"
    } 
    console.log("someshit", point[0], point[1])
    ctx.beginPath()
    ctx.arc(point[0], point[1], 5, 0, 2*Math.PI)
    ctx.closePath()
    ctx.fill()
  }
})


function loadCanvas(viewport) {
  canvas = document.createElement('canvas')
  canvas.width = viewport.offsetWidth
  canvas.height = viewport.offsetHeight
  return canvas
}
