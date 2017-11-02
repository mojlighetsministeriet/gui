const proxyHttpRequest = require('proxy-http-request');
const http = require('http');

const apiRegExp = /\/api\//;

http.createServer(function (request, response) {
  if (request.url.match(apiRegExp)) {
    proxyHttpRequest('https://internt.mojlighetsministeriet.se' + request.url, request, response);
  } else {
    proxyHttpRequest('http://localhost:1323' + request.url, request, response);
  }
}).listen(3000);
