package prompts

import "fmt"

func ChooseBetweenFunctionsSingle(
	opt struct {
		Prompt    string
		Functions string
	},
) string {
	construction := `
		Your role is to choose between functions which is the most suitable for the given prompt:

		%s

		You have to choose between the following functions:
		%s

		The rules are:
		1. You have to choose the function which is the most suitable for the given prompt.
		2. You have to answer in a JSON format like this:
		{
			"function": "function_name",
			"parameters": {
				"param1": "value1",
				"param2": value2
			}
		}
		3. If the type of the parameter is a string, you have to put the value in double quotes and translate it to the language of the prompt.
	`

	// Construct the final prompt
	prompt := fmt.Sprintf(construction, opt.Prompt, opt.Functions)

	return prompt
}

func ChooseBetweenFunctionsMultiple(
	opt struct {
		prompt    string
		functions string
	},
) string {
	construction := `
		Consider the following prompt:

		%s

		Now consider the following function:
		%s

		Give an extimation of how well this function fits the prompt. Consider the following scale:

		1 - Very bad: the function has nothing to do with the prompt
		2 - Bad: the function is not suitable for the prompt
		3 - Good: the function can be used for the prompt, but it is not the best option
		4 - Very good: the function is suitable for the prompt
		5 - Perfect: the function is the best option for the prompt


		The rules are:
		1. You have to decide how well the function fits the prompt.
		2. You have to answer in a JSON format like this:
		{
			"score": 1,
			"parameters": {
				"param1": "value1",
				"param2": null
			}
		}
		3. If the type of the parameter is a string, you have to put the value in double quotes and translate it to the language of the prompt.
		4. If the parameter is not found in the prompt, you have to put the value as null.
	`

	prompt := fmt.Sprintf(construction, opt.prompt, opt.functions)

	return prompt
}
