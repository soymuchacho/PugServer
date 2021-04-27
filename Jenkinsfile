pipeline {
  agent any
  stages {
    stage('Build') {
      parallel {
        stage('Build') {
          steps {
            echo 'build'
          }
        }

        stage('') {
          steps {
            echo 'banch'
          }
        }

      }
    }

  }
}