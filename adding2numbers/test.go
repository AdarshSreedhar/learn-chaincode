package main
import (
    "errors"
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)
type SimpleChaincode struct {
}

func main() {
	err:= shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting simlpe chaincode:%s",err)
	}
}
func (t *SimpleChaincode) Init (stub shim.ChaincodeStubInterface, function string, args []string) ([]byte,error) {
	fmt.Printf("Init called, initializing chaincode")
	var err error
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments.Expecting 2")
	}
	// Write the state to the ledger
	err=stub.PutState(args[0],[]byte(args[1]))
	if(err!=nil) {
		return nil,err
	}
	if function=="add" {
		return t.add(stub,args)
	}
	
	fmt.Println("query did not find function: "+function)
	return nil, errors.New("Received unknown function query: "+function)
}
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface,function string,args []string) ([]byte, error) {
	fmt.Println("invoke is running "+function)
	if function=="add" { 
		return t.Init(stub,"add",args)
	}
	fmt.Println("invoke did not function: "+function)
	return nil, errors.New("Received unknown function invocation: "+function)
}
func (t* SimpleChaincode) add(stub shim.ChaincodeStubInterface,args []string)([]byte,error) {
	var A,B int
	var err error
	fmt.Println("running add()")
	A,err=strconv.Atoi(args[0])
	B,err=strconv.Atoi(args[1])
	if(len(args)!=2){
		return nil,errors.New("Incorrect number of argumetns.expecting two numbers")
	}
	err=stub.PutState("Sum",[]byte(strconv.Itoa(A+B)))
	if err!=nil{
		return nil,err
	}
	return nil,nil
}
func (t* SimpleChaincode) Query(stub shim.ChaincodeStubInterface,function string,args []string) ([]byte,error) {
	fmt.Println("query is running"+function)
	if function=="getsum"{
		return t.getsum(stub,args)
	}
	fmt.Println("query did not find func: "+function)
	return nil, errors.New("Received unknown function query: "+function)
}
func (t *SimpleChaincode) getsum(stub shim.ChaincodeStubInterface,args []string) ([]byte, error) {
	
	var err error
	var jsonResp string
	if len(args)!=1{
		return nil,errors.New("Incorrect number of arguments.")
	}
	sum, err:=stub.GetState("Sum")
	if err!=nil{
		jsonResp="{\"Error\":\"Failed to get state for\"}"
		return nil,errors.New(jsonResp)
	}
	return sum,nil
}