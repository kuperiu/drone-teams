pipeline {
  agent any
  stages {
    stage('input') {
      parallel {
        stage('input') {
          steps {
            input(message: 'Do you want to proceed', id: 'proceed')
          }
        }

        stage('test') {
          steps {
            echo 'test'
          }
        }

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