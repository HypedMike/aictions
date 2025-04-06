package functions

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
)

//go:embed config.json
var configFile []byte

type FunctionsHandler struct {
	Functions []FunctionModel `json:"functions"`
}

// NewFunctionsHandler creates a new FunctionsHandler with the given functions
func NewFunctionsHandler() *FunctionsHandler {
	var functions []FunctionModel
	// Unmarshal the config file into the functions slice
	if err := json.Unmarshal(configFile, &functions); err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}
	return &FunctionsHandler{
		Functions: functions,
	}
}

// PrintFunctions prints the functions to the console
func (fh *FunctionsHandler) PrintFunctions() {
	for _, function := range fh.Functions {
		arrayOfParams := ""
		for _, param := range function.Parameters {
			arrayOfParams += fmt.Sprintf("%s: %s, ", param.Name, param.Type)
		}
		// Remove the trailing comma
		if len(arrayOfParams) > 0 {
			arrayOfParams = arrayOfParams[:len(arrayOfParams)-2]
		}
		fmt.Printf("%s(%s)\n", function.Name, arrayOfParams)
	}
}

func (fh *FunctionsHandler) SprintFunction() (string, error) {
	res := ""
	for _, function := range fh.Functions {
		arrayOfParams := ""
		for _, param := range function.Parameters {
			arrayOfParams += fmt.Sprintf("%s: %s, ", param.Name, param.Type)
		}
		// Remove the trailing comma
		if len(arrayOfParams) > 0 {
			arrayOfParams = arrayOfParams[:len(arrayOfParams)-2]
		}

		res += fmt.Sprintf("%s(%s)\n", function.Name, arrayOfParams)
	}

	return res, nil
}

// GetFunctionByName returns the function with the given name
func (fh *FunctionsHandler) GetFunctionByName(name string) (*FunctionModel, error) {
	for _, function := range fh.Functions {
		if function.Name == name {
			return &function, nil
		}
	}
	return nil, fmt.Errorf("function %s not found", name)
}
