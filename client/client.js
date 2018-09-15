var express = require('express');
var app = express();

var PROTO_PATH = "/Users/manraajnijjar/Desktop/Screen/screenpb/screen.proto"

var grpc = require('grpc');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });

var sampleDef = grpc.loadPackageDefinition(packageDefinition).screen;
var stub = new sampleDef.ScreenService('localhost:50051', grpc.credentials.createInsecure());

app.use(express.static(__dirname + '/public'));

app.listen('5000'); 

function calculate(){
    var calculat = {first_int: 4, second_int: 5};
    var sampCalc = {
        calculation: calculat
    }

    
    //Console.log(sampCalc)
    stub.sample(sampCalc, function(err, sampRes) {
    if (err) {
        // process error
        console.log(err)
    } else {
        // process feature
        console.log(sampRes);
    }
    });
}

function speak(){
    stub.speak({speakPhrase: "Hellsfsfso"}, function(err, sampRes) {
        if (err) {
            // process error
            console.log(err)
        } else {
            // process feature
            console.log(sampRes);
        }
    });
}