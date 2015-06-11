package pages

type HomePage struct {
}

func (p *HomePage) GetContent() []byte {
	return []byte(p.content())
}

func (p *HomePage) content() string {
	return `
        <html>
            <head>
                <title>Betwixt</title>
                <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css">
                <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap-theme.min.css">
                <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/js/bootstrap.min.js"></script>
                <script src="https://ajax.googleapis.com/ajax/libs/angularjs/1.3.14/angular.min.js"></script>
                <script>

                </script>
            </head>
            <body role="document" ng-app="betwixt">
                <!-- Fixed navbar -->
                <nav class="navbar navbar-inverse navbar-fixed-top">
                  <div class="container">
                    <div class="navbar-header">
                      <a class="navbar-brand" href="#">Betwixt LWM2M Server</a>
                    </div>
                    <div id="navbar" class="navbar-collapse collapse">
                      <ul class="nav navbar-nav">
                        <li class="active"><a href="#">Home</a></li>
                      </ul>
                    </div><!--/.nav-collapse -->
                  </div>
                </nav>

                <div class="container theme-showcase" role="main">
                    <br /><br /><br />

                <div class="row" style="text-align: center">
                    <div class="col-sm-2">&nbsp;</div>

                    <div class="col-sm-2" ng-model="MemUsage">
                        <div class="panel panel-primary">
                            <div class="panel-heading">
                                <h3 class="panel-title">Memory</h3>
                            </div>
                            <div class="panel-body">
                              <h1>{{.MemUsage}} KB</h1>
                            </div>
                        </div>
                    </div>

                    <div class="col-sm-2">
                        <div class="panel panel-info">
                            <div class="panel-heading">
                                <h3 class="panel-title">Clients</h3>
                            </div>
                            <div class="panel-body">
                              <h1>{{.ClientsCount}}</h1>
                            </div>
                        </div>
                    </div>

                    <div class="col-sm-2">
                        <div class="panel panel-success">
                            <div class="panel-heading">
                                <h3 class="panel-title">Requests</h3>
                            </div>
                            <div class="panel-body">
                              <h1>{{.RequestCount}}</h1>
                            </div>
                        </div>
                    </div>

                    <div class="col-sm-2">
                        <div class="panel panel-danger">
                            <div class="panel-heading">
                                <h3 class="panel-title">Errors</h3>
                            </div>
                            <div class="panel-body">
                              <h1>{{.ErrorsCount}}</h1>
                            </div>
                        </div>
                    </div>

                    <div class="col-sm-2">&nbsp;</div>
                  </div>

                  <div class="page-header">
                    <h3>Registered Clients</h3>
                  </div>
                  <div class="row">
                    <div class="col-md-12">
                      <table class="table">
                        <thead>
                          <tr>
                            <th>Endpoint</th>
                            <th>Registration ID</th>
                            <th>Registration Date</th>
                            <th>Last Update</th>
                            <th>Actions</th>
                          </tr>
                        </thead>
                        <tbody>

                          {{ range .Clients }}
                          <tr>
                            <td><a href="/client/{{.Endpoint}}/view">{{.Endpoint}}</a></td>
                            <td>{{.RegistrationID}}</td>
                            <td>{{.RegistrationDate}}</td>
                            <td>{{.LastUpdate}}</td>
                            <td>
                              <h4>
                                <button type="button" class="btn btn-xs btn-info"><a href="/client/{{.Endpoint}}/view">view</a></button>
                                <button type="button" class="btn btn-xs btn-danger"><a href="/client/{{.Endpoint}}/delete">delete</a></button>
                              </h4>
                            </td>
                          </tr>
                          {{ end }}

                        </tbody>
                      </table>
                    </div>

                  </div>

                </div> <!-- /container -->

                <!-- Bootstrap core JavaScript
                ================================================== -->
                <!-- Placed at the end of the document so the pages load faster -->
                <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
            </body>
        </html>
    `
}
