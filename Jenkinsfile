node {



  stage ('Prepare the build') {

    echo 'Preparing the build'

  }
  
  
  stage ('test') {

  
   echo 'bob has changed 10.18'
   sh 'date'

  }

  stage ('Checkout') {

  
   checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[url: 'git@github.com:qianlifu12345/remoteServer.git']]])

  }


  stage ('Compile Go files') {
    sh 'pwd'
    sh 'ls -la'
    
    sh 'export PATH="$PATH:/usr/local/bin/go/bin"'
    sh 'export GOPATH="/var/lib/jenkins/goworkspace"'
    sh 'export GOROOT="/usr/local/bin/go"'

    sh 'env'
    sh 'bash $GOROOT' 
    sh 'bash $GOPATH'
    sh 'bash $GOROOT'    
    sh 'bash go build'

  }





}

