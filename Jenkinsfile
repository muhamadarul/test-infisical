pipeline {
    agent any

    environment {
        INFISICAL_HOST_URL = 'http://host.docker.internal:8080'
    }

    stages {
        stage('Detect Branch & Set Environment') {
            steps {
                script {
                    // Ambil nama branch yang sedang di-build
                    // GIT_BRANCH biasanya bernilai 'origin/main' atau 'origin/staging'
                    def currentBranch = env.GIT_BRANCH ?: env.BRANCH_NAME ?: ''
                    
                    echo "Branch yang sedang di-build: ${currentBranch}"

                    if (currentBranch.contains('main') || currentBranch.contains('master')) {
                        env.INFISICAL_PROJECT_ID = 'a779a216-3ecd-47fb-b552-061e78c3cd26'
                        env.TARGET_ENV           = 'prod'
                    } else if (currentBranch.contains('staging')) {
                        env.INFISICAL_PROJECT_ID = 'ID_PROJECT_STAGING_ANDA'
                        env.TARGET_ENV           = 'staging'
                    } else {
                        // Default / Fallback ke Dev
                        env.INFISICAL_PROJECT_ID = 'ID_PROJECT_DEV_ANDA'
                        env.TARGET_ENV           = 'dev'
                    }

                    echo "Menggunakan Infisical Project ID: ${env.INFISICAL_PROJECT_ID} [Env: ${env.TARGET_ENV}]"
                }
            }
        }

        stage('Run Go App with Infisical') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'INFISICAL_CREDENTIALS', usernameVariable: 'CLIENT_ID', passwordVariable: 'CLIENT_SECRET')]) {
                    sh '''
                        # 1. Login Universal Auth
                        export INFISICAL_TOKEN=$(infisical login --method=universal-auth \
                            --client-id="$CLIENT_ID" \
                            --client-secret="$CLIENT_SECRET" \
                            --domain="$INFISICAL_HOST_URL" \
                            --silent --plain)

                        # 2. Run Go app sesuai Project ID dan Env dinamis
                        infisical run \
                            --projectId="$INFISICAL_PROJECT_ID" \
                            --env="$TARGET_ENV" \
                            --domain="$INFISICAL_HOST_URL" \
                            -- go run main.go
                    '''
                }
            }
        }
    }
}