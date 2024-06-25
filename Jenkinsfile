pipeline {
    agent any

    environment {
        // Переменные окружения, если необходимо
    }

    stages {
        stage('Checkout Repositories') {
            steps {
                script {
                    // Клонирование репозитория с docker-compose
                    dir('back') {
                        git url: 'git@github.com:ShawaDev/back.git', branch: 'main'
                    }
                    
                    dir('auth'){
                        git url: 'git@github.com:ShawaDev/auth.git' branch: 'main'
                    }

                    dir('react'){
                        git url: 'git@github.com:ShawaDev/front.git' branch: 'main'
                    }
                }
            }
        }

        stage('Build Docker Images') {
            steps {
                script {
                    // Сборка Docker-образа для первого приложения на Go
                    dir('back') {
                        sh 'docker build -t back .'
                    }

                    // Сборка Docker-образа для второго приложения на Go
                    dir('auth') {
                        sh 'docker build -t auth .'
                    }

                    // Сборка Docker-образа для приложения на React
                    dir('react') {
                        sh 'docker build -t react .'
                    }
                }
            }
        }

        stage('Deploy with Docker Compose') {
            steps {
                script {
                    dir('back') {
                        sh 'docker-compose down' // Остановка предыдущих контейнеров, если необходимо
                        sh 'docker-compose up -d --build' // Поднятие контейнеров в фоне
                    }
                }
            }
        }
    }

    post {
        success {
            echo 'Pipeline completed successfully'
        }
        failure {
            echo 'Pipeline failed'
        }
    }
}
