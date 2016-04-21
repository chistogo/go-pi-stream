<?php 
function getNavbar($selection){

?>
<nav class="navbar navbar-default remove-rc">
  <div class="container-fluid">
    <!-- Brand and toggle get grouped for better mobile display -->
    <div class="navbar-header">
      <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
        <span class="sr-only">Toggle navigation</span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <a class="navbar-brand" href="index.php">M.A.R.V.I.N</a>
    </div>

    <!-- Collect the nav links, forms, and other content for toggling -->
    <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
      <ul class="nav navbar-nav">
        <li <?php if($selection =='1'){?>class="active"<?php } ?>><a href="index.php">Home <span class="sr-only">(current)</span></a></li>
        <li <?php if($selection =='2'){?>class="active"<?php } ?>><a href="#">Fullscreen Video</a></li>
        <li <?php if($selection =='3'){?>class="active"<?php } ?>><a href="websocket.php">Websocket Test</a></li>
        <li <?php if($selection =='4'){?>class="active"<?php } ?>><a href="#">About Marvin</a></li>
      </ul>
  </div><!-- /.container-fluid -->
</nav>

<?php
  # code...
}

?>