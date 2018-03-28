package config

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
