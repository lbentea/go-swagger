// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/go-swagger/examples/cli/client/operations"
	"github.com/go-swagger/go-swagger/examples/cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationOperationsPutTest2766Cmd returns a command to handle operation putTest2766
func makeOperationOperationsPutTest2766Cmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "PutTest2766",
		Short: ``,
		RunE:  runOperationOperationsPutTest2766,
	}

	if err := registerOperationOperationsPutTest2766ParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationOperationsPutTest2766 uses cmd flags to call endpoint api
func runOperationOperationsPutTest2766(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := operations.NewPutTest2766Params()
	if err, _ = retrieveOperationOperationsPutTest2766Plus1Flag(params, "", cmd); err != nil {
		return err
	}
	if err, _ = retrieveOperationOperationsPutTest2766Minus1Flag(params, "", cmd); err != nil {
		return err
	}
	if err, _ = retrieveOperationOperationsPutTest2766BodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationOperationsPutTest2766Result(appCli.Operations.PutTest2766(params, nil))
	if err != nil {
		return err
	}

	if !debug {
		fmt.Println(msgStr)
	}

	return nil
}

// registerOperationOperationsPutTest2766ParamFlags registers all flags needed to fill params
func registerOperationOperationsPutTest2766ParamFlags(cmd *cobra.Command) error {
	if err := registerOperationOperationsPutTest2766Plus1ParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationOperationsPutTest2766Minus1ParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationOperationsPutTest2766BodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationOperationsPutTest2766Plus1ParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	flagPlus1Description := ``

	var flagPlus1Name string
	if cmdPrefix == "" {
		flagPlus1Name = "+1"
	} else {
		flagPlus1Name = fmt.Sprintf("%v.+1", cmdPrefix)
	}

	var flagPlus1Default int64

	_ = cmd.PersistentFlags().Int64(flagPlus1Name, flagPlus1Default, flagPlus1Description)

	return nil
}

func registerOperationOperationsPutTest2766Minus1ParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	flagMinus1Description := ``

	var flagMinus1Name string
	if cmdPrefix == "" {
		flagMinus1Name = "-1"
	} else {
		flagMinus1Name = fmt.Sprintf("%v.-1", cmdPrefix)
	}

	var flagMinus1Default int64

	_ = cmd.PersistentFlags().Int64(flagMinus1Name, flagMinus1Default, flagMinus1Description)

	return nil
}

func registerOperationOperationsPutTest2766BodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var flagBodyName string
	if cmdPrefix == "" {
		flagBodyName = "body"
	} else {
		flagBodyName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(flagBodyName, "", `Optional json string for [body]. `)

	// add flags for body
	if err := registerModelGithubReactionsFlags(0, "githubReactions", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationOperationsPutTest2766Plus1Flag(m *operations.PutTest2766Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("+1") {

		var flagPlus1Name string
		if cmdPrefix == "" {
			flagPlus1Name = "+1"
		} else {
			flagPlus1Name = fmt.Sprintf("%v.+1", cmdPrefix)
		}

		flagPlus1Value, err := cmd.Flags().GetInt64(flagPlus1Name)
		if err != nil {
			return err, false
		}
		m.Plus1 = &flagPlus1Value

	}

	return nil, retAdded
}

func retrieveOperationOperationsPutTest2766Minus1Flag(m *operations.PutTest2766Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("-1") {

		var flagMinus1Name string
		if cmdPrefix == "" {
			flagMinus1Name = "-1"
		} else {
			flagMinus1Name = fmt.Sprintf("%v.-1", cmdPrefix)
		}

		flagMinus1Value, err := cmd.Flags().GetInt64(flagMinus1Name)
		if err != nil {
			return err, false
		}
		m.Minus1 = &flagMinus1Value

	}

	return nil, retAdded
}

func retrieveOperationOperationsPutTest2766BodyFlag(m *operations.PutTest2766Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		flagBodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		flagBodyValue := models.GithubReactions{}
		if err := json.Unmarshal([]byte(flagBodyValueStr), &flagBodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.GithubReactions: %v", err), false
		}
		m.Body = &flagBodyValue
	}
	flagBodyModel := m.Body
	if swag.IsZero(flagBodyModel) {
		flagBodyModel = &models.GithubReactions{}
	}
	err, added := retrieveModelGithubReactionsFlags(0, flagBodyModel, "githubReactions", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = flagBodyModel
	}

	if dryRun && debug {
		flagBodyValueDebugBytes, err := json.Marshal(m.Body)
		if err != nil {
			return err, false
		}
		logDebugf("Body dry-run payload: %v", string(flagBodyValueDebugBytes))
	}

	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationOperationsPutTest2766Result parses request result and return the string content
func parseOperationOperationsPutTest2766Result(resp0 *operations.PutTest2766OK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*operations.PutTest2766OK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
