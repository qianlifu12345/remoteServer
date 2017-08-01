node {



  stage ('Prepare the build') {

    echo 'Preparing the build'

  }

  stage ('Checkout') {


  checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[credentialsId: 'aa99aa0f-dbd1-42b9-8ecd-616860dd3336', url: 'git@gogs.miz.so:data/glaucus.git']]])
  }


  stage ('Compile Go files') {
    sh 'go build'

  }





}

