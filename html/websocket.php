<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
		<title>Marvin</title>

		<!-- Bootstrap -->
		<link href="css/bootstrap.min.css" rel="stylesheet">
		<link rel="stylesheet" href="css/main.css">

		<style>
			#imghere img {
				width: 100%;
			}
		</style>
		
	</head>
	<body>
			<!-- Navbar -->
			<?php 
				include_once("elements/navbar.php");
				getNavbar(3);
			?>

			<div class="container">
				 <div class="row">
						<button type="submit" class="btn btn-default" onclick="WebSocketTest();" style="margin: 0px auto;">Start Websocket Test</button>
						<div class="video-picture">
							<div class="text-center" id="img-here"></div>
							<div class="text-center" id="img-here-old"></div>
						</div>

				 </div>
				 

			</div>

		<!-- jQuery first, then Bootstrap JS. -->
		<script src="js/jquery.min.js"></script>
		<script src="js/bootstrap.min.js" i></script>
		<!-- MY SCRIPTS -->
		<script type="text/javascript">
			var time = new Date();
			var interval = 1000;
			var oldTime = new Date();

			$('#test').click(function () {

				WebSocketTest();

			});
				 function WebSocketTest()
				 {
						console.log("WebSocketTest");
						if ("WebSocket" in window)
						{
							console.log("socket in window");
							 
							 
							// Let us open a web socket
							var ws = new WebSocket("ws://wkuacm.org:9696");
								
							console.log("ws created");

							ws.onopen = function()
							{
								console.log("ws opened");
								// Web Socket is connected, send data using send()
								ws.send("Initiate Stream");
								//alert("Message is sent...");
							};
						
							ws.onmessage = function (evt) 
							{ 
								console.log("message received");
								time = new Date();
								console.log(time.getTime());
								if (time.getTime() - oldTime.getTime() > interval) {
									oldTime = new Date();
									console.log('PASS');
									var received_blob = evt.data;
									//alert("Message is received..."+received_blob);
									var img = new Image();
									img.onload = function (e) {
										console.log("PNG Loaded");
										window.URL.revokeObjectURL(img.src);    
										img = null;  
									};

									img.onerror = img.onabort = function () {         
										img = null;
									};
									img.src = window.URL.createObjectURL(received_blob);
									document.getElementById("img-here-old").innerHTML = "<img style='height:100%' src='"+img.src+"'>"

								} else {
									console.log("message NOT received");
									var received_blob = evt.data;
									//alert("Message is received..."+received_blob);
									var img = new Image();
									img.onload = function (e) {
										console.log("PNG Loaded");
										window.URL.revokeObjectURL(img.src);    
										img = null;  
									};

									img.onerror = img.onabort = function () {         
										img = null;
									};
									img.src = window.URL.createObjectURL(received_blob);
									document.getElementById("img-here").innerHTML = "<img style='height:100%' src='"+img.src+"'>"
								}
							};
						
							ws.onclose = function()
							{ 
								// websocket is closed.
								alert("Connection is closed..."); 
							};

						}

						else 

						{
							// The browser doesn't support WebSocket
							alert("WebSocket NOT supported by your Browser!");
						}
				 }

				 // window.onload = function(){
					// WebSocketTest();
				 // }
			</script>
	 </body>
</html>
