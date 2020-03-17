node('haimaxy-jnlp') {
  stage('Prepare Stage') {
    echo "1. Prepare"
    checkout scm
    script {
        build_tag = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
        echo "####################### ${build_tag} ########${env.BRANCH_NAME}####${env.branch}####${env}#####"
    }
  }
  stage('Test') {
    echo "2.Test Stage"
  }
  stage('Build') {
    echo "3.Build Stage"
    sh "docker build -t whatlong7/jenkins-demo:${build_tag} ."
  }
  stage('Push') {
    echo "4. Push Docker image stage"
    withCredentials([usernamePassword(credentialsId: 'dockerHub', passwordVariable: 'dockerHubPassword', usernameVariable: 'dockerHubUser')]) {
        sh "docker login  -u ${dockerHubUser} -p ${dockerHubPassword}"
        sh "docker push whatlong7/jenkins-demo:${build_tag}"
    }
  }
  stage('Deploy'){
        echo "5. Deploy Stage"
        def userInput = input(
            id: 'userInput',
            message: 'Choose a deploy environment',
            parameters: [
                [
                    $class: 'ChoiceParameterDefinition',
                    choices: "Dev\nQA\nProd",
                    name: 'Env'
                ]
            ]
        )
        echo "This is a deploy step to ${userInput}"
        sh "sed -i 's/<BUILD_TAG>/${build_tag}/' k8s.yaml"
        sh "sed -i 's/<BRANCH_NAME>/${env.BRANCH_NAME}/' k8s.yaml"

        if (userInput == "Dev") {
         // deploy dev stuff
           echo "*************************use dev yaml*************************"
        } else if (userInput == "QA"){
         // deploy qa stuff
           echo "*************************use QA yaml*************************"
        } else {
           // deploy prod stuff
           echo "*************************use stuff yaml*************************"
        }
        sh "kubectl apply -f k8s.yaml"
  }
}
