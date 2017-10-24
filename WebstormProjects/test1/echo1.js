var http;
http = require('http');
var url = require('url');
function func (req, res) {
    var urlParsed = url.parse(req.url, true);
    if (urlParsed.query.message){
        res.writeHead(200, {"Content-Type": "text/plain"});
        res.write("Haa\n");
        res.end(urlParsed.query.message );
        console.log(urlParsed.query.message)
    } else {
        res.statusCode = 404;
        res.end("Page not found");
    }
};

http.createServer(func).listen(8888);