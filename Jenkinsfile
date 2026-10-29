pipeline {
    agent any

    parameters {
        choice(name: 'ENVIRONMENT', choices: ['dev', 'staging', 'prod'], description: 'Pilih Environment Deployment')
    }

    environment {
        INFISICAL_HOST_URL = 'http://host.docker.internal:8080'
    }

    stages {
        stage('Set Environment Config') {
            steps {
                script {
                    if (params.ENVIRONMENT == 'prod') {
                        env.INFISICAL_PROJECT_ID = 'a779a216-3ecd-47fb-b552-061e78c3cd26'
                    } else if (params.ENVIRONMENT == 'staging') {
                        env.INFISICAL_PROJECT_ID = 'ID_PROJECT_AUTH_STAGING'
                    } else {
                        env.INFISICAL_PROJECT_ID = 'ID_PROJECT_AUTH_DEV'
                    }
                }
            }
        }

        stage('Run Go App with Infisical') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'INFISICAL_CREDENTIALS', usernameVariable: 'CLIENT_ID', passwordVariable: 'CLIENT_SECRET')]) {
                    sh '''
                        export INFISICAL_TOKEN=$(infisical login --method=universal-auth \
                            --client-id="$CLIENT_ID" \
                            --client-secret="$CLIENT_SECRET" \
                            --domain="$INFISICAL_HOST_URL" \
                            --silent --plain)

                        infisical run \
                            --projectId="$INFISICAL_PROJECT_ID" \
                            --env="$ENVIRONMENT" \
                            --domain="$INFISICAL_HOST_URL" \
                            -- go run main.go
                    '''
                }
            }
        }
    }
}