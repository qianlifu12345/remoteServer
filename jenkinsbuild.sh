pwd;
ls -la;
export PATH="$PATH:/usr/local/bin/go/bin";
export GOPATH="/var/lib/jenkins/goworkspace";
export GOROOT="/usr/local/bin/go";
env;
echo $GOPATH;
echo $GOROOT;
echo ;
echo ;
echo "Start build";
sudo docker login --username=bob --password=mAIZUOWANG3011 reg.miz.so
sudo bash ./autobuild.sh build;
sudo bash ./autobuild.sh docker;
