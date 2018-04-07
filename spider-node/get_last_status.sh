SPIDER_HOME=/path/to/spider-node
cat $SPIDER_HOME/logs/status.json | sed 's/.*\("status":".*","status_code"\).*/\1/' | sed 's/,".*//'
