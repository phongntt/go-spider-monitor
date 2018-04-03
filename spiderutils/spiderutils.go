package spiderutils

import (
	"context"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/phongntt/crypto_helper"
	"os"
	"os/exec"
	"strings"
	"time"
)

const EVN_MASTER_KEY = "GO_SPIDER_MASTERKEY"
const DEFAULTKEY = "gospidergospidergospidergospider"
const CHECK_TOOLS_ROOT_PATH = "."

func GetMasterKey() string {
	if os.Getenv(EVN_MASTER_KEY) == "" {
		return DEFAULTKEY
	}
	return os.Getenv(EVN_MASTER_KEY)
}

func SpiderEncrypt(text string) (string, error) {
	ciphertext, err := crypto_helper.Encrypt_Base64(text, GetMasterKey())
	if err != nil {
		return "", err
	}

	return ciphertext, nil
}

func SpiderDecrypt(text64 string) (string, error) {
	ciphertext, err := crypto_helper.Decrypt_Base64(text64, GetMasterKey())
	if err != nil {
		return "", err
	}

	return ciphertext, nil
}

func RunCheckCommand(commandStr string) (int, string) {
	commandTimeoutSec := 120 * time.Second //Seconds

	commandWords := strings.Fields(commandStr)

	app := CHECK_TOOLS_ROOT_PATH + "/" + commandWords[0]
	args := commandWords[1:]

	// Create a new context and add a timeout to it
	ctx, cancel := context.WithTimeout(context.Background(), commandTimeoutSec)
	defer cancel() // The cancel should be deferred so resources are cleaned up

	// Create the command with our context
	cmd := exec.CommandContext(ctx, app, args...)
	//cmd := exec.Command(app, arg0, arg1, arg2, arg3)

	stdout, err := cmd.Output()

	// We want to check the context error to see if the timeout was executed.
	// The error returned by cmd.Output() will be OS specific based on what
	// happens when a process is killed.
	if ctx.Err() == context.DeadlineExceeded {
		println("Command timed out")
		return 1, "FAIL| Check command time-out"
	}

	if err != nil {
		println("Check command running fail")
		return 1, "FAIL Check command running fail"
	}

	return 0, string(stdout)
}

func EvalBoolExpr(expr string, params map[string]int) (bool, error) {
	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return false, err
	}

	parameters := make(map[string]interface{}, len(params))
	for k, v := range params {
		parameters[k] = v
	}

	ret, err := expression.Evaluate(parameters)
	if err != nil {
		fmt.Println(err)
	}

	if ret == true {
		return true, nil
	}
	return false, nil
}

func StatusNumToStr(statusNum int) string {
	switch statusNum {
	case 0:
		return "OK"
	case 1:
		return "Warning"
	case 2:
		return "Critical"
	}
	return "Unknown"
}
