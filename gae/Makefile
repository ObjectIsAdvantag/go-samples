PROJECT_ID=gae-helloworld-1193

local:
	goapp serve . & start http://localhost:8080/

gae:
	appcfg.py -A $(PROJECT_ID) update . & start http://gae-helloworld-1193.appspot.com