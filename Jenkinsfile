pipeline {
    agent any
    tools {
        go 'go1.23.2'
    }
    environment {
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('Pre Test') {
            steps {
                echo 'installing dependencies'
                sh 'go version'
            }
        }
        stage('Build') {
            steps {
                echo 'compiling and building'
                sh 'go build'
            }
        }
        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]) {
                    echo 'running vetting'
                    sh 'go vet .'
                    echo 'running fmt'
                    sh 'go fmt .'
                    echo 'running test'
                    sh 'go test -v ./...'
                }
            }
        }
    }
}
