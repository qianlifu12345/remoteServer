pwd;
ls -la;
export PATH="$PATH:/usr/local/bin/go/bin";
export GOPATH="/var/lib/jenkins/goworkspace";
export GOROOT="/usr/local/bin/go";
env;
echo $GOPATH;
echo $GOROOT;
docker login --username=bob --password=mAIZUOWANG3011 reg.miz.so
bash ./autobuild.sh build;
bash ./autobuild.sh docker;
