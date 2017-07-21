var request = require('request');
var express = require('express');
var app = express();

var options1 = 
{
	method: 'POST',
	url:"https://6a2ca48808dc43a195d8c18dfb063edb-vp0.us.blockchain.ibm.com:5003/chaincode",
	json:{
  "jsonrpc": "2.0",
  "method": "invoke",
  "params": {
    "type": 1,
    "chaincodeID": {
      "name": "27d8aeb7e5d3e357e62bf06040e9d65d2c7dd89a1fca7247b985936bf8d0e4a2ab0436bec186dd659c73930678c5ec3cbeb567dafa2f33bcafcfab08c307bd44"
    },
    "ctorMsg": {
      "function": "add",
      "args": [
        "67","66"
      ]
    },
    "secureContext": "user_type2_2"
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
