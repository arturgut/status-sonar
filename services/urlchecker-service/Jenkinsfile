pipeline{

  agent any
  
  environment {
    GIT_COMMIT_SHA = sh(returnStdout: true, script: "git log -n 1 --pretty=format:'%h'").trim()
    DOCKER_USER = credentials('docker-hub-user')
    DOCKER_PASS = credentials('docker-hub-pass')
  }
 
  stages{
    
    stage('Build Dev') {
      steps {
        echo 'Building...'
        sh '''
          make docker-build-dev 
        '''
      }
    }

    stage('Test') {
      steps {
        echo 'Testing...'
        sh ''' make docker-run-test '''
      }
    }

    stage('Docker Push Dev') {
      steps {
        echo 'Push Dev version of container to Docker Hub'
        sh ''' make docker-push-dev '''
      }
    }
  }
  post {
        failure {
            echo 'Push Dev version of container to Docker Hub'
            sh 'printenv'
        }
    }
}