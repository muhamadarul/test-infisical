pipeline {
    agent any

    // Parameter diambil dari UI Jenkins yang Anda atur di gambar
    parameters {
        choice(name: 'ENVIRONMENT', choices: ['dev', 'staging', 'prod'], description: 'Pilih Environment Target')
    }

    environment {
        // Port Infisical disesuaikan dengan setup lokal Anda (8078 / 8080)
        INFISICAL_HOST_URL = 'http://host.docker.internal:8078'
    }

    stages {
        stage('Set Environment Config') {
            steps {
                script {
                    // Membaca langsung dari parameter UI Jenkins (params.ENVIRONMENT)
                    def selectedEnv = params.ENVIRONMENT ?: 'dev'
                    echo "Environment terpilih dari UI: ${selectedEnv}"

                    switch(selectedEnv) {
                        case 'prod':
                            env.INFISICAL_PROJECT_ID = 'a779a216-3ecd-47fb-b552-061e78c3cd26'
                            env.TARGET_ENV           = 'prod'
                            break
                        case 'staging':
                            env.INFISICAL_PROJECT_ID = 'ID_PROJECT_STAGING_ANDA'
                            env.TARGET_ENV           = 'staging'
                            break
                        default: // dev
                            env.INFISICAL_PROJECT_ID = 'ID_PROJECT_DEV_ANDA'
                            env.TARGET_ENV           = 'dev'
                            break
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

                        # 2. Run Go app dengan Project ID dan Target Env sesuai pilihan UI
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