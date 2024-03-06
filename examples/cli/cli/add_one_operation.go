// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/go-swagger/examples/cli/client/todos"
	"github.com/go-swagger/go-swagger/examples/cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationTodosAddOneCmd returns a command to handle operation addOne
func makeOperationTodosAddOneCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "addOne",
		Short: ``,
		RunE:  runOperationTodosAddOne,
	}

	if err := registerOperationTodosAddOneParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTodosAddOne uses cmd flags to call endpoint api
func runOperationTodosAddOne(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := todos.NewAddOneParams()
	if err, _ = retrieveOperationTodosAddOneBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTodosAddOneResult(appCli.Todos.AddOne(params, nil))
	if err != nil {
		return err
	}

	if !debug {
		fmt.Println(msgStr)
	}

	return nil
}

// registerOperationTodosAddOneParamFlags registers all flags needed to fill params
func registerOperationTodosAddOneParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTodosAddOneBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTodosAddOneBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var flagBodyName string
	if cmdPrefix == "" {
		flagBodyName = "body"
	} else {
		flagBodyName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(flagBodyName, "", `Optional json string for [body]. `)

	// add flags for body
	if err := registerModelItemFlags(0, "item", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationTodosAddOneBodyFlag(m *todos.AddOneParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		flagBodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		flagBodyValue := models.Item{}
		if err := json.Unmarshal([]byte(flagBodyValueStr), &flagBodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.Item: %v", err), false
		}
		m.Body = &flagBodyValue
	}
	flagBodyModel := m.Body
	if swag.IsZero(flagBodyModel) {
		flagBodyModel = &models.Item{}
	}
	err, added := retrieveModelItemFlags(0, flagBodyModel, "item", cmd)
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

// parseOperationTodosAddOneResult parses request result and return the string content
func parseOperationTodosAddOneResult(resp0 *todos.AddOneCreated, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*todos.AddOneDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*todos.AddOneCreated)
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
