var request = require('request');
var express = require('express');
var app = express();

var options1 = 
{
	method: 'POST',
	url:"https://6a2ca48808dc43a195d8c18dfb063edb-vp0.us.blockchain.ibm.com:5003/chaincode",
	json:{
  "jsonrpc": "2.0",
  "method": "deploy",
  "params": {
    "type": 1,
    "chaincodeID": {
      "path": "https://github.com/AdarshSreedhar/learn-chaincode/adding2numbers"
    },
    "ctorMsg": {
      "function": "init",
      "args": [
        "4","77"
      ]
    },
    "secureContext": "admin"
  },
  "id": 0
},
	headers:
	{
			'Accept': 'application/json',
        	'Accept-Charset': 'utf-8'
	}
};

//The first one----DEPLOY
request(options1, function(err,res,body)
{
	if(err)
	{
		console.log("ERROR" + err);
	}

	console.log(body);
	console.log("_____________________");
	//console.log(res);
});
