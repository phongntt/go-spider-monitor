/******************************************************
* Program:
*
* Huong dan su dung:
* - Can lam truoc khi chay ung dung:
*   -- Chuan bi app SqlCL + Script cap nhat CMND co ten "check-db-online.sql"
*     --- SqlCL --> dat tai thu muc co duong dan tuong doi tu chuong trinh nay nhu sau '../shelper' (Sql Helper)
*			--- check-db-online.sql --> dat tai thu muc '../shelper/script/check-db-online.sql'
*   -- Thiet lap cac bien moi truong
*     --- ENV_GO_CHECK_ON_DB --> DB can cap nhat
*     --- ENV_GO_SQLCL_DIR --> Thu muc chua chuong trinh SqlCL va script cap nhat CMND (test_upd_idnum.sql)
* - Su dung:
*   <ten app> [thong ting ket noi ebankkhdn/ebankkhdn@//ebank-db-test.eab.com.vn:1521/ebanking]
******************************************************/

package main

import (
	"fmt"
	"os"
	"os/exec"
)

const RESULT_OUT_BEGIN = "____PROCESS_RESULT_BEGIN____"
const RESULT_OUT_END = "____PROCESS_RESULT_END____"
const SQL_HELPER_DIR = "../shelper"

func main() {

	println("CHECK SELECT ON DB --> START")

	app := SQL_HELPER_DIR + "/bin/sql"
	arg0 := "-L"       // exit (exit code = 1) if Logon fail
	arg1 := os.Args[1] //os.Getenv("ENV_GO_CHECK_ON_DB")
	arg2 := "@" + SQL_HELPER_DIR + "/script/check-db-online.sql"

	// Run SQLCL
	cmd := exec.Command(app, arg0, arg1, arg2)
	stdout, err := cmd.Output()

	println(RESULT_OUT_BEGIN)
	defer println(RESULT_OUT_END)

	if err != nil {
		////fmt.Println(err)
		fmt.Print("1- WARNING")
		os.Exit(1)
		//return
	}

	println(string(stdout))

	fmt.Print("0- OK")
	// exit 0
}
