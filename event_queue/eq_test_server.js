var http = require('http'),
	url = require('url')
    httpProxy = require('http-proxy');

//
// Create a proxy server with custom application logic
//
httpProxy.createServer(function (req, res, proxy) {
  //
  // Put your custom server logic here
  //
  proxy.proxyRequest(req, res, {
    host: 'localhost',
    port: 9988
  });
}).listen(8000);

request_queue = []
request_attempt = 0

http.createServer(function (req, res) {

	console.log(req.url);
	url_parts = url.parse(req.url)
	console.log(url_parts.pathname)
	switch (url_parts.pathname) {
		case '/get':
			if (req.method == 'GET'){
				console.log('receive get call from ' + req.url)
				console.log(request_queue)
				res.write(JSON.stringify(request_queue));
				res.end();
			}
			break;
		case '/post':
			if (req.method == 'POST'){
				body = ''
				req.on('data', function(chunk) {
      				console.log("Received body data:");
				    console.log(chunk.toString());
				    body += chunk
				});
    
    			req.on('end', function() {
    				console.log('receive post call from ' + req.url)
      				// empty 200 OK response for now
      				request_queue.push(body)
      				res.writeHead(200, "OK", {'Content-Type': 'text/html'});
      				res.write(body);
      				res.end();
    			});
			}
			break;
		case '/delete':
			if (req.method == 'GET'){
				console.log('receive delete call from ' + req.url)
				request_queue = []
				request_attempt = 0
				res.write(JSON.stringify(request_queue));
				res.end();
			}
			break;

		case '/fail':
			if (req.method == 'POST'){
				console.log('receive post fail call from ' + req.url)
				request_attempt++
				body = request_attempt.toString() + " attempt failed"
				res.writeHead(500, {
 					'Content-Length': body.length,
  					'Content-Type': 'text/plain' }
  				);
  				res.write(body)
				res.end();
			}else if (req.method == 'GET'){
				console.log('receive get fail call from ' + req.url)
				res.write(JSON.stringify(request_attempt));
				res.end();
			}
			break;
		default:
			console.log('unexpected call with ' + req.url + ' ' + req.method)
			break;

	}
}).listen(9988);

date = new Date(1360011904*1000);

console.log(date)