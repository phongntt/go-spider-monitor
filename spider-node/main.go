package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/phongntt/go-spider-monitor/config"
	"github.com/phongntt/go-spider-monitor/spiderutils"
)

const NO_RESULT_INT = -999999
const NO_RESULT_STR = "[no_result]"
const STATUS_FILE_PATH = "../logs/"
const STATUS_FILE_PREFIX = "status"
const STATUS_FILE_EXT = ".json"
const FILENAME_SEPERATOR = "_"
const CONF_PATH = "../conf"

type CheckTaskResult struct {
	Name       string `json:"name"`
	ResultCode int    `json:"result_code"`
	ResultDesc string `json:"result_desc"`
}

type StatusSum struct {
	NodeName     string            `json:"node_name"`
	NodeType     string            `json:"node_type"`
	Status       string            `json:"status"`
	StatusCode   int               `json:"status_code"`
	CheckResults []CheckTaskResult `json:"check_results"`
}

func emptyResultFromCheckTask(checkTask config.CheckTask) CheckTaskResult {
	taskResult := CheckTaskResult{checkTask.Name, NO_RESULT_INT, NO_RESULT_STR}
	return taskResult
}

func main() {
	println("Start >>>")

	allConfig, err := config.LoadAllConfig(CONF_PATH)
	if err != nil {
		println("Cannot load config files from config Dir")
	}

	if len(allConfig) == 0 {
		fmt.Println("1- WARNING| None config file")
		os.Exit(1)
	}

	for confName, config := range allConfig {
		println(fmt.Sprintf("---- RUN FOR: %s -----", confName))

		if config.IsErrorConfig() {
			println("3- UNKNOWN| Cannot read config `%s`", confName)
			println()
		} else {
			results := runTasks(config.CheckTasks)
			fmt.Println(results)

			status := processCheckResult(config, results)
			println("==== RESULT ====")
			fmt.Println(status)

			writeStatusToFile(status, confName)

			println("\n")
		}
	}
}

/*
func main() {
	confFile := CONF_PATH + "/config.json"

	config, err := config.ReadFromFile(confFile)
	if err != nil {
		println(fmt.Sprintln(err))
		fmt.Print("3- UNKNOWN| Cannot read config file")
		println()
		os.Exit(3)
	}

	// DEBUG
	////println(config)

	results := runTasks(config.CheckTasks)
	fmt.Println(results)

	status := processCheckResult(config, results)
	println("==== RESULT ====")
	fmt.Println(status)

	writeStatusToFile(status)
}
*/

func writeStatusToFile(status StatusSum, confName string) {
	statusBytes, err := json.Marshal(status)
	if err != nil {
		fmt.Print("3- UNKNOWN| Cannot convert status to JSON format")
		println()
		os.Exit(3)
	}

	timepart := time.Now().Format("2006-01-02_150405.999999")
	statusFilename1 := STATUS_FILE_PATH + confName + FILENAME_SEPERATOR +
		STATUS_FILE_PREFIX + STATUS_FILE_EXT
	statusFilename2 := STATUS_FILE_PATH + confName + FILENAME_SEPERATOR +
		STATUS_FILE_PREFIX + FILENAME_SEPERATOR + status.Status +
		FILENAME_SEPERATOR + timepart + STATUS_FILE_EXT

	err = ioutil.WriteFile(statusFilename1, statusBytes, 0644)
	if err != nil {
		fmt.Print("3- UNKNOWN| Cannot write status file")
		println()
		os.Exit(3)
	}

	err = ioutil.WriteFile(statusFilename2, statusBytes, 0644)
	if err != nil {
		fmt.Print("3- UNKNOWN| Cannot write status file")
		println()
		os.Exit(3)
	}
}

func runTasks(tasks []config.CheckTask) []CheckTaskResult {
	checkResutlArr := make([]CheckTaskResult, 0, len(tasks))

	for _, task := range tasks {
		checkTaskResult := runOneTask(task)
		checkResutlArr = append(checkResutlArr, checkTaskResult)
	}

	return checkResutlArr
}

func runOneTask(task config.CheckTask) CheckTaskResult {
	checkTaskResult := emptyResultFromCheckTask(task)
	checkTaskResult.ResultCode, checkTaskResult.ResultDesc = spiderutils.RunCheckCommand(task.Command)
	return checkTaskResult
}

func processCheckResult(conf config.ConfigData, checkResults []CheckTaskResult) StatusSum {
	status := StatusSum{conf.NodeName, conf.NodeType, NO_RESULT_STR, NO_RESULT_INT, checkResults}

	// Add checking results to params
	params := make(map[string]int)
	for _, checkResult := range checkResults {
		params[checkResult.Name] = checkResult.ResultCode
	}

	// Add date time to param
	dNow := time.Now()
	year, month, day := dNow.Date()
	hour, min, sec := dNow.Clock()
	params["year"] = year
	params["month"] = int(month)
	params["day"] = day
	params["hour"] = hour
	params["minute"] = min
	params["second"] = sec

	statusValueMap := make(map[int]bool)
	for _, statusExpr := range conf.CheckExpressions {
		exprVal, err := spiderutils.EvalBoolExpr(statusExpr.Expression, params)
		if err != nil {
			panic(err)
		}
		statusValueMap[statusExpr.Status] = exprVal
		////fmt.Println(statusExpr.Status, "---", statusExpr.Expression, "--->", exprVal)
	}

	// Get Final status
	for _, st := range [4]int{0, 1, 2, 3} {
		if statusValueMap[st] {
			status.StatusCode = st
			status.Status = spiderutils.StatusNumToStr(st)
			return status
		}
	}

	status.StatusCode = 3
	status.Status = spiderutils.StatusNumToStr(3)
	return status
}
