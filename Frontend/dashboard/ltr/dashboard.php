<?php 
error_reporting(0);
$server = "localhost";
$user = "securednsapi";
$pass = "rSiEN95k%6^NhKjcsScD7wsn";
$database = "securedns";

$conn = mysqli_connect($server, $user, $pass, $database);

if (!$conn) {
    die("<script>alert('Connection Failed.')</script>");
}

?>

<html dir="ltr" lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <!-- Tell the browser to be responsive to screen width -->
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="keywords"
        content="">
    <meta name="description"
        content="Dashboard for CPF DNS">
    <meta name="robots" content="noindex,nofollow">
    <title>DNS Dashboard</title>
    <link rel="canonical" href="https://www.cyberpeacefoundation.com/" />
    <!-- Favicon icon -->
    <link rel="icon" type="image/png" sizes="16x16" href="../../assets/images/favicon.png">
    <!-- Custom CSS -->
    <link href="../../assets/libs/chartist/dist/chartist.min.css" rel="stylesheet">
    <!-- Custom CSS -->
    <link href="../../dist/css/style.min.css" rel="stylesheet">
    <script>
			window.onload = function () {
			var dps = []; // dataPoints
			var chart = new CanvasJS.Chart("chartContainer", {
				title :{
					text: ""
				},
				data: [{
					type: "line",
					dataPoints: dps
				}]
				});

				var xVal = 0;
				var yVal = 100; 
				var updateInterval = 2000;
				var dataLength = 20; // number of dataPoints visible at any point

				var updateChart = function (count) {

					count = count || 1;

					for (var j = 0; j < count; j++) {
						yVal = yVal +  Math.round(5 + Math.random() *(-5-5));
						dps.push({
							x: xVal,
							y: yVal
						});
						xVal++;
					}

					if (dps.length > dataLength) {
						dps.shift();
					}

					chart.render();
				};

				updateChart(dataLength);
				setInterval(function(){updateChart()}, updateInterval);

				}
</script>
</head>

