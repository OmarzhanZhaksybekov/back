pipeline {
    agent any

    environment {
        DOCKER_COMPOSE_VERSION = '2.27.1' // Укажите версию Docker Compose, если необходимо
        DB_CONTAINER_NAME = 'db'
        API_SERVER_CONTAINER_NAME = 'web-app'
        AUTH_SERVER_CONTAINER_NAME = 'auth-server'
        FRONTEND_CONTAINER_NAME = "react-app"
    }

    stages {
        stage('Checkout') {
            steps {
                script{
                    git branch: 'main', url: 'https://github.com/ShawaDev/back'
                }

                script{
                    dir('auth'){
                        git branch: 'main', url: 'https://github.com/ShawaDev/auth'
                    }
                    dir('front'){
                        git branch: 'main', url: 'https://github.com/ShawaDev/front'
                    }
                }
            }
        }

        stage('Build and Push Docker Images') {
            steps {
                script {
                    // Сборка и пуш Docker образов для каждого сервера
                    dir('back') {
                        docker.build("shawadeveloper/back-web-app:latest").push()
                    }
                    dir('auth') {
                        docker.build("shawadeveloper/back-auth-server:latest").push()
                    }
                    dir('front') {
                        docker.build("shawadeveloper/back-react-app:latest").push()
                    }
                }
            }
        }

        stage('Run Docker Compose') {
            steps {
                script {
                    // Устанавливаем Docker Compose, если не установлено
                    sh """
                    if ! docker-compose --version; then
                        curl -L "https://github.com/docker/compose/releases/download/${DOCKER_COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
                        chmod +x /usr/local/bin/docker-compose
                    fi
                    """
                    // Запускаем Docker Compose
                    dir('back') {
                        sh 'docker-compose down'
                        sh 'docker-compose up -d --build'
                    }
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    // Здесь можно определить ваши тесты
                    // Например, можно запустить тесты для вашего API сервера
                    // Важно подождать некоторое время, чтобы контейнеры полностью запустились
                    sleep 30

                    // Пример теста: проверяем, что контейнеры работают
                    sh "docker ps | grep ${DB_CONTAINER_NAME}"
                    sh "docker ps | grep ${API_SERVER_CONTAINER_NAME}"
                    sh "docker ps | grep ${AUTH_SERVER_CONTAINER_NAME}"
                    sh "docker ps | grep ${FRONTEND_CONTAINER_NAME}"
                }
            }
        }

    post {
        success {
            echo 'Pipeline completed successfully.'
        }
        failure {
            echo 'Pipeline failed.'
        }
    }
}
