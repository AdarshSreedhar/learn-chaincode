/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
/*
	in order to conver a piece of go code into chaincode, all you need to do 
	is implement the chaincode shim interface.the three functions are init,invoke and query
	all three functions have the same prototype they take in a stub(what we use to read and write to/from ledger)
	a function name and an array of strings
	this tutorial we will be building chaincode to create generic assets
*/
package main

import (
	"errors"//standard go error format
	"fmt"//contains println for debugging/logging
	"github.com/hyperledger/fabric/core/chaincode/shim"//contains the defnition for the chaincode interface and the chaincode stub
	//whcih we will need to interact with the ledger
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
//this function will execcute when each peer deploys thier instance of the chaincode
//it just calls shim.Start() which sets up communication between this chaincode and the peer that deployed it.
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
/*
	as the name implies should be used to do any initialisation the chaincode needs.
	here we use this Init to configure the initial state of a single key/value pair on the ledger
*/
// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}
	/*
	here we change the Init function so that it stores the first element in the args argument
	to the key hello_world.This is done by using the stub function stub.putState
	the function interprets the first argument sent in the deployment request 
	as the value to be stored under the key hello_world in the ledger
	if we send wrong number of arguments error will be returned else nothing, exits cleanly
	*/
	//err:=stub.PutState("hello_world",[]byte(args[0]))
	var key,value string
	var err error
	key=args[0]
	value=args[1]
	err=stub.PutState(key,[]byte(value))
	if(err!=nil){
		return nil,err
	}
	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
/*
incoke is called when you want chaincode functions to do real work
invocations will be captured as transactions which will be grouped into blocks on the chain
invoking chaincode updates the ledger
*/
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}else if function=="write" {
		return t.write(stub,args)
	}
	
	fmt.Println("invoke did not find func: " + function)					//error
	return nil, errors.New("Received unknown function invocation: " + function)
	
}
/*
notice that both Init and write functions are very similar, i.e. they check for a certain number of arguments 
and write a key/value pair to the ledger.
but in write function the putSate allows me to pass in both the key and the value for the call to putstate
so it allows us to insert whatever key value pair we want into the blockchain
*/
func (t* SimpleChaincode) write(stub shim.ChaincodeStubInterface,args []string)([]byte,error){
	var key,value string
	var err error
	fmt.Println("running write()")
	if(len(args)!=2){
		return nil,errors.New("Incorrect number of arguments.expecting name of the key and value to set")
	}
	//renaming for fun
	key=args[0]
	value=args[1]
	err=stub.PutState(key,[]byte(value))//write the variable into the chaincode state
	if err!=nil{
		return nil,err
	}
	return nil,nil
}
/*
query is called whenever you query your chaincode's state.queries dont result in blocks being 
added amd we cant use functions like putquery inside query.
this is just to read the value of chaincode key/value pairs
*/
// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	/*if function == "dummy_query" {											//read a variable
		fmt.Println("hi there " + function)						//error
		return nil, nil;
	}*/
	//we will have to write a read function as well now
	if function== "read"{
		return t.read(stub,args)
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query: " + function)
}
//getState is complementary to putState.So just like how putState inserts a key value pair, getState 
//lets us read a value for a previously written key
//so this function returns back as an array of bytes to Query() which send it back to the REST handler
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key,jsonResp string
	var err error
	if len(args)!=1{
		return nil,errors.New("Incorrect number of arguments.expecting the name of the key to query")
	}
	key=args[0]
	valasbytes,err:=stub.GetState(key)
	if err!=nil{
		jsonResp="{\"Error\":\"Failed to get state for"+key+"\"}"
		return nil,errors.New(jsonResp)
	}
	return valasbytes,nil
}
/*
the fastest way to test chaincode is to use REST interface onn Blumix or Postman
*/