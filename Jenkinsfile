pipeline {
    agent any

    environment {
        INFISICAL_HOST_URL = 'http://host.docker.internal:8078'
    }

    stages {
        stage('Run Go App with Infisical') {
            steps {
                // Membaca Username (Client ID) & Password (Client Secret) sekaligus
                withCredentials([usernamePassword(credentialsId: 'INFISICAL_CREDENTIALS', usernameVariable: 'CLIENT_ID', passwordVariable: 'CLIENT_SECRET')]) {
                    sh '''
                        # Login Universal Auth
                        export INFISICAL_TOKEN=$(infisical login --method=universal-auth \
                            --client-id="$CLIENT_ID" \
                            --client-secret="$CLIENT_SECRET" \
                            --domain="$INFISICAL_HOST_URL" \
                            --silent --plain)

                        # Run Go App
                        infisical run --env=dev --domain="$INFISICAL_HOST_URL" -- go run main.go
                    '''
                }
            }
        }
    }
}