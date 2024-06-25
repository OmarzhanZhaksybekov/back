pipeline {
    agent any
    environment {
        GO111MODULE = 'on'
        GOPATH = '/go'
    }
    stages {
        stage('Checkout Repositories') {
            steps {
                dir('back') {
                    git 'https://github.com/ShawaDev/back'
                }
                dir('auth') {
                    git 'https://github.com/ShawaDev/auth'
                }
                dir('react') {
                    git 'https://github.com/ShawaDev/front'
                }
            }
        }
        stage('Build Go Applications') {
            steps {
                dir('back') {
                    sh 'go build -o back'
                }
                dir('auth') {
                    sh 'go build -o auth'
                }
            }
        }
        stage('Build Docker Images') {
            steps {
                dir('back') {
                    sh 'docker build -t back .'
                }
                dir('auth') {
                    sh 'docker build -t auth .'
                }
                dir('react') {
                    sh 'docker build -t react .'
                }
            }
        }
        stage('Deploy with Docker Compose') {
            steps {
                dir('back') {
                    sh 'docker-compose down'
                    sh 'docker-compose up -d --build'
                }
            }
        }
    }
}
