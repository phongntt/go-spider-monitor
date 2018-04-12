package config

import (
	"github.com/phongntt/go-spider-monitor/spiderutils"
	"strings"
)

const NONAME = "[NONE]"
const ERROR_TYPE = "ERROR"

type CheckTask struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}

type CheckExpression struct {
	Status     int    `json:"status"`
	Expression string `json:"expression"`
}

type ConfigData struct {
	NodeName         string            `json:"node_name"`
	NodeDescription  string            `json:"node_description"`
	NodeType         string            `json:"node_type"`
	CheckTasks       []CheckTask       `json:"check_tasks"`
	CheckExpressions []CheckExpression `json:"check_expressions"`
}

func (c *ConfigData) IsErrorConfig() bool {
	return (c.NodeName == NONAME && c.NodeType == ERROR_TYPE)
}

func createErrorConfig() ConfigData {
	return ConfigData{NONAME, "", ERROR_TYPE, nil, nil}
}

func processAndDecryptText(text string) (string, error) {
	if strings.HasPrefix(text, "ENC:") {
		encText := strings.TrimPrefix(text, "ENC:")
		return spiderutils.SpiderDecrypt(encText)
	}
	return text, nil
}

/****************************
* Decyript all encrypted value in ConfigData
* Note: decrypt on the ConfigData-object it self
****************************/
func (c *ConfigData) DecryptAll() error {
	decText, err := processAndDecryptText(c.NodeName)
	if err != nil {
		return err
	}
	c.NodeName = decText

	c.NodeDescription, err = processAndDecryptText(c.NodeDescription)
	if err != nil {
		return err
	}

	c.NodeType, err = processAndDecryptText(c.NodeType)
	if err != nil {
		return err
	}

	for ic, _ := range c.CheckTasks {
		c.CheckTasks[ic].Name, err = processAndDecryptText(c.CheckTasks[ic].Name)
		if err != nil {
			return err
		}

		c.CheckTasks[ic].Command, err = processAndDecryptText(c.CheckTasks[ic].Command)
		if err != nil {
			return err
		}
	}

	for ie, _ := range c.CheckExpressions {
		c.CheckExpressions[ie].Expression, err = processAndDecryptText(c.CheckExpressions[ie].Expression)
		if err != nil {
			return err
		}
	}

	return nil
}
