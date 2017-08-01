node {



  stage ('Prepare the build') {

    echo 'Preparing the build'

  }

  stage ('Checkout') {

  
   checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[url: 'git@github.com:qianlifu12345/remoteServer.git']]])

  }


  stage ('Compile Go files') {
    sh 'bash go build'

  }





}

