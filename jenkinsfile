node(label: 'master') {
    def appName
    def branch
    def version
    def targetEnv
    def dockerTag
    def commitAuthorName
    def commitMessage
    serverInfoData = null
try{
    // Check Target
    stage('STEP1: Checkout SCM') {
            println "Running Stage: Checkout SCM"
            checkout scm
            appName = scm.getUserRemoteConfigs()[0].getUrl().tokenize('/').last().split("\\.")[0]
            branch = env.BRANCH_NAME
            version = readFile('version').trim()
            targetEnv = ''
            if (targetEnv == '' && branch != '') {
                if (branch.startsWith('test') || branch.startsWith('hotfix/')) {
                    targetEnv = 'dev'
                    dockerTag = "${version}-${targetEnv}"
                }else if (branch == 'develop') {
                    targetEnv = 'uat'
                    dockerTag = "${version}-${targetEnv}"
                } else if (branch == 'main') {
                    targetEnv = 'prd'
                    dockerTag = "${version}"
                }
            }
    }
    if (targetEnv == 'dev' || targetEnv == 'uat' || targetEnv == 'prd') {
            stage('STEP2: Get Git Information') {
                    echo "Running Stage: Get Commit Information"
                    def commitInfo = sh(script: 'git log -n 1 --pretty=format:"%an %ae %s"', returnStdout: true).trim().split(' ', 3)
                    commitAuthorName = commitInfo[0]
                    commitMessage = commitInfo[2]
                }

            stage('STEP3: Dockerize') {
                    println "Running Stage : Dockerize"
                    def DOCKER_HUB_USERNAME = env.DOCKER_HUB_USERNAME
                    def DOCKER_IMAGE = ''
                    DOCKER_IMAGE = "${DOCKER_HUB_USERNAME}/${appName}:${dockerTag}"
                    println "DOCKER_IMAGE: ${DOCKER_IMAGE}"
                    script {
                        def registryUrl = env.DOCKER_REGISTRY_URL
                        def registryCredentials = env.DOCKER_CREDENTIALS_ID
                        def dockerfile = 'Dockerfile'
                        def dockerImage
                        docker.withRegistry(registryUrl, registryCredentials) {
                            dockerImage = docker.build("${DOCKER_IMAGE}", "-f ${dockerfile} .")
                            dockerImage.push()
                            // def buildDockerImage = "${DOCKER_HUB_USERNAME}/${appName}:${version}"
                            // sh "docker tag ${DOCKER_IMAGE} ${buildDockerImage}"
                            // dockerImage = docker.build("${buildDockerImage}")
                            // dockerImage.push()
                        }
                    }
                }
        }
    
    if (targetEnv == 'dev' || targetEnv == 'uat') {
        stage('STEP4: Pipeline Deploy') {
            build job: 'DEVOPS-PIPELINE/ithq-app-deployer',
            parameters: [
                string(name: 'APP_NAME', value: "${appName}"),
                string(name: 'ENV_NAME', value: "${targetEnv}"),
                string(name: 'VERSION', value: "${dockerTag}"),
                string(name: 'DETAIL', value: "\n\nService Name: \n${appName} \n\nVersion: \n${dockerTag} \n\nDeployer: \n${commitAuthorName} \n\nMessage: \n${commitMessage} to ${branch}\n\n")
            ]
        }
    }
    if (targetEnv == 'prd'){
        stage('STEP4: Alert Buil Main') {
            build job: 'DEVOPS-PIPELINE/ithq-telegram-jenkins',
            parameters: [
                string(name: 'NAME', value: "\n\nBuild Success!\nPlease go to Jenkins and initiate the deployment pipeline (ithq-production-deployer)\n\nService Name: \n${appName} \n\nVersion: \n${dockerTag} \n\nMerged By: \n${commitAuthorName} \n\nMessage: \nBuild images  to ${branch}\n\n"),
                string(name: 'STATUS', value: "SUCCESS"),
            ]

        }
    }

    } catch (err){
        stage('FAILURE: Pipeline Alert') {
            build job: 'DEVOPS-PIPELINE/ithq-telegram-jenkins',
            parameters: [
                string(name: 'NAME', value: "\n\nJob:\nContinuous Integration (CI)\n\nService Name: \n${appName} \n\nVersion: \n${dockerTag} \n\nDeployer: \n${commitAuthorName} \n\nError: \n${err}\n\n"),
                string(name: 'STATUS', value: "FAILURE"),
            ]
        }
        currentBuild.result = 'FAILURE'
    }
}