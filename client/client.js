//var PROTO_PATH = __dirname + '/../screenpb/screen.proto';
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
//console.log(packageDefinition)
var sampleDef = grpc.loadPackageDefinition(packageDefinition).screen;

//console.log(sampleDef)

function main(){
    var stub = new sampleDef.ScreenService('localhost:50051', grpc.credentials.createInsecure());

    var calculation = {first_int: 4, second_int: 5};
    var sampCalc = {
        Calculation: calculation
    }

    
    //Console.log(sampCalc)
    /*stub.sample(sampCalc, function(err, sampRes) {
    if (err) {
        // process error
        console.log(err)
    } else {
        // process feature
        console.log(sampRes);
    }
    });*/
    stub.speak({SpeakPhrase: "Hello"}, function(err, sampRes) {
        if (err) {
            // process error
            console.log(err)
        } else {
            // process feature
            console.log(sampRes);
        }
    });
}

main();
