let userid;
  let color;

  let ws;
  if (window.WebSocket === undefined) {
      alert("Your browser does not support WebSockets");
  } else {
      ws = initWS();
  }

  function initWS() {
      var socket = new WebSocket("ws://localhost:6000/ws"),
          container = $("#chatTemplate")
      socket.onopen = function() {
          console.log("Socket is open");
      };

      socket.onmessage = function (e) {
        var wsJs = JSON.parse(e.data);            
        container.append("<p style='color:"+ wsJs.color +"'>"+ wsJs.username +" : " + wsJs.message + "</p>");
      }
      socket.onclose = function () {
          console.log("Socket closed");
      }
      return socket;
  }

    prompUserId();
    $("#sendMsg").on("click",function(){
      $("#chatTemplate").append("<p style='color:"+ color +";'>"+ userid +" : " + $("#message").val() + "</p>");
     
        ws.send(
            JSON.stringify({
                Username : userid,
                Message : $("#message").val(),
                Color : color
            }
        ));
        $("#message").val("");        
    });



function prompUserId(){    
    var name= prompt("Please enter your name","username");
    if (name!=null){
       var back = ["black","gray","#333","blue","red","pink","orange"];
       var rand = back[Math.floor(Math.random() * back.length)];
       color = rand;
       userid = name;
   }
}

$('#message').on("keypress", function(e) {
        if (e.keyCode == 13) {
            $("#sendMsg").trigger("click");            
        }
});

