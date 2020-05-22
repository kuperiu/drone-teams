pipeline {
  agent any
  stages {
    stage('input') {
      parallel {
        stage('input') {
          steps {
            input(message: 'Do you want to proceed', id: 'proceed', submitter: ' choice(name: \'APP_VERSION\', choices: "v1.1\\nv1.2\\nv1.3", description: \'What to deploy?\')')
          }
        }

        stage('test') {
          steps {
            echo 'test'
          }
        }

        stage('print') {
          steps {
            echo 'ddd'
          }
        }

      }
    }

    stage('Deploy') {
      options {
        timeout(time: 30, unit: 'SECONDS')
      }
      input {
        message 'Which Version?'
        id 'Deploy'
        parameters {
          choice(name: 'APP_VERSION', choices: '''v1.1
v1.2
v1.3''', description: 'What to deploy?')
        }
      }
      steps {
        echo "Deploying ${APP_VERSION}."
      }
    }

  }
  environment {
    env = 'dev'
  }
}