<body>
    <div class="preloader">
        <div class="lds-ripple">
            <div class="lds-pos"></div>
            <div class="lds-pos"></div>
        </div>
    </div>
    <div id="main-wrapper" data-layout="vertical" data-navbarbg="skin5" data-sidebartype="full"
        data-sidebar-position="absolute" data-header-position="absolute" data-boxed-layout="full">
        <aside class="left-sidebar" data-sidebarbg="skin6">
            <div class="scroll-sidebar">
                <nav class="sidebar-nav">
                    <ul id="sidebarnav">
                        <li>
                            <div class="user-profile d-flex no-block dropdown m-t-20">
                                <div class="user-pic"><img src="../../assets/images/users/1.jpg" alt="users"
                                        class="rounded-circle" width="40" /></div>
                                <div class="user-content hide-menu m-l-10">
                                	<?php 
                                   $result = mysqli_query($conn,"SELECT * FROM mentor where men_email='test@gmail.com'");
                                   if (mysqli_num_rows($result) > 0) 
									{
										$i=0;
									while($row = mysqli_fetch_array($result)) {
                                    ?>
                                    <a href="#" class="" id="Userdd" role="button"
                                        data-bs-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                                        <h5 class="m-b-0 user-name font-medium"><?php echo $row['men_email'];?></h5>
                                    </a>
                                </div>
                            </div>
                        </li>
                        <li class="p-10 m-t-15"><a href="javascript:void(0)"
                                class="btn d-block w-100 create-btn text-white no-block d-flex align-items-center"><i
                                    class="fa fa-plus-square"></i> <span class="hide-menu m-l-5">More Details</span> </a>
                        </li>
                        <li class="p-10 m-t-15">
                                <h5 class="hide-menu m-l-5">Created at :</h5>
                                <label class="label label-success" style='font-size:15px;'><?php echo $row['created_at'];?></label>
                        </li>
			<li class="p-10 m-t-15">
                                <h5 class="hide-menu m-l-5">Client Downloads</h5>
                                <h5>-----------------------------</h5>
                                <h5 class="hide-menu m-l-5">Download for Windows</h5>
                                <a class="hide-menu m-l-5" href="downloads/client.exe">window_client</a>
                                <h5 class="hide-menu m-l-5">Download for Linux</h5>
                                <a class="hide-menu m-l-5" href="downloads/setup.sh">linux_client</a>
                        </li>
                        <?php
						$i++;
							}
						}
						?>
                    </ul>

                </nav>
            </div>
        </aside>
        <div class="page-wrapper">
            <div class="page-breadcrumb">
                <div class="row align-items-center">
                    <div class="col-5">
                        <h4 class="page-title">Dashboard</h4>
                        <div class="d-flex align-items-center">
                            <nav aria-label="breadcrumb">
                                <ol class="breadcrumb">
                                    <li class="breadcrumb-item"><a href="#">Home</a></li>
                                    <li class="breadcrumb-item"><a href="#">Admin</a></li>
                                    <li class="breadcrumb-item"><a href="http://101.53.147.32:8000/en-GB/app/search/search?q=search%20*&sid=1618046567.63&display.page.search.mode=smart&dispatch.sample_ratio=1&workload_pool=&earliest=-24h%40h&latest=now">Live View of Analytics</a></li>
                                </ol>
                            </nav>
                        </div>
                    </div>
                </div>
            </div>
            <div class="container-fluid">
                <div class="row">
                    <div class="col-md-8">
                        <div class="card">
                            <div class="card-body">
                                <div class="d-md-flex align-items-center">
                                    <div>
                                        <h4 class="card-title">Website Status</h4>
                                        <h5 class="card-subtitle">Site Status</h5>
                                    </div>
                                    <div class="ms-auto d-flex no-block align-items-center">
                                        <ul class="list-inline font-12 dl m-r-15 m-b-0">
                                            <li class="list-inline-item text-primary"><i class="fa fa-circle"></i>Normal
                                            </li>
                                        </ul>
                                    </div>
                                </div>
                                <div class="row">
                                    <div class="col-lg-12">
                                        <!--<div class="campaign ct-charts"></div>-->
                                        <div id="chartContainer" style="height: 300px; width: 100%;"></div>

                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>


                    <div class="col-md-4">
                        <div class="card">
                            <div class="card-body">
                                <h4 class="card-title">Blocked URL</h4>
                                <div class="feed-widget">
                                    <ul class="list-style-none feed-body m-0 p-b-20">
                                   <?php 
                                   $result = mysqli_query($conn,"SELECT * FROM mentee");
                                   if (mysqli_num_rows($result) > 0) 
									{
										$i=0;
									while($row = mysqli_fetch_array($result)) {
                                    ?>
                                        <li class="feed-item">
                                            <div class="feed-icon bg-danger"></div><?php echo $row['blacklist'];?>
                                        </li>
                                         <?php
									$i++;
									}
									}
								  ?>
                                    </ul>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="row">
                    <div class="col-12">
                        <div class="card">
                            <div class="card-body">
                                <div class="d-md-flex">
                                    <div>
                                        <h4 class="card-title">Visited URL</h4>
                                        <h5 class="card-subtitle"></h5>
                                    </div>
                                    <div class="ms-auto">
                                        <div class="dl">
                                            <select class="form-select shadow-none">
                                                <option value="0" selected>Monthly</option>
                                                <option value="1">Daily</option>
                                                <option value="2">Weekly</option>
                                                <option value="3">Yearly</option>
                                            </select>
                                        </div>
                                    </div>
                                </div>
                                <!-- title -->
                            </div>
                            <div class="table-responsive">
                                <table class="table v-middle">
                                    <thead>
                                        <tr class="bg-light">
                                            <th class="border-top-0">Name</th>
                                            <th class="border-top-0">Email</th>
                                            <th class="border-top-0">Age Group</th>
                                            <th class="border-top-0">MAC</th>
                                            <th class="border-top-0">Whitelist</th>
                                            <th class="border-top-0">Blacklist</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                    	<?php 
                                    		$result = mysqli_query($conn,"SELECT * FROM mentee");
                                    		if (mysqli_num_rows($result) > 0) 
											{
												$i=0;
												while($row = mysqli_fetch_array($result)) {
                                    	?>
                                        <tr>
                                            <td>
                                                <div class="d-flex align-items-center">
                                                    <div class="">
                                                        <h4 class="m-b-0 font-16"><?php echo $row['name'];?></h4>
                                                    </div>
                                                </div>
                                            </td>
                                            <td>
                                                <label class="label label-success" style='font-size:15px;'><?php echo $row['email'];?></label>
                                            </td>
                                            <td>
                                            	<div class="d-flex align-items-center">
                                                    <div class="">
                                                        <h4 class="m-b-0 font-16"><?php echo $row['age_grp'];?></h4>
                                                    </div>
                                                </div>
                                            </td>
                                            <td>
                                                <h5 class="label label-success"style='font-size:15px;'><?php echo $row['mac'];?></h5>
                                            </td>
                                            <td>
                                                <h5 class="m-b-0 font-16"><?php echo $row['whitelist'];?></h5>
                                            </td>
                                            <td>
                                                <h5 class="m-b-0 font-16"><?php echo $row['blacklist'];?></h5>
                                            </td>
                                        </tr>
                                        <?php
											$i++;
										}
										?>
                                    </tbody>
                                </table>
                            </div>
                        </div>
                    </div>
                </div>
            <footer class="footer text-center">
                All Rights Reserved by <a
                    href="https://www.cyberpeacefoundation.com">CPF</a>.
            </footer>
        </div>
    </div>
    <script src="../../assets/libs/jquery/dist/jquery.min.js"></script>
    <script src="../../assets/libs/bootstrap/dist/js/bootstrap.bundle.min.js"></script>
    <script src="../../dist/js/app-style-switcher.js"></script>
    <script src="../../dist/js/waves.js"></script>
    <script src="../../dist/js/sidebarmenu.js"></script>
    <script src="../../dist/js/custom.js"></script>
    <script src="../../assets/libs/chartist/dist/chartist.min.js"></script>
    <script src="../../assets/libs/chartist-plugin-tooltips/dist/chartist-plugin-tooltip.min.js"></script>
    <script src="../../dist/js/pages/dashboards/dashboard1.js"></script>
    <script src="canvasjs.min.js"></script>
</body>
</html>
<?php

}
?>
