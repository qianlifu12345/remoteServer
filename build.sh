pwd;
ls -la;
export PATH="$PATH:/usr/local/bin/go/bin";
export GOPATH="/var/lib/jenkins/goworkspace";
export GOROOT="/usr/local/bin/go";
env;
echo $GOPATH;
echo $GOROOT;  
go build main.go;
