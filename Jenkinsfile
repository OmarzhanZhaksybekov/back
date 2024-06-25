pipeline {
    agent any

    stages {
        stage('Checkout Repositories') {
            steps {
                script {
                    // Клонирование репозитория с docker-compose
                    dir('back') {
                        git url: 'https://github.com/ShawaDev/back', branch: 'main'
                    }
                    
                    dir('auth'){
                        git url: 'https://github.com/ShawaDev/auth', branch: 'main'
                    }

                    dir('front'){
                        git url: 'https://github.com/ShawaDev/front', branch: 'main'
                    }
                }
            }
        }

        stage('Build Docker Images') {
            steps {
                script {
                    // Сборка Docker-образа для первого приложения на Go
                    dir('back') {
                        bat 'docker build -t back .'
                    }

                    // Сборка Docker-образа для второго приложения на Go
                    dir('auth') {
                        bat 'docker build -t auth .'
                    }

                    // Сборка Docker-образа для приложения на React
                    dir('front') {
                        bat 'docker build -t react .'
                    }
                }
            }
        }

        stage('Deploy with Docker Compose') {
            steps {
                script {
                    dir('back') {
                        bat 'docker-compose down' // Остановка предыдущих контейнеров, если необходимо
                        bat 'docker-compose up' // Поднятие контейнеров в фоне
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