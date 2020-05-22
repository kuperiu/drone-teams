pipeline {
  agent any
  stages {
    stage('input') {
      steps {
        input(message: 'Do you want to proceed', id: 'proceed')
      }
    }

  }
  environment {
    env = 'dev'
  }
}