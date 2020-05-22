pipeline {
  agent any
  stages {
    stage('input') {
      steps {
        input(message: 'Do you want to proceed', id: 'proceed')
      }
    }

    stage('another') {
      steps {
        input(message: 'Which version', id: 'version', submitterParameter: 'string', submitter: 'string')
      }
    }

  }
  environment {
    env = 'dev'
  }
}