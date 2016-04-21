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

    
  </head>
  <body>
   	<!-- Navbar -->
   	<?php include_once("elements/navbar.php");
    getNavbar('1');
    ?>

   	<div class="container">
   		<div class="row">
   			<div class="col-md-6">
   				<img src="img/cam.jpg" style="width:100%">
   			</div>
   			<div class="col-md-6">
   				<h1>Sensor Data</h1>
   			</div>
   		</div>
   		<div class="row">
   			<div class="col-md-6">
   				<h1>Command Line</h1>
   			</div>
   			<div class="col-md-6">
   				<h1>Graphs</h1>
   			</div>
   		</div>

   	</div>

    <!-- jQuery first, then Bootstrap JS. -->
    <script src="js/jquery.min.js"></script>
    <script src="js/bootstrap.min.js" i></script>
    <!-- MY SCRIPTS -->

  </body>
</html>